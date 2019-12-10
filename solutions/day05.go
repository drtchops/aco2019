package solutions

import (
	"fmt"
)

// Day05 ...
func Day05(input string) {
	inputs05 := parseProg(input)
	progInput := make(chan int64, 5)
	progInput <- 5
	progOutput := make(chan int64, 5)
	progTerm := make(chan termSig, 5)
	runIntcode(inputs05, 0, progInput, progOutput, progTerm)
	term := <-progTerm
	fmt.Println(term.output)
}
