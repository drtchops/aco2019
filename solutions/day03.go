package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

// Day03 ...
func Day03(input string) {
	inputs03 := strings.Split(input, "\n")
	wire1 := strings.Split(inputs03[0], ",")
	wire2 := strings.Split(inputs03[1], ",")

	points1 := make(map[point]int)
	p := point{0, 0}
	steps := 0
	for _, cmd := range wire1 {
		points := follow(cmd, p)
		for i, pp := range points {
			if _, ok := points1[pp]; !ok {
				points1[pp] = steps + i + 1
			}
		}
		p = points[len(points)-1]
		steps += len(points)
	}

	intersections := make(map[point]int)
	points2 := make(map[point]int)
	p = point{0, 0}
	steps = 0
	for _, cmd := range wire2 {
		points := follow(cmd, p)
		for i, pp := range points {
			d2 := steps + i + 1
			if _, ok := points2[pp]; !ok {
				points2[pp] = d2
			}
			if d1, ok := points1[pp]; ok {
				intersections[pp] = d1 + d2
			}
		}
		p = points[len(points)-1]
		steps += len(points)
	}

	best := 99999
	for _, d := range intersections {
		if d < best {
			best = d
		}
	}

	fmt.Println(best)
}

func follow(cmd string, p point) []point {
	d := string(cmd[0])
	n64, _ := strconv.ParseInt(cmd[1:], 10, 32)
	n := int(n64)
	points := make([]point, n)

	if d == "U" {
		for i := 1; i <= n; i++ {
			points[i-1] = point{p.x, p.y + i}
		}
	} else if d == "D" {
		for i := 1; i <= n; i++ {
			points[i-1] = point{p.x, p.y - i}
		}
	} else if d == "R" {
		for i := 1; i <= n; i++ {
			points[i-1] = point{p.x + i, p.y}
		}
	} else if d == "L" {
		for i := 1; i <= n; i++ {
			points[i-1] = point{p.x - i, p.y}
		}
	}

	return points
}
