package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/poccariswet/m3u8"
)

func main() {
	file, err := os.Open("sample path")
	if err != nil {
		log.Fataln(err)
	}

	playlist, listtype, err := m3u8.DecodeFrom(bufio.NewReader(file))
	if err != nil {
		log.Fataln(err)
	}

	fmt.Println(playlist)
	fmt.Println(listtype)
}
