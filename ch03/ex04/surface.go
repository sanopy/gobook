// surfaceは3-D面の関数のSVGレンダリングを計算します
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // キャンバスの大きさ（画素数）
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲（-xyrange..+xyrange）
	xyscale       = width / 2 / xyrange // x単位およびy単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度（=30度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")

		funcType := r.FormValue("f") // 描画する関数の種別
		render(w, funcType)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func render(out io.Writer, funcType string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, isAFinite := corner(i+1, j, funcType)
			bx, by, bz, isBFinite := corner(i, j, funcType)
			cx, cy, cz, isCFinite := corner(i, j+1, funcType)
			dx, dy, dz, isDFinite := corner(i+1, j+1, funcType)
			if isAFinite && isBFinite && isCFinite && isDFinite {
				z := (az + bz + cz + dz) / 4.
				r, g, b := preudoColor(z)
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke='#%02x%02x%02x' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, funcType string) (float64, float64, float64, bool) {
	// マス目 (i, j) の角の点 (x, y) を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := f(x, y, funcType)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, false
	}

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に投影
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64, funcType string) float64 {
	switch funcType {
	case "egg":
		return f2(x, y)
	case "mogul":
		return f3(x, y)
	case "saddle":
		return f4(x, y)
	default:
		return f1(x, y)
	}
}

func f1(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0) からの距離
	return math.Sin(r) / r
}

func f2(x, y float64) float64 {
	return (-math.Abs(math.Sin(x/3.)) - math.Abs(math.Cos(y/3.))) / 5.
}

func f3(x, y float64) float64 {
	return (math.Abs(math.Sin(x/3.)) + math.Abs(math.Cos(y/3.))) / 5.
}

func f4(x, y float64) float64 {
	return (math.Pow(x, 2) - math.Pow(y, 2)) / 350.
}

func preudoColor(v float64) (int, int, int) {
	max, min := .7, -.7
	v = math.Min(math.Max((v-min)/(max-min), 0.), 1.) // 0..+1

	r := int(v * 255)
	g := 0
	b := int((1 - v) * 255)
	return r, g, b
}
