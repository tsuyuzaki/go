/**
 * echo プログラムを修正して、個々の引数のインデックスと値の組を1行ごとに表示しなさい。
 */
package main

import (
    "fmt"
    "os"
)

func main() {
    for index, arg := range os.Args {
        fmt.Printf("%s [index:%d]\n", arg, index)
    }
}
