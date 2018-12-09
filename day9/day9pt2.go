package main

import (
	"fmt"
	"container/ring"
	"io/ioutil"
	"os"
	"strconv"
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

func PrintCircle(circle *ring.Ring) {
	var output []string
	for x := 0; x < circle.Len() ; x++{
		output = append(output, strconv.Itoa(circle.Value.(int)))
		circle = circle.Next()
	}
	fmt.Println(output)

}

func PlayTurn(circle *ring.Ring, marble int) (outcircle *ring.Ring, score int){
	if marble%23 == 0 {
		score += marble
		circle = circle.Move(-8)
		removed := circle.Unlink(1)
		circle = circle.Next()
		score += removed.Value.(int)

		
	} else {
		circle = circle.Move(1)
		circle = circle.Link(&ring.Ring{Value: marble})
		circle = circle.Move(-1)

	}

	outcircle = circle
//	PrintCircle(outcircle)
	return
}

func PlayGame(players, maxmarble int) (result int) {
	scores := make(map[int]int)
	circle := ring.New(1)
	circle.Value = 0
	
	for i := 1; i<maxmarble + 1; i++ {
		newscore := 0
		circle, newscore = PlayTurn(circle,i)
		scores[i%players] += newscore
	}
	return(Max(scores))

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

	for _,l := range(input) {
		players,maxmarble := LineToInts(l)
		fmt.Println(players, "players; last marble is worth", maxmarble, "points: high score is", PlayGame(players, maxmarble))
		maxmarble = maxmarble * 100
		fmt.Println(players, "players; last marble is worth", maxmarble, "points: high score is", PlayGame(players, maxmarble))
	}
	fmt.Println(time.Since(start))
}
