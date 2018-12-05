package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"time"
)

func LoadFile(filename string)[]string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	lines := strings.Split(str, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return lines

}

func PushSlice(sl []string, addme string) []string {
	output := append(sl, addme)
	return(output)
}

func PopSlice(sl []string) []string {
	return(sl[:len(sl)-1])
}

func PeekSlice(sl []string) string {
	return(sl[len(sl)-1])
}

func StackReduce(input string, ignore string) int {
	var filtered []string
	for _,i := range input {
		if strings.ToLower(ignore) == strings.ToLower(string(i)) {
			continue
		}

		if len(filtered) == 0 {
			filtered = PushSlice(filtered, string(i))
			continue
		}
		if string(i) != PeekSlice(filtered) && strings.ToLower(string(i)) == strings.ToLower(string(PeekSlice(filtered))) {
			filtered = PopSlice(filtered)
		} else {
			filtered = PushSlice(filtered, string(i))
		}
		
	}
	return(len(filtered))
}

func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := lines[0]
	finallength := StackReduce(input, " ")
	return(finallength)
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := lines[0]
	minlength := len(input)
	for _,l := range "abcdefghijklmnopqrstuvwxyz" {
		testlength := StackReduce(input, string(l))
		if testlength < minlength {
			minlength = testlength
		}
	}

	return(minlength)
}

func main() {
	start := time.Now()
	inputfile := "input"
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}		
	fmt.Println(PartOne(inputfile))
	fmt.Println(time.Since(start))		
	fmt.Println(PartTwo(inputfile))
	fmt.Println(time.Since(start))	
}

