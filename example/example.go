package main

import (
	"fmt"
	"github.com/k0kubun/go-keybind"
)

func main() {
	bind := keybind.Open()
	defer bind.Close()

	fmt.Println("Input some keys (hit 'q' to quit):")

	for {
		ch, _ := bind.ReadRune()
		fmt.Print("input = ")

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
