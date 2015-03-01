# go-keybind

Terminal key input reader for go application integrated with utf-8.  
You can handle all terminal input with this library.

## Installation

```bash
$ go get github.com/k0kubun/go-keybind
```

## Usage

```go
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
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/go-keybind
