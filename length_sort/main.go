package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand/v2"

	"github.com/PerformLine/go-stockutil/colorutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var mainList = []float32{}

type Vector2 struct {
	x float32
	y float32
}

var win_dim = Vector2{800, 600}
var highest float32 = 10
var length = 4
var waitTime float32 = 200
var ticks float32 = 0
var idx = 0
var speed = 1
var selectMenu = true
var timer float32 = 0

type Game struct{}

func move(array []float32, oldI int, newI int) []float32 {
	if oldI < 0 || oldI >= len(array) || newI < 0 || newI >= len(array) {
		return array
	}
	item := array[oldI]
	//Remove the item from the array
	array = append(array[:oldI], array[oldI+1:]...)
	if newI > oldI {
		newI--
	}
	//Adds the item back in
	array = append(array[:newI], append([]float32{item}, array[newI:]...)...)

	return array
}

var stage = 1

func numPressed() uint8 {
	var finalNum uint8 = 10
	if ebiten.IsKeyPressed(ebiten.Key0) {
		finalNum = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key1) {
		finalNum = 1
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		finalNum = 2
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		finalNum = 3
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		finalNum = 4
	}
	if ebiten.IsKeyPressed(ebiten.Key5) {
		finalNum = 5
	}
	if ebiten.IsKeyPressed(ebiten.Key6) {
		finalNum = 6
	}
	if ebiten.IsKeyPressed(ebiten.Key7) {
		finalNum = 7
	}
	if ebiten.IsKeyPressed(ebiten.Key8) {
		finalNum = 8
	}
	if ebiten.IsKeyPressed(ebiten.Key9) {
		finalNum = 9
	}
	return finalNum
}

func isComplete() bool {
	var nextIsHigher = false
	for i := 0; i < len(mainList); i++ {
		var curItem = mainList[i]
		var lastItem float32
		if i-1 >= 0 {
			lastItem = mainList[i-1]
		} else {
			lastItem = -2
		}
		if lastItem < curItem {
			nextIsHigher = true
		} else {
			nextIsHigher = false
			break
		}
	}
	return nextIsHigher
}

var done = false

func (g *Game) Update() error {
	var w, h = ebiten.WindowSize()
	win_dim.x, win_dim.y = float32(w), float32(h)
	ticks += 1 / float32(ebiten.ActualFPS())
	if selectMenu {
		if ticks >= 0.093125 {
			switch stage {
			case 1:
				if numPressed() < 10 {
					length = length*10 + int(numPressed())
				}
				if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
					length = int(length/10 - (length/10 - int(math.Floor(float64(length/10)))))
				}
				if ebiten.IsKeyPressed(ebiten.KeyEnter) {
					stage += 1
				}
			case 2:
				if numPressed() < 10 {
					highest = highest*10 + float32(numPressed())
				}
				if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
					highest = (highest/10 - (highest/10 - float32(math.Floor(float64(highest/10)))))
				}
				if ebiten.IsKeyPressed(ebiten.KeyEnter) {
					stage += 1
				}
			case 3:
				if numPressed() < 10 {
					waitTime = waitTime*10 + float32(numPressed())
				}
				if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
					waitTime = (waitTime/10 - (waitTime/10 - float32(math.Floor(float64(waitTime/10)))))
				}
				if ebiten.IsKeyPressed(ebiten.KeyEnter) {
					waitTime /= 1000
					stage += 1
				}
			default:
				for idx := 0; idx < length; idx++ {
					var num = rand.Float32() * highest
					mainList = append(mainList, num)
				}
				selectMenu = false
			}
			ticks = 0
		}
	} else {
		if !done {
			timer += 1 / float32(ebiten.ActualFPS())
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			idx++
			if waitTime <= 1/float32(ebiten.ActualFPS()) && speed < 0 {
				idx++
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			idx--
			if waitTime <= 1/float32(ebiten.ActualFPS()) && speed > 0 {
				idx--
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyPeriod) {
			speed = 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyComma) {
			speed = -1
		}
		if ticks >= waitTime {
			if idx > len(mainList)-1 {
				idx = 0
				done = isComplete()
			}
			if idx < 0 {
				idx = len(mainList) - 1
				done = isComplete()
			}
			ticks = 0
			if !done {
				var thisNum = mainList[idx]
				var finIdx = -1
				var smallRange = highest * 1.01
				for i := 0; i < len(mainList); i++ {
					var numBef float32
					if i > 0 {
						numBef = mainList[i-1]
					} else {
						numBef = float32(-2)
					}
					var numCur = mainList[i]
					if numCur-numBef < smallRange && (numBef < thisNum && thisNum < numCur) {
						finIdx = i
						smallRange = numCur - numBef
					}
				}
				mainList = move(mainList, idx, finIdx)
			}
			idx += speed
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if selectMenu {
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Length of the list: ", length), 370, 280)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Highest Number: ", highest), 370, 300)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Wait time (in milliseconds): ", waitTime), 370, 320)
	} else {
		var width = win_dim.x / float32(len(mainList))
		for i := float32(0); int(i) < len(mainList); i++ {
			var height float32 = win_dim.y * (mainList[int(i)] / highest)
			var brightness float64 = 0.5
			if int(i) == idx {
				brightness = 1
			}
			var r, g, b = colorutil.HslToRgb(float64(mainList[int(i)]/highest)*360, 1, brightness)
			var light_int = uint8(int(i) % 2)
			vector.DrawFilledRect(screen, float32(i)*width, 0,
				width, win_dim.y, color.RGBA{uint8(float32(light_int)*40 + 40), 0, light_int*40 + 120, 255}, true)
			vector.DrawFilledRect(screen, float32(i)*width, win_dim.y-height,
				width, height, color.RGBA{r, g, b, 255}, true)
		}
		ebitenutil.DebugPrint(screen, fmt.Sprint("FPS: ",
			ebiten.ActualTPS(), " Time it took to compute: ", timer, "seconds"))
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
func main() {
	ebiten.SetWindowSize(int(win_dim.x), int(win_dim.y))
	ebiten.SetWindowTitle("Length Sorter")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
