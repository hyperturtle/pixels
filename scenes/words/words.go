package main

import (
	"github.com/hyperturtle/pixels/lib"
)

var bits []uint = []uint{
	0,
	8338,
	45,
	11962,
	32223,
	17057,
	11898,
	18,
	17556,
	5265,
	2746,
	1488,
	5120,
	448,
	8192,
	5812,
	15214,
	29850,
	29991,
	31143,
	18921,
	31183,
	31695,
	18727,
	31727,
	18927,
	1040,
	9232,
	17492,
	3640,
	5393,
	8615,
	0,
	24426,
	15083,
	25166,
	15211,
	29391,
	4815,
	31567,
	23533,
	29847,
	11047,
	23277,
	29257,
	23421,
	24573,
	31599,
	8047,
	12143,
	22379,
	30863,
	9367,
	31597,
	12141,
	24429,
	23213,
	9389,
	29351,
	25750,
	19609,
	13459,
	42,
	28672,
}

func drawLetter(screen lib.Screen, sx, sy int, letter rune) {
	for y := 0; y < 5; y++ {
		for x := 0; x < 3; x++ {
			r := bits[int(letter)-32] & (1 << uint(y*3+x))
			if r == 0 {
				screen.Set(sx+x, sy+y, 0, 0, 0)
			} else {
				screen.Set(sx+x, sy+y, 255, 255, 255)
			}
		}
	}
}

func drawWord(screen lib.Screen, x, y int, word string) {
	sx := x
	for _, letter := range word {
		drawLetter(screen, sx, y, letter)
		sx += 4
	}
}

func main() {
	screen := lib.NewScreen()
	for {
		drawWord(screen, 6, 2, "!?!?")
		screen.Dump()
	}
}
