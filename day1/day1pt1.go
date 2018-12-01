package main

import (
       "fmt"
       "io/ioutil"
       "strings"
       "strconv"
       )


func main() {
b, err := ioutil.ReadFile("input")
if err != nil {
   fmt.Print(err)
   }

str := string(b)
lines := strings.Split(str, "\n")

freq := 0

for _, element := range lines {
inc, _ := strconv.ParseInt(element,10,64)
freq += int(inc)
}

fmt.Println(freq)

}