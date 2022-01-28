package main

import (
    "fmt"
    "os"
    "github.com/lightbearer42/wordcount/processor"
)

func main() {
    fmt.Println(processor.Count(os.Args[1]))
}
