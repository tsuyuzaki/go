/**
 * 絶対温度 (Kelvin) のサポートを tempflag へ追加しなさい。
 */
package main

import (
	"./tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.KelvinFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
