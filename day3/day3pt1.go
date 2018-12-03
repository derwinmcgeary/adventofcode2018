package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func listsquares(claim string)[]string{
	// takes a line like #1 @ 1,3: 4x4
	// and converts it to coords 1,3 2,3 3,3 4,4 1,4 2,4 3,4 4,4 etc
	var squareslist []string
//	fmt.Println(claim)
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
//	fmt.Println(squareslist)
	return squareslist
}

func main() {
	b, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)
	lines := strings.Split(str, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	usedfabric := make(map[string]string)

	for _, claim := range lines {
		
		for _, element := range listsquares(claim) {
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
	
	takeninches := 0

	for _,element := range usedfabric {
		if element == "double" {
			takeninches++
		}
	}

	fmt.Println(takeninches)

	for _, claim := range lines {
		nooverlap := true
		for _, element := range listsquares(claim) {
//			fmt.Println(element)
			if usedfabric[element] == "double" {
				nooverlap = false
			}
		}
		if nooverlap {
			fmt.Println(strings.Split(claim," ")[0])
		}
	}

	
}
