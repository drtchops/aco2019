package solutions

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var width13 int64 = 43
var height13 int64 = 23

type tile13 struct {
	x  int64
	y  int64
	id int64
}

// Day13 ...
func Day13(input string) {
	prog := parseProg(input)
	prog[0] = 2
	in := make(chan int64, 100)
	out := make(chan int64, 100)
	term := make(chan termSig, 100)
	state := make([]int64, width13*height13+1)

	go runIntcode(prog, 0, in, out, term)

	var needsInput bool
	for {
		var done bool
		changed := updateState13(state, out)

		if changed {
			needsInput = true
			printState13(state)
		} else if needsInput {
			needsInput = false
			if quit := doInput13(in); quit {
				done = true
			}
		} else {
			select {
			case _ = <-term:
				done = true
			default:
				done = false
			}
		}

		if done {
			break
		}
	}

	for {
		changed := updateState13(state, out)
		if !changed {
			break
		}
		printState13(state)
		fmt.Println("did not drain")
	}
}

func updateState13(state []int64, output chan int64) bool {
	var changed bool

	for {
		var done bool
		var x int64
		select {
		case x = <-output:
			changed = true
		default:
			done = true
		}

		if done {
			break
		}

		y := <-output
		tid := <-output

		if x == -1 && y == 0 {
			state[len(state)-1] = tid
		} else {
			state[idx13(x, y)] = tid
		}
	}

	return changed
}

func printState13(state []int64) {
	fmt.Println("Score:", state[len(state)-1])

	var x int64
	var y int64
	for y = 0; y < height13; y++ {
		for x = 0; x < width13; x++ {
			tile := ""
			switch state[idx13(x, y)] {
			case 0:
				tile = " "
			case 1:
				tile = "X"
			case 2:
				tile = "#"
			case 3:
				tile = "_"
			case 4:
				tile = "o"
			}
			fmt.Print(tile)
		}
		fmt.Println("")
	}
}

func doInput13(input chan int64) bool {
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}

	switch char {
	case 'q':
		return true
	case 'a':
		input <- -1
	case 'd':
		input <- 1
	default:
		input <- 0
	}

	return false
}

func idx13(x, y int64) int64 {
	return (y * width13) + x
}
