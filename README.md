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

for {
	// Each key input will be detected as soon as possible
	ch := <-receiver
	fmt.Printf("input = %c\n", ch)

	if ch == 'q' {
		break
	}
}
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/keyring
