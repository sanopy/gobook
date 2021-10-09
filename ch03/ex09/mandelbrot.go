package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x, y := 0., 0.
		scale := 1.
		var err error

		if formval := r.FormValue("x"); formval != "" {
			if x, err = strconv.ParseFloat(formval, 64); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		if formval := r.FormValue("y"); formval != "" {
			if y, err = strconv.ParseFloat(formval, 64); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		if formval := r.FormValue("scale"); formval != "" {
			if scale, err = strconv.ParseFloat(formval, 64); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		render(w, x, y, scale)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func render(out io.Writer, x, y, scale float64) {
	const width, height = 1024, 1024
	ymin, ymax := -2.*(1./scale)+y, 2.*(1./scale)+y
	xmin, xmax := -2.*(1./scale)+x, 2.*(1./scale)+x

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // 注意：エラーを無視
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
