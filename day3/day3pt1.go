package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type claim struct {
	id   int
	x, y int
	w, h int
	overlaps bool
}

func ListSquares(record string)[]string{
	// takes a line like #1 @ 1,3: 4x4
	// and converts it to coords 1,3 2,3 3,3 4,4 1,4 2,4 3,4 4,4 etc
	var squareslist []string
	var c claim
	
	fmt.Sscanf(record, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
	
	for i := 0; i < c.w; i++ {
		for j := 0; j < c.h; j++ {
			x := strconv.Itoa(c.x + i)
			y := strconv.Itoa(c.y + j)
			squareslist = append(squareslist, strings.Join([]string{x,y},","))
		}
		
	}
	return squareslist
}
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

func CountOverlap(usedfabric map[string]string)int {
	
	takeninches := 0

	for _,element := range usedfabric {
		if element == "double" {
			takeninches++
		}
	}
	return takeninches
}


func MakeMap(claims []string)map[string]string {
	usedfabric := make(map[string]string)

	for _, claim := range claims {
		
		for _, element := range ListSquares(claim) {
			if usedfabric[element] == "taken" {
				usedfabric[element] = "double"
			} else if usedfabric[element] == "double" {
				usedfabric[element] = "double"
			} else {
				usedfabric[element] = "taken"
			}
		}
	}
	return usedfabric
}

func FindNonOverlap(claims []string, usedfabric map[string]string)string {
	for _, claim := range claims {
		nooverlap := true
		for _, element := range ListSquares(claim) {
			if usedfabric[element] == "double" {
				nooverlap = false
			}
		}
		if nooverlap {
			return(strings.Split(claim," ")[0])
		}
	}
	return("Not Found")
}

func PartOne(filename string) int {
	lines := LoadFile(filename)
	usedfabric := MakeMap(lines)
	takeninches := CountOverlap(usedfabric)
	return(takeninches)


}

func PartTwo(filename string) string {
	lines := LoadFile(filename)
	usedfabric := MakeMap(lines)
	return(FindNonOverlap(lines, usedfabric))
}

func main() {
	inputfile := "input"
	fmt.Println(PartOne(inputfile))
	fmt.Println(PartTwo(inputfile))
	
	
}
