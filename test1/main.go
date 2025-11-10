package main

import (
	"fmt"
	"github.com/Kalyle-68/Golang/extras"
)

func main() {
	//This was a test of importing my own packages
	var vector extras.Vector2 = extras.Vector2{8, 6}
	println(fmt.Sprint(vector.X) + ", " + fmt.Sprint(vector.Y))
}
