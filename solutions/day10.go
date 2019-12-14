package solutions

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Day10 ...
func Day10(input string) {
	rocks := parse10(input)
	rock := Point{11, 11}
	angles := make(map[Point][]Point)
	keys := make([]Point, 0)

	for _, other := range rocks {
		if other == rock {
			continue
		}

		dx := other.x - rock.x
		dy := other.y - rock.y
		gcd := gcd10(intAbs10(dx), intAbs10(dy))
		dx /= gcd
		dy /= gcd

		a := Point{dx, dy}

		group, ok := angles[a]
		if !ok {
			group = make([]Point, 0)
			keys = append(keys, a)
		}
		group = append(group, other)
		sort.Slice(group, func(i, j int) bool {
			return magnitude10(group[i]) < magnitude10(group[j])
		})
		angles[a] = group
	}

	sort.Slice(keys, func(i, j int) bool {
		return angle10(keys[i]) > angle10(keys[j])
	})

	count := 0
	last := Point{}

	for count < 200 {
		for _, k := range keys {
			rocks := angles[k]
			if len(rocks) == 0 {
				continue
			}

			last = rocks[0]
			angles[k] = rocks[1:]
			count++
			if count >= 200 {
				break
			}
		}
	}

	fmt.Println(last)
}

func parse10(input string) []Point {
	lines := strings.Split(input, "\n")
	rocks := make([]Point, 0)

	for y, line := range lines {
		points := strings.Split(line, "")
		for x, p := range points {
			if p == "#" {
				rocks = append(rocks, Point{x, y})
			}
		}
	}

	return rocks
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd10(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func angle10(p Point) float64 {
	return math.Atan2(float64(p.x), float64(p.y))
}

func magnitude10(p Point) int {
	return intAbs10(p.x) + intAbs10(p.y)
}

func intAbs10(n int) int {
	return int(math.Abs(float64(n)))
}
