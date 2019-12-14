package solutions

import "fmt"

// Day01 ...
func Day01(input string) {
	inputs01 := ParseInputInts(input, "\n")

	var totalFuel int64
	for _, mass := range inputs01 {
		lastFuel := fuel01(mass)

		for lastFuel > 0 {
			totalFuel += lastFuel
			lastFuel = fuel01(lastFuel)
		}
	}

	fmt.Println(totalFuel)
}

func fuel01(mass int64) int64 {
	fuel := mass/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}
