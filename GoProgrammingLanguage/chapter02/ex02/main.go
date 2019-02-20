package main

import (
    "fmt"
    "os"
    "strconv"
    "bufio"
    "./tempconv"
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
