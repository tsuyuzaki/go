package main

import (
    "fmt"
    "os"
    "strings"
    "testing"
)

func BenchmarkJoin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Println(strings.Join(os.Args[:], " "))
    }
}

func BenchmarkNotJoin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := ""
        for _, arg := range os.Args[:] {
            s += arg
            s += " "
        }
        fmt.Println(s)
    }
}
