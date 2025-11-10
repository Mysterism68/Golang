package main

import (
	"fmt"
	"math"

	"github.com/Kalyle-68/Golang/extras"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"github.com/hajimehoshi/ebiten/v2/vector"
)

var nearPlane float64 = 0
var farPlane float64 = 32

func persProj(oriPoint extras.Vector3) extras.Vector2 {
	return extras.PersProj(oriPoint, WinDim)
}

func rotatePoint(point extras.Vector3) extras.Vector3 {
	return extras.RotatePoint(point, CamPos, CamRot, nearPlane, farPlane, false)
}

func drawTriangles(screen *ebiten.Image) {
	var NewEnv = extras.ReorderTriangles(CamPos, Env)
	for triangle := 0; triangle < len(NewEnv); triangle++ {
		var NewEnv2D []ebiten.Vertex
		var indices []uint16
		var smallDistToCen float64 = 18_446_744_073_709_551_615
		var pointsBehindCam = 0
		for vertex := 0; vertex < len(NewEnv[triangle].Vertices); vertex++ {
			var point3D extras.Vector3 = rotatePoint(NewEnv[triangle].Vertices[vertex])
			var point2D extras.Vector2 = persProj(point3D)
			smallDistToCen = math.Min(smallDistToCen, math.Hypot(point2D.X-WinDim.W/2, point2D.Y-WinDim.H/2))
			if point3D.Z < 0 {
				pointsBehindCam++
			}
		}
		if smallDistToCen < math.Max(WinDim.W/2, WinDim.H/2) || pointsBehindCam < 3 {
			for vertex := 0; vertex < len(NewEnv[triangle].Vertices); vertex++ {
				var point extras.Vector2 = persProj(rotatePoint(NewEnv[triangle].Vertices[vertex]))
				//var R, G, B, A = NewEnv[triangle].color.RGBA()
				smallDistToCen = math.Min(smallDistToCen, math.Hypot(point.X-WinDim.W/2, point.Y-WinDim.H/2))
				var srcX float32 = 0
				var srcY float32 = 0
				NewEnv2D = append(NewEnv2D, ebiten.Vertex{DstX: float32(point.X), DstY: float32(point.Y),
					SrcX: srcX, SrcY: srcY, ColorR: 1, ColorG: 1, ColorB: 1, ColorA: 1})
				indices = append(indices, uint16(vertex))
			}
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Triangle %+v:%t", triangle+1,
				smallDistToCen < math.Max(WinDim.W/2, WinDim.H/2)), (100*triangle)%600, int(math.Floor(float64((100*triangle)/600))*40))
			texture := ebiten.NewImage(1, 1)
			texture.Fill(NewEnv[triangle].Color)
			options := &ebiten.DrawTrianglesOptions{}
			//polygon := ebiten.NewImage(int(WinDim.W), int(WinDim.H))
			screen.DrawTriangles(NewEnv2D, indices, texture, options)
			//screen.DrawImage(polygon, nil)
		}
	}
}
