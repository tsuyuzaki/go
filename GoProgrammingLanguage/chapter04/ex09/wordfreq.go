package main
/**
 * 入力テキストファイル内のそれぞれの単語の出現頻度を報告するプログラムwordfreqを書きなさい。
 * 入力を行ではなく単語へ分割するために、最初のScan呼び出しの前にinput.Split(bufio.ScanWords)を呼び出しなさい。
 */
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please input file path")
        return
    }
    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Fprintf(os.Stderr, "os.Open() error [%v]\n", err)
        os.Exit(1)
    }
    defer f.Close()

    s := bufio.NewScanner(f)
    s.Split(bufio.ScanWords)

    counts := make(map[string]int)
    for s.Scan() {
        counts[s.Text()]++
    }

    fmt.Println(counts)
}