package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
    "math"
)

const (
    xmin, ymin, xmax, ymax = -2, -2, +2, +2
    width, height          = 1024, 1024
)

func main() {
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    colors := []color.Color{}
    for py := 0; py < height; py++ {
        y := float64(py)/height*(ymax-ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px)/width*(xmax-xmin) + xmin
            z := complex(x, y)
            if len(colors) < 4 {
                colors = append(colors, mandelbrot(z))
            } else {
                colors = append(colors[1:], mandelbrot(z))
            }
            img.Set(px, py, toAveragedColor(colors))
        }
    }
    png.Encode(os.Stdout, img)
}

func toAveragedColor(colors []color.Color) color.Color {
    sumR, sumG, sumB, sumA := 0, 0, 0, 0
    for _, color := range colors {
        r, g, b, a := color.RGBA()
        sumR += int(r)
        sumG += int(g)
        sumB += int(b)
        sumA += int(a)
    }
    return color.RGBA{
        uint8(sumR / len(colors)), 
        uint8(sumG / len(colors)), 
        uint8(sumB / len(colors)), 
        uint8(sumA / len(colors))}
}

func mandelbrot(z complex128) color.Color {
    const iterations = 200
    const contrast = 15

    var v complex128
    for n := uint8(0); n < iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            px := ((real(z) - xmin) * width / (xmax - xmin))
            r := uint8(px / 4)
            g := uint8(math.MaxUint8 * cmplx.Abs(v) / math.Hypot(xmax, ymax))
            b := uint8(math.MaxUint8 * cmplx.Abs(v) / math.Hypot(xmax, ymax))
            return color.RGBA{r, g, b, 128}
        }
    }
    return color.Black
}
