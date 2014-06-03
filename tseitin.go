package main

import "flag"
import "fmt"
import "strconv"
import "strings"

type formula struct {
	value string
	left  *formula
	right *formula
}

var connectives = []string{"^", "v", ">"}

func parseFormula(f string) *formula {
	if strings.HasPrefix(f, "(") && strings.HasSuffix(f, ")") {
		for _, connective := range connectives {
			if strings.Contains(f, connective) {
				endLeft := strings.Index(f, ")"+connective+"(") + 1
				startRight := strings.Index(f, ")"+connective+"(") + 2
				fmt.Println("TEST ONE: " + f[1:endLeft] + " from " + f)
				leftFormula := parseFormula(f[1:endLeft])
				rightFormula := parseFormula(f[startRight : len(f)-1])
				result := formula{
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
		result := formula{
			value: "~",
			left:  leftFormula,
			right: nil,
		}
		return &result
	}
	result := formula{
		value: f,
		left:  nil,
		right: nil,
	}
	return &result
}

func printFormula(formula formula) string {
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

func gatherNames(names map[*formula]string, formula *formula) {
	if formula == nil {
		return
	}
	gatherNames(names, formula.left)
	gatherNames(names, formula.right)
	_, ok := names[formula]
	if !ok {
		names[formula] = getLiteralName(len(names))
	}
}

func shortenFormula(names map[*formula]string, f *formula) *formula {
	if f.left == nil && f.right == nil {
		literal_formula := formula{
			value: names[f],
			left:  nil,
			right: nil,
		}
		return &literal_formula
	}
	result := formula{
		value: f.value,
		left:  nil,
		right: nil,
	}
	if f.left != nil {
		left_formula := formula{
			value: names[f.left],
			left:  nil,
			right: nil,
		}
		result.left = &left_formula
	}
	if f.right != nil {
		right_formula := formula{
			value: names[f.right],
			left:  nil,
			right: nil,
		}
		result.right = &right_formula
	}
	return &result
}

func main() {
	formulaString := flag.String("formula", "", "The formula in propositional logic")
	flag.Parse()
	f := parseFormula(*formulaString)

	fmt.Println("Shorter names:")
	names := map[*formula]string{}
	gatherNames(names, f)
	shortNames := map[*formula]*formula{}
	for subformula, name := range names {
		shortNames[subformula] = shortenFormula(names, subformula)
		fmt.Println(name + ": " +
			printFormula(*shortNames[subformula]) + ": " +
			printFormula(*subformula))
	}

	fmt.Println("Components:")
	fmt.Println(names[f])

	for subformula, shortSubformula := range shortNames {
		if subformula.left != nil || subformula.right != nil {
			fmt.Println(printFormula(*shortSubformula))
		}
	}

	fmt.Println("Formula: " + printFormula(*f))
}
