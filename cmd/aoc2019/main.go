package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/drtchops/aoc2019/solutions"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("wrong number of args")
	}

	d, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil || d < 1 || d > 25 {
		panic("bad number")
	}

	label := strconv.FormatInt(d, 10)
	if len(label) == 1 {
		label = "0" + label
	}

	inputBytes, err := ioutil.ReadFile(fmt.Sprintf("./inputs/day%s.txt", label))
	if err != nil {
		panic(err)
	}
	input := string(inputBytes)

	t := time.Now()

	switch d {
	case 1:
		solutions.Day01(input)
	case 2:
		solutions.Day02(input)
	case 3:
		solutions.Day03(input)
	case 4:
		solutions.Day04(input)
	case 5:
		solutions.Day05(input)
	case 6:
		solutions.Day06(input)
	case 7:
		solutions.Day07(input)
	case 8:
		solutions.Day08(input)
	case 9:
		solutions.Day09(input)
	case 10:
		solutions.Day10(input)
	case 11:
		solutions.Day11(input)
	case 12:
		solutions.Day12(input)
	case 13:
		solutions.Day13(input)
	case 14:
		solutions.Day14(input)
	case 15:
		solutions.Day15(input)
	case 16:
		solutions.Day16(input)
	case 17:
		solutions.Day17(input)
	case 18:
		solutions.Day18(input)
	case 19:
		solutions.Day19(input)
	case 20:
		solutions.Day20(input)
	case 21:
		solutions.Day21(input)
	case 22:
		solutions.Day22(input)
	case 23:
		solutions.Day23(input)
	case 24:
		solutions.Day24(input)
	case 25:
		solutions.Day25(input)
	}

	fmt.Printf("took %.2f seconds", time.Since(t).Seconds())
}
