package main

import (
	"../lib"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"math/rand"
)

type walker struct {
	x, y    int
	d       int
	r, g, b byte
}

const count = 10

func main() {
	screen := lib.NewScreen()

	walkers := make([]walker, count)

	for i := range walkers {
		walkers[i].x = rand.Int() % 28
		walkers[i].y = rand.Int() % 10
		d := float64(i) / float64(count)
		r, g, b := colorful.Hcl(d*360, 1.0, 0.1).RGB255()

		walkers[i].r = r
		walkers[i].g = g
		walkers[i].b = b
	}

	log.Println(walkers)

	for {

		for p := 1; p <= 280*3; p++ {
			screen.Data[p] = byte(int(screen.Data[p]) * 90 / 100)
		}

		for i := range walkers {
			if rand.Int()%30 == 0 {
				walkers[i].d = rand.Int() % 4
			}

			switch walkers[i].d {
			case 0:
				walkers[i].x += 1
			case 1:
				walkers[i].x -= 1
			case 2:
				walkers[i].y += 1
			case 3:
				walkers[i].y -= 1
			}

			if walkers[i].x < 0 {
				walkers[i].x = 27
			} else if walkers[i].x >= 28 {
				walkers[i].x = 0
			}

			if walkers[i].y < 0 {
				walkers[i].y = 9
			} else if walkers[i].y >= 10 {
				walkers[i].y = 0
			}

			r, g, b := screen.Get(walkers[i].x, walkers[i].y)
			r = (r + walkers[i].r)
			g = (g + walkers[i].g)
			b = (b + walkers[i].b)

			screen.Set(walkers[i].x, walkers[i].y, r, g, b)
		}

		screen.Dump()
	}
}
