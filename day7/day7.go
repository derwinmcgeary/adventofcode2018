package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

type Task struct {
	antecedents []string
	descendents []string
}

type Elf struct {
	task string
	timetogo int
	busy bool
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

func LinesToTaskGraph(lines []string) map[string]Task {
	output := make(map[string]Task)
	for _,line := range lines {
		var name string
		var antecedent string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.",&antecedent, &name)
		
		if val, ok := output[name]; ok {
			val.antecedents = append(val.antecedents, antecedent)
			output[name] = val
			
		} else {
			var t Task
			t.antecedents = append(t.antecedents, antecedent)
			output[name] = t
		}
		
		if val, ok := output[antecedent]; ok {
			val.descendents = append(val.descendents, name)
			output[antecedent] = val
			
		} else {
			var t Task
			t.descendents = append(t.descendents, name)
			output[antecedent] = t
		}
		
	}
	return(output)
}

func RemoveByName(list []string, name string) []string {
	index := -1
	for i, n := range list {
		if n == name {
			index = i	
		}
	}

	if index == -1 {
		return(list)
	} else if index == len(list) - 1 {
		list = list[:len(list)-1]
	} else {
		list = append(list[:index], list[index+1:]...)
	}
	return(list)
}

func GetMinimum(list []string) string {
	sort.Strings(list)
	return(list[0])
}

func Add(list []string, names []string) []string {
	for _,name := range names {
		if NotIn(list, name) {
		list = append(list, name)
		}
	}
	return(list)
}

func NotIn(list []string, name string) bool {
	result := true
	for _,i := range list {
		if i == name {
			result = false
		}
	}
	return(result)
}

func Ready(graph map[string]Task, done []string, candidates []string) []string {
	var output []string
	for _,c := range candidates {
		if AllIn(graph[c].antecedents, done) {
			output = append(output, c)
		}
	}
	return(output)
}

func AllIn(sub []string, pool []string) bool {
	output := true
	for _, test := range sub {
		if NotIn(pool, test){
			output = false
		}
	}
	return(output)
}

func GetOrder(graph map[string]Task) string {
	var output string
	var outputlist []string
	var nextones []string
	
	for k, v := range graph {
		if len(v.antecedents) == 0 {
			nextones = append(nextones, k)
		}
	}

	for i:= 0; i < len(graph); i++ {
		var a string
		a = GetMinimum(Ready(graph, outputlist, nextones))
		outputlist = append(outputlist,a)
		nextones = RemoveByName(nextones,a)
		nextones = Add(nextones, graph[a].descendents)
		

	}
	output = strings.Join(outputlist, "")
	return(output)
}

func GetTime(graph map[string]Task, elves string, offset int) int{
	var output int
	var outputlist []string
	var nextones []string
	team := make(map[string]Elf)
	for _,x := range elves {
		var e Elf
		team[string(x)] = e
	}
	
	for k, v := range graph {
		if len(v.antecedents) == 0 {
			nextones = append(nextones, k)
		}
	}

	for seconds := 0; seconds < 2000 ; seconds ++ {
		for i,elf := range team {
			
			if elf.busy {
				elf.timetogo--
				if elf.timetogo == 0 {
					elf.busy = false

					nextones = Add(nextones, graph[elf.task].descendents)
					outputlist = append(outputlist, elf.task)
					elf.task = ""					
				}
			}
			
			if !elf.busy {
				if len(nextones) > 0 {
					readies := Ready(graph, outputlist, nextones)
					if len(readies) > 0 {
						a := GetMinimum(readies)

						elf.timetogo = int(rune(a[0])) - offset
						elf.busy = true
						elf.task = a
						nextones = RemoveByName(nextones,a)
					}
				}
			}
			team[i] = elf
			
		}

		if len(outputlist) == len(graph) {
			output = seconds
			break
		}

	}
	return(output)
}
	
func PartOne(filename string) string {
	lines := LoadFile(filename)
	input := LinesToTaskGraph(lines)
	return(GetOrder(input))
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := LinesToTaskGraph(lines)
	var output int
	
	if filename == "input.test" {
		output = GetTime(input, "AB", 64)
	} else {
		output = GetTime(input, "ABCDE", 4)
	}

	return(output)
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
