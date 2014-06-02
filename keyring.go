/*
Terminal key input receiver for go application.
keyring.Bind() returns channel which returns each key input.
*/
package keyring

import (
	"github.com/pkg/term/termios"
	"syscall"
	"unicode/utf8"
)

// Disable canonical mode, input echo and signal.
// Then generate goroutine to receive all key input.
func Bind() chan rune {
	// Get current terminal parameters
	orgTerm := getCurrentTerm()

	// Change terminal parameters
	rawTerm := rawModeTerm(orgTerm)
	setTerm(&rawTerm)

	receiver := make(chan rune)
	go keyringRoutine(receiver, &orgTerm)

	return receiver
}

// Terminal input reader by syscall.Read().
// This method is for use of goroutine.
func keyringRoutine(receiver chan rune, orgTerm *syscall.Termios) {
	defer setTerm(orgTerm)
	readBuf := make([]byte, 1)
	runeBuf := []byte{}

	for {
		_, err := syscall.Read(syscall.Stdin, readBuf)
		if err != nil {
			panic(err)
		}

		// Send char only when runeBuf is valid utf-8 byte sequence
		runeBuf = append(runeBuf, readBuf[0])
		if utf8.FullRune(runeBuf) {
			ch, _ := utf8.DecodeRune(runeBuf)
			receiver <- ch
			runeBuf = []byte{}
		} else if len(runeBuf) > utf8.UTFMax {
			panic("unexpected utf-8 byte sequence")
		}
	}
}

func getCurrentTerm() syscall.Termios {
	var term syscall.Termios
	if err := termios.Tcgetattr(uintptr(syscall.Stdin), &term); err != nil {
		panic(err)
	}
	return term
}

func setTerm(term *syscall.Termios) {
	if err := termios.Tcsetattr(uintptr(syscall.Stdin), termios.TCSAFLUSH, term); err != nil {
		panic(err)
	}
}

// returns non-canonical mode term for keyring
func rawModeTerm(term syscall.Termios) syscall.Termios {
	term.Iflag &= syscall.IGNCR  // ignore received CR
	term.Lflag ^= syscall.ICANON // disable canonical mode
	term.Lflag ^= syscall.ECHO   // disable echo of input
	term.Lflag ^= syscall.ISIG   // disable signal
	term.Cc[syscall.VMIN] = 1    // number of bytes to read()
	term.Cc[syscall.VTIME] = 0   // timeout of read()
	return term
}
