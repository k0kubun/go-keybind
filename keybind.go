// Terminal key input receiver for go application.
// keybind.Bind() returns channel which returns each key input.
package keybind

import (
	"github.com/pkg/term/termios"
	"syscall"
	"unicode/utf8"
)

const (
	// ASCII Control Characters
	NUL = iota
	SOH
	STX
	ETX
	EOT
	ENQ
	ACK
	BEL
	BS
	HT
	LF
	VT
	FF
	CR
	SO
	SI
	DLE
	DC1
	DC2
	DC3
	DC4
	NAK
	SYN
	ETB
	CAN
	EM
	SUB
	ESC
	FS
	GS
	RS
	US
	DEL = 127
)

const (
	// Major Key input for ASCII control characters
	CTRL_Space = iota // or Ctrl+@
	CTRL_A
	CTRL_B
	CTRL_C
	CTRL_D
	CTRL_E
	CTRL_F
	CTRL_G
	BACKSPACE // or Ctrl+H
	TAB       // or Ctrl+I
	CTRL_J
	CTRL_K
	CTRL_L
	CTRL_M
	CTRL_N
	CTRL_O
	CTRL_P
	CTRL_Q
	CTRL_R
	CTRL_S
	CTRL_T
	CTRL_U
	CTRL_V
	CTRL_W
	CTRL_X
	CTRL_Y
	CTRL_Z
	ESCAPE             // or Ctrl+[
	Ctrl_BACKSLASH     // Ctrl+\
	Ctrl_RIGHT_BRACKET // Ctrl+]
	CTRL_HAT           // Ctrl+^
	CTRL_UNDERSCORE    // Ctrl+_
	DELETE             = 127
)

const (
	// control characters range (except DEL)
	controlCharactersMin = NUL
	controlCharactersMax = US
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
	go keybindRoutine(receiver, &orgTerm)

	return receiver
}

// If given character is ASCII control characters, this function returns false.
func IsPrintable(ch rune) bool {
	if controlCharactersMin <= int(ch) && int(ch) <= controlCharactersMax {
		return false
	} else if ch == DEL {
		return false
	}
	return true
}

// Terminal input reader by syscall.Read().
// This method is for use of goroutine.
func keybindRoutine(receiver chan rune, orgTerm *syscall.Termios) {
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
