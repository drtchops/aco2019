package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

type resource14 struct {
	name   string
	amount int64
}

type reaction14 struct {
	reagents []resource14
	result   resource14
}

// Day14 ...
func Day14(input string) {
	reactions := parse14(input)
	resourceTotals := make(map[string]int64)
	resourceUsed := make(map[string]int64)
	resourceProducers := make(map[string]reaction14)
	for _, r := range reactions {
		resourceTotals[r.result.name] = 0
		resourceUsed[r.result.name] = 0
		resourceProducers[r.result.name] = r
	}

	requirements := []resource14{
		resource14{"FUEL", 1},
	}

	for {
		var done bool
		newReqs := make([]resource14, 0)
		for _, req := range requirements {
			if req.name == "ORE" {
				resourceTotals["ORE"] += req.amount
				resourceUsed["ORE"] += req.amount

				if resourceTotals["ORE"] >= 1000000000000 {
					done = true
					break
				}

				continue
			}

			if done {
				break
			}

			total := resourceTotals[req.name]
			used := resourceUsed[req.name] + req.amount
			if used < total {
				resourceUsed[req.name] = used
				continue
			}

			reaction := resourceProducers[req.name]
			for total < used {
				total += reaction.result.amount
				newReqs = append(newReqs, reaction.reagents...)
			}
			resourceTotals[req.name] = total
			resourceUsed[req.name] = used
		}

		if done {
			break
		}

		if len(newReqs) == 0 {
			newReqs = append(newReqs, resource14{"FUEL", 1})
		}

		requirements = newReqs
	}

	fmt.Println(resourceTotals["FUEL"] - 1)
}

func parse14(input string) []reaction14 {
	lines := strings.Split(input, "\n")
	reactions := make([]reaction14, len(lines))

	for i, line := range lines {
		formulaParts := strings.Split(line, " => ")
		reactionParts := strings.Split(formulaParts[0], ", ")
		reagents := make([]resource14, len(reactionParts))
		for j, r := range reactionParts {
			reagents[j] = parseResource14(r)
		}
		reactions[i] = reaction14{reagents, parseResource14(formulaParts[1])}
	}

	return reactions
}

func parseResource14(input string) resource14 {
	parts := strings.Split(input, " ")
	a, _ := strconv.ParseInt(parts[0], 10, 64)
	return resource14{parts[1], a}
}
