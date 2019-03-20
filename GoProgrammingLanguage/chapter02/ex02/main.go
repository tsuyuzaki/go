/**
 * コマンドライン引数、もしくはコマンドライン引数が指定されなかった場合には標準入力から数値を読み込む、cf に似た汎用単位変換プログラムを書きなさい。
 * 各数値は、温度なら華氏と摂氏で、長さならフィートとメートルで、重さならポンドとキログラムでといった具合に各種単位へ変換しなさい。
 */
package main

import (
	"./tempconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		printTempByStdin()
	} else {
		printTempByArgs()
	}
}

func printTempByArgs() {
	for _, arg := range os.Args[1:] {
		printTemp(arg)
	}
}

func printTempByStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("Please input float value")
		input := scanner.Text()
		printTemp(input)
	}
}

func printTemp(input string) {
	t, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}
