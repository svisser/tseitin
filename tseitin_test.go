package main

import (
	"fmt"
	"testing"
)

func TestParseFormulaLiteral(t *testing.T) {
	result := parseFormula("(a)")
	if result.value != "(a)" {
		t.Error("TestParseFormulaLiteral failed: " + result.value)
	}
	fmt.Println(result.value)
}

func TestParseFormulaLiteralWithConnective(t *testing.T) {
	result := parseFormula("(a^)")
	if result.value != "(a^)" {
		t.Error("TestParseFormulaLiteralWithConnective failed: " + result.value)
	}
	fmt.Println(result.value)
}

func TestParseFormulaConjunction(t *testing.T) {
	result := parseFormula("((a)^(b))")
	if result.value != "^" {
		t.Error("TestParseFormulaConjunction failed for connective: " + result.value)
	}
	if result.left == nil || result.left.value != "(a)" {
		t.Error("TestParseFormulaConjunction failed for left formula.")
	}
	if result.right == nil || result.right.value != "(b)" {
		t.Error("TestParseFormulaConjunction failed for right formula.")
	}
	fmt.Println(result.value)
}
