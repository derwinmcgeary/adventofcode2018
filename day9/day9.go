package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Circle struct{
	marbles []int
	position int
}

func Modular(circ *Circle, position int) int {
	if len(circ.marbles) == 1 {
		return(0)
	}

	if len(circ.marbles) == 2 {
		return(1)
	}
	return position%len(circ.marbles)
}

func InsertMarble(circ *Circle, marble int) *Circle {

	if len(circ.marbles) == 1 {
		circ.marbles = append(circ.marbles,marble)
		circ.position = 1
		return circ
	}
	circ.position = Modular(circ, circ.position + 2)
	// fmt.Println("Insert at position",circ.position)
	circ.marbles = append(circ.marbles, 0)
	copy(circ.marbles[circ.position+1:], circ.marbles[circ.position:])
	circ.marbles[circ.position] = marble

	return circ
	
}

func RemoveMarbleMinusSeven(circ *Circle) *Circle {
	circ.position = circ.position - 7
	if circ.position < 0 {
		circ.position = circ.position + len(circ.marbles)
	}

	circ.marbles = append(circ.marbles[:circ.position], circ.marbles[circ.position + 1:]...)

	return circ	
}

func GetMarbleMinusSeven(circ *Circle) int {
	i := circ.position - 7
	if i < 0 {
		i = i + len(circ.marbles)
	}
	return(circ.marbles[i])
}

func LoadFile(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return([]string{})
	}

	str := string(b)
	lines := strings.Split(str, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return lines

}

func LineToInts(input string) (players, finalmarble int) {
	fmt.Sscanf(input,"%d players; last marble is worth %d points", &players, &finalmarble)
	
	return
}

func Max(scores map[int]int) (max int) {
	for _,v := range scores {
		if v > max {
			max = v
		}
	}
	return
}

func main() {
	start := time.Now()
	inputfile := "input.test"
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}

	input := LoadFile(inputfile)
	if len(input) < 1 {
		fmt.Println("No input!")
		os.Exit(-1)
	}

	players, maxmarble :=LineToInts(input[0])
	maxmarble = maxmarble * 100

	testcirc := new(Circle)
	testcirc.position = 0
	testcirc.marbles = append(testcirc.marbles,0)

	scores := make(map[int]int)
	for i := 1; i<maxmarble; i++ {

		if i%23 == 0 {
			scores[i%players] += i
			scores[i%players] += GetMarbleMinusSeven(testcirc)
			testcirc = RemoveMarbleMinusSeven(testcirc)

		} else {
			testcirc = InsertMarble(testcirc,i)
		}
	}
//	fmt.Println(testcirc)
	fmt.Println(Max(scores))
	fmt.Println(time.Since(start))
}
