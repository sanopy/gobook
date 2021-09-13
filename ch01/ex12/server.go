package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // パレットの最初の色
	blackIndex = 1 // パレットの次の色
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := 5   // 発振器 x が完了する周回の回数
		res := 0.001  // 回転の分解能
		size := 100   // 画像キャンバスは [-size..+size] の範囲を扱う
		nframes := 64 // アニメーションフレーム数
		delay := 8    // 10ms 単位でのフレーム間の遅延
		var err error

		if formval := r.FormValue("cycles"); formval != "" {
			if cycles, err = strconv.Atoi(formval); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		if formval := r.FormValue("res"); formval != "" {
			if res, err = strconv.ParseFloat(formval, 64); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		if formval := r.FormValue("size"); formval != "" {
			if size, err = strconv.Atoi(formval); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		if formval := r.FormValue("nframes"); formval != "" {
			if nframes, err = strconv.Atoi(formval); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		if formval := r.FormValue("delay"); formval != "" {
			if delay, err = strconv.Atoi(formval); err != nil {
				fmt.Fprintf(w, "%q", err)
				return
			}
		}
		lissajous(w, cycles, res, size, nframes, delay)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, res float64, size, nframes, delay int) {
	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：エンコードエラーを無視
}
