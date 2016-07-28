package main

import (
	"github.com/hyperturtle/pixels/lib"
	"log"
	"math/rand"
)

var screen lib.Screen

var gridX [16]float32
var gridY [4]float32

func init() {
	for ii := 0; ii < 16; ii++ {
		gridX[ii] = float32(1) / float32(16)
	}
	for ii := 0; ii < 4; ii++ {
		gridY[ii] = float32(1) / float32(4)
	}
}

var colors [4][16]int = [4][16]int{
	[16]int{1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1},
	[16]int{1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1},
	[16]int{1, 0, 1, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 1},
	[16]int{1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 0, 0, 1, 1, 1},
}

func getPixel(x, y int) (r, g, b byte) {
	if colors[y][x] != 0 {
		return 255, 255, 255
	}
	return 0, 0, 0
}

var counter = 0

func update() {
	draw()

	sum := float32(0)
	for ii := 0; ii < 16; ii++ {
		sum += gridX[ii]
	}

	for ii := 0; ii < 16; ii++ {
		if sum > rand.Float32()*1.4 {
			gridX[ii] *= 0.99
		}
		if sum < rand.Float32()*1.0 {
			gridX[ii] += rand.Float32() * 0.01
		}
	}

	sum = 0
	for ii := 0; ii < 4; ii++ {
		sum += gridY[ii]
	}
	for ii := 0; ii < 4; ii++ {
		if sum > rand.Float32()*1.4 {
			gridY[ii] *= 0.99
		}
		if sum < rand.Float32()*1.0 {
			gridY[ii] += rand.Float32() * 0.01
		}
	}

}

func draw() {

	indexX := 0
	indexY := 0

	nextX := gridX[0]
	nextY := gridY[0]

	for y := 0; y < 10; y++ {
		indexX = 0
		nextX = gridX[0]
		for x := 0; x < 28; x++ {
			r, g, b := getPixel(indexX, indexY)
			screen.Set(x, y, r, g, b)

			if float32(x)/float32(28) > nextX {
				if indexX < 15 {
					indexX++
					nextX += gridX[indexX]
				}
			}
		}
		if float32(y)/float32(10) > nextY {
			if indexY < 3 {
				indexY++
				nextY += gridY[indexY]
			}
		}
	}

	screen.Dump()
}

func main() {

	log.Println("starting")
	screen = lib.NewScreen()
	for {
		update()
	}
}
