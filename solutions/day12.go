package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type match12 struct {
	id    int
	start int64
	stop  int64
}

// Day12 ...
func Day12(input string) {
	posX, posY, posZ := parse12(input)
	matches := make(chan match12, 100)
	perms := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{0, 3},
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}

	go findLoop12(0, posX, perms, matches)
	go findLoop12(1, posY, perms, matches)
	go findLoop12(2, posZ, perms, matches)

	stops := make([]int64, 0)
	for {
		m := <-matches
		stops = append(stops, m.stop)
		if len(stops) == 3 {
			break
		}
	}

	fmt.Println(LCM(stops[0], stops[1], stops[2]))
}

func parse12(input string) ([]int64, []int64, []int64) {
	re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
	lines := strings.Split(input, "\n")
	posX := make([]int64, len(lines))
	posY := make([]int64, len(lines))
	posZ := make([]int64, len(lines))

	for i, line := range lines {
		pos := re.FindStringSubmatch(line)[1:]
		x, _ := strconv.ParseInt(pos[0], 10, 64)
		y, _ := strconv.ParseInt(pos[1], 10, 64)
		z, _ := strconv.ParseInt(pos[2], 10, 64)
		posX[i] = x
		posY[i] = y
		posZ[i] = z
	}

	return posX, posY, posZ
}

func findLoop12(id int, points []int64, perms [][]int, matches chan match12) {
	vels := make([]int64, len(points))

	history := make(map[string]int64)
	history[key12(points, vels)] = 0

	var step int64
	for step = 1; ; step++ {
		for _, pair := range perms {
			i1 := pair[0]
			i2 := pair[1]
			p1 := points[i1]
			p2 := points[i2]

			if p1 > p2 {
				vels[i1]--
				vels[i2]++
			} else if p1 < p2 {
				vels[i1]++
				vels[i2]--
			}
		}

		for i := range points {
			points[i] += vels[i]
		}

		s := key12(points, vels)
		if start, ok := history[s]; ok {
			matches <- match12{id, start, step}
			break
		} else {
			history[s] = step
		}
	}
}

func key12(pos []int64, vel []int64) string {
	return fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d", pos[0], pos[1], pos[2], pos[3], vel[0], vel[1], vel[2], vel[3])
}

func hasMatch(steps [][]int64) bool {
	for _, x := range steps[0] {
		for _, y := range steps[0] {
			if x != y {
				continue
			}

			for _, z := range steps[0] {
				if x == y && y == z {
					return true
				}
			}
		}
	}

	return false
}
