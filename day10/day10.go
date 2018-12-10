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
		fmt.Sscanf(line, "position=<%d,  %d> velocity=<%d,  %d>", &dot.x, &dot.y, &dot.vx, &dot.vy)
		output = append(output, dot)
	}
	return output
}

func FindArea(dots []Point) int {
	bounds := FindBounds(dots)
	return((bounds[1].x - bounds[0].x) * (bounds[1].y - bounds[0].y))
}

func FindDot(dots []Point, x int, y int) string {
	output := "."
	for _, dot := range dots {
		if dot.x == x && dot.y == y {
			output = "#"
		}
	}
	return(output)
}

func PrintView(dots []Point) {
	bounds := FindBounds(dots)
	for y := bounds[0].y; y <= bounds[1].y; y++ {
		var line []string
		for x := bounds[0].x; x <= bounds[1].x; x++ {
			line = append(line,FindDot(dots,x,y))
		}
		fmt.Println(strings.Join(line,""))
	}
}

func DoStep(dots []Point) []Point {
	for i,dot := range dots {
		dots[i].x += dot.vx
		dots[i].y += dot.vy
	}
	return(dots)
}

func FindBounds(dots []Point) []Point {
	var max Point
	var min Point

	max.x = 0
	max.y = 0
	min.x = dots[0].x
	min.y = dots[0].y

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

func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := LinesToPoints(lines)
	minarea := FindArea(input)
	finalstep := 0
	gettingbigger := 0
	lastarea := minarea
	output := make([]Point, len(input))
	for i := 1; i < 1000000; i++ {
		input = DoStep(input)
		area := FindArea(input)
		if area > lastarea {
			gettingbigger++
		}

		if area < lastarea {
			gettingbigger--
		}
		if area < minarea {
			minarea = area
			copy(output,input)
			finalstep = i
		}
		lastarea = area
		if gettingbigger > 10 {
			break
		}
	}
	PrintView(output)

	return (finalstep)
}


func main() {
	start := time.Now()
	inputfile := "input"
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}
	fmt.Println(PartOne(inputfile))
	fmt.Println(time.Since(start))
}
