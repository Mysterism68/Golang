package extras

import (
	"image/color"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

type Dim2 struct {
	W float64
	H float64
}

type Dim3 struct {
	W float64
	H float64
	D float64
}

func MouseRot(x float64, y float64, ancX float64, ancY float64) float64 {
	var pos = Vector2{float64(x), float64(y)}
	var rot float64 = float64(math.Atan2(float64(pos.Y-300), float64(pos.X-400)) + math.Pi/4)
	return rot
}

func IsoVec(x float64, y float64, dir float64, xSqu float64, ySqu float64, ancX float64, ancY float64) Vector2 {
	return Vector2{
		(x*xSqu*math.Cos(dir*math.Pi/180) - y*xSqu*math.Sin(dir*math.Pi/180)) + ancX,
		(x*ySqu*math.Sin(dir*math.Pi/180) + y*ySqu*math.Cos(dir*math.Pi/180)) + ancY,
	}
}

func DegToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func RadToDeg(rad float64) float64 {
	return rad / (math.Pi / 180)
}

func RotatePoint(oriPoint Vector3, CamPos Vector3, CamRot Vector3, nearPlane float64, farPlane float64, clip bool) Vector3 {
	point := Vector3{
		X: oriPoint.X - CamPos.X,
		Y: oriPoint.Y - CamPos.Y,
		Z: oriPoint.Z - CamPos.Z,
	}
	camRad := Vector3{
		X: DegToRad(-CamRot.X),
		Y: DegToRad(CamRot.Y),
		Z: DegToRad(CamRot.Z),
	}
	tempX := point.X*math.Cos(camRad.X) + point.Z*math.Sin(camRad.X)
	tempZ := -point.X*math.Sin(camRad.X) + point.Z*math.Cos(camRad.X)
	point.X = tempX
	point.Z = tempZ
	tempY := point.Y*math.Cos(camRad.Y) - point.Z*math.Sin(camRad.Y)
	tempZ = point.Y*math.Sin(camRad.Y) + point.Z*math.Cos(camRad.Y)
	point.Y = tempY
	point.Z = tempZ
	tempX = point.X*math.Cos(camRad.Z) - point.Y*math.Sin(camRad.Z)
	tempY = point.X*math.Sin(camRad.Z) + point.Y*math.Cos(camRad.Z)
	point.X = tempX
	point.Y = tempY
	if clip && (point.Z < nearPlane || point.Z > farPlane) {
		point.Z = 0
	}
	return point
}

func PersProj(point Vector3, WinDim Dim2) Vector2 {
	scale := Dim2{
		W: math.Min(WinDim.W/2, WinDim.H/2),
		H: math.Min(WinDim.W/2, WinDim.H/2),
	}
	screenX := (point.X/point.Z)*scale.W + WinDim.W/2
	screenY := -(point.Y/point.Z)*scale.H + WinDim.H/2
	return Vector2{X: screenX, Y: screenY}
}

type Triangle struct {
	Vertices []Vector3
	Color    color.Color
}

func ReorderTriangles(anchor Vector3, triangles []Triangle) []Triangle {
	var finalTriangles []Triangle
	for tri := 0; tri < len(triangles); tri++ {
		if tri == 0 {
			finalTriangles = append(finalTriangles, triangles[tri])
		} else {
			var averPoint = Vector3{0, 0, 0}
			for vertex := 0; vertex < len(triangles[tri].Vertices); vertex++ {
				averPoint.X += triangles[tri].Vertices[vertex].X
				averPoint.Y += triangles[tri].Vertices[vertex].Y
				averPoint.Z += triangles[tri].Vertices[vertex].Z
			}
			averPoint.X /= float64(len(triangles[tri].Vertices))
			averPoint.Y /= float64(len(triangles[tri].Vertices))
			averPoint.Z /= float64(len(triangles[tri].Vertices))
			var newX = math.Pow(averPoint.X-anchor.X, 2)
			var newY = math.Pow(averPoint.Y-anchor.Y, 2)
			var newZ = math.Pow(averPoint.Z-anchor.Z, 2)
			var relDist = math.Sqrt(newX + newY + newZ)
			for idx := 0; idx < len(triangles[tri].Vertices); idx++ {
				var averPoint2 = Vector3{0, 0, 0}
				for vertex := 0; vertex < len(finalTriangles[idx].Vertices); vertex++ {
					averPoint2.X += finalTriangles[idx].Vertices[vertex].X
					averPoint2.Y += finalTriangles[idx].Vertices[vertex].Y
					averPoint2.Z += finalTriangles[idx].Vertices[vertex].Z
				}
				averPoint2.X /= float64(len(finalTriangles[idx].Vertices))
				averPoint2.Y /= float64(len(finalTriangles[idx].Vertices))
				averPoint2.Z /= float64(len(finalTriangles[idx].Vertices))
				var newX2 = math.Pow(averPoint2.X-anchor.X, 2)
				var newY2 = math.Pow(averPoint2.Y-anchor.Y, 2)
				var newZ2 = math.Pow(averPoint2.Z-anchor.Z, 2)
				var relDist2 = math.Sqrt(newX2 + newY2 + newZ2)
				if relDist <= relDist2 {
					finalTriangles = append(finalTriangles[:idx], append([]Triangle{triangles[tri]}, finalTriangles[idx:]...)...)
					break
				}
			}
		}
	}
	return finalTriangles
}
