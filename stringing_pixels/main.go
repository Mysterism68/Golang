package main

import (
	"fmt"
	"image/color"
	"math"

	//Ebiten
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	//"log"
	//"image/Image"
	//"time"
	//"github.com/go-vgo/screenshot"
)

var WinDim Dim2 = Dim2{800, 600}

type Game struct{}

func (g *Game) Update() error {
	WinDim.w, WinDim.h = ebiten.WindowSize()
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	vector.DrawFilledCircle(screen, 400, 300, 32, color.White, true)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World %f", math.Pi))
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WinDim.w, WinDim.h
}
func main() {
	println("Loaded")
	/*
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Window Assist")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Print(err)
	}*/
	for ;; {
		print(coloredStr("\rHello, World", 100, 0, 255))
	}
}
