package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
	"os"
	"time"
)

type guardevent struct {
	date   time.Time
	guard int
	event string
}

type shift struct {
	guard int
	sleepminutes []int
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

func LinesToRecords(lines []string) []guardevent {
	sort.Strings(lines)
	
	var outputlist []guardevent

	currentguard := 0
	for _,element := range lines {
		var output guardevent
		var d string
		var t string
		var g int
		var e string
		records := strings.Split(element, "] ")
		fmt.Sscanf(records[0], "[%s", &d)
		fmt.Sscanf(strings.Split(records[0]," ")[1], "%s", &t)
		eventstring := records[1]

		if strings.Contains(eventstring, "Guard"){
			fmt.Sscanf(eventstring, "Guard #%d", &g)
			currentguard = g
			e = "begins"
		}
		if strings.Contains(eventstring, "wakes"){
			g = currentguard
			e = "wakes"
		}
		if strings.Contains(eventstring, "asleep") {
			g = currentguard
			e = "sleeps"
		}

		output.guard = g
		output.event = e
		output.date, _ = time.Parse("2006-01-02 15:04",strings.Join([]string{d,t}, " "))
		outputlist = append(outputlist, output)
	}
	return outputlist
}


func EventsToShift (guardevents []guardevent) shift {
	// parse the events for one day into a shift
	// takes one full day of events

	startminute := 0
	endminute := 0
	var sleepminutes []int
	var output shift
	fmt.Println("Here!")
	fmt.Println(guardevents)
	guard := guardevents[0].guard
	for _, event := range guardevents {
		if event.event == "sleeps" {
			startminute = event.date.Minute()
		}
		if event.event == "wakes" {
			endminute = event.date.Minute()
			for i := startminute; i < endminute; i++ {
				sleepminutes = append(sleepminutes,i)
			}
		}
	}
	output.sleepminutes = sleepminutes
	output.guard = guard
	return output
}

func EventsToShifts(guardevents []guardevent) []shift {
	// parse the entire event log and return a shift log with one entry per day
	minDate := guardevents[0].date
	maxDate := guardevents[len(guardevents)-1].date
	var output []shift
	for i := minDate; i.YearDay() <= maxDate.YearDay() && i.Year() <= maxDate.Year(); i = i.AddDate(0,0,1) {

		var oneDay []guardevent
		for _, event := range guardevents {
			if event.date.YearDay() == i.YearDay() && event.date.Year()==i.Year() {
				oneDay = append(oneDay, event)
			}
		}
		if len(oneDay) > 0 {
			output = append(output, EventsToShift(oneDay))
		}
	}

	return(output)

}

func FindMax(someMap map[int]int) int {
	maximumIndex := 0
	for index,value := range someMap {
		if value > someMap[maximumIndex] {
			maximumIndex = index
		}
		
	}
	return maximumIndex
}

func SleepiestGuard(shifts []shift) int {
	guards := make(map[int]int)
	for _,event := range shifts {
		guards[event.guard]+=len(event.sleepminutes)
	}
	sleepiest := FindMax(guards)
	return(sleepiest)
}

func SleepiestMinute(shifts []shift, guard int) int {
	sleepmins := make(map[int]int)
	
	for _,duty := range shifts {
		if duty.guard == guard {
			for _,min := range duty.sleepminutes {
				sleepmins[min]++
			}
		}
	}

	return(FindMax(sleepmins))
}

func PartOne(filename string) int {
	lines := LoadFile(filename)
	myevents := LinesToRecords(lines)
	shifts := EventsToShifts(myevents)
	sg := SleepiestGuard(shifts)
	sm := SleepiestMinute(shifts,sg)
	return(sg*sm)
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	myevents := LinesToRecords(lines)
	shifts := EventsToShifts(myevents)
	guardsleepminutes := make(map[int]map[int]int)

	for _, event := range shifts {
		if guardsleepminutes[event.guard] == nil {
			guardsleepminutes[event.guard] = make(map[int]int)
		}
		for _, min := range event.sleepminutes {
			guardsleepminutes[event.guard][min]++
		}
	}

	maxg := 0
	maxmin := 0
	maxtimes :=0
	
	for g,ms := range guardsleepminutes {
		for min, times := range ms {
			if times > maxtimes {
				maxg = g
				maxtimes = times
				maxmin = min
			}
		}
	}
	return(maxg*maxmin)
}

func main() {
	inputfile := "input"
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}		
	fmt.Println(PartOne(inputfile))
	fmt.Println(PartTwo(inputfile))
	
	
}
