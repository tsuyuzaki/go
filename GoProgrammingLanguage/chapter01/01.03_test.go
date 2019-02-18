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

// testingをimportしBenchmarkしたい関数の接頭辞にBenchmarkを付与し、引数にb *testing.Bを指定。
// ループの数は b.N を指定すると適切なベンチマークの回数ループしてくれる。
// ソースコードの接尾辞を "_test.go" をとし、以下実行するとベンチマークしてくれる。
//  $ go test -bench .