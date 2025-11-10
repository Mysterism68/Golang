package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Vector2 struct {
	x int
	y int
}
type Dim2 struct {
	w int
	h int
}

func RGBToHex(r, g, b int) string {
	if r < 0 {
		r = 0
	} else if r > 255 {
		r = 255
	}
	if g < 0 {
		g = 0
	} else if g > 255 {
		g = 255
	}
	if b < 0 {
		b = 0
	} else if b > 255 {
		b = 255
	}
	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

var tick int = 0

func coloredStr(str string, red int, green int, blue int) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(RGBToHex(red, green, blue))).Render(str)
}
