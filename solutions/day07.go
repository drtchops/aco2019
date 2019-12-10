package solutions

import (
	"fmt"
)

// Day07 ...
func Day07(input string) {
	var phaseSettings = []int64{5, 6, 7, 8, 9}
	var numProgs int64 = 5

	var best int64
	inputs07 := parseProg(input)

	for _, phases := range permutations(phaseSettings) {
		fmt.Println(phases)
		term := make(chan termSig, numProgs)
		inputs := make([]chan int64, numProgs)
		var n int64
		for n = 0; n < numProgs; n++ {
			inputs[n] = make(chan int64, 1000)
		}

		var i int64
		for i = 0; i < numProgs; i++ {
			prog := make([]int64, len(inputs07))
			copy(prog, inputs07)

			input := inputs[i]
			var output chan int64
			if i == numProgs-1 {
				output = inputs[0]
			} else {
				output = inputs[i+1]
			}
			input <- phases[i]
			if i == 0 {
				input <- 0
			}

			go runIntcode(prog, i, input, output, term)
		}

		for n = 0; n < numProgs; n++ {
			out := <-term
			if out.err != nil {
				fmt.Println(out.err)
				break
			}
			if out.id == numProgs-1 && out.output > best {
				best = out.output
			}
		}
	}

	fmt.Println(best)
}
