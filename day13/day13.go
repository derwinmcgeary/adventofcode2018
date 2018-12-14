package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
//	"time"
)

type coord struct {
	// really a complex int
	x int
	y int
}

type Cart struct {
	position coord
	direction coord
	nextTurn string
	moved bool
}

func SortCarts(carts []Cart) (outcarts []Cart) {
	mx, my := MaxXY(carts)
	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			testcoord := coord{x,y}
			for _,cart := range carts {
				if cart.position == testcoord {
					outcarts = append(outcarts,cart)
				}
			}
		}
	}
	return
}

func CartAtPosition(carts []Cart, pos coord) bool {
	for _,cart := range carts {
		if cart.position == pos {
			return true
		}
	}
	return false
}

func RemoveCartAtPosition(carts []Cart, pos coord) []Cart {
	for i, cart := range carts {
		if cart.position == pos {
			carts = append(carts[:i], carts[i+1:]...)
		}
	}
	return carts
}

func GetCartAtPosition(carts []Cart, pos coord) (outcart int) {
	for i, cart := range carts {
		if cart.position == pos {
			outcart = i
		}
	}
	return outcart

}

func MaxXY(carts []Cart) (maxx, maxy int) {
	for _,cart := range carts {
		if cart.position.x > maxx {
			maxx = cart.position.x
		}
		if cart.position.y > maxy {
			maxy = cart.position.y
		}
	}
	return
}

func (a coord) Add(b coord) (c coord) {
	c.x = a.x + b.x
	c.y = a.y + b.y
	return
}

func TurnLeft(a coord) (b coord) {
	b = a.Times(coord{0,-1})
	return
}

func TurnRight(a coord) (b coord) {
	b = a.Times(coord{0,1})
	return
}

func (a coord) Times(b coord) (c coord) {
	c.x = a.x * b.x - a.y * b.y
	c.y = a.y * b.x + b.y *a.x
	return
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

func LinesToBoard (lines []string) map[coord]string {
	board := make(map[coord]string)
	
	for y, line := range lines {
		for x, char := range line {
			var pos coord
			pos.x = x
			pos.y = y
			t := string(char)
			if t == "^" || t == "v" {
				board[pos] = "|"
			} else if t == "<" || t == ">" {
				board[pos] = "-"
			}  else {
				board[pos] = t
			}
		}
	}
	return(board)
}

func PrintBoard(carts []Cart, board map[coord]string) {
	dirchar := make(map[coord]string)
	dirchar[coord{0,-1}] = "^"
	dirchar[coord{0,1}] = "v"
	dirchar[coord{-1,0}] = "<"
	dirchar[coord{1,0}] = ">"

	mx,my := MaxXY(carts)
	for y := 0; y <= my + 3; y++ {
		var outstring []string
		for x :=0; x <= mx + 3; x++ {

			if CartAtPosition(carts,coord{x,y}) {
				outstring = append(outstring, dirchar[carts[GetCartAtPosition(carts,coord{x,y})].direction])
			} else if val, ok := board[coord{x,y}]; ok {
				outstring = append(outstring, val)
			} else {
				outstring = append(outstring, " ")
			}
		}
		fmt.Println(strings.Join(outstring,""))
	}
}

func GetCarts(lines []string) []Cart {
	var carts []Cart

	dir := make(map[string]coord)
	dir["^"] = coord{0,-1}
	dir["v"] = coord{0,1}
	dir["<"] = coord{-1,0}
	dir[">"] = coord{1,0}
	
	
	for y, line := range lines {
		for x, char := range line {
			var pos coord
			pos.x = x
			pos.y = y
			t := string(char)
			if v,ok := dir[t]; ok {
				carts = append(carts,Cart{pos,v,"l",false})
			} 
		}
	}
	return(carts)
	
}

func MoveCarts(carts []Cart, board map[coord]string) []Cart {
	mx,my := MaxXY(carts)
	for y := 0; y <= my; y++ {
		for x :=0; x <= mx; x++ {
			curpos := coord{x,y}
			if CartAtPosition(carts, curpos) && carts[GetCartAtPosition(carts,curpos)].moved == false {
				mover := carts[GetCartAtPosition(carts,curpos)]
				nextpos := mover.position.Add(mover.direction)
				if CartAtPosition(carts, nextpos) {
					carts = RemoveCartAtPosition(carts,curpos)
					carts = RemoveCartAtPosition(carts,nextpos)
					fmt.Println("Collision at",nextpos)
				} else {
					mover.position = nextpos
					mover.moved = true
					
					if board[nextpos] == "+" {
						if mover.nextTurn == "l" {
							mover.direction = TurnLeft(mover.direction)
							mover.nextTurn = "s"
						} else if mover.nextTurn == "s" {
							mover.nextTurn = "r"
						} else if mover.nextTurn == "r" {
							mover.direction = TurnRight(mover.direction)
							mover.nextTurn = "l"
						}
					}
					if board[nextpos] == "/" {
						mover.direction = coord{-mover.direction.y,-mover.direction.x}
					}
					
					if board[nextpos] == "\\" {
						mover.direction = coord{mover.direction.y,mover.direction.x}
					}
					carts = RemoveCartAtPosition(carts,curpos)
					carts = append(carts,mover)
				}
			}
		}
	}
	for i,cart := range carts {
		cart.moved = false
		carts[i] = cart
	}
	return(carts)
}


func PartOne() {

}

func main() {
	inputfile := "input"
	if len(os.Args) > 1 {
		inputfile = os.Args[1]
	}


	lines := LoadFile(inputfile)
	carts := GetCarts(lines)
	board := LinesToBoard(lines)
	for n := 0; n < 200000000;n++ {
		carts = MoveCarts(carts, board)
		if len(carts) == 1 {
			//PrintBoard(carts,board)
			fmt.Println("last cart:", carts)
			os.Exit(0)
		}
	}

}
