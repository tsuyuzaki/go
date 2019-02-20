/**
 * URLのqueryとしてcycles, res, size, nframes, delayを指定可能。
 * 以下が指定の例。
 * http://localhost:8000/?cycles=10,size=200,res=0.0001,nframes=128,delay=4
 * 指定がなければデフォルト値が利用される。
    cycles  = 5
    res     = 0.001
    size    = 100
    nframes = 64
    delay   = 8
 */
package main

import (
    "fmt"
    "os"
    "strconv"
    "image"
    "image/color"
    "image/gif"
    "io"
    "log"
    "math"
    "math/rand"
    "net/http"
    "net/url"
    "strings"
)

type lissajousProps struct {
    cycles float64
    res float64
    size int
    nframes int
    delay int
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        lissajous(w, r.URL)
    })
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var palette = []color.Color{color.White, color.Black}

const (
    whiteIndex = 0
    blackIndex = 1
)

func lissajous(out io.Writer, url *url.URL) {
    freq := rand.Float64() * 3.0
    phase := 0.0

    props := parseQuery(url.RawQuery)
    anim := gif.GIF{LoopCount: props.nframes}
    for i := 0; i < props.nframes; i++ {
        rect := image.Rect(0, 0, 2 * props.size + 1, 2 * props.size + 1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < props.cycles * 2 * math.Pi; t += props.res {
            x := math.Sin(t)
            y := math.Sin(t * freq + phase)
            img.SetColorIndex(props.size + int(x * (float64)(props.size) + 0.5), 
                props.size + int(y * (float64)(props.size) + 0.5), blackIndex)
        }

        phase++
        anim.Delay = append(anim.Delay, props.delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}

const (
    cycles  = 5
    res     = 0.001
    size    = 100
    nframes = 64
    delay   = 8
)

func parseQuery(query string) *lissajousProps {
    props := lissajousProps{cycles, res, size, nframes, delay}
    for _, namedValueStr := range strings.Split(query, ",") {
        fillProps(namedValueStr, &props)
    }
    return &props
}

func fillProps(namedValueStr string, props *lissajousProps) {
    namedValue := strings.Split(namedValueStr, "=")
    if len(namedValue) != 2 {
        fmt.Printf("[%s] is not supported\n", namedValueStr);
        return
    }
    if namedValue[0] == "cycles" {
        props.cycles = parseFloat(namedValue[1], cycles)
    } else if namedValue[0] == "res" {
        props.res = parseFloat(namedValue[1], res)
    } else if namedValue[0] == "size" {
        props.size = parseInt(namedValue[1], size)
    } else if namedValue[0] == "nframes" {
        props.nframes = parseInt(namedValue[1], nframes)
    } else if namedValue[0] == "delay" {
        props.delay = parseInt(namedValue[1], delay)
    } else {
        fmt.Printf("[%s] is not supported\n", namedValueStr);
    }
}

func parseInt(s string, defaultValue int) int {
    i, err := strconv.ParseInt(s, 10, 32)
    if err != nil {
        fmt.Fprintf(os.Stderr, "strconv.ParseInt: %v\n", err)
        return defaultValue
    }
    return (int)(i)
}

func parseFloat(s string, defaultValue float64) float64 {
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        fmt.Fprintf(os.Stderr, "strconv.ParseInt: %v\n", err)
        return defaultValue
    }
    return f
}