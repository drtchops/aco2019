package solutions

import "fmt"

// Day01 ...
func Day01(input string) {
	inputs01 := parseInputInts(input, "\n")

	var totalFuel int64
	for _, mass := range inputs01 {
		lastFuel := fuel(mass)

		for lastFuel > 0 {
			totalFuel += lastFuel
			lastFuel = fuel(lastFuel)
		}
	}

	fmt.Println(totalFuel)
}

func fuel(mass int64) int64 {
	fuel := mass/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}
