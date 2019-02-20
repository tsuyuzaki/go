package main

import (
    "math"
    "testing"
)

func BenchmarkPopCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCount(0)
        PopCount(math.MaxUint64 / 2)
        PopCount(math.MaxUint64)
    }
}

func BenchmarkPopCountEx03(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCountEx03(0)
        PopCountEx03(math.MaxUint64 / 2)
        PopCountEx03(math.MaxUint64)
    }
}

func BenchmarkPopCountEx04(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCountEx04(0)
        PopCountEx04(math.MaxUint64 / 2)
        PopCountEx04(math.MaxUint64)
    }
}

func BenchmarkPopCountEx05(b *testing.B) {
    for i := 0; i < b.N; i++ {
        PopCountEx05(0)
        PopCountEx05(math.MaxUint64 / 2)
        PopCountEx05(math.MaxUint64)
    }
}
