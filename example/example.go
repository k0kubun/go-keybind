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
		ch, err := bind.ReadRune()
		if err != nil {
			panic(err)
		}
		fmt.Print("input = ")

		if keybind.IsPrintable(ch) {
			fmt.Printf("%c\n", ch)
		} else {
			switch ch {
			case keybind.Escape:
				fmt.Println("ESCAPE")
			case keybind.Delete:
				fmt.Println("DELETE")
			case keybind.Tab:
				fmt.Println("TAB")
			default:
				fmt.Printf("Ctrl+%c\n", '@'+ch)
			}
		}

		if ch == 'q' {
			bind.Close()
			break
		}
	}
}
