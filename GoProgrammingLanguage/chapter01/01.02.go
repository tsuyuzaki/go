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