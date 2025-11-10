package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var win_dim = Vector2{800, 600}
var mouse_pos = Vector2{0, 0}
var pixel_size = 1
var moon_mass = math.Pow(4096, 2)

type Game struct {
	zoom    float32
	cam_pos Vector2
}

var particles = []Particle{
	newParticle(Vector2{400, 300}, 64, color.RGBA{255, 255, 0, 255}, float64(moon_mass)*81.27, 0, 0, true),
}

func camHandling(cam *Game, speed float32, magn float32) {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cam.cam_pos.x -= speed / cam.zoom
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		cam.cam_pos.x += speed / cam.zoom
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		cam.cam_pos.y -= speed / cam.zoom
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		cam.cam_pos.y += speed / cam.zoom
	}
	if ebiten.IsKeyPressed(ebiten.KeyMinus) {
		cam.zoom -= magn
	}
	if ebiten.IsKeyPressed(ebiten.KeyEqual) {
		cam.zoom += magn
	}
	cam.zoom = max(cam.zoom, 0.0025)
}

func (g *Game) Update() error {
	var x, y = ebiten.CursorPosition()
	mouse_pos = Vector2{((float32(x) + g.cam_pos.x) * float32(pixel_size)) / g.zoom, ((float32(y) + g.cam_pos.y) * float32(pixel_size)) / g.zoom}
	camHandling(g, 2, 0.0125)
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		//This adds to the previously added movement
		camHandling(g, 8, 0)
	}
	for idx := 0; idx < len(particles); idx++ {
		particles[idx].Update(particles, 6.674*math.Pow(10, -11))
	} /*
		g.cam_pos.x = particles[0].pos.x - 400
		g.cam_pos.y = particles[0].pos.y - 300*/
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return ebiten.Termination
	}
	return nil
}

func (g *Game) DrawOutline(screen *ebiten.Image, black float32, white float32) {
	for color_int := 1; color_int > -1; color_int-- {
		for _, part := range particles {
			fmt.Print( /*"Normal Pos: ", part.pos.x, "Camera Pos: ", (part.pos.x-g.cam_pos.x)*g.zoom, "\r"*/ )
			vector.FillCircle(screen, ((part.pos.x-g.cam_pos.x)*g.zoom)/float32(pixel_size), ((part.pos.y-g.cam_pos.y)*g.zoom)/float32(pixel_size),
				((part.diam+black+white*float32(color_int))/2*g.zoom)/float32(math.Abs(float64(pixel_size))), color.RGBA{uint8(255 * color_int),
					uint8(255 * color_int), uint8(255 * color_int), 255}, true)
		}
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{8, 0, 20, 255})
	g.DrawOutline(screen, 6, 6)
	for i := 0; i < len(particles); i++ {
		particles[i].Draw(screen, *g)
	}
	ebitenutil.DebugPrint(screen, "Zoom: x"+fmt.Sprint(g.zoom))
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var x, y = ebiten.WindowSize()
	return int(math.Abs(float64(x / pixel_size))), int(math.Abs(float64(y / pixel_size)))
}
func main() {
	ebiten.SetWindowSize(int(win_dim.x), int(win_dim.y))
	ebiten.SetWindowTitle("Gravity Simulator")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	for rot := float64(0); rot < 360; rot += 180 {
		particles = append(particles, newParticle(Vector2{float32(math.Sin(rot*float64(
			math.Pi/180))*192 + 400), float32(math.Cos(rot*float64(math.Pi/180))*192 + 300)},
			32, color.RGBA{200, 200, 255, 255}, float64(moon_mass), float32(rot+90), 14, false))
	}
	if err := ebiten.RunGame(&Game{1, Vector2{0, 0}}); err != nil {
		log.Fatal(err)
	}
}
