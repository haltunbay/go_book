package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
)

var img = image.NewRGBA(image.Rect(0, 0, 200, 200))
var c color.Color

func main() {
	args := os.Args
	radius := 15
	degree := 360
	step := 10.0
	if len(args) > 1 {
		radius, _ = strconv.Atoi(args[1])
	}
	if len(args) > 2 {
		degree, _ = strconv.Atoi(args[2])
	}
	if len(args) > 3 {
		step, _ = strconv.ParseFloat(args[3], 64)
	}
	c = color.RGBA{255, 0, 0, 255}
	//circle(50, 50, float64(radius), float64(degree))
	spiral(100, 100, float64(radius), float64(degree), step)
	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func circle(cx, cy int, r float64, d float64) {
	for i := 0.00; i < d; i += 1 {
		x := cx + int(math.Round(math.Cos(toRadian(i))*r))
		y := cy - int(math.Round(math.Sin(toRadian(i))*r))
		img.Set(x, y, c)
	}
}

func spiral(cx, cy int, r float64, d float64, step float64) {

	for i := 0.00; i < d; i += 1 {
		radius := math.Exp2(r + i/step)
		x := cx + int(math.Round(math.Cos(toRadian(i))*radius))
		y := cy - int(math.Round(math.Sin(toRadian(i))*radius))
		img.Set(x, y, c)
	}
}

func toRadian(d float64) float64 {
	return d * math.Pi / 180
}
