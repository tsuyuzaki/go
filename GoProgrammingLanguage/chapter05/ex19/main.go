/**
 * return 文を含んでいないのに、ゼロ値ではない値を返す関数を panic と recover を使って書きなさい。
 */
package main

import (
	"fmt"
)

func main() {
	fmt.Println(getValueByPanicAndRecover())
}

func getValueByPanicAndRecover() (ret int) {
	defer func() {
		if p := recover(); p != nil {
			ret = p.(int)
		}
	}()
	panic(1)
}
