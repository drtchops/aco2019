package solutions

import (
	"fmt"
	"strconv"
)

type termSig struct {
	id     int64
	output int64
	err    error
}

var paramCounts = map[int64]int64{
	1:  3,
	2:  3,
	3:  1,
	4:  1,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
	99: 0,
}
var writeParams = map[int64]int64{
	1: 2,
	2: 2,
	3: 0,
	7: 2,
	8: 2,
}

func runIntcode(prog []int64, id int64, input, output chan int64, term chan termSig) {
	var ptr int64
	var lastOutput int64
	var err error

	for {
		jump := false
		opt, pms := parseOptcode(prog[ptr])
		count, ok := paramCounts[opt]
		if !ok {
			err = fmt.Errorf("unknown opt, ptr=%d line=%v opt=%d pms=%v", ptr, prog[ptr:ptr+count+1], opt, pms)
			break
		}
		if int64(len(prog)) < ptr+count {
			err = fmt.Errorf("not enough values, ptr=%d line=%v opt=%d pms=%v", ptr, prog[ptr:ptr+count+1], opt, pms)
			break
		}

		if opt == 99 {
			break
		}

		params := make([]int64, count)
		var i int64
		for i = 0; i < count; i++ {
			var pm int64
			if i < int64(len(pms)) {
				pm = pms[i]
			}

			writeParam, ok := writeParams[opt]
			val := prog[ptr+i+1]
			if pm == 0 && !(ok && writeParam == i) {
				val = prog[val]
			}
			params[i] = val
		}

		// fmt.Printf("ptr=%d line=%v opt=%d params=%v pms=%v\n", ptr, prog[ptr:ptr+count+1], opt, params, pms)

		if opt == 1 {
			writePtr := params[2]
			val := params[0] + params[1]
			if writePtr >= int64(len(prog)) {
				err = fmt.Errorf("address %d out of range, ptr=%d line=%v opt=%d params=%v pms=%v", writePtr, ptr, prog[ptr:ptr+count+1], opt, params, pms)
				break
			}
			prog[writePtr] = val
		} else if opt == 2 {
			writePtr := params[2]
			val := params[0] * params[1]
			if writePtr >= int64(len(prog)) {
				err = fmt.Errorf("address %d out of range, ptr=%d line=%v opt=%d params=%v pms=%v", writePtr, ptr, prog[ptr:ptr+count+1], opt, params, pms)
				break
			}
			prog[writePtr] = val
		} else if opt == 3 {
			writePtr := params[0]
			if writePtr >= int64(len(prog)) {
				err = fmt.Errorf("address %d out of range, ptr=%d line=%v opt=%d params=%v pms=%v", writePtr, ptr, prog[ptr:ptr+count+1], opt, params, pms)
				break
			}
			fmt.Printf("%d input\n", id)
			in := <-input
			prog[writePtr] = in
		} else if opt == 4 {
			lastOutput = params[0]
			fmt.Printf("%d output %d\n", id, lastOutput)
			output <- lastOutput
		} else if opt == 5 {
			if params[0] != 0 {
				ptr = params[1]
				jump = true
			}
		} else if opt == 6 {
			if params[0] == 0 {
				ptr = params[1]
				jump = true
			}
		} else if opt == 7 {
			var val int64
			if params[0] < params[1] {
				val = 1
			}
			writePtr := params[2]
			if writePtr >= int64(len(prog)) {
				err = fmt.Errorf("address %d out of range, ptr=%d line=%v opt=%d params=%v pms=%v", writePtr, ptr, prog[ptr:ptr+count+1], opt, params, pms)
				break
			}
			prog[writePtr] = val
		} else if opt == 8 {
			var val int64
			if params[0] == params[1] {
				val = 1
			}
			writePtr := params[2]
			if writePtr >= int64(len(prog)) {
				err = fmt.Errorf("address %d out of range, ptr=%d line=%v opt=%d params=%v pms=%v", writePtr, ptr, prog[ptr:ptr+count+1], opt, params, pms)
				break
			}
			prog[writePtr] = val
		}

		if !jump {
			ptr += count + 1
		}
	}

	term <- termSig{id: id, output: lastOutput, err: err}
}

func parseOptcode(inst int64) (int64, []int64) {
	is := strconv.FormatInt(inst, 10)

	opts := ""
	if len(is) == 1 {
		opts = is
	} else {
		opts = is[len(is)-2:]
	}
	opt, _ := strconv.ParseInt(opts, 10, 64)

	pms := make([]int64, 0)
	for i := len(is) - 3; i >= 0; i-- {
		p, _ := strconv.ParseInt(string(is[i]), 10, 64)
		pms = append(pms, p)
	}

	return opt, pms
}

func parseProg(input string) []int64 {
	return parseInputInts(input, ",")
}
