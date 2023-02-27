package ansi

import (
	"fmt"
)

const ANSIReset string = "\x1b[m"

type Color interface {
	Foreground() string
	Background() string
}

var _ Color = ColorANSI(0)

type ColorANSI uint8

const (
	ColorANSIBlack   ColorANSI = 0
	ColorANSIRed     ColorANSI = 1
	ColorANSIGreen   ColorANSI = 2
	ColorANSIYellow  ColorANSI = 3
	ColorANSIBlue    ColorANSI = 4
	ColorANSIMagenta ColorANSI = 5
	ColorANSICyan    ColorANSI = 6
	ColorANSIWhite   ColorANSI = 7
)

func (c ColorANSI) Foreground() string {
	return fmt.Sprintf("\x1b[3%dm", uint8(c))
}

func (c ColorANSI) Background() string {
	return fmt.Sprintf("\x1b[4%dm", uint8(c))
}

var _ Color = Color256(0)

type Color256 uint8

func (c Color256) Foreground() string {
	return fmt.Sprintf("\x1b[38;5;%dm", uint8(c))
}

func (c Color256) Background() string {
	return fmt.Sprintf("\x1b[48;5;%dm", uint8(c))
}
