package main

import (
	"github.com/hyperturtle/pixels/lib"
)

func main() {
	screen := lib.NewScreen()
	x := 0
	y := 0
	for {

		screen.Set(x, y, 0, 0, 0)

		x += 1
		if x >= 28 {
			x = 0
			y += 1
		}

		if y >= 10 {
			y = 0
			x = 0
		}

		screen.Set(x, y, 32, 32, 32)
		screen.Dump()
	}
}
