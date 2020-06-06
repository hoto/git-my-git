package main

import (
    "fmt"
    "github.com/logrusorgru/aurora"
)

func main() {
    fmt.Printf("hello %s\n", aurora.Cyan("world"))
    ParseArgsAndFlags()
}
