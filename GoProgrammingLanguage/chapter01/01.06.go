﻿package main

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
    "strconv"
)

var palette = []color.Color{
    color.Black, 
    color.RGBA{255, 0, 0, 255}, 
    color.RGBA{0, 255, 0, 255}, 
    color.RGBA{0, 0, 255, 255}}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    if len(os.Args) != 2 {
        fmt.Printf("Please specify colorIndex\n")
        return
    }
    
    index, err := strconv.Atoi(os.Args[1])
    if err != nil || (index < 0 || index >= len(palette)) {
        fmt.Printf("Please specify valid colorIndex [0 =< index < %d]\n", len(os.Args))
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
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), (uint8)(index))
        }

        phase += 1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}