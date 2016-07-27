package main

import (
	"../lib"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"math/rand"
	"time"
)

var screen lib.Screen

const HEIGHT = 10
const WIDTH = 30

var pop [HEIGHT][WIDTH][2]float32

var palette [16]colorful.Color

func init() {
	reset()

	c1, _ := colorful.Hex("#2a1831")
	c2, _ := colorful.Hex("#000000")

	for i := 0; i < 16; i++ {
		d := float64(i) / 16.0
		palette[i] = c2.BlendHsv(c1, d*d)
	}
}

var lastResetTime time.Time

func reset() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			pop[y][x][0] = 1.0
			pop[y][x][1] = 0.0
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if rand.Float32() > 0.9 {
				pop[y][x][1] = 1.0
			}
		}
	}

	switch rand.Int() % 5 {
	case 0:
		RD_F = 0.07
		RD_K = 0.057
	case 1:
		RD_F = 0.11
		RD_K = 0.053
	case 2:
		RD_F = 0.046
		RD_K = 0.067
	case 3:
		RD_F = 0.066
		RD_K = 0.065
	case 4:
		RD_F = 0.05
		RD_K = 0.061
	}

	log.Println(RD_F, RD_K)

	lastResetTime = time.Now()
}

var sum [HEIGHT][WIDTH][2]float32

var RD_F float32 = 0.07
var RD_K float32 = 0.057

var RD_DA float32 = 0.5
var RD_DB float32 = 0.25

var RD_DT float32 = 1.0

func resolve(x, y int) (int, int) {
	if y >= HEIGHT {
		y -= HEIGHT
	} else if y < 0 {
		y += HEIGHT
	}
	if x >= WIDTH {
		x -= WIDTH
	} else if x < 0 {
		x += WIDTH
	}

	return x, y
}

func update() {

	// for y := 0; y < HEIGHT; y++ {
	// 	for x := 0; x < WIDTH; x++ {
	// 		if rand.Float32() > 0.999 {
	// 			pop[y][x][1] = 1.0
	// 		}
	// 		if rand.Float32() > 0.999 {
	// 			pop[y][x][0] = 1.0
	// 		}
	// 	}
	// }

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			for i := 0; i < 2; i++ {
				sum[y][x][i] = 0
			}
		}
	}
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			for i := 0; i < 2; i++ {
				v := pop[y][x][i]
				var dx, dy int
				dx, dy = resolve(x-1, y-1)
				sum[dy][dx][i] += v * 0.05

				dx, dy = resolve(x, y-1)
				sum[dy][dx][i] += v * 0.2

				dx, dy = resolve(x+1, y-1)
				sum[dy][dx][i] += v * 0.05

				dx, dy = resolve(x-1, y)
				sum[dy][dx][i] += v * 0.2

				dx, dy = resolve(x, y)
				sum[dy][dx][i] += v * -1.0

				dx, dy = resolve(x+1, y)
				sum[dy][dx][i] += v * 0.2

				dx, dy = resolve(x-1, y+1)
				sum[dy][dx][i] += v * 0.05

				dx, dy = resolve(x, y+1)
				sum[dy][dx][i] += v * 0.2

				dx, dy = resolve(x+1, y+1)
				sum[dy][dx][i] += v * 0.05
			}
		}
	}
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			a := pop[y][x][0]
			b := pop[y][x][1]
			pop[y][x][0] = a + (RD_DA*sum[y][x][0]-a*b*b+RD_F*(1.0-a))*RD_DT
			pop[y][x][1] = b + (RD_DB*sum[y][x][1]+a*b*b-(RD_K+RD_F)*b)*RD_DT
		}
	}

	draw()
}

func draw() {
	var max, min float32
	max = pop[0][0][0]
	min = pop[0][0][0]
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			if pop[y][x][0] > max {
				max = pop[y][x][0]
			}
			if pop[y][x][0] < min {
				min = pop[y][x][0]
			}
		}
	}

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			dx, dy := resolve(x, y)
			diff := (pop[dy][dx][0] - min) / (max - min)
			color := int(diff * 16)
			if color > 15 {
				color = 15
			}
			r, g, b := palette[color].RGB255()
			screen.Set(x, y, r, g, b)
		}
	}

	screen.Dump()

	if max-min < 0.0001 {
		reset()
	} else if time.Now().Sub(lastResetTime) > time.Minute {
		reset()
	}
}

func main() {
	log.Println("starting")
	screen = lib.NewScreen()
	for {
		update()
	}
}
