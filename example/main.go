package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/poccariswet/m3u8-decoder"
)

func main() {
	path := os.Getenv("GOPATH")
	if path == "" {
		panic("$GOPATH is empty")
	}

	// run関数に分離するのが結構一般的
	master, err := os.Open(path + "/src/github.com/poccariswet/m3u8-decoder/example/playlist/master.m3u8")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	media, err := os.Open(path + "/src/github.com/poccariswet/m3u8-decoder/example/playlist/media.m3u8")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	playlist, err := m3u8.DecodeFrom(bufio.NewReader(master))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(playlist)

	playlist, err = m3u8.DecodeFrom(bufio.NewReader(media))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(playlist)
}
