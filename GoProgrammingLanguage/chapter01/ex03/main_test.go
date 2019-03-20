/**
 * 非効率な可能性のあるバージョンと strings.Join を使ったバージョンとで、実行時間の差を計測しなさい
 * (1.6節は time パッケージの一部を説明していますし、11.4節では体系的に性能評価を行うためのベンチマークテストの書き方を説明しています。)
 */
package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(strings.Join(os.Args, " "))
	}
}

func BenchmarkNotJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for _, arg := range os.Args {
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
