package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var WinWid, WinHei int = ebiten.WindowSize()
var cameraPos Vector2 = Vector2{0, 0}

type Vector2 struct {
	x float64
	y float64
}
type Dimens struct {
	w int
	h int
}
type Platform struct {
	pos Vector2
	dim Dimens
	col color.Color
}

func (this *Platform) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, this.pos.x-cameraPos.x*float64(WinWid),
		this.pos.y-cameraPos.y*float64(WinHei), float64(this.dim.w),
		float64(this.dim.h), this.col)
}
func (this *Platform) Main() int {
	return 0
}

type Player struct {
	pos Vector2
	dim Dimens
	acc Vector2
	vel Vector2
	col color.Color
}

/*
func (this *Player) CreatePlayer(vecInput *Vector2, dimInput *Dimens,

			accInput *Vector2, colorInput *color.Color) Player {
		return Player{
			pos: vecInput *Vector2,
			dim: dimInput *Vector2,
			acc: accInput *Vector2,
			vel: Vector2{0, 0},
			col: colorInput *color.Color,
		}
	}
*/
func (this *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, this.pos.x-cameraPos.x*float64(WinWid),
		this.pos.y-cameraPos.y*float64(WinHei),
		float64(this.dim.w), float64(this.dim.h), this.col)
}
func (this *Player) Main() {
	var maxHeight int = WinHei - 64
	for idx := 0; idx < len(scene); idx++ {
		var curObj Platform = scene[idx]
		var leftBound bool = this.pos.x <= curObj.pos.x+float64(curObj.dim.w)
		var rightBound bool = this.pos.x+float64(this.dim.w) >= curObj.pos.x
		var topBound bool = this.pos.y <= curObj.pos.y+float64(curObj.dim.h)
		var bottomBound bool = this.pos.y+float64(this.dim.h) >= curObj.pos.y
		if leftBound && rightBound && topBound && bottomBound {
			if this.pos.y+float64(this.dim.h) <= curObj.pos.y+float64(curObj.dim.h) {
				maxHeight = int(curObj.pos.y) - this.dim.w
			}
			if this.pos.y >= curObj.pos.y {
				this.vel.y = 0.75 * math.Abs(this.vel.y)
			}
		}
	}
	if this.pos.y >= float64(maxHeight) {
		this.vel.y = 0
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		this.vel.x += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		this.vel.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) && this.pos.y >= float64(maxHeight) {
		this.vel.y = -16.6
	}
	cameraPos.x = math.Floor(this.pos.x / float64(WinWid))
	cameraPos.y = math.Floor(this.pos.y / float64(WinHei))
	this.pos.x = this.pos.x + this.vel.x
	this.vel.x = max(min(this.vel.x*0.85, 16), -16)
	this.pos.y = min(this.pos.y+this.vel.y, float64(maxHeight))
	this.vel.y = max(min(this.vel.y+0.6, math.Inf(0) /*16*/), -16)
}

var scene [5]Platform
var player Player = Player{
	pos: Vector2{0, 0},
	dim: Dimens{64, 64},
	acc: Vector2{1, 1},
	vel: Vector2{0, 0},
	col: color.RGBA{0, 0, 255, 255},
}

type Game struct{}

func (g *Game) Update() error {
	WinWid, WinHei = ebiten.WindowSize()
	player.Main()
	for i := 0; i < len(scene); i++ {
		scene[i].Main()
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 25})
	player.Draw(screen)
	for i := 0; i < len(scene); i++ {
		scene[i].Draw(screen)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}
func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Window Assist")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	scene[0] = Platform{
		pos: Vector2{600, 600},
		dim: Dimens{96, 32},
		col: color.RGBA{255, 0, 255, 255},
	}
	scene[1] = Platform{
		pos: Vector2{400, 400},
		dim: Dimens{96, 32},
		col: color.RGBA{255, 0, 255, 255},
	}
	scene[2] = Platform{
		pos: Vector2{200, 240},
		dim: Dimens{96, 32},
		col: color.RGBA{255, 0, 255, 255},
	}
	scene[3] = Platform{
		pos: Vector2{400, 0},
		dim: Dimens{WinWid, 64},
		col: color.RGBA{255, 0, 255, 255},
	}
	scene[4] = Platform{
		pos: Vector2{float64(400 - WinWid), 0},
		dim: Dimens{
			w: 0,
			h: 64},
		col: color.RGBA{255, 0, 255, 255},
	}
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
