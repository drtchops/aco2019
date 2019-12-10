package solutions

import (
	"fmt"
	"strconv"
)

// Day04 ...
func Day04(input string) {
	inputs04 := parseInputInts(input, "-")
	min := inputs04[0]
	max := inputs04[1]

	numPassed := 0
	for n := min; n <= max; n++ {
		if valid(n) {
			numPassed++
		}
	}

	fmt.Println(numPassed)
}

func valid(pass int64) bool {
	digits := 6
	last := -1
	streaks := make([]int, 0)
	curStreak := 1
	for p := 0; p < digits; p++ {
		d := digit(pass, p)
		if d < last {
			return false
		}
		if d == last {
			curStreak++
		} else {
			streaks = append(streaks, curStreak)
			curStreak = 1
		}
		last = d
	}
	streaks = append(streaks, curStreak)
	for _, s := range streaks {
		if s == 2 {
			return true
		}
	}
	return false
}

func digit(num int64, place int) int {
	ds := string(strconv.FormatInt(int64(num), 10)[place])
	d, _ := strconv.ParseInt(ds, 10, 64)
	return int(d)
}
