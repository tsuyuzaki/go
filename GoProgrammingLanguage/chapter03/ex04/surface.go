/**
 * top/bottomRGBとしてred, green, blueと、height, widthを指定可能。
 * http://localhost:8000/?topRGB=green,bottomRGB=red,height=1300,width=300
 */

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		writeSurface(w, r.URL)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func writeSurface(out io.Writer, url *url.URL) {
	query := parseQuery(url.RawQuery)
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", query.width, query.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, ok := corner(i+1, j, query)
			if !ok {
				continue
			}
			bx, by, bz, ok := corner(i, j, query)
			if !ok {
				continue
			}
			cx, cy, cz, ok := corner(i, j+1, query)
			if !ok {
				continue
			}
			dx, dy, dz, ok := corner(i+1, j+1, query)
			if !ok {
				continue
			}
			c := getColor(((az + bz + cz + dz) / 4), query)
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, query *surfaceQuery) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}

	sx := float64(query.width)/2 + (x-y)*cos30*query.xyscale
	sy := float64(query.height)/2 + (x+y)*sin30*query.xyscale - z*query.zscale
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
