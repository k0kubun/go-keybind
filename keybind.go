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
	Delete           = 127
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
