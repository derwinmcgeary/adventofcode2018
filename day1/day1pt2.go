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
seenfreq := make(map[int]bool)
seenfreq[0] = true
escape := false

for {
for _, element := range lines {
inc, _ := strconv.ParseInt(element,10,64)
freq += int(inc)
if seenfreq[freq] && int(inc) != 0{
fmt.Println(freq)
escape = true
break
}
seenfreq[freq] = true
}
if escape {
break
}
}

}