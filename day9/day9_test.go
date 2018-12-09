package main

import (
	"testing"

)

func TestLoadFile(t *testing.T) {
	expected := []string{"10 players; last marble is worth 1618 points",
"13 players; last marble is worth 7999 points",
"17 players; last marble is worth 1104 points",
"21 players; last marble is worth 6111 points",
"30 players; last marble is worth 5807 points"}

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
	input := "17 players; last marble is worth 1104 points"
	expectedplayers := 17
	expectedmaxmarble := 1104

	ps, maxmarble := LineToInts(input)
	
	if ps != expectedplayers {
		t.Errorf("Players was %v, should be %v", ps, expectedplayers)
	}

	if maxmarble != expectedmaxmarble {
		t.Errorf("Maximum marble was %v, should be %v", maxmarble, expectedmaxmarble)
	}

	

}

