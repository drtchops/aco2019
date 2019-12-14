package solutions

import (
	"strconv"
	"strings"
)

// Point ...
type Point struct {
	x int
	y int
}

// Point3D ...
type Point3D struct {
	x int
	y int
	z int
}

// Point64 ...
type Point64 struct {
	x int64
	y int64
}

// Point3D64 ...
type Point3D64 struct {
	x int64
	y int64
	z int64
}

// Permutations ...
func Permutations(arr []int64) [][]int64 {
	var helper func([]int64, int64)
	res := [][]int64{}

	helper = func(arr []int64, n int64) {
		if n == 1 {
			tmp := make([]int64, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			var i int64
			for i = 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, int64(len(arr)))
	return res
}

// ParseInputInts ...
func ParseInputInts(input, sep string) []int64 {
	lines := strings.Split(input, sep)
	parsed := make([]int64, len(lines))
	for i, line := range lines {
		num, _ := strconv.ParseInt(line, 10, 64)
		parsed[i] = num
	}
	return parsed
}
