package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"time"
)

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

func RemoveLetter(input string, letter string) string {
	replaceone := strings.ToUpper(letter)
	replacetwo := strings.ToLower(letter)
	r := strings.NewReplacer(replaceone, "", replacetwo, "")
	output := r.Replace(input)
	return(output)
}

func deDupe (input string) string {
	r := strings.NewReplacer(
		"Aa","",
		"Bb","",
		"Cc","",
		"Dd","",
		"Ee","",
		"Ff","",
		"Gg","",
		"Hh","",
		"Ii","",
		"Jj","",
		"Kk","",
		"Ll","",
		"Mm","",
		"Nn","",
		"Oo","",
		"Pp","",
		"Qq","",
		"Rr","",
		"Ss","",
		"Tt","",
		"Uu","",
		"Vv","",
		"Ww","",
		"Xx","",
		"Yy","",
		"Zz","",
		"aA","",
		"bB","",
		"cC","",
		"dD","",
		"eE","",
		"fF","",
		"gG","",
		"hH","",
		"iI","",
		"jJ","",
		"kK","",
		"lL","",
		"mM","",
		"nN","",
		"oO","",
		"pP","",
		"qQ","",
		"rR","",
		"sS","",
		"tT","",
		"uU","",
		"vV","",
		"wW","",
		"xX","",
		"yY","",
		"zZ","")
	input = r.Replace(input)
	return(input)
}

func boilDown(input string) string {
	currentlength := len(input)
	for {
		input = deDupe(input)
		newlength := len(input)
		if newlength == currentlength {
			break
		}
		currentlength = newlength
	}
	return(input)
}

func PartOne(filename string) int {
	lines := LoadFile(filename)
	input := lines[0]
	input = boilDown(input)
	return(len(input))
}

func PartTwo(filename string) int {
	lines := LoadFile(filename)
	input := lines[0]
	minlength := len(input)
	for _,l := range "abcdefghijklmnopqrstuvwxyz" {
		testlength := len(boilDown(RemoveLetter(input,string(l))))
		if testlength < minlength {
			minlength = testlength
		}
	}
	return(minlength)
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

