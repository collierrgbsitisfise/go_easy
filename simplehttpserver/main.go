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
	"sync"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex    = 0
	blackIndex    = 1
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angel         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angel), math.Cos(angel)
var mu sync.Mutex
var count int

func getSvg() string {
	var res string
	res += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d' >", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			res += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	res += fmt.Sprintf("</svg>")
	return res
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x*y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", getLissajousAnimationHandler)
	http.HandleFunc("/svg", getSvgHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getSvgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "image/svg+xml")
	fmt.Fprintf(w, getSvg())
}

func getLissajousAnimationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get cycle params")
	cycleParam := r.URL.Query().Get("cycles")
	cycles, err := strconv.Atoi(cycleParam)

	if err != nil {
		cycles = 20
	}

	lissajous(w, float64(cycles))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycles float64) {
	const (
		// cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
