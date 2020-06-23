package main

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320                    // canvas size in pixels
	minxy, maxxy  = -20.0, 20.0                 //range of coordinates to print
	xyscale       = width / (maxxy - minxy) / 2 // pixels per x or y unit
	zscale        = height * 0.4                // pixels per z unit
	angle         = math.Pi / 8                 // angle of x, y axes (=30°)
	defaultRes    = 0.3                         // resolution to increase x and y used to calculate z
)

func startServer() {
	http.HandleFunc("/svg-3d-function", svg3DFunction)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func svg3DFunction(w http.ResponseWriter, r *http.Request) {
	params, _ := r.URL.Query()["res"]
	w.Header().Set("Content-Type", "image/svg+xml")
	res := defaultRes
	if len(params) > 0 {
		res, _ = strconv.ParseFloat(params[0], 8)
	}
	printFunction(w, res, functionToPrint)
}

func printFunction(out io.Writer, res float64, f func(x, y float64) (float64, color.RGBA)) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke-width: 0.05' "+
		"width='%d' height='%d'>\n", width, height)
	for i := minxy; i < maxxy; i += res {
		for j := minxy; j < maxxy; j += res {
			ax, ay, color := corner(i+res, j, f)
			bx, by, _ := corner(i, j, f)
			cx, cy, _ := corner(i, j+res, f)
			dx, dy, _ := corner(i+res, j+res, f)
			colorString := fmt.Sprintf("#%.2x%.2x%.2x", color.R, color.G, color.B)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s; stroke: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colorString, "grey")
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func corner(x, y float64, f func(x, y float64) (float64, color.RGBA)) (float64, float64, color.RGBA) {
	z, color := f(x, y)

	sx := width/2 + (x-y)*math.Cos(angle)*xyscale
	sy := height/2 + (x+y)*math.Sin(angle)*xyscale - z*zscale
	return sx, sy, color
}

func functionToPrint(x, y float64) (float64, color.RGBA) {
	r := math.Hypot(x, y)
	z := math.Sin(r) / r
	redValue := 255 * z * 5
	rr := uint8(redValue)
	gg := uint8(0)
	if redValue > 255 {
		rr = 255
		if redValue > 255*2 {
			gg = 255
		} else {
			gg = uint8(redValue)
		}
	} else if redValue < 0 {
		rr = 0
	}

	return z, color.RGBA{rr, gg, 100, 255}
}
