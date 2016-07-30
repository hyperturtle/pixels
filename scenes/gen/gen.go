package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func main() {
	reader, _ := os.Open("./scenes/words/ascii.png")
	m, _, _ := image.Decode(reader)
	bounds := m.Bounds()
	for c := 0; c < 64; c++ {
		current := 0
		for y := 4; y >= 0; y-- {
			for x := 2; x >= 0; x-- {
				r, _, _, _ := m.At(bounds.Min.X+(c%32)*4+x, (bounds.Min.Y+c/32)*6+y).RGBA()
				if r == 0 {
					current = current << 1
				} else {
					current = current<<1 | 1
				}
			}
		}
		fmt.Printf("%d,\n", current)
	}
}
