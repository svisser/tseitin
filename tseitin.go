package main

import (
    "flag"
    "fmt"
    "strconv"
    "strings"
)

type formula struct {
	value string
	left  *formula
	right *formula
}

var connectives = []string{"^", "v", ">"}

func parseFormula(f string) *formula {
	if strings.HasPrefix(f, "(") && strings.HasSuffix(f, ")") {
		for _, connective := range connectives {
			if strings.Contains(f, ")"+connective+"(") {
				endLeft := strings.Index(f, ")"+connective+"(") + 1
				startRight := strings.Index(f, ")"+connective+"(") + 2
				fmt.Println(f + " " + strconv.Itoa(endLeft))
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

func (f *formula) String() string {
	leftString := ""
	if f.left != nil {
		leftString = (*f.left).String()
	}
	rightString := ""
	if f.right != nil {
		rightString = (*f.right).String()
	}
	openString := "("
	closeString := ")"
	if f.left == nil && f.right == nil {
		openString = ""
		closeString = ""
	}
	return openString + leftString + f.value + rightString + closeString
}

func getLiteralName(number int) string {
	return "p" + strconv.Itoa(number+1)
}

func gatherNames(names map[*formula]string, f *formula) {
	if f == nil {
		return
	}
	gatherNames(names, f.left)
	gatherNames(names, f.right)
	_, ok := names[f]
	if !ok {
		names[f] = getLiteralName(len(names))
	}
}

func shortenFormula(names map[*formula]string, f *formula) *formula {
	if f.left == nil && f.right == nil {
		literalFormula := formula{
			value: names[f],
			left:  nil,
			right: nil,
		}
		return &literalFormula
	}
	result := formula{
		value: f.value,
		left:  nil,
		right: nil,
	}
	if f.left != nil {
		leftFormula := formula{
			value: names[f.left],
			left:  nil,
			right: nil,
		}
		result.left = &leftFormula
	}
	if f.right != nil {
		rightFormula := formula{
			value: names[f.right],
			left:  nil,
			right: nil,
		}
		result.right = &rightFormula
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
			(*shortNames[subformula]).String() + ": " +
			(*subformula).String())
	}

	fmt.Println("Components:")
	fmt.Println(names[f])

	for subformula, shortSubformula := range shortNames {
		if subformula.left != nil || subformula.right != nil {
			fmt.Println((*shortSubformula).String())
		}
	}

	fmt.Println("Formula: " + (*f).String())
}
