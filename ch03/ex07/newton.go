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
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // 注意：エラーを無視
}

/*
	f(x) = x^4 - 1
	f'(x) = 4x^3

	f(x) = 0
	x = ±1, ±i

	z_{n+1} = z_n - f(x) / f'(x)
*/
func newton(z complex128) color.Color {
	const iterations = 64

	for n := uint8(0); n < iterations; n++ {
		z -= (cmplx.Pow(z, 4) - 1) / (4 * cmplx.Pow(z, 3))
		fz := cmplx.Pow(z, 4) - 1
		if cmplx.Abs(fz) < 1e-6 {
			val := float64(n) / iterations
			return pseudoColor(z, val)
		}
	}
	return color.Black
}

// v: 0..1
func pseudoColor(z complex128, v float64) color.Color {
	var r, g, b uint8
	if cmplx.Abs(z-complex(-1, 0)) < 1e-6 {
		r = uint8(v * 255)
		g = uint8(v * 255)
		b = uint8(0)
	} else if cmplx.Abs(z-complex(0, -1)) < 1e-6 {
		r = uint8(v * 255)
		g = uint8(0)
		b = uint8(v * 255)
	} else if cmplx.Abs(z-complex(1, 0)) < 1e-6 {
		r = uint8(0)
		g = uint8(v * 255)
		b = uint8(v * 255)
	} else {
		r = uint8(v * 255)
		g = uint8(v * 255)
		b = uint8(v * 255)
	}
	return color.RGBA{r, g, b, 0xff}
}
