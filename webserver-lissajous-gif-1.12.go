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
	"sync"
	"strings"
	"regexp"
	"strconv"
	"fmt"
)

var Red     = color.RGBA{0xff, 0x00, 0x00, 0xff}
var Blue    = color.RGBA{0x00, 0x00, 0x99, 0xff}
var Yellow  = color.RGBA{0xff, 0xff, 0x00, 0xff}
var Orange  = color.RGBA{0xff, 0x33, 0x00, 0xff}
var Aqua    = color.RGBA{0x00, 0x99, 0xff, 0xff}
var Brown   = color.RGBA{0x99, 0x33, 0x00, 0xff}
var Rose    = color.RGBA{0xff, 0x00, 0x66, 0xff}
var Green   = color.RGBA{0x00, 0xcc, 0x00, 0xff}
var palette = []color.Color{color.White, Red, Blue, Yellow, Orange, Aqua, Brown, Green, Rose, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)
var mu sync.Mutex

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var result = strings.Split(r.URL.RawQuery,"=")
		if result[0] == "cycles" {
			re := regexp.MustCompile("[0-9]+")
			digit:=re.FindAllString(result[1], -1)
			f,err := strconv.ParseFloat(digit[0], 64)
			if err == nil {
				if  f > 0 {
					lissajous(w,f)
				} else {
					fmt.Fprintf(w, "f <=0: %v",r.URL.RawQuery)
				}
			} else {
				fmt.Fprintf(w, "error: %v",r.URL.RawQuery)
			}
		} else {
			fmt.Fprintf(w, "query wrong format: %v\nfull path: %v\nraw path: %v",r.URL.RawQuery, r.URL.Path, r.URL.RawPath)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.0001 // angular resolution
		size    = 300    // image canvas covers [-size..+size]
		nframes = 512    // number of animation frames
		delay   = 16     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y scillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			i := uint((x * x) * 8)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5),
				(blackIndex + uint8(i)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
