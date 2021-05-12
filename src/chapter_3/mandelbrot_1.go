package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin   = -2
		ymin   = -2
		xmax   = 2
		ymax   = 2
		width  = 1024
		height = 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py += 1 {
		y := float64(py)/height*(ymax-ymin) + ymin

		for px := 0; px < width; px += 1 {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		constrast  = 15
	)

	var v complex128 = complex(0, 0)

	for n := uint8(0); n < iterations; n += 1 {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - constrast*n}
		}
	}

	return color.Black

}
