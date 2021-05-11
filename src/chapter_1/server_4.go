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
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles int) {
	const (
		resolution = 0.001
		size       = 100
		frames     = 64
		delay      = 8
	)

	relativeYFrequency := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: frames}
	phase := 0.0

	for i := 0; i < frames; i += 1 {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2.0*math.Pi; t += resolution {
			x := math.Sin(t)
			y := math.Sin(t*relativeYFrequency + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}

		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)
	}

	gif.EncodeAll(out, &animation)
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		cyclesString := request.URL.Query().Get("cycles")

		cycles := 5

		if cyclesString != "" {
			if cyclesInt, err := strconv.Atoi(cyclesString); err == nil {
				cycles = cyclesInt
			}

		}

		lissajous(writer, cycles)
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
