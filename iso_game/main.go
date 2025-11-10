package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var mouse_pos Vector2

type Vector2 struct {
	x float64
	y float64
}
type Dimens2 struct {
	w float64
	h float64
}
type Game struct{}

func mouseRot() float64 {
	var mX, mY int = ebiten.CursorPosition()
	mouse_pos = Vector2{float64(mX), float64(mY)}
	var dirOfMouse float64 = float64(math.Atan2(float64(mouse_pos.y-300), float64(mouse_pos.x-400)) + math.Pi/4)
	return dirOfMouse
}
func isoVec(x float64, y float64, dir float64, xSqu float64, ySqu float64, ancX float64, ancY float64) Vector2 {
	return Vector2{
		(x*xSqu*math.Cos(dir*math.Pi/180) - y*xSqu*math.Sin(dir*math.Pi/180)) + ancX,
		(x*ySqu*math.Sin(dir*math.Pi/180) + y*ySqu*math.Cos(dir*math.Pi/180)) + ancY,
	}
}

type tile struct {
	pos    Vector2
	dim    Dimens2
	color  color.Color
	rot    float64
	squish Vector2
}

func (this *tile) Draw(screen *ebiten.Image) {
	var pointA Vector2 = isoVec(this.pos.x-this.dim.w/2, this.pos.y-this.dim.h/2, this.rot, this.squish.x, this.squish.y, 400, 300)
	var pointB Vector2 = isoVec(this.pos.x+this.dim.w/2, this.pos.y-this.dim.h/2, this.rot, this.squish.x, this.squish.y, 400, 300)
	var pointC Vector2 = isoVec(this.pos.x+this.dim.w/2, this.pos.y+this.dim.h/2, this.rot, this.squish.x, this.squish.y, 400, 300)
	var pointD Vector2 = isoVec(this.pos.x-this.dim.h/2, this.pos.y+this.dim.h/2, this.rot, this.squish.x, this.squish.y, 400, 300)
	vector.StrokeLine(screen, float32(pointA.x), float32(pointA.y), float32(pointB.x), float32(pointB.y), 8, this.color, true)
	vector.StrokeLine(screen, float32(pointB.x), float32(pointB.y), float32(pointC.x), float32(pointC.y), 8, this.color, true)
	vector.StrokeLine(screen, float32(pointC.x), float32(pointC.y), float32(pointD.x), float32(pointD.y), 8, this.color, true)
	vector.StrokeLine(screen, float32(pointD.x), float32(pointD.y), float32(pointA.x), float32(pointA.y), 8, this.color, true)
	vector.DrawFilledCircle(screen, float32(pointA.x), float32(pointA.y), 4, this.color, true)
	vector.DrawFilledCircle(screen, float32(pointB.x), float32(pointB.y), 4, this.color, true)
	vector.DrawFilledCircle(screen, float32(pointC.x), float32(pointC.y), 4, this.color, true)
	vector.DrawFilledCircle(screen, float32(pointD.x), float32(pointD.y), 4, this.color, true)
}

var ground tile = tile{
	Vector2{0, 0},
	Dimens2{768, 768},
	color.RGBA{0, 255, 0, 255},
	0,
	Vector2{1, 0.5},
}

func (g *Game) Update() error {
	const speed = 8
	var movementVector = Vector2{math.Sin(ground.rot*(math.Pi/180)) * speed,
		math.Cos(ground.rot*(math.Pi/180)) * -speed} //, ground.rot, 1, 0.5, 0, 0)
	print(fmt.Printf("\rx:%d, y:%d", int(movementVector.x), int(movementVector.y)))
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ground.rot += 6
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ground.rot -= 6
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ground.pos.x += movementVector.x
		ground.pos.y += movementVector.y
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ground.pos.x -= movementVector.x
		ground.pos.y -= movementVector.y
	}
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	ground.Draw(screen)
	var pointE Vector2 = isoVec(0, 0, ground.rot, 1, 0.5, 400, 300)
	vector.DrawFilledCircle(screen, float32(pointE.x), float32(pointE.y), 8, color.White, true)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}
func main() {

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Isometric Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSizeLimits(400, 300, 1920, 1080)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
