package main

import (
	"github.com/hyperturtle/pixels/lib"
	"github.com/lucasb-eyer/go-colorful"
)

var bits []uint = []uint{
	0,
	137577500804,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	0,
	502097700654,
	485469851844,
	1066297213486,
	533130101263,
	567490364977,
	533130085439,
	499878691902,
	35469271583,
	499878676014,
	567490332206,
	1065151889408,
	1065151889408,
	1065151889408,
	1065151889408,
	1065151889408,
	137582102062,
	0,
	602958448174,
	534238447151,
	499861472814,
	534238512687,
	1066260726847,
	35468575807,
	499879150654,
	602958448177,
	485469851790,
	499844137503,
	611274962737,
	1066260268065,
	602957993841,
	611820686961,
	499878774318,
	35468592687,
	840590476846,
	593963304495,
	533127529534,
	141872468127,
	499878774321,
	142080525873,
	355062040113,
	602950216241,
	499878151729,
	1066333004319,
	0,
	0,
	0,
	0,
	0,
}

const height = 8
const width = 5

func drawLetter(screen lib.Screen, sx, sy int, letter rune) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if sx+x >= lib.WIDTH || sy+y >= lib.HEIGHT || sx+x < 0 || sy+y < 0 {
				continue
			}
			r := bits[int(letter)-32] & (1 << uint(y*width+x))
			if r == 0 {
				screen.Set(sx+x, sy+y, 0, 0, 0)
			} else {
				r, g, b := palette[y].RGB255()
				screen.Set(sx+x, sy+y, r, g, b)
			}
		}
	}
}

func drawWord(screen lib.Screen, x, y int, word string) {
	sx := x
	for _, letter := range word {
		if sx >= lib.WIDTH {
			return
		}
		drawLetter(screen, sx, y, letter)

		if sx+width >= 0 && sx+width < lib.WIDTH {
			screen.Set(sx+width, y, 0, 0, 0)
			screen.Set(sx+width, y+1, 0, 0, 0)
			screen.Set(sx+width, y+2, 0, 0, 0)
			screen.Set(sx+width, y+3, 0, 0, 0)
			screen.Set(sx+width, y+4, 0, 0, 0)
			screen.Set(sx+width, y+5, 0, 0, 0)
			screen.Set(sx+width, y+6, 0, 0, 0)
			screen.Set(sx+width, y+7, 0, 0, 0)
		}

		sx += width + 1
	}
}

var palette [height]colorful.Color

func init() {
	c1, _ := colorful.Hex("#FFFF00")
	c2, _ := colorful.Hex("#FF0000")

	for i := 0; i < height; i++ {
		d := float64(i) / float64(height)
		palette[i] = c2.BlendHsv(c1, d)
	}
}

func main() {
	screen := lib.NewScreen()
	i := 0
	for {
		drawWord(screen, -i, 1, "SHOOT ALEX SHOOT")
		i = (i + 1) % (11 * 6)
		screen.Dump()
	}
}
