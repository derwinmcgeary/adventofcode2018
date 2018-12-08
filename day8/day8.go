package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"strconv"
)

func LoadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return([]string{})
	}

	str := string(b)
	lines := strings.Split(str, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return lines

}

func ParseNode(input []int, position int) (NewPosition, sum, value int) {
	children := input[position]
	position++
	metadata := input[position]
	position++

	var childvalue []int
	for x := 0; x < children; x++ {
		inc := 0
		var valueinc int
		position, inc, valueinc = ParseNode(input, position)
		sum += inc
		childvalue = append(childvalue,valueinc)
	}

	for x := 0; x < metadata; x++ {
		metadatavalue := input[position]
		sum += metadatavalue
		if metadatavalue > 0 && metadatavalue <= len(childvalue) {
			value += childvalue[metadatavalue - 1]
		}

		if len(childvalue) == 0 {
			value += metadatavalue
		}
		
		position++
	}
	
	return position, sum, value
}

func LineToInts(input string) []int {
	var output []int
	
	entries := strings.Split(input, " ")
	for _, entry := range entries {
		entryint, _ := strconv.Atoi(entry)
		output = append(output, entryint)
	}
	
	return(output)
}

func Solve(filename string) (sum, value int) {
	lines := LoadFile(filename)
	
	input := LineToInts(lines[0])
	_, sum,value = ParseNode(input,0)
	return sum, value
}


func main() {
	start := time.Now()
	inputfile := "input.test"
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}
	fmt.Println(Solve(inputfile))
	fmt.Println(time.Since(start))
}
