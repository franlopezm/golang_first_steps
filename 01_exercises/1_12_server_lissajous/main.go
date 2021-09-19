// Modify the Lissajous server to read parameter values from the URL.
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
	"net/url"
	"strconv"
)

var palette = []color.Color{
	color.White,
	color.RGBA{21, 182, 48, 255},   // Green
	color.RGBA{255, 87, 51, 255},   // Red
	color.RGBA{21, 50, 182, 255},   // Blue
	color.RGBA{227, 223, 12, 255},  // yellow
	color.RGBA{0, 0, 0, 255},       // black
	color.RGBA{144, 144, 144, 255}, // Gray
}

type LissajousData struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Print("Otro error")
	}
	if len(query) == 0 {
		return
	}
	fmt.Println("query", query, err)

	cycles, err := strconv.Atoi(query["cycles"][0])
	fmt.Println("cycles", cycles, err)
	if err != nil {
		cycles = 5
	}

	res, err := strconv.ParseFloat(query["res"][0], 64)
	fmt.Println("res", res, err)
	if err != nil {
		res = 0.001
	}

	size, err := strconv.Atoi(query["size"][0])
	fmt.Println("size", size, err)
	if err != nil {
		size = 100
	}

	nframes, err := strconv.Atoi(query["nframes"][0])
	fmt.Println("nframes", nframes, err)
	if err != nil {
		nframes = 64
	}

	delay, err := strconv.Atoi(query["delay"][0])
	fmt.Println("delay", delay, err)
	if err != nil {
		delay = 8
	}

	lissajous(w, LissajousData{
		cycles:  cycles,
		res:     res,
		size:    size,
		nframes: nframes,
		delay:   delay,
	})
}

func lissajous(out io.Writer, payload LissajousData) {
	var (
		cycles  = payload.cycles  // number of complete x oscillator revolutions
		res     = payload.res     // angular resolution
		size    = payload.size    // image canvas covers [-size..+size]
		nframes = payload.nframes // number of animation frames
		delay   = payload.delay   // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(phase) + 1

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
