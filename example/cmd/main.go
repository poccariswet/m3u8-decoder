package main

import (
	"fmt"
	"os"

	"github.com/poccariswet/m3u8-decoder"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Args is not 2\nexample command <.m3u8>")
	}
	f := os.Args[1]

	playlist, err := m3u8.ReadFile(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(playlist)
}
