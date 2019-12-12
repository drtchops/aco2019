package solutions

import "fmt"

// Day11 ...
func Day11(input string) {
	prog := parseProg(input)
	in := make(chan int64, 100)
	out := make(chan int64, 100)
	term := make(chan termSig, 10)

	go runIntcode(prog, 0, in, out, term)

	colors := make(map[point]bool)
	pos := point{0, 0}
	colors[pos] = true
	dir := "up"
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	done := false
	sentInput := false

	for {
		if !sentInput {
			color, ok := colors[pos]
			if !ok || !color {
				in <- 0
			} else {
				in <- 1
			}
			sentInput = true
		}

		var newColor int64
		select {
		case newColor = <-out:
		default:
			newColor = -1
		}

		if newColor != -1 {
			sentInput = false
			colors[pos] = newColor == 1

			turn := <-out
			if turn == 0 {
				if dir == "up" {
					dir = "left"
					pos.x--
				} else if dir == "down" {
					dir = "right"
					pos.x++
				} else if dir == "left" {
					dir = "down"
					pos.y--
				} else {
					dir = "up"
					pos.y++
				}
			} else {
				if dir == "up" {
					dir = "right"
					pos.x++
				} else if dir == "down" {
					dir = "left"
					pos.x--
				} else if dir == "left" {
					dir = "up"
					pos.y++
				} else {
					dir = "down"
					pos.y--
				}
			}

			if pos.x < minX {
				minX = pos.x
			}
			if pos.x > maxX {
				maxX = pos.x
			}
			if pos.y < minY {
				minY = pos.y
			}
			if pos.y > maxY {
				maxY = pos.y
			}
		}

		select {
		case _ = <-term:
			done = true
		default:
			done = false
		}

		if done {
			break
		}
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			tile := " "
			c, ok := colors[point{x, y}]
			if ok && c {
				tile = "X"
			}
			fmt.Print(tile)
		}
		fmt.Println("")
	}
}
