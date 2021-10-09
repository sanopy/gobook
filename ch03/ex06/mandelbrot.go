// mandelbrotはマンデルブロフラクタルのPNG画像を生成します
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -.5, -1, +.5, 0
	width, height          = 1024, 1024
	iterations             = 200
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, supersampling(x, y))
		}
	}
	png.Encode(os.Stdout, img) // 注意：エラーを無視
}

func supersampling(x, y float64) color.Color {
	const (
		kernelsize = 3
		dx, dy     = 1 / (height * (ymax - ymin) * kernelsize), 1 / (width * (xmax - xmin) * kernelsize)
	)
	var sum int

	n := kernelsize / 2
	for i := -n; i <= n; i++ {
		for j := -n; j <= n; j++ {
			z := complex(x+dx*float64(j), y+dy*float64(i))
			sum += mandelbrot(z)
		}
	}

	val := float64(sum) / float64(iterations*kernelsize*kernelsize)
	return pseudoColor(val)
}

func mandelbrot(z complex128) int {
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return int(n)
		}
	}
	return 0
}

// v: 0..1
func pseudoColor(v float64) color.Color {
	r := uint8(0)
	g := uint8(v * 255)
	b := uint8(v * 255)
	return color.RGBA{r, g, b, 0xff}
}
