package puzzle1

import "fmt"

const (
	Done  = 99
	Add   = 1
	Mul   = 2
	Store = 3
	Put   = 4
)

type ParameterInstruction struct {
	opcode int
	param1 int
	param2 int
	param3 int
}

func parseParameterInstruction(instruction int) ParameterInstruction {
	opcode := instruction % 100
	firstParamMode := instruction / 100 % 10
	secondParamMode := instruction / 1000 % 10
	thirdParamMode := instruction / 10000 % 10

	return ParameterInstruction{opcode, firstParamMode, secondParamMode, thirdParamMode}
}

func executeParameter(pi ParameterInstruction, opcodes []int, i int, input int) int {
	var x, y, steps int
	if pi.param1 == 1 {
		x = opcodes[i+1]
	} else if pi.param1 == 0 {
		x = opcodes[opcodes[i+1]]
	}

	if pi.param2 == 1 {
		y = opcodes[i+2]
	} else if pi.param2 == 0 && pi.opcode != 4 {
		y = opcodes[opcodes[i+2]]
	}

	switch pi.opcode {
	case Add:
		opcodes[opcodes[i+3]] = x + y
		steps = 3
	case Mul:
		opcodes[opcodes[i+3]] = x * y
		steps = 3
	case Store:
		opcodes[opcodes[i+1]] = input
		steps = 1
	case Put:
		fmt.Println(x)
		steps = 1
	}

	return steps
}

func ParseOpcodes(opcodes []int, input int) []int {
	for i := 0; i < len(opcodes); i++ {
		if opcodes[i] == Done {
			break
		}
		pi := parseParameterInstruction(opcodes[i])
		i += executeParameter(pi, opcodes, i, input)

	}

	return opcodes
}
