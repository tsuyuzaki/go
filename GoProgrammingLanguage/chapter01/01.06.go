package main

import (
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

type namedColor struct {
    name  string
    value color.Color
}

var namedColors = []namedColor{
namedColor{"white", color.White},
namedColor{"black", color.Black},
namedColor{"red",   color.RGBA{255, 0,     0, 255}},
namedColor{"green", color.RGBA{0,   255,   0, 255}},
namedColor{"blue",  color.RGBA{0,   0,   255, 255}}}

var palette = colors()

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func colors() []color.Color {
    colors := []color.Color{}
    for _, namedColor := range namedColors {
        colors = append(colors, namedColor.value)
    }
    return colors
}

func indexOf(name string) uint8 {
    for i, namedColor := range namedColors {
        if name == namedColor.name {
        	return (uint8)(i)
        }
    }
    return 1 // use black
}

func lissajous(out io.Writer) {
    if len(os.Args) != 2 {
        fmt.Printf("Please specify colorIndex\n")
        return
    }
    
    const (
        cycles  = 5
        res     = 0.001
        size    = 100
        nframes = 64
        delay   = 8
    )
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nframes}
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), (uint8)(indexOf(os.Args[1])))
        }

        phase += 1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}