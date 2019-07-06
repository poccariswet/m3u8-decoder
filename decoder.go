package m3u8

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// decode parses a playlist
func decode(buf *bytes.Buffer) (*Playlist, error) {
	playlist := NewPlaylist()
	var end bool
	states := new(States)

	for !end {
		line, err := buf.ReadString('\n')
		// ここはEOFの場合、lineは帰ってこないはずなので、breakすべき、もし他の理由があるのならば、
		// 無駄に複雑にする必要性はない、nil出ないものの処理の集合として表現したほうが、その中で、
		// errがnil出ない処理であることが表現できる
		if err != nil {
			if err == io.EOF {
				end = true
			}

			return nil, err
		}

		if len(line) < 1 || line == "\r" {
			continue
		}

		line = strings.TrimSpace(line)
		// ここの処理に行く前に純粋な形に直すべき、decodeLineという名前なので、ここがわかりにくくなっているが、この中で行われているのは、
		// 文字列のパースと型の判定、そして、目的の構造化の三つの責務を背負ってしまっているため、内部的な責務がでかすぎる
		// 関数は常に一つのことに集中すべきだ、なので、個人的には、ココそmapにすべきだと思う、
		//
		if err := decodeLine(playlist, line, states); err != nil {
			return playlist, err
		}
	}

	return playlist, nil
}

// DecodeFrom read a playlist passed from the io.Reader
func DecodeFrom(r io.Reader) (*Playlist, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	//bufio.NewReader(r)でこいつを受け取ったほうがいい、
	// buf,ReadFromしてるのに、次の関数でbytes.Bufferとして使っているので、ここで、意味がないので（コピーが走っているから）直したほうがよさそう
	return decode(buf)
}

// ReadFile reads contents from filepath and return Playlist
func ReadFile(path string) (*Playlist, error) {
	// ここでは二つの責務ReadFileとDecodeを行ってしまっているので、分離した方が良い
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "ReadFile err")
	}

	return decode(bytes.NewBuffer(file))
}

func decodeLine(p *Playlist, line string, s *States) error {
	// OSS的にはErrorは予測される全て外部定義されているべき
	// ガード節を適用
	// ここの条件式もstrings.HasPrefixがほとんどでおそらく、意味合い的にもこれを使うところが正しいので、
	// 条件のbool値を返す関数を用意した方が良い
	// decodeLine decodes a line of playlist and parses
	// lineをまずは構造体にした方が良い気がする, そして、lineにStatusを渡すと適切なものが帰ってくる方が自然な書き方な様に見える
	// さらに、それぞれのフィールドに値が入っているかどうかで、bool値を返す様にする
	// それぞれのNewメソッドの中にパースの責務も混じってしまっているので、itemから生成できる様にした方が自然に見える
	// Statusをポインタにして、内部で扱っているのは悪くはないけど、参照と値の区別がわかりにくくなり、バグの元になりことが多いので、個人的には
	// ここは毎回returnしてあげた方が良さそうそして、内部で値が変更されるものは直接触らない方がいい、参照だったとしても内部では別の変数が触り、
	// 最後に入れる様にした方がいい
	// 条件がそれぞれ、そのまま関数を使っているので、これらは全てラップした関数にしたほうが良い、なぜなら、処理の分離が測ることができ、変更が容易になるので、
	// 現状副作用がでかすぎるので、直していきたい

	if !s.m3u8 && line != EXTM3U {
		return errors.New("invalid playlist, not exist #EXTM3U")
	}

	switch {
	case line == EXTM3U:
		s.m3u8 = true
	case strings.HasPrefix(line, ExtENDList):
		p.live = false
	case strings.HasPrefix(line, ExtVersion):
		p.hasVersion = true
		// ここのScanも明確な名称をつけて別にするべき、
		_, err := fmt.Sscanf(line, ExtVersion+":%d", &p.Version)
		if err != nil {
			return errors.Wrap(err, "invalid scan version")
		}
	case strings.HasPrefix(line, EXTINF):
		inf, err := NewExtInf(line)
		if err != nil {
			return errors.Wrap(err, "new extinf err")
		}
		p.master = false
		s.segment = inf
		s.segmentTag = true

	case strings.HasPrefix(line, ExtMedia):
		m, err := NewMedia(line)
		if err != nil {
			return errors.Wrap(err, "new media err")
		}
		p.AppendSegment(m)
	case strings.HasPrefix(line, ExtStreamInf):
		p.master = true
		s.segmentTag = true
		// ここも別でlineを返すのと再代入されているので、意味合いがわかりずらくなっているのと、バグの元になってしまっている
		line = line[len(ExtStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		s.segment = v
	case strings.HasPrefix(line, ExtFrameStreamInf):
		p.master = true
		s.segmentTag = false
		// ここも別でlineを返すのと再代入されているので、意味合いがわかりずらくなっているのと、バグの元になってしまっている
		line = line[len(ExtFrameStreamInf+":"):]
		v, err := NewVariant(line)
		if err != nil {
			return errors.Wrap(err, "new variant err")
		}
		v.IFrame = true
		s.segment = v
		p.AppendSegment(v)
	case strings.HasPrefix(line, ExtByteRange):
		br, err := NewByteRange(line)
		if err != nil {
			return errors.Wrap(err, "new byte range err")
		}
		br.Extflag = true
		// ガード節
		if m, has := s.segment.(*MapSegment); has {
			m.ByteRange = br
			s.segment = m
			br.Extflag = false
		} else if inf, has := s.segment.(*InfSegment); has {
			inf.ByteRange = br
			s.segment = inf
		}
	case strings.HasPrefix(line, ExtMap):
		m, err := NewMap(line)
		if err != nil {
			return errors.Wrap(err, "new map err")
		}
		p.AppendSegment(m)
	case strings.HasPrefix(line, ExtKey):
		key, err := NewKey(line)
		if err != nil {
			return errors.Wrap(err, "new key err")
		}
		p.AppendSegment(key)
	case strings.HasPrefix(line, ExtProgramDateTime):
		dt, err := NewProgramDateTime(line)
		if err != nil {
			return errors.Wrap(err, "new program date time err")
		}
		p.AppendSegment(dt)
	case strings.HasPrefix(line, ExtDateRange):
		dr, err := NewDateRange(line)
		if err != nil {
			return errors.Wrap(err, "new date range err")
		}
		p.AppendSegment(dr)

		/* session tags */
	case strings.HasPrefix(line, ExtSessionKey):
		sk, err := NewSessionKey(line)
		if err != nil {
			return errors.Wrap(err, "new session key err")
		}
		p.AppendSegment(sk)
	case strings.HasPrefix(line, ExtSessionData):
		sd, err := NewSessionData(line)
		if err != nil {
			return errors.Wrap(err, "new session data err")
		}
		p.AppendSegment(sd)

	case strings.HasPrefix(line, ExtStart):
		start, err := NewStart(line)
		if err != nil {
			return errors.Wrap(err, "new start err")
		}
		p.AppendSegment(start)
	case strings.HasPrefix(line, ExtIndependentSegments):
		p.IndependentSegments = true

		/* playlist tags */
	case strings.HasPrefix(line, ExtPlaylistType):
		_, err := fmt.Sscanf(line, ExtPlaylistType+":%s", &p.PlaylistType)
		if err != nil {
			return errors.Wrap(err, "invalid scan playlist type")
		}
	case strings.HasPrefix(line, ExtIFramesOnly):
		p.IFrameOnly = true
	case strings.HasPrefix(line, ExtTargetDutation):
		_, err := fmt.Sscanf(line, ExtTargetDutation+":%f", &p.TargetDuration)
		if err != nil {
			return errors.Wrap(err, "invalid scan TargetDuration")
		}
	case strings.HasPrefix(line, ExtDiscontinuitySequence):
		_, err := fmt.Sscanf(line, ExtTargetDutation+":%d", &p.DiscontinuitySequence)
		if err != nil {
			return errors.Wrap(err, "invalid scan DiscontinuitySequence")
		}
	case strings.HasPrefix(line, ExtAllowCache):
		p.AllowCache = parseBool(line[len(ExtAllowCache+":"):])
	case strings.HasPrefix(line, ExtMediaSequence):
		// ここも関数で切り出すべき
		_, err := fmt.Sscanf(line, ExtTargetDutation+":%d", &p.MediaSequence)
		if err != nil {
			return errors.Wrap(err, "invalid scan MediaSequence")
		}
	default:
		line = strings.Trim(line, "\n")
		uri := strings.TrimSpace(line)
		// ガード節とドモルガンで浅くできそう

		if s.segment != nil && s.segmentTag {
			if p.master {
				v, has := s.segment.(*VariantSegment)
				if !has {
					return errors.New("invalid variant playlist")
				}
				v.URI = uri
				p.AppendSegment(v)
			} else {
				i, has := s.segment.(*InfSegment)
				if !has {
					return errors.New("invalid EXTINF segment")
				}
				i.URI = uri
				p.AppendSegment(i)
			}
			s.segmentTag = false

			return nil
		}
	}
	return nil
}
