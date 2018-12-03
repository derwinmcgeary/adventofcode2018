package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func ListSquares(claim string)[]string{
	// takes a line like #1 @ 1,3: 4x4
	// and converts it to coords 1,3 2,3 3,3 4,4 1,4 2,4 3,4 4,4 etc
	var squareslist []string
	components := strings.Split(claim, " ")
	startcoords := strings.Split(components[2], ":")[0]
	startx, _ := strconv.Atoi(strings.Split(startcoords,",")[0])
	starty, _ := strconv.Atoi(strings.Split(startcoords,",")[1])

	size := components[3]
	sizex, _ := strconv.Atoi(strings.Split(size, "x")[0])
	sizey, _ := strconv.Atoi(strings.Split(size, "x")[1])

	for i := 0; i < sizex; i++ {
		for j := 0; j < sizey; j++ {
			xnumber := strconv.Itoa(startx + i)
			ynumber := strconv.Itoa(starty + j)
			squareslist = append(squareslist, strings.Join([]string{xnumber,ynumber},","))
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
//			fmt.Println(element)
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
//			fmt.Println(element)
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
