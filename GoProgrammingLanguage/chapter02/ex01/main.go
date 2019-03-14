/**
 * 絶対温度 (Kelvin scale) で温度を処理するために tempconv に型、定数、関数を追加しなさい。
 * 0K は -273.15℃ であり、1K の差と 1℃ の差は同じ大きさです。
 */
package main

import (
    "fmt"
    "./tempconv"
)

func main() {
    fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
    fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
    fmt.Println(0)
    fmt.Println(0)
    fmt.Println(0)
    fmt.Println(0)
}
