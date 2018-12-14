package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func SumAndReturnDigits(a, b int) []int {
	var recipes []int
	s := strconv.Itoa(a+b)
	for i := 0; i < len(s); i++ {
		digit,_ := strconv.Atoi(string(s[i]))
	recipes = append(recipes, digit)
	}
	return(recipes)
}

func Matches(recipes []byte, targstring string) bool {
	if len(recipes) < len(targstring) {
		return(false)
	}
	
	for i := 0; i < len(recipes); i++ {
		//fmt.Println(int(targstring[i]-'0'),int(recipes[i]-'0'))
		if targstring[i] != recipes[i] {
			return(false)
		}
	}
	return(true)
}
func Max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	var recipes []byte

	recipes = append(recipes,'3')
	recipes = append(recipes,'7')
	target := "890691"
//	target = "59414" // should yield 2018
//	target = "51589" // should yield 9
//	target = "01245" // hsould yiled 5
//	target = "59414"
	p1 := 0
	p2 := 1
	for x := 0; x < 317000000; x++ {
		s := []byte(strconv.Itoa(int(recipes[p1]-'0')+int(recipes[p2]-'0')))
		recipes = append(recipes, s...)
		
//		recipes = append(recipes,SumAndReturnDigits(recipes[p1],recipes[p2])...)
		p1 = (p1 + int(recipes[p1])-int('0') + 1)%len(recipes)
		p2 = (p2 + int(recipes[p2])-int('0') + 1)%len(recipes)
		
//		if len(recipes) > len(target) {
//			if Matches(recipes[len(recipes)-len(target):],target) {
//				fmt.Println(len(recipes)-len(target)) // this yields badness
//				os.Exit(0)
//			}
//		}
	}
	fmt.Println(strings.Index(string(recipes),target)) // this works
	os.Exit(0)
	
}
