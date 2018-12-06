package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func LoadFile(filename string) []string {
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
	return (output)
}

func PopSlice(sl []string) []string {
	return (sl[:len(sl)-1])
}

func PeekSlice(sl []string) string {
	return (sl[len(sl)-1])
}

func StackReduce(input string, ignore string) []string {
	var filtered []string
	ignore = strings.ToLower(ignore)
	for _, i := range input {
		if ignore == strings.ToLower(string(i)) {
			continue
		}

		if len(filtered) == 0 {
			filtered = PushSlice(filtered, string(i))
			continue
		}
		if (int(i) - int(PeekSlice(filtered)[0])) * (int(i) - int(PeekSlice(filtered)[0])) == 1024 {
			filtered = PopSlice(filtered)
		} else {
			filtered = PushSlice(filtered, string(i))
		}

	}
	return (filtered)
}

func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := lines[0]
	finallength := len(StackReduce(input, " "))
	return (finallength)
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := strings.Join(StackReduce(lines[0], " "),"")
	minlength := len(input)
	for _, l := range "abcdefghijklmnopqrstuvwxyz" {
		testlength := len(StackReduce(input, string(l)))
		if testlength < minlength {
			minlength = testlength
		}
	}

	return (minlength)
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
