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

func FindIntersectTime(a, b Point) int {
	// If a, b ever intersect, return the time, else return -1
	if b.vx != a.vx && b.vy != a.vy {
		xt := (a.x - b.x)/(b.vx - a.vx)
		yt := (a.y - b.y)/(b.vy - a.vy)
		if xt == yt {
			return xt
		}
	}
	return(-1)
}

func FindIntersects(dots []Point) int {

	freqs := make(map[int]int)
	for i,a := range dots {
		for j,b := range dots {
			if i != j {
				t := FindIntersectTime(a,b)
				if t > 0 {
					freqs[t]++
				}
			}
		}
	}
	maxfreq := 0
	maxtime := 0
	for t,f := range freqs {
		if f > maxfreq {
			maxfreq = f
			maxtime = t
		}
	}
	
	return(maxtime)
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

func DoStep(dots []Point, t int) []Point {
	for i,dot := range dots {
		dots[i].x += t * dot.vx
		dots[i].y += t * dot.vy
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
	finalstep := FindIntersects(input)
	PrintView(DoStep(input,finalstep))
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
