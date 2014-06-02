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
		print("input = ")
		if keyring.IsPrintable(ch) {
			fmt.Printf("%c\n", ch)
		} else {
			switch ch {
			case keyring.ESCAPE:
				fmt.Println("ESCAPE")
			case keyring.DELETE:
				fmt.Println("DELETE")
			case keyring.TAB:
				fmt.Println("TAB")
			default:
				fmt.Printf("Ctrl+%c\n", '@'+ch)
			}
		}

		if ch == 'q' {
			break
		}
	}
}
