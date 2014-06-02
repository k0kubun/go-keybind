package main

import (
	"fmt"
	"github.com/k0kubun/keyring"
)

func main() {
	var ch rune

	receiver := keyring.Bind()
	println("Input some keys (hit 'q' to quit):")
	for ch != 'q' {
		ch = <-receiver
		fmt.Printf("input = %c\n", ch)
	}
}
