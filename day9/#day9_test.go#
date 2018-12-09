package main

import (
	"testing"

)

func TestLoadFile(t *testing.T) {
	expected := []string{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"}

	result := LoadFile("input.test")

	if len(result) != len(expected) || result[0] != expected[0] {
		t.Errorf("Loaded lines were %v, should be %v", result, expected)
	}

}

func TestLoadFileBad(t *testing.T) {
	expected := []string{}

	result := LoadFile("inpoot.test")

	if len(result) != len(expected) {
		t.Errorf("Loaded lines were %v, should be %v", result, expected)
	}

}

func TestLineToInts(t *testing.T) {
	input :=
		expectedplayers :=
		expectedscore := 
	input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

	ps, score := LineToInts(input)
	
	if resultSum != expectedSum {
		t.Errorf("Part One was %v, should be %v", resultSum, expectedSum)
	}

	if resultValue != expectedValue {
		t.Errorf("Part Two was %v, should be %v", resultValue, expectedValue)
	}

	

}

func TestParseNode (t *testing.T) {
	input := []int{2, 3, 0, 3, 10 ,11, 12 ,1 ,1 ,0 ,1 ,99, 2, 1, 1, 2}
	expectedSum := 138
	expectedValue := 66

	_,resultSum, resultValue := ParseNode(input, 0)

	if resultSum != expectedSum {
		t.Errorf("Part One was %v, should be %v", resultSum, expectedSum)
	}

	if resultValue != expectedValue {
		t.Errorf("Part Two was %v, should be %v", resultValue, expectedValue)
	}

}

func TestSolve(t *testing.T) {
	
	expectedSum := 138
	expectedValue := 66
	resultSum, resultValue := Solve("input.test")

	if resultSum != expectedSum {
		t.Errorf("Part One was %v, should be %v", resultSum, expectedSum)
	}

	if resultValue != expectedValue {
		t.Errorf("Part Two was %v, should be %v", resultValue, expectedValue)
	}

	

}
