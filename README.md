# go-keybind

Multi-platform terminal key input reader for go language integrated with utf-8.  
You can handle all terminal input with this library.

## Installation

```bash
$ go get github.com/k0kubun/go-keybind
```

## Usage

You can enable reading all key input with `keybind.Open()`.  
`ReadRune()` reads one rune or control input.

```go
bind := keybind.Open()
defer bind.Close()

for {
  ch, _ := bind.ReadRune()
  switch ch {
  case keybind.Escape:
    fmt.Println("ESCAPE")
  case keybind.Delete:
    fmt.Println("DELETE")
  case keybind.Tab:
    fmt.Println("TAB")
  default:
    if keybind.IsPrintable(ch) {
      fmt.Printf("%c\n", ch)
    } else {
      fmt.Printf("Ctrl+%c\n", '@'+ch)
    }
  }
}
```

### Keys

You can check what is the input with following `keybind.**`

```go
CtrlSpace = iota // or Ctrl+@
CtrlA
CtrlB
CtrlC
CtrlD
CtrlE
CtrlF
CtrlG
BackSpace // or Ctrl+H
Tab       // or Ctrl+I
CtrlJ
CtrlK
CtrlL
Return // or Ctrl+M
CtrlN
CtrlO
CtrlP
CtrlQ
CtrlR
CtrlS
CtrlT
CtrlU
CtrlV
CtrlW
CtrlX
CtrlY
CtrlZ
Escape           // or Ctrl+[
CtrlBackSlash    // Ctrl+\
CtrlRightBracket // Ctrl+]
CtrlHat          // Ctrl+^
CtrlUnderscore   // Ctrl+_
Delete
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/k0kubun/go-keybind
