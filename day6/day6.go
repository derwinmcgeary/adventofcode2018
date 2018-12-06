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

func ManhattanDistance(one Point, two Point) int {
	return(Abs(one.x - two.x) + Abs(one.y - two.y))
}

func Abs(n int) int {
	if n < 0 {
		return(-n)
	}
	return(n)
}

func LinesToPoints (input []string) []Point {

	var output []Point
	for _,line := range input {
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

	for i,dot := range dots {
		testdistance := ManhattanDistance(dot,testloc)
		if testdistance < nearestdistance {
			closest = i
			nearestdistance = testdistance
			continue
		}
		if testdistance == nearestdistance {
			closest = -1
		}
	}
	return(closest)
	
}

func FindBounds(dots []Point) []Point {
	var max Point
	var min Point
	
	max.x = 0
	max.y = 0
	min.x = 1000
	min.y = 1000
	
	var output []Point
	
	for _,dot := range dots {
		if dot.x > max.x {
			max.x = dot.x
		}
		if dot.y > max.y {
			max.y = dot.y
		}
		if dot.x < min.x {
			min.x = dot.x
		}
		if dot.y < min.y {
			min.y = dot.y
		}

	}
	output = append(output, min)
	output = append(output, max)
	return(output)
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
	return(field)
}

func MakeDistanceField(input []Point, xs, ys, xe, ye int) map[Point]int {
	field := make(map[Point]int)
	for x := xs; x < xe; x++ {
		for y := ys; y < ye; y++ {
			for _,inputpoint := range input {
				var dot Point
				dot.x = x
				dot.y = y
				field[dot] += ManhattanDistance(inputpoint, dot)}
		}
	}
	return(field)
}

func FreqTable(field map[Point]int) map[int]int {
	freq := make(map[int]int)
	for _,val := range field {
		freq[val]++
	}
	return(freq)
}
func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := LinesToPoints(lines)
	bounds := FindBounds(input)
	var testpoint Point
	testpoint.x, testpoint.y = 10,10
	field := MakeField(input, bounds[0].x, bounds[0].y, bounds[1].x, bounds[1].y)
	freq := FreqTable(field)

	bigfield := MakeField(input, -1000, -1000, 1000, 1000)
	bigfreq := FreqTable(bigfield)

	for k,v := range freq {
		if v != bigfreq[k] {
			delete(freq,k)
		}
	}
	maxFreq := 0

	for k,v := range freq {
		if v > maxFreq {
			if k != -1 {
				maxFreq = v
			}
		}
	}
	return (maxFreq)
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := LinesToPoints(lines)
	bigfield := MakeDistanceField(input, -100, -100, 1000, 1000)

	area := 0
	var testpoint Point
	testpoint.x = 4
	testpoint.y = 3
	for _,total := range bigfield {
		if total < 10000 {
			area++
		}
	}
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
