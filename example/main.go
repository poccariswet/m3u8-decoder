package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/poccariswet/m3u8"
)

func main() {
	file, err := os.Open("./master.m3u8")
	if err != nil {
		log.Fatalln(err)
	}

	playlist, listtype, err := m3u8.DecodeFrom(bufio.NewReader(file))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(playlist)
	fmt.Println(listtype)
}
