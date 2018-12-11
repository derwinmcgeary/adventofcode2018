package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
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

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return(a)
}

func PrintView(grid map[Point]int, xlim, ylim int) {
	for y := 1; y <= ylim; y++ {
		var line []string
		for x := 1; x <= xlim; x++ {
			var z Point
			z.x = x
			z.y = y
			line = append(line,strconv.Itoa(grid[z]))
		}
		fmt.Println(strings.Join(line,""))
	}
}

func ExtractHundredsDigit(input int) int {
	if input < 100 {
		return(0)
	}
	digits := strconv.Itoa(input)
	digit := string(digits[len(digits)-3])
	output, _ := strconv.Atoi(digit)
	return(output)
}

func CoordsToPowerLevel(x,y,s int) int {
	rackid := x + 10
	power := (rackid * y) + s
	power = power * rackid
	power = ExtractHundredsDigit(power)
	power = power - 5

	return(power)
}

func CreateGrid(s,xlim,ylim int) map[Point]int {
	output := make(map[Point]int)

	for x := 1; x <= xlim; x++ {
		for y :=1; y <= ylim; y++ {
			var tmp Point
			tmp.x = x
			tmp.y = y
			output[tmp] = CoordsToPowerLevel(x,y,s)
		}
	}
	return(output)
}

func FindOptimalCoords(grid map[Point]int, xlim, ylim int) Point {

	var max Point
	var maxval int

	sst := GenerateSummedSquaresTable(grid, xlim, ylim)
	

	
	for x := 1; x <= xlim-2; x++ {
		for y :=1; y <= ylim-2; y++ {
			curval := AreaFromSummedSquaresTable(sst,x,y,3)
			if curval > maxval {
				maxval = curval
				max = Point{x: x,y: y}
			}
		}
	}
	return(max)
	
}

func FindOptimalSquare(grid map[Point]int, xlim, ylim int) (Point, int) {
// This method takes literally more than an hour but I'm leaving it in as a warning to others
	var max Point
	var maxval int
	var maxsize int
	
	for size := 1; size <= xlim; size++ {
		fmt.Println(size)
		for x := 1; x <= xlim - size + 1; x++ {
			for y :=1; y <= ylim - size + 1; y++ {
				var tmp Point
				var curval int
				tmp.x = x
				tmp.y = y
				
				for xi := 0; xi < size; xi++ {
					for yi := 0; yi < size; yi++ {
						var temp Point
						temp.x = tmp.x + xi
						temp.y = tmp.y + yi
						curval += grid[temp]
					}
				}
				if curval > maxval {
					maxval = curval
					max = tmp
					maxsize = size
				}
			}
		}
	}
	return max, maxsize
	
}

func MaxSquareSize(dot Point, xlim, ylim int) int {
	xspace := Min(xlim - dot.x,dot.x)
	yspace:= Min(ylim - dot.y,dot.y)

	return Min(xspace,yspace)
}

func GenerateSummedSquaresTable(grid map[Point]int, xlim, ylim int) map[Point]int {
	output := make(map[Point]int)

	for y := 0; y <= ylim; y++ {
		horizontalsum := 0
		for x := 0; x <= xlim; x++ {
			if x == 00 || y == 0 {
				output[Point{x,y}] = 0
			}
			if x == 1 && y == 1 {
				output[Point{x,y}] = grid[Point{x,y}]
				horizontalsum += grid[Point{x,y}]
			}
			if y == 1 && x > 1 {
				horizontalsum += grid[Point{x,y}]
				output[Point{x,y}] = horizontalsum
			}
			if y > 1 {
				horizontalsum += grid[Point{x,y}]
				output[Point{x,y}] = output[Point{x,y-1}] + horizontalsum
			}
		}
	}
	return(output)
}

func AreaFromSummedSquaresTable(sst map[Point]int, x, y, s int) int {
	area := sst[Point{x-1,y-1}] + sst[Point{x + s - 1, y + s - 1}] -  sst[Point{x-1, y + s - 1}] - sst[Point{x + s - 1, y-1}]

	return(area)
}

func FindOptimalSquareOptimally(grid map[Point]int, xlim, ylim int) (Point, int) {
// 1 second
	var max Point
	var maxval int
	var maxsize int
	var curval int
	sst := GenerateSummedSquaresTable(grid,xlim,ylim)
	for x := 1; x <= xlim ; x++ {
		for y :=1; y <= ylim ; y++ {
			var dot Point
			dot.x = x
			dot.y = y
			curval = 0
			for size := 0; size < MaxSquareSize(dot,xlim,ylim); size++ {
				curval = AreaFromSummedSquaresTable(sst,x,y,size + 1)

				if curval > maxval {
					maxval = curval
					max = dot
					maxsize = size + 1
				}
			}
		}
	}
	return max, maxsize
	
}

func PartOne(input int) Point {
	var output Point
	output = FindOptimalCoords(CreateGrid(input,300,300),300,300)
	return (output)
}

func PartTwo(input int) (Point, int) {
	var output Point
	var outputsize int
	output, outputsize = FindOptimalSquareOptimally(CreateGrid(input,300,300),300,300)
	return output,outputsize
}

func main() {
	start := time.Now()
	input := 9110
	ptone := PartOne(input)
	fmt.Printf("%d,%d\n",ptone.x,ptone.y)
	fmt.Println(time.Since(start))
	pttwo, size := PartTwo(input)
	fmt.Printf("%d,%d,%d\n", pttwo.x, pttwo.y, size)
	fmt.Println(time.Since(start))
}
