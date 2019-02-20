package main

import (
    "fmt"
    "./tmpconv"
)

func main() {
    fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
    fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
    fmt.Println(tempconv.FToC((tempconv.Fahrenheit)(tempconv.AbsoluteZeroC)))
    fmt.Println(tempconv.FToK((tempconv.Fahrenheit)(tempconv.AbsoluteZeroC)))
    fmt.Println(tempconv.KToC((tempconv.Kelvin)(tempconv.AbsoluteZeroC)))
    fmt.Println(tempconv.KToF((tempconv.Kelvin)(tempconv.AbsoluteZeroC)))
}
