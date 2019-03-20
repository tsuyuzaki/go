/**
 * ex: http://localhost:8000/?x=1&y=1&contrast=3
 */
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type query struct {
	centerX, centerY, contrast int
}

func (q *query) set(values url.Values) {
	if x, ok := getInt("x", values); ok {
		q.centerX = x
	}
	if y, ok := getInt("y", values); ok {
		q.centerY = y
	}
	if contrast, ok := getInt("contrast", values); ok {
		q.contrast = contrast
	}
}

func getInt(name string, values url.Values) (int, bool) {
	value := values.Get(name)
	if value == "" {
		return 0, false
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "strconv.Atoi [%v]\n", err)
		return 0, false
	}
	return i, true
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := query{}
		q.set(r.URL.Query())
		plot(w, &q)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot(out io.Writer, q *query) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		contrast               = 15
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin + float64(q.centerY)
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin + float64(q.centerX)
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z, q))
		}
	}
	png.Encode(out, img)
}

func mandelbrot(z complex128, q *query) color.Color {
	const iterations = 200

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - uint8(q.contrast)*n}
		}
	}
	return color.Black
}
