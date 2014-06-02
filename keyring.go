/*
Terminal key input manager for go application
*/
package keyring

import (
	"github.com/pkg/term/termios"
	"syscall"
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
	buf := make([]byte, 1)

	for buf[0] != 'q' {
		_, err := syscall.Read(syscall.Stdin, buf)
		if err != nil {
			panic(err)
		}
		receiver <- rune(buf[0])
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
