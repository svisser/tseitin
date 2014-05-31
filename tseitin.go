package main

import "flag"
import "fmt"
import "strconv"
import "strings"

type Formula struct {
	value string
	left  *Formula
	right *Formula
}

func parseFormula(f string) *Formula {
	if strings.HasPrefix(f, "(") && strings.HasSuffix(f, ")") {
		for _, connective := range []string{"^", "v", ">"} {
			if strings.Contains(f, connective) {
				endLeft := strings.Index(f, ")"+connective+"(") + 1
				startRight := strings.Index(f, ")"+connective+"(") + 2
				leftFormula := parseFormula(f[1:endLeft])
				rightFormula := parseFormula(f[startRight : len(f)-1])
				result := Formula{
					value: connective,
					left:  leftFormula,
					right: rightFormula,
				}
				return &result
			}
		}
	}
	if strings.HasPrefix(f, "~") {
		leftFormula := parseFormula(f[1:])
		result := Formula{
			value: "~",
			left:  leftFormula,
			right: nil,
		}
		return &result
	}
	result := Formula{
		value: f,
		left:  nil,
		right: nil,
	}
	return &result
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
	openString := "("
	closeString := ")"
	if formula.left == nil && formula.right == nil {
		openString = ""
		closeString = ""
	}
	return openString + leftString + formula.value + rightString + closeString
}

func getLiteralName(number int) string {
	return "p" + strconv.Itoa(number+1)
}

func gatherNames(names map[string]string, formula *Formula) {
	if formula == nil {
		return
	}
	gatherNames(names, formula.left)
	gatherNames(names, formula.right)
	displayFormula := printFormula(*formula)
	_, ok := names[displayFormula]
	if !ok {
		names[displayFormula] = getLiteralName(len(names))
	}
}

func main() {
	formulaString := flag.String("formula", "", "The formula in propositional logic")
	flag.Parse()
	formula := parseFormula(*formulaString)

	names := map[string]string{}
	gatherNames(names, formula)
	for key, value := range names {
		fmt.Println(key + ": " + value)
	}

	fmt.Println("Formula: " + printFormula(*formula))
}
