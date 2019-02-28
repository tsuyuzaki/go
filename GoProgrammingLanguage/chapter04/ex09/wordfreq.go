package main
/**
 * 入力テキストファイル内のそれぞれの単語の出現頻度を報告するプログラムwordfreqを書きなさい。
 * 入力を行ではなく単語へ分割するために、最初のScan呼び出しの前にinput.Split(bufio.ScanWords)を呼び出しなさい。
 */
import (
    "bufio"
    "fmt"
    "io"
    "os"
    "unicode"
    "unicode/utf8"
)

func main() {
    counts := make(map[string]map[rune]int)
    var utflen [utf8.UTFMax + 1]int
    invalid := 0

    in := bufio.NewReader(os.Stdin)
    for {
        r, n, err := in.ReadRune()
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
            os.Exit(1)
        }
        if r == unicode.ReplacementChar && n == 1 {
            invalid++
            continue
        }
        increment(r, counts)
        utflen[n]++
    }
    fmt.Printf("rune\tcount\n")
    printCounts(counts)
    fmt.Print("\nlen\tcount\n")
    for i, n := range utflen {
        if i > 0 {
            fmt.Printf("%d\t%d\n", i, n)
        }
    }
    if invalid > 0 {
        fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
    }
}