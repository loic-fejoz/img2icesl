// Because IceSL cannot yet load an image, this one convert image to Lua
package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

var filename = flag.String("input", "", "image to convert")

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	flag.Parse()
	reader, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("-- Converted by https://github.com/loic-fejoz/img2icesl")
	fmt.Printf(
		"tex_dim = v(%v, %v, %v)\n",
		bounds.Max.X-bounds.Min.X,
		bounds.Max.X-bounds.Min.Y,
		1,
	)
	fmt.Printf("tex = tex3d_rgb8f(tex_dim.x,tex_dim.y,tex_dim.z)\n")
	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// So shift 8 to range [0, 255]
			fmt.Printf(
				"tex:set(%v, %v, 0, v(%v, %v, %v))\n",
				x-bounds.Min.X,
				y-bounds.Min.Y,
				float64(r)/65535.0,
				float64(g)/65535.0,
				float64(b)/65535.0,
			)
		}
	}
}
