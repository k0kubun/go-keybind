# go-keybind

Multi-platform terminal key input reader for go language integrated with utf-8.  
You can handle all terminal input with this library.

## Installation

```bash
$ go get github.com/k0kubun/go-keybind
```

## Usage

```go
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
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/go-keybind
