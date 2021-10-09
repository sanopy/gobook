package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"

	"github.com/sanopy/gobook/ch03/ex08/mycmplx"
)

const (
	iterations             = 200
	xmin, ymin, xmax, ymax = -.2, -1, 0, -.8
	width, height          = 1024, 1024
)

var calcType = flag.String("t", "", "specify type ('complex64', 'complex128', 'float' or 'rat')")

func main() {
	flag.Parse()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			switch *calcType {
			case "complex64":
				z := complex(float32(x), float32(y))
				img.Set(px, py, mandelbrotComplex64(z))
			case "complex128":
				z := complex(x, y)
				img.Set(px, py, mandelbrotComplex128(z))
			case "float":
				z := mycmplx.NewBigFloat(x, y)
				img.Set(px, py, mandelbrotFloat(z))
			case "rat":
				z := mycmplx.NewBigRat(x, y)
				img.Set(px, py, mandelbrotRat(z))
			default:
				return
			}
		}
	}
	png.Encode(os.Stdout, img) // 注意：エラーを無視
}

func mandelbrotComplex64(z complex64) color.Color {
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			val := float64(n) / iterations
			return pseudoColor(val)
		}
	}
	return color.Black
}

func mandelbrotComplex128(z complex128) color.Color {
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

func mandelbrotFloat(z *mycmplx.BigFloat) color.Color {
	v := mycmplx.NewBigFloat(0, 0)
	for n := uint8(0); n < iterations; n++ {
		v = v.Mul(v).Add(z)
		if v.Abs().Cmp(big.NewFloat(2)) > 0 {
			val := float64(n) / iterations
			return pseudoColor(val)
		}
	}
	return color.Black
}

func mandelbrotRat(z *mycmplx.BigRat) color.Color {
	v := mycmplx.NewBigRat(0, 0)
	for n := uint8(0); n < iterations; n++ {
		fmt.Println(n, 1)
		v = v.Mul(v).Add(z)
		fmt.Println(n, 2)
		if v.Abs() > 2 {
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
