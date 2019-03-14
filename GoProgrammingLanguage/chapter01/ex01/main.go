/**
 * echo プログラムを修正して、そのプログラムを起動したコマンド名である os.Args[0] も表示するようにしなさい。
 */
package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {
    fmt.Println(strings.Join(os.Args, " "))
}
