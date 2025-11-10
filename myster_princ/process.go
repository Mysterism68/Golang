package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Vector2 struct {
	x int
	y int
}
type Dim2 struct {
	w int
	h int
}

func drawMatrix(screen *ebiten.Image, matrix [][]int32, posX float32, posY float32, scale float32) {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			var point Vector2 = Vector2{int(-float32(len(matrix[y])/2-x)*scale + posX),
				int(-float32(len(matrix)/2-y)*scale + posY)}
			if (mouse.x >= point.x && mouse.x <= point.x+int(scale)) &&
				(mouse.y >= point.y && mouse.y <= point.y+int(scale)) {
				selTile = Vector2{x, y}
			}
			if matrix[y][x] > 0 {
				vector.DrawFilledRect(screen, float32(point.x),
					float32(point.y), scale, scale, color.RGBA{200, 100, 255, 255}, true)
			}
			if isDefiningEnv && (selTile.x == x && selTile.y == y) {
				vector.DrawFilledRect(screen, float32(point.x),
					float32(point.y), scale, scale, color.RGBA{150, 75, 255, 26}, true)
			}
		}
	}
	ebitenutil.DebugPrint(screen, fmt.Sprint(selTile.x)+" ,"+
		fmt.Sprint(selTile.y)+" "+fmt.Sprint(center.x)+" ,"+fmt.Sprint(center.y))
}

func convFileToMtrx(fileToRead string, matrix *[][]int32) {
	file, err := os.Open(fileToRead)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line string = string(scanner.Bytes())
		var row []int32
		for idx := 0; idx < len(line); idx++ {
			row = append(row, int32(line[idx])-48)
		}
		*matrix = append(*matrix, row)
	}
}
func convMtrxToFile(fileToWrite string, matrix [][]int32) {
	var fileContents string = ""
	for row := 0; row < len(matrix); row++ {
		var line string = ""
		for col := 0; col < len(matrix[row]); col++ {
			line = line + string(matrix[row][col]+48)
		}
		fileContents = fileContents + line + "\n"
	}
	os.WriteFile(fileToWrite, []byte(fileContents), 0644)
}
var center Vector2 = Vector2{0, 0}
func updateOrg() {
	//var prevOrgs [][]int32 = organisms
	if center.x == center.y && center.x < 1 {
		for y := 0; y < int(len(neighborhood)); y++ {
			for x := 0; x < int(len(neighborhood[y])); x++ {
				if neighborhood[y][x] < 0 {
					center.x = x
					center.y = y
					break
				}
			}
		}
	}
	for y := 0; y < len(organisms); y++ {
		for x := 0; x < len(organisms[y]); x++ {
			//var neighbors int = 0
		}
	}
}