package main

import (
	//"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Vector2 struct {
	x float32
	y float32
}

type Particle struct {
	pos   Vector2
	vel   Vector2
	diam  float32
	color color.Color
	mass  float64
	anc   bool
}

func newParticle(pos Vector2, diam float32, color color.Color, mass float64, vel_rot float32, magn float32, anc bool) Particle {
	return Particle{pos, Vector2{float32(math.Sin(float64(vel_rot*float32(math.Pi/180)))) * magn, float32(math.Cos(
		float64(vel_rot*float32(math.Pi/180)))) * magn}, diam, color, mass, anc}
}

func (this *Particle) Draw(screen *ebiten.Image, cam Game) {
	vector.FillCircle(screen, ((this.pos.x-cam.cam_pos.x)*cam.zoom)/float32(pixel_size),
		((this.pos.y-cam.cam_pos.y)*cam.zoom)/float32(pixel_size), float32(math.Max(float64(this.diam/2*cam.zoom), 4)/
			math.Abs(float64(pixel_size))), this.color, true)
}
func (this *Particle) Update(allParticles []Particle, grav float64) {
	this.pos.x += this.vel.x
	this.pos.y += this.vel.y
	for idx := 0; idx < len(allParticles); idx++ {
		if &allParticles[idx] == this || this.anc {
			continue
		}
		var dist = math.Hypot(float64(this.pos.x-allParticles[idx].pos.x), float64(this.pos.y-allParticles[idx].pos.y))
		var rel_dir = math.Atan2(float64(allParticles[idx].pos.y-this.pos.y), float64(allParticles[idx].pos.x-this.pos.x))
		var speed = grav * (this.mass * allParticles[idx].mass) / math.Pow(dist, 2) / 60
		if dist > float64((allParticles[idx].diam+this.diam)/2) {
			this.vel.x += float32(math.Cos(rel_dir) * speed)
			this.vel.y += float32(math.Sin(rel_dir) * speed)
			//allParticles[idx].vel.x -= float32(math.Cos(rel_dir) * speed*0)
			//allParticles[idx].vel.y -= float32(math.Sin(rel_dir) * speed*0)
		} else {
			this.vel.x = 0
			this.vel.y = 0
			this.pos.x = allParticles[idx].pos.x - float32(math.Cos(rel_dir)*float64((allParticles[idx].diam+this.diam)/2))
			this.pos.y = allParticles[idx].pos.y - float32(math.Sin(rel_dir)*float64((allParticles[idx].diam+this.diam)/2))
		}
	}
	var dist = float32(math.Hypot(float64(this.pos.x-mouse_pos.x), float64(this.pos.y-mouse_pos.y)))
	if dist <= this.diam/2 && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		this.vel.x = mouse_pos.x - this.pos.x
		this.vel.y = mouse_pos.y - this.pos.y
	}
}
