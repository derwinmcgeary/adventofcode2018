package main

import (
	"testing"

)


func ExpectedMap() map[string]string {
	return(map[string]string{
		"5,6":"taken",
		"4,3":"double",
		"3,2":"taken",
		"5,1":"taken",
		"5,4":"taken",
		"1,3":"taken",
		"1,4":"taken",
		"2,6":"taken",
		"4,2":"taken",
		"3,1":"taken",
		"6,4":"taken",
		"2,5":"taken",
		"5,5":"taken",
		"6,5":"taken",
		"1,5":"taken",
		"4,6":"taken",
		"4,1":"taken",
		"6,3":"taken",
		"6,6":"taken",
		"3,5":"taken",
		"3,6":"taken",
		"4,5":"taken",
		"5,2":"taken",
		"2,3":"taken",
		"4,4":"double",
		"5,3":"taken",
		"6,1":"taken",
		"6,2":"taken",
		"1,6":"taken",
		"2,4":"taken",
		"3,3":"double",
		"3,4":"double"})

}

func ExpectedClaims () []string {
	return([]string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4","#3 @ 5,5: 2x2"})
}

func TestListsquares (t *testing.T) {
	expected := []string{"0,0", "0,1", "1,0", "1,1"}
	testinput := "#1 @ 0,0: 2x2"
	
	result := ListSquares(testinput)
	for i, resultelement := range result {
		if resultelement != expected[i] {
			t.Errorf("List was incorrect, got: %s, want: %s.", expected, result)
		}
	}
}

func TestLoadFile (t *testing.T) {
	expected := ExpectedClaims()
	result := LoadFile("input.test")

	if len(expected) != len(result) {
		t.Errorf("List length was incorrect, got: %d, want: %d.", len(result), len(expected))
	}

	
	for i, resultelement := range result {
		if resultelement != expected[i] {
			t.Errorf("List was incorrect, got: %s, want: %s.", result, expected)
		}
	}	
}

func TestLoadFileError (t *testing.T) {
	result := LoadFile("inputlkadfglkjerthg")
	if len(result) != 0 {
		t.Errorf("Non-existent file should not produce result")
	}
}
func TestMakeMap (t *testing.T) {
	expected := ExpectedMap()
	input := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4","#3 @ 5,5: 2x2"}
	result := MakeMap(input)

	if len(expected) != len(result) {
		t.Errorf("List length was incorrect, got: %d, want: %d.", len(result), len(expected))
	}

	
	for key, value := range result {
		if value != expected[key] {
			t.Errorf("List was incorrect, got: %s, want: %s.", result, expected)
		}
	}	
}

func TestCountOverlap (t *testing.T) {
	input := ExpectedMap()
	expected := 4

	result := CountOverlap(input)

	if expected != result {
		t.Errorf("Overlap count was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestFindNonOverlap (t *testing.T) {
	inputmap := ExpectedMap()
	inputclaims := ExpectedClaims()
	expected := "#3"

	result := FindNonOverlap(inputclaims, inputmap)

	if expected != result {
		t.Errorf("Nonoverlap was incorrect, got: %s, want: %s.", result, expected)
	}
}
func TestNoOverlapFound (t *testing.T) {
	
	expected := "Not Found"
	inputclaims := ExpectedClaims()
	inputclaims = append(inputclaims, "#4 @ 5,5: 2x2")
	inputclaims = append(inputclaims, "#5 @ 5,5: 2x2")
	inputmap := MakeMap(inputclaims)
	result := FindNonOverlap(inputclaims, inputmap)
	if expected != result {
		t.Errorf("Nonoverlap was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestPartOne (t *testing.T) {
	input := "input.test"
	expected := 4

	result := PartOne(input)

	if expected != result {
		t.Errorf("Overlap count was incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPartTwo (t *testing.T) {
	input := "input.test"
	expected := "#3"

	result := PartTwo(input)

	if expected != result {
		t.Errorf("Overlap count was incorrect, got: %s, want: %s.", result, expected)
	}
}


func ExampleMainFunction () {
	main()
	// output:
	// 111630
	// #724
	
}
