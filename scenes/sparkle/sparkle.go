package main

import (
	"github.com/hyperturtle/pixels/lib"
	"math/rand"
)

func main() {
	screen := lib.NewScreen()

	for {

		for p := 1; p <= 280*3; p++ {
			screen.Data[p] = byte(int(screen.Data[p]) * 9 / 10)

			if (p-1)%3 == 0 {

				if rand.Int()%1000 == 0 {
					screen.Data[p] = 255
					screen.Data[p+1] = 255
					screen.Data[p+2] = 255
				}
			}
		}

		screen.Dump()
	}
}
