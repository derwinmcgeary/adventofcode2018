package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(input string) string {
	target := []rune(input)
	sort.Sort(sortRunes(target))
	return string(target)
}

func SplitString(input string) []string {
	// return an array of strings of identical characters in the input string
	var output []string
	startPos := int(0)
	for index, element := range input {
		if index + 1 == len(input) {
			output = append(output, input[startPos:])
			break
		}
		if element != rune(input[index + 1]) {
			output = append(output, input[startPos:index + 1])
			startPos = index + 1
		}
	}
	return output
}

func isRepeat(s string, n int) int {
	// Takes a string and returns 1 iff there are exactly n repeated letters
	foundRepeat := int(0)
	sS := SortString(s)
	sSplit := SplitString(sS)
	for _,element := range sSplit {
		if len(element) == n {
			foundRepeat = 1
		}
	}
	return foundRepeat	
}

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Print(err)
	}
	
	str := string(b)
	lines := strings.Split(str, "\n")
	threes := int(0)
	twos := int(0)
	
	for _, element := range lines {
		twos += isRepeat(element, 2)
		threes += isRepeat(element, 3)
	}
	fmt.Println(twos*threes)
}
