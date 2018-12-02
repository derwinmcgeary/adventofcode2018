package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func skipLetter(s string, n int) string {
	// return the input s minus the specified character (0-indexed)
	if len(s) == 0 {
		return s
	}
	if n == len(s) - 1 {
		return s[:n]
	}
	if n == 0 {
		return s[1:]
	}

	return s[:n] + s [n + 1:]
}

func skipLetters(s []string, n int) []string {
	var output []string
	for _ , element := range s {
		output = append(output, skipLetter(element, n))
	}
	return output
}

func maxLen(s []string) int {
	maxLength := 0
	for _, element := range s {
		if len(element) > maxLength {
			maxLength = len(element)
		}
	}
	return maxLength
}

func findMatch(s []string) string {

	seenStrings := make(map[string]bool)
	for _ , element := range s {
		if seenStrings[element] {
			return element
		}
		seenStrings[element] = true
	}
	return "false"
}

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Print(err)
	}
	
	str := string(b)
	lines := strings.Split(str, "\n")

	for i := 0; i < maxLen(lines); i++ {
		match := findMatch(skipLetters(lines,i))
		if match != "false" {
			fmt.Println(match)
		}
	}
}
