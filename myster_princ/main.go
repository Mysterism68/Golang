package main

import (
	//"fmt"
	//"bufio"
	//"os"
	"log"
	"math"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var organisms [][]int32
var neighborhood [][]int32
var timer int
var WinDim Dim2 = Dim2{800, 600}
var mouse Vector2 = Vector2{0, 0}
var selTile Vector2
var scale float32 = 12

type Game struct{}

var isDefiningEnv bool = true

func (g *Game) Update() error {
	timer++
	mouse.x, mouse.y = ebiten.CursorPosition()
	WinDim.w, WinDim.h = ebiten.WindowSize()
	if isDefiningEnv {
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			convMtrxToFile("layout.txt", organisms)
		}
		if ebiten.IsKeyPressed(ebiten.KeyI) {

		}
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			isDefiningEnv = false
		}
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) && timer >= 20 {
			timer = 0
			organisms[selTile.y][selTile.x] = int32(math.Abs(float64(organisms[selTile.y][selTile.x] - 1)))
		}
	} else {
		updateOrg()
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(WinDim.w), float32(WinDim.h), color.RGBA{0, 0, 40, 26}, true)
	drawMatrix(screen, organisms, 400, 300, scale)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WinDim.w, WinDim.h
}

func main() {
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)*/
	convFileToMtrx("layout.txt", &organisms)
	convFileToMtrx("nb.txt", &neighborhood)
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Mystery Prinples")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
