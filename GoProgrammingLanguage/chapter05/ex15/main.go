/**
 * sum のように、可変個引数関数である max と min を書きなさい。
 * 引数なしで呼び出されたら、これらの関数は何をすべきでしょうか。
 * 少なくとも一つの引数が必要なバージョンも書きなさい。
 */
package main

import (
	"fmt"
)

func max(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	max := vals[0]
	for _, val := range vals[1:] {
		if max < val {
			max = val
		}
	}
	return max, true
}

func min(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	min := vals[0]
	for _, val := range vals[1:] {
		if min > val {
			min = val
		}
	}
	return min, true
}

func test(vals []int) {
	fmt.Println(vals)
	max, ok := max(vals...)
	if !ok {
		fmt.Println("Error has occurred.")
	} else {
		fmt.Println("Max:", max)
	}

	min, ok := min(vals...)
	if !ok {
		fmt.Println("Error has occurred.")
	} else {
		fmt.Println("Min:", min)
	}
}

func main() {
	test([]int{})
	test([]int{1, 2})
	test([]int{-1, 5, 6})
	test([]int{2, 4, 6})
	test([]int{-2, -6})
}
