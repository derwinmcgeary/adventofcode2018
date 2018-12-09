package main

import (
	"fmt"
	"container/ring"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

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
	scores := make(map[int]int)
	circle := ring.New(1)
	circle.Value = 0
//	players,maxmarble = 13,7999
//	players,maxmarble = 5,25

	for i := 1; i<maxmarble; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}
		if i%23 == 0 {
			scores[i%players] += i
			circle = circle.Move(-8)
			removed := circle.Unlink(1)
			circle = circle.Next()
			scores[i%players] += removed.Value.(int)
//			fmt.Println("Current", circle.Value)
		} else {
//			var output []int
//			for i := 0; i<circle.Len(); i++ {
//				output = append(output, circle.Value.(int))
//				circle = circle.Next()
//			}
//			fmt.Println(output)
			tmp := ring.New(1)
			tmp.Value = i
			circle = circle.Move(1)
			circle = circle.Link(tmp)
			circle = circle.Move(-1)
//			if circle.Value.(int) == i {
//				fmt.Println("OK")
//			}


		}
	}
//	fmt.Println(testcirc)
	fmt.Println(Max(scores))
	fmt.Println(time.Since(start))
}
