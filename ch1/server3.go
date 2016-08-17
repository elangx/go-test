package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handle)
	rand.Seed(time.Now().UTC().UnixNano())
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	cycles := 5
	if err := r.ParseForm(); err == nil {
		v, err := strconv.Atoi(r.Form.Get("cycles"))
		if err == nil {
			cycles = v
		}
	}
	lissajous(w, cycles)
	//fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	//for k, v := range r.Header {
	//	fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
	//}
	//fmt.Fprintf(w, "Host: %q\nRemoteAddr: %q\n", r.Host, r.RemoteAddr)
	//if err := r.ParseForm(); err != nil {
	//	log.Print(err)
	//}
	//for k, v := range r.Form {
	//	fmt.Fprintf(w, "Form[%q]: %q\n", k, v)
	//}
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phas := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phas)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phas += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
