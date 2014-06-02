# keyring

Terminal key input receiver library for go application integrated with utf-8.  
You can handle all terminal input with this library.

## Installation

```bash
$ go get github.com/k0kubun/keyring
```

## Usage

```go
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
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/keyring
