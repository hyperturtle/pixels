package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"os"
)

var w, h int
var filename *string

func init() {
	width := flag.Int("w", 3, "width")
	height := flag.Int("h", 5, "height")
	filename = flag.String("f", "", "filename")
	flag.Parse()
	w = *width
	h = *height
}

func main() {
	reader, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	m, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	bounds := m.Bounds()
	for c := 0; c < 64; c++ {
		current := uint64(0)
		for y := h - 1; y >= 0; y-- {
			for x := w - 1; x >= 0; x-- {
				r, _, _, _ := m.At(bounds.Min.X+(c%32)*(w+1)+x, (bounds.Min.Y+c/32)*(h+1)+y).RGBA()
				if r == 0 {
					current = current << 1
				} else {
					current = current<<1 | 1
				}
			}
			if y == 4 || y == 0 {
				fmt.Printf("%d,\n", current)
				current = 0
			}
		}
	}
}
