package solutions

import (
	"regexp"
	"strconv"
	"strings"
)

type point3d struct {
	x int
	y int
	z int
}

type moon struct {
	pos point3d
	vel point3d
}

type match struct {
	id   int
	step int64
}

// Day12 ...
func Day12(input string) {
	poses := parse12(input)
	moons := make([]moon, len(poses))
	for i, pos := range poses {
		moons[i] = moon{pos: pos}
	}

	perms := [][]moon{
		[]moon{
			moons[0],
			moons[1],
		},
		[]moon{
			moons[0],
			moons[2],
		},
		[]moon{
			moons[0],
			moons[3],
		},
		[]moon{
			moons[1],
			moons[2],
		},
		[]moon{
			moons[1],
			moons[3],
		},
		[]moon{
			moons[2],
			moons[3],
		},
	}

	matches := make(chan match, 10)

	go loop12(0, moons, perms, matches)
	go loop12(1, moons, perms, matches)
	go loop12(2, moons, perms, matches)
}

func parse12(input string) []point3d {
	re := regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
	lines := strings.Split(input, "\n")
	moons := make([]point3d, len(lines))

	for i, line := range lines {
		pos := re.FindStringSubmatch(line)[1:]
		x, _ := strconv.ParseInt(pos[0], 10, 64)
		y, _ := strconv.ParseInt(pos[1], 10, 64)
		z, _ := strconv.ParseInt(pos[2], 10, 64)
		moons[i] = point3d{int(x), int(y), int(z)}
	}

	return moons
}

func loop12(id int, moons []moon, perms [][]moon, matches chan match) {
	history := make(map[int]bool)
	history[sum12(moons, id)] = true

	var step int64
	for step = 0; ; step++ {
		for _, pair := range perms {
			r1 := pair[0]
			r2 := pair[1]

			switch id {
			case 0:
				if r1.pos.x > r2.pos.x {
					r1.vel.x--
					r2.vel.x++
				} else if r1.pos.x < r2.pos.x {
					r1.vel.x++
					r2.vel.x--
				}
			case 1:
				if r1.pos.y > r2.pos.y {
					r1.vel.y--
					r2.vel.y++
				} else if r1.pos.y < r2.pos.y {
					r1.vel.y++
					r2.vel.y--
				}
			case 2:
				if r1.pos.z > r2.pos.z {
					r1.vel.z--
					r2.vel.z++
				} else if r1.pos.z < r2.pos.z {
					r1.vel.z++
					r2.vel.z--
				}
			}
		}

		for _, m := range moons {
			switch id {
			case 0:
				m.pos.x += m.vel.x
			case 1:
				m.pos.y += m.vel.y
			case 2:
				m.pos.z += m.vel.z
			}
		}

		s := sum12(moons, id)
		if _, ok := history[s]; ok {
			matches <- match{id, step + 1}
		} else {
			history[s] = true
		}
	}
}

func sum12(moons []moon, id int) int {
	total := 0
	for _, m := range moons {
		switch id {
		case 0:
			total += m.pos.x
		case 1:
			total += m.pos.y
		case 2:
			total += m.pos.z
		}
	}

	return total
}
