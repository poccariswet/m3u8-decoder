package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/poccariswet/m3u8-decoder"
)

func main() {
	media, err := os.Open("./media.m3u8")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	master, err := os.Open("./master.m3u8")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	playlist, listtype, err := m3u8.DecodeFrom(bufio.NewReader(media))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch listtype {
	case m3u8.MASTER:
		fmt.Println(playlist)
	case m3u8.MEDIA:
		fmt.Println(playlist)
	default:
	}

	playlist, listtype, err = m3u8.DecodeFrom(bufio.NewReader(master))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch listtype {
	case m3u8.MASTER:
		fmt.Println(playlist)
	case m3u8.MEDIA:
		fmt.Println(playlist)
	default:
	}
}
