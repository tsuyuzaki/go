/**
 * デフォルトで標準入力のSHA256ハッシュを表示するプログラムを書きなさい。
 * ただし、SHA384ハッシュやSHA512ハッシュを表示するコマンドラインのフラグもサポートしなさい。
 */
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Please input:")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		fmt.Fprintf(os.Stderr, "Scan error\n")
		return
	}

	if len(os.Args) == 1 || os.Args[1] == "sha256" {
		c := sha256.Sum256([]byte(scanner.Text()))
		fmt.Printf("%x\n", c)
	} else if os.Args[1] == "sha384" {
		c := sha512.Sum384([]byte(scanner.Text()))
		fmt.Printf("%x\n", c)
	} else if os.Args[1] == "sha512" {
		c := sha512.Sum512([]byte(scanner.Text()))
		fmt.Printf("%x\n", c)
	}
}
