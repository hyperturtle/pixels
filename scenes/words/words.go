package main

import (
	"github.com/hyperturtle/pixels/lib"
	"github.com/lucasb-eyer/go-colorful"
)

var bits []uint = []uint{
	0,
	0,
	131204,
	135300,
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
	478837,
	714542,
	462980,
	135364,
	1016900,
	279086,
	508432,
	508431,
	541200,
	1033777,
	508432,
	492607,
	476721,
	492606,
	33826,
	139807,
	476721,
	476718,
	541200,
	1001006,
	1015808,
	0,
	1015808,
	0,
	1015808,
	0,
	1015808,
	0,
	1015808,
	0,
	131208,
	542254,
	0,
	0,
	575025,
	1033774,
	509489,
	509487,
	476705,
	50734,
	509489,
	575023,
	1016865,
	492607,
	33825,
	492607,
	476721,
	951358,
	575025,
	1033777,
	462980,
	135310,
	476688,
	541215,
	582957,
	243505,
	1016865,
	33825,
	575025,
	579441,
	583477,
	708209,
	476721,
	575022,
	33825,
	509487,
	801649,
	575022,
	566447,
	575023,
	508430,
	33854,
	135300,
	135327,
	476721,
	575025,
	135498,
	575025,
	338613,
	575025,
	575018,
	141873,
	476720,
	1001009,
	1016934,
	418335,
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
}

const height = 8
const width = 5

func drawLetter(screen lib.Screen, sx, sy int, letter rune) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if sx+x >= lib.WIDTH || sy+y >= lib.HEIGHT || sx+x < 0 || sy+y < 0 {
				continue
			}
			index := (int(letter)-32)*2 + 1
			yy := y
			if y >= 4 {
				index -= 1
				yy -= 4
			}
			r := bits[index] & (1 << uint(yy*width+x))
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
