package colors

import (
	"github.com/fatih/color"
)

var (
	M  = color.New(color.FgHiMagenta, color.Bold)
	W  = color.New(color.FgHiWhite, color.Bold)
	Y  = color.New(color.FgHiYellow, color.Bold)
	G  = color.New(color.FgHiGreen, color.Bold)
	Bc = color.New(color.FgHiCyan, color.BlinkSlow, color.Bold)
	Bg = color.New(color.FgHiGreen, color.BlinkSlow, color.Bold)
	Br = color.New(color.FgHiRed, color.BlinkSlow, color.Bold)

	R   = color.New(color.FgHiRed, color.Bold)
	B   = color.New(color.FgHiBlue, color.Bold)
	FgR = color.New(color.FgRed, color.Bold)
	FgY = color.New(color.FgYellow, color.Bold)
)

func Ms(s string) string {
	ms := color.HiMagentaString(s)
	return ms
}

func Gs(s string) string {
	ms := color.HiGreenString(s)
	return ms
}
