/**
 * 二つのSHA256ハッシュで異なるビットの数を数える関数を書きなさい。(2.6.2 説のPopCountを参照)。
 */
package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%b\n%b\n%d\n\n", c1, c2, countBitDiff(&c1, &c2))

	for i := 0; i < 32; i++ {
		c1[i] = byte(i)
		c2[i] = byte(i)
	}
	fmt.Printf("%b\n%b\n%d\n\n", c1, c2, countBitDiff(&c1, &c2))

	c1[0] = 2 // 10
	fmt.Printf("%b\n%b\n%d\n\n", c1, c2, countBitDiff(&c1, &c2))

	c1[0] = 7 // 111
	fmt.Printf("%b\n%b\n%d\n\n", c1, c2, countBitDiff(&c1, &c2))
}

func countBitDiff(lhs, rhs *[32]byte) int {
	diffCnt := 0
	for i := 0; i < 32; i++ {
		xor := (lhs[i] ^ rhs[i]) // 1010 ^ 1100 => 0110
		diffCnt += int(pc[xor])  // xorの値のビット数を加算。
	}
	return diffCnt
}
