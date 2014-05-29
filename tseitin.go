package main

import "flag"
import "fmt"

func main() {
    formula := flag.String("formula", "", "The formula in propositional logic")
    flag.Parse()
    fmt.Println("Formula: " + *formula)
}
