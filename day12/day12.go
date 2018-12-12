package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
	vx int
	vy int
}

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

func LineToPlants(line string) map[int]string {

	plants := make(map[int]string)
	input := strings.Split(line,": ")[1]
	for i,j := range input {
		plants[i] = string(j)
	}
	return(plants)
}

func LinesToRules(lines []string) map[string]string {

	rules := make(map[string]string)
	for _,line := range lines {
		rule := strings.Split(line, " => ")
		rules[rule[0]] = rule[1]
	}
	return rules
}

func Bounds(plants map[int]string) (min, max int) {
	for i,_ := range plants {
		if i < min {
			min = i
		}
		if i > max {
			max =i
		}
	}
	return
}

func DoStep(plants map[int]string, rules map[string]string) map[int]string {

	output := make(map[int]string)
	min, max := Bounds(plants)
	for i := min - 2; i <= max + 2; i++ {
		var testcase []string
		for j := -2; j <= 2; j++ {
			if val,ok := plants[i + j]; ok {
				testcase = append(testcase,val)
			} else {
				testcase = append(testcase, ".")
			}
		}

		if val,ok := rules[strings.Join(testcase,"")];ok { 
			output[i] = val
		} else {
			output[i] = "."
		}
	}
	return(output)
}

func PrintPlants(plants map[int]string) {
	var outline []string
	min, max := Bounds(plants)
	for i := min; i <= max; i++ {
		outline = append(outline,plants[i])
	}
}

func CountUp(plants map[int]string) (counter int) {
	min, max := Bounds(plants)
	for i := min; i <= max; i++ {
		if plants[i] == "#" {
			counter += i
		}
	}
	return(counter)
}

func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := LineToPlants(lines[0])
	rules := LinesToRules(lines[2:])

	finalcount := 0
	curcount := 0
	limit := 20
	for i := 1; i <= limit; i++ {
		input = DoStep(input,rules)
		curcount = CountUp(input)

		finalcount = curcount
	}
	return (finalcount)
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := LineToPlants(lines[0])
	rules := LinesToRules(lines[2:])

	lastcount := 0
	finalcount := 0
	lastdelta := 0
	curcount := 0
	cycling := 0
	limit := 50000000000
	for i := 1; i <= limit; i++ {
		input = DoStep(input,rules)
		curcount = CountUp(input)
		delta := curcount - lastcount
		if lastdelta == delta {
			cycling++
		} else {
			cycling = 0
		}
		if cycling > 10 {
			finalcount = curcount + (limit - i) * delta
			break
		}
		lastcount = curcount
		lastdelta = delta
	}
	return (finalcount)
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
