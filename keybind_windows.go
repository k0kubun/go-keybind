package keybind

import (
	"os"
	"syscall"
	"unsafe"

	"github.com/mattn/go-isatty"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle     = kernel32.NewProc("GetStdHandle")
	procSetConsoleMode   = kernel32.NewProc("SetConsoleMode")
	procGetConsoleMode   = kernel32.NewProc("GetConsoleMode")
	procReadConsoleInput = kernel32.NewProc("ReadConsoleInputW")
)

const (
	enableLineInput       = 2
	enableEchoInput       = 4
	enableProcessedInput  = 1
	enableWindowInput     = 8
	enableMouseInput      = 16
	enableInsertMode      = 32
	enableQuickEditMode   = 64
	enableExtendedFlags   = 128
	enableAutoPosition    = 256
	enableProcessedOutput = 1
	enableWrapAtEolOutput = 2

	keyEvent              = 0x1
	mouseEvent            = 0x2
	windowBufferSizeEvent = 0x4
)

type Term struct {
	in      uintptr
	orgTerm uint32
}

type wchar uint16
type short int16
type dword uint32
type word uint16

type inputRecord struct {
	eventType word
	padding   [2]byte
	event     [16]byte
}

type keyEventRecord struct {
	keyDown         int32
	repeatCount     word
	virtualKeyCode  word
	virtualScanCode word
	unicodeChar     wchar
	controlKeyState dword
}

func Open() *Term {
	var in uintptr
	if isatty.IsTerminal(os.Stdin.Fd()) {
		in = getStdHandle(syscall.STD_INPUT_HANDLE)
	} else {
		conin, _ := os.Open("CONIN$")
		in = conin.Fd()
	}

	var orgTerm uint32
	procGetConsoleMode.Call(in, uintptr(unsafe.Pointer(&orgTerm)))

	rawTerm := orgTerm
	rawTerm &^= (enableEchoInput | enableLineInput)
	procSetConsoleMode.Call(in, uintptr(rawTerm))

	return &Term{
		in:      in,
		orgTerm: orgTerm,
	}
}

func (t *Term) ReadRune() (rune, error) {
	for {
		var ir inputRecord
		var w uint32
		ok, _, err := procReadConsoleInput.Call(t.in, uintptr(unsafe.Pointer(&ir)), 1, uintptr(unsafe.Pointer(&w)))
		if ok == 0 && err != nil {
			return NUL, err
		}

		switch ir.eventType {
		case keyEvent:
			kr := (*keyEventRecord)(unsafe.Pointer(&ir.event))
			if kr.keyDown != 0 {
				ch := rune(kr.unicodeChar)
				if ch > 0 {
					return ch, nil
				}
			}
		}
	}
}

func (t *Term) Close() {
	procSetConsoleMode.Call(t.in, uintptr(t.orgTerm))
}

func getStdHandle(stdhandle int32) uintptr {
	handle, _, _ := procGetStdHandle.Call(uintptr(stdhandle))
	return handle
}
