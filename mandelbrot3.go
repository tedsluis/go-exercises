// Mandelbrot emits a PNG image of the Mandelbrot fractal.
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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 10048, 10048
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.RGBA {
	const iterations = 4000
	const contrast = 1000000
	var v complex128
	for n := uint32(0); n < iterations; n++ {
		v = v*v*v*v + z*z
		if cmplx.Abs(v) > 2 {
			c := uint32(contrast * n)
			r := uint8(c / (256 * 256 * 256))
			c = uint32(c - uint32(r))
			g := uint8(c / (256 * 256))
			c = uint32(c - uint32(g))
			b := uint8(c / 256)
			c = uint32(c - uint32(b))
			a := uint8(c)
			return color.RGBA{r, g, b, a}
		}
	}
	return color.RGBA{255, 255, 255, 255}
}
