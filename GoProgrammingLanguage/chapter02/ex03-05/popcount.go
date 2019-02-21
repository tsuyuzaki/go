package main

import (
    "fmt"
    "math"
    "unsafe"
)

var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i & 1)
    }
}

func PopCount(x uint64) int {
    return int(pc[byte(x >> (0 * 8))] +
        pc[byte(x >> (1 * 8))] +
        pc[byte(x >> (2 * 8))] +
        pc[byte(x >> (3 * 8))] +
        pc[byte(x >> (4 * 8))] +
        pc[byte(x >> (5 * 8))] +
        pc[byte(x >> (6 * 8))] +
        pc[byte(x >> (7 * 8))])
}

// ex02.03
func PopCountEx03(x uint64) int {
    cnt := 0
    byteSize := uint64(unsafe.Sizeof(x))
    for i := uint64(0); i < byteSize; i++ {
        cnt += int(pc[byte(x >> (i * 8))])
    }
    return cnt
}

// ex02.04
func PopCountEx04(x uint64) int {
    cnt := 0
    bitSize := uint64(unsafe.Sizeof(x) * 8)
    for i := uint64(0); i < bitSize; i++ {
        cnt += int((x >> i) & 1)
    }
    return cnt
}

// ex02.05
func PopCountEx05(x uint64) int {
    cnt := 0
    for x > 0 {
        x &= (x - 1)
        cnt++
    }
    return cnt
}

func main() {
    printAll(0)
    printAll(math.MaxUint64 / 2)
    printAll(math.MaxUint64)
}

func printAll(x uint64) {
    fmt.Println("----")
    fmt.Printf("%d\n", PopCount(x))
    fmt.Printf("%d\n", PopCountEx03(x))
    fmt.Printf("%d\n", PopCountEx04(x))
    fmt.Printf("%d\n", PopCountEx05(x))
}