package main

import (
	"testing"

)

func TestAllInTrue (t *testing.T) {
	test := []string{"A", "B", "C", "D"}
	pool := []string{"A", "B", "C", "D", "E"}
	expected := true
	
	result := AllIn(test, pool)
	if result != expected {
		t.Errorf("AllIn was incorrect, should be true")
	}
}

func TestAllInFalse (t *testing.T) {
	test := []string{"A", "B", "C", "D"}
	pool := []string{"A", "B", "C", "D", "E"}
	expected := false
	
	result := AllIn(pool, test)
	if result != expected {
		t.Errorf("AllIn was %t, should be %t", result, expected)
	}
}

func TestNotInTrue(t *testing.T) {
	pool := []string{"A", "B", "C", "D", "E"}
	test := "F"
	expected := true
	result := NotIn(pool, test)
	if result != expected {
		
	}
	
}

func TestNotInFalse(t *testing.T) {
	pool := []string{"A", "B", "C", "D", "E"}
	test := "D"
	expected := false
	result := NotIn(pool, test)
	if result != expected {
		t.Errorf("NotIn was %t, should be %t", result, expected)
	}
	
}

func TestAddEqual(t *testing.T) {
	A := []string{"A", "B", "C", "D", "E"}
	B := []string{"A", "B", "C", "D", "E"}
	expected := []string{"A", "B", "C", "D", "E"}
	result := Add(A,B)
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
}

func TestAdd(t *testing.T) {
	A := []string{"A", "B", "C"}
	B := []string{"D", "E"}
	expected := []string{"A", "B", "C", "D", "E"}
	result := Add(A,B)
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
}

func TestRemoveByName(t *testing.T) {
	A := []string{"A", "B", "C"}
	B := "B"
	expected := []string{"A", "C"}
	result := RemoveByName(A,B)
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
}

func TestRemoveByNameLast(t *testing.T) {
	A := []string{"A", "B", "C"}
	B := "C"
	expected := []string{"A", "B"}
	result := RemoveByName(A,B)
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
}

func TestRemoveByNameAbsent(t *testing.T) {
	A := []string{"A", "B", "C"}
	B := "D"
	expected := []string{"A","B","C"}
	result := RemoveByName(A,B)
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
}

func TestGetMinimum(t *testing.T) {
	A := []string{"Z" ,"A", "B", "C"}
	expected := "A"
	result := GetMinimum(A)
	if result != expected {
		t.Errorf("Add was %v, should be %v", result, expected)
	}
}

func TestReady(t *testing.T) {
	examplelines := []string{"Step C must be finished before step A can begin.", "Step C must be finished before step F can begin.", "Step A must be finished before step B can begin.", "Step A must be finished before step D can begin.", "Step B must be finished before step E can begin.", "Step D must be finished before step E can begin.", "Step F must be finished before step E can begin."}
	examplegraph := LinesToTaskGraph(examplelines)

		// map[string]main.Task{"D":main.Task{antecedents:[]string{"A"}, descendents:[]string{"E"}}, "E":main.Task{antecedents:[]string{"B", "D", "F"}, descendents:[]string(nil)}, "A":main.Task{antecedents:[]string{"C"}, descendents:[]string{"B", "D"}}, "C":main.Task{antecedents:[]string(nil), descendents:[]string{"A", "F"}}, "F":main.Task{antecedents:[]string{"C"}, descendents:[]string{"E"}}, "B":main.Task{antecedents:[]string{"A"}, descendents:[]string{"E"}}}

	next := []string{"A","F", "B", "D"}
	outputlist := []string{"C"}
	result := Ready(examplegraph, outputlist,next)
	expected := []string{"A","F"}
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
	

}

func TestReadyTwo(t *testing.T) {
	examplelines := []string{"Step C must be finished before step A can begin.", "Step C must be finished before step F can begin.", "Step A must be finished before step B can begin.", "Step A must be finished before step D can begin.", "Step B must be finished before step E can begin.", "Step D must be finished before step E can begin.", "Step F must be finished before step E can begin."}
	examplegraph := LinesToTaskGraph(examplelines)

		// map[string]main.Task{"D":main.Task{antecedents:[]string{"A"}, descendents:[]string{"E"}}, "E":main.Task{antecedents:[]string{"B", "D", "F"}, descendents:[]string(nil)}, "A":main.Task{antecedents:[]string{"C"}, descendents:[]string{"B", "D"}}, "C":main.Task{antecedents:[]string(nil), descendents:[]string{"A", "F"}}, "F":main.Task{antecedents:[]string{"C"}, descendents:[]string{"E"}}, "B":main.Task{antecedents:[]string{"A"}, descendents:[]string{"E"}}}

	next := []string{"B","F", "E"}
	outputlist := []string{"C", "A", "D"}
	result := Ready(examplegraph, outputlist,next)
	expected := []string{"B","F"}
	for i, e := range expected {
		if result[i] != e {
			t.Errorf("Add was %v, should be %v", result, expected)
		}
	}
	

}

func TestGetTime(t *testing.T) {
	examplelines := []string{"Step C must be finished before step A can begin.", "Step C must be finished before step F can begin.", "Step A must be finished before step B can begin.", "Step A must be finished before step D can begin.", "Step B must be finished before step E can begin.", "Step D must be finished before step E can begin.", "Step F must be finished before step E can begin."}
	examplegraph := LinesToTaskGraph(examplelines)

	elves := "AB"
	offset := 64
	result := GetTime(examplegraph, elves, offset)
	expected := 15

	if result != expected {
		t.Errorf("Time was %v, should be %v", result, expected)
	}

	

}

func TestGetOrder(t *testing.T) {
	examplelines := []string{"Step C must be finished before step A can begin.", "Step C must be finished before step F can begin.", "Step A must be finished before step B can begin.", "Step A must be finished before step D can begin.", "Step B must be finished before step E can begin.", "Step D must be finished before step E can begin.", "Step F must be finished before step E can begin."}
	examplegraph := LinesToTaskGraph(examplelines)
	result := GetOrder(examplegraph)
	expected := "CABDFE"

	if result != expected {
		t.Errorf("Time was %v, should be %v", result, expected)
	}

	

}

func TestLoadFile(t *testing.T) {
	expected := []string{"Step C must be finished before step A can begin.", "Step C must be finished before step F can begin.", "Step A must be finished before step B can begin.", "Step A must be finished before step D can begin.", "Step B must be finished before step E can begin.", "Step D must be finished before step E can begin.", "Step F must be finished before step E can begin."}

	result := LoadFile("input.test")

	if !AllIn(result,expected) || !AllIn(expected,result) {
		t.Errorf("Loaded lines were %v, should be %v", result, expected)
	}

	

}

func TestPartOne(t *testing.T) {
	
	expected := "CABDFE"
	result := PartOne("input.test")

	if result != expected {
		t.Errorf("PartOne was %v, should be %v", result, expected)
	}

	

}

func TestPartTwo(t *testing.T) {
	expected := 15

	result := PartTwo("input.test")

	if result != expected {
		t.Errorf("Time was %v, should be %v", result, expected)
	}

	

}

/*
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
*/
