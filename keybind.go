// Terminal key input receiver for go application.
// keybind.Bind() returns channel which returns each key input.
package keybind

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

// If given character is ASCII control characters, this function returns false.
func IsPrintable(ch rune) bool {
	if controlCharactersMin <= int(ch) && int(ch) <= controlCharactersMax {
		return false
	} else if ch == DEL {
		return false
	}
	return true
}
