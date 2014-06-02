package main

import (
	"fmt"
	"github.com/k0kubun/keyring"
)

func main() {
	receiver := keyring.Bind()
	println("Input some keys (hit 'q' to quit):")

	for {
		ch := <-receiver
		fmt.Printf("input = %c\n", ch)

		if ch == 'q' {
			break
		}
	}
}
