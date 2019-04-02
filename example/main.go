package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/poccariswet/m3u8-decoder"
)

func main() {
	file, err := os.Open("./media.m3u8")
	if err != nil {
		log.Fatalln(err)
	}

	playlist, listtype, err := m3u8.DecodeFrom(bufio.NewReader(file))
	if err != nil {
		log.Fatalln(err)
	}

	switch listtype {
	case m3u8.MASTER:
		fmt.Println(playlist)
	case m3u8.MEDIA:
		fmt.Println(playlist)
	default:
	}
}
