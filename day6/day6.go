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

func ManhattanDistance(one Point, two Point) int {
	return (Abs(one.x-two.x) + Abs(one.y-two.y))
}

func Abs(n int) int {
	if n < 0 {
		return (-n)
	}
	return (n)
}

func Max(a,b int) int {
	if a > b {
		return(a)
	}
	return(b)
}

func Min(a,b int) int {
	if a < b {
		return(a)
	}
	return(b)
}




func LinesToPoints(input []string) []Point {

	var output []Point
	for _, line := range input {
		var dot Point
		// 1, 1
		fmt.Sscanf(line, "%d, %d", &dot.x, &dot.y)
		output = append(output, dot)
	}
	return output
}

func WhichClosest(dots []Point, testloc Point) int {
	var closest int
	nearestdistance := 10000

	for i, dot := range dots {
		testdistance := ManhattanDistance(dot, testloc)
		if testdistance < nearestdistance {
			closest = i
			nearestdistance = testdistance
			continue
		}
		if testdistance == nearestdistance {
			closest = -1
		}
	}
	return (closest)

}

func FindBounds(dots []Point) []Point {
	var max Point
	var min Point

	max.x = 0
	max.y = 0
	min.x = 1000
	min.y = 1000

	var output []Point

	for _, dot := range dots {
		max.x = Max(max.x,dot.x)
		max.y = Max(max.y,dot.y)
		min.x = Min(min.x,dot.x)
		min.y = Min(min.y,dot.y)
	}
	output = append(output, min)
	output = append(output, max)
	return (output)
}

func MakeField(input []Point, xs, ys, xe, ye int) map[Point]int {
	field := make(map[Point]int)
	for x := xs; x < xe; x++ {
		for y := ys; y < ye; y++ {
			var dot Point
			dot.x = x
			dot.y = y
			field[dot] = WhichClosest(input, dot)
		}
	}
	return (field)
}

func MakeDistanceField(input []Point, xs, ys, xe, ye int) map[Point]int {
	field := make(map[Point]int)
	for x := xs; x < xe; x++ {
		for y := ys; y < ye; y++ {
			for _, inputpoint := range input {
				var dot Point
				dot.x = x
				dot.y = y
				field[dot] += ManhattanDistance(inputpoint, dot)
			}
		}
	}
	return (field)
}

func FreqTable(field map[Point]int) map[int]int {
	freq := make(map[int]int)
	for _, val := range field {
		freq[val]++
	}
	return (freq)
}
func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := LinesToPoints(lines)
	bounds := FindBounds(input)
	var testpoint Point
	testpoint.x, testpoint.y = 10, 10
	field := MakeField(input, bounds[0].x, bounds[0].y, bounds[1].x, bounds[1].y)
	freq := FreqTable(field)

	bigfield := MakeField(input, bounds[0].x-1, bounds[0].y-1, bounds[1].x+1, bounds[1].y+1)
	bigfreq := FreqTable(bigfield)
	maxFreq := 0
	for k, v := range freq {
		if v != bigfreq[k] {
			delete(freq, k)
			continue
		}
		if v > maxFreq {
			if k != -1 {
				maxFreq = v
			}
		}
		
	}

	return (maxFreq)
}

func CountArea(field map[Point]int, threshold int) int {
	area := 0
	var testpoint Point
	testpoint.x = 4
	testpoint.y = 3
	for _, total := range field {
		if total < threshold {
			area++
		}
	}
	return (area)
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := LinesToPoints(lines)
	bounds := FindBounds(input)
	bigfield := MakeDistanceField(input, bounds[0].x, bounds[0].y, bounds[1].x, bounds[1].y)
	threshold := 10000
	if filename == "input.test" {
		threshold = 32
	}
	area := CountArea(bigfield, threshold)
	return (area)
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
