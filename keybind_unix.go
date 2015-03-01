// +build !windows

package keybind

import (
	"errors"
	"syscall"
	"unicode/utf8"

	"github.com/pkg/term/termios"
)

type Term struct {
	orgTerm syscall.Termios
}

// To receive all input, disable canonical mode, input echo and signal.
func Open() *Term {
	// Get current terminal parameters
	orgTerm := getTerm()

	rawTerm := rawModeTerm(orgTerm)
	setTerm(rawTerm)

	return &Term{
		orgTerm: orgTerm,
	}
}

// Read one rune or control sequence.
func (t *Term) ReadRune() (rune, error) {
	readBuf := make([]byte, 1)
	runeBuf := []byte{}

	for {
		_, err := syscall.Read(syscall.Stdin, readBuf)
		if err != nil {
			return NUL, err
		}

		// Send char only when runeBuf is valid utf-8 byte sequence
		runeBuf = append(runeBuf, readBuf[0])
		if utf8.FullRune(runeBuf) {
			ch, _ := utf8.DecodeRune(runeBuf)
			return ch, nil
		} else if len(runeBuf) > utf8.UTFMax {
			return NUL, errors.New("invalid byte sequence as utf-8")
		}
	}
}

// Reset terminal to canonical mode.
func (t *Term) Close() {
	setTerm(t.orgTerm)
}

func getTerm() syscall.Termios {
	var term syscall.Termios
	if err := termios.Tcgetattr(uintptr(syscall.Stdin), &term); err != nil {
		panic(err)
	}
	return term
}

func setTerm(term syscall.Termios) {
	if err := termios.Tcsetattr(uintptr(syscall.Stdin), termios.TCSAFLUSH, &term); err != nil {
		panic(err)
	}
}

// returns non-canonical mode term for keybind
func rawModeTerm(term syscall.Termios) syscall.Termios {
	term.Iflag &= syscall.IGNCR  // ignore received CR
	term.Lflag ^= syscall.ICANON // disable canonical mode
	term.Lflag ^= syscall.ECHO   // disable echo of input
	term.Lflag ^= syscall.ISIG   // disable signal
	term.Cc[syscall.VMIN] = 1    // number of bytes to read()
	term.Cc[syscall.VTIME] = 0   // timeout of read()
	return term
}
