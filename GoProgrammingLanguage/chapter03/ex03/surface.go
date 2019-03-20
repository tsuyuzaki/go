package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, bz, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, cz, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, dz, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, getColor((az+bz+cz+dz)/4))
		}
	}
	fmt.Println("</svg>")
}

func getColor(z float64) string {
	color := int(0xff * (math.Abs(z)))
	if color > 0xff {
		color = 0xff
	}
	strColor := fmt.Sprintf("%02x", color)

	if z > 0 {
		return ("#" + strColor + "0000")
	} else {
		return ("#0000" + strColor)
	}
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y)
	if math.IsNaN(r) || math.IsInf(r, 0) {
		fmt.Fprintf(os.Stderr, "f(%f, %f) returns invalid value.\n", x, y)
		return 0, false
	}

	ret := math.Sin(r) / r
	if math.IsNaN(ret) || math.IsInf(ret, 0) {
		fmt.Fprintf(os.Stderr, "f(%f, %f) returns invalid value.\n", x, y)
		return 0, false
	}
	return ret, true
}
