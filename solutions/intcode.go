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

type optcode int64
type paramMode int64

const (
	optcodeAdd       optcode = 1
	optcodeMul       optcode = 2
	optcodeInput     optcode = 3
	optcodeOutput    optcode = 4
	optcodeJumpTrue  optcode = 5
	optcodeJumpFalse optcode = 6
	optcodeLessThan  optcode = 7
	optcodeEqual     optcode = 8
	optcodeRelative  optcode = 9
	optcodeTerm      optcode = 99

	paramPosition  paramMode = 0
	paramImmediate paramMode = 1
	paramRelative  paramMode = 2
)

var paramCounts = map[optcode]int64{
	optcodeAdd:       3,
	optcodeMul:       3,
	optcodeInput:     1,
	optcodeOutput:    1,
	optcodeJumpTrue:  2,
	optcodeJumpFalse: 2,
	optcodeLessThan:  3,
	optcodeEqual:     3,
	optcodeRelative:  1,
	optcodeTerm:      0,
}
var writeParams = map[optcode]int64{
	optcodeAdd:      2,
	optcodeMul:      2,
	optcodeInput:    0,
	optcodeLessThan: 2,
	optcodeEqual:    2,
}

func runIntcode(prog []int64, id int64, input, output chan int64, term chan termSig) {
	var ptr int64
	var relAddr int64
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

		if opt == optcodeTerm {
			// fmt.Printf("ptr=%d line=%v opt=%d\n", ptr, prog[ptr:ptr+count+1], opt)
			break
		}

		params := make([]int64, count)
		var i int64
		for i = 0; i < count; i++ {
			var pm paramMode
			if i < int64(len(pms)) {
				pm = pms[i]
			}

			writeIdx, ok := writeParams[opt]
			var val int64
			param := prog[ptr+i+1]

			if ok && i == writeIdx {
				if pm == paramRelative {
					val = relAddr + param
				} else {
					val = param
				}
			} else {
				if pm == paramPosition {
					val = prog[param]
				} else if pm == paramImmediate {
					val = param
				} else {
					val = prog[relAddr+param]
				}
			}
			params[i] = val
		}

		// fmt.Printf("ptr=%d line=%v opt=%d params=%v pms=%v\n", ptr, prog[ptr:ptr+count+1], opt, params, pms)

		if opt == optcodeAdd {
			writePtr := params[2]
			val := params[0] + params[1]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == optcodeMul {
			writePtr := params[2]
			val := params[0] * params[1]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == optcodeInput {
			writePtr := params[0]
			// fmt.Printf("%d input\n", id)
			val := <-input
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == optcodeOutput {
			lastOutput = params[0]
			// fmt.Printf("%d output %d\n", id, lastOutput)
			output <- lastOutput
		} else if opt == optcodeJumpTrue {
			if params[0] != 0 {
				ptr = params[1]
				jump = true
			}
		} else if opt == optcodeJumpFalse {
			if params[0] == 0 {
				ptr = params[1]
				jump = true
			}
		} else if opt == optcodeLessThan {
			var val int64
			if params[0] < params[1] {
				val = 1
			}
			writePtr := params[2]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == optcodeEqual {
			var val int64
			if params[0] == params[1] {
				val = 1
			}
			writePtr := params[2]
			if err := writeVal(prog, writePtr, val); err != nil {
				break
			}
		} else if opt == optcodeRelative {
			relAddr += params[0]
		}

		if !jump {
			ptr += count + 1
		}
	}

	term <- termSig{id: id, output: lastOutput, err: err}
}

func writeVal(prog []int64, ptr, val int64) error {
	if ptr < 0 || ptr >= int64(len(prog)) {
		return fmt.Errorf("address %d out of range", ptr)
	}
	prog[ptr] = val
	return nil
}

func parseOptcode(inst int64) (optcode, []paramMode) {
	is := strconv.FormatInt(inst, 10)

	opts := ""
	if len(is) == 1 {
		opts = is
	} else {
		opts = is[len(is)-2:]
	}
	opt, _ := strconv.ParseInt(opts, 10, 64)

	pms := make([]paramMode, 0)
	for i := len(is) - 3; i >= 0; i-- {
		p, _ := strconv.ParseInt(string(is[i]), 10, 64)
		pms = append(pms, paramMode(p))
	}

	return optcode(opt), pms
}

func parseProg(input string) []int64 {
	prog := ParseInputInts(input, ",")
	prog = append(prog, make([]int64, len(prog)*10)...)
	return prog
}
