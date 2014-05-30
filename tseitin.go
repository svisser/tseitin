package main

import "flag"
import "fmt"

type Formula struct {
	value string
	left *Formula
	right *Formula
}

func parseFormula(f string) Formula {
	result := Formula{
		value: f,
		left: nil,
		right: nil,
	}
	return result
}

func printFormula(formula Formula) string {
    leftString := ""
    if formula.left != nil {
        leftString = printFormula(*formula.left)
    }
    rightString := ""
    if formula.right != nil {
        rightString = printFormula(*formula.right)
    }
    return "(" + leftString + formula.value + rightString + ")"
}

func main() {
	formulaString := flag.String("formula", "", "The formula in propositional logic")
	flag.Parse()
	formula := parseFormula(*formulaString)
	fmt.Println("Formula: " + formula.value)
}
