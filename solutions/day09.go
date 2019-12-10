package solutions

import "fmt"

// Day09 ...
func Day09(input string) {
	prog := parseProg(input)
	in := make(chan int64, 5)
	in <- 2
	out := make(chan int64, 500)
	term := make(chan termSig, 5)
	go runIntcode(prog, 0, in, out, term)
	t := <-term
	fmt.Println(t.output)
}
