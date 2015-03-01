package main

import (
	"fmt"
	"github.com/k0kubun/go-keybind"
)

func main() {
	receiver := keybind.Bind()
	println("Input some keys (hit 'q' to quit):")

	for {
		ch := <-receiver
		print("input = ")
		if keybind.IsPrintable(ch) {
			fmt.Printf("%c\n", ch)
		} else {
			switch ch {
			case keybind.ESCAPE:
				fmt.Println("ESCAPE")
			case keybind.DELETE:
				fmt.Println("DELETE")
			case keybind.TAB:
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
