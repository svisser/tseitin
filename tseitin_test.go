package main

import "fmt"
import "testing"

func TestParseFormulaLiteral(t *testing.T) {
	result := parseFormula("(a)")
	if result.value != "(a)" {
		t.Error("Test_parseFormulaLiteral failed: " + result.value)
	}
	fmt.Println(result.value)
}

func TestParseFormulaConjunction(t *testing.T) {
	result := parseFormula("((a)^(b))")
	if result.value != "^" {
		t.Error("Test_parseFormulaConjunction failed for connective: " + result.value)
	}
	if result.left == nil || result.left.value != "(a)" {
		t.Error("Test_parseFormulaConjunction failed for left formula.")
	}
	if result.right == nil || result.right.value != "(b)" {
		t.Error("Test_parseFormulaConjunction failed for right formula.")
	}
	fmt.Println(result.value)
}