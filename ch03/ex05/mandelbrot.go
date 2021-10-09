// mandelbrotはマンデルブロフラクタルのPNG画像を生成します
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
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // 注意：エラーを無視
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			val := float64(n) / iterations
			return pseudoColor(val)
		}
	}
	return color.Black
}

// v: 0..1
func pseudoColor(v float64) color.Color {
	r := uint8(0)
	g := uint8(v * 255)
	b := uint8(v * 255)
	return color.RGBA{r, g, b, 0xff}
}
