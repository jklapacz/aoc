package computer

import (
	"fmt"
	"log"
)

type Opcode = int

const (
	OpcodeAdd Opcode = iota + 1
	OpcodeMultiply
	OpcodeSave
	OpcodeOutput
	OpcodeJIT
	OpcodeJIF
	OpcodeLT
	OpcodeEq
	OpcodeUnknown Opcode = 98
	OpcodeErr     Opcode = 99
)

var OpLengths = map[Opcode]int{
	OpcodeAdd:      4,
	OpcodeMultiply: 4,
	OpcodeSave:     2,
	OpcodeOutput:   2,
	OpcodeJIT:      3,
	OpcodeJIF:      3,
	OpcodeLT:       4,
	OpcodeEq:       4,
}

func Decode(encodedOp int) Opcode {
	tens := nthdigit(encodedOp, 1)
	ones := nthdigit(encodedOp, 0)
	if tens != 0 && tens != 9 {
		return OpcodeUnknown
	}
	if tens == 9 && ones == 9 {
		return OpcodeErr
	}
	if ones < 1 || ones > 8 {
		return OpcodeUnknown
	}
	return ones
}

func nthdigit(x, n int) int {
	powersof10 := []int{1, 10, 100, 1000, 10000}
	return ((x / powersof10[n]) % 10)
}

type Operation struct {
	opcode          Opcode
	encoded         int
	output          int
	params          []int
	nextInstruction int
	exec            func(c *Computer)
}

type opParam struct {
}

type operable interface {
	exec()
}

func (c *Computer) inputValueForPosition(op Operation, pos, paramIdx int) int {
	val := nthdigit(op.encoded, pos)
	if val == 1 {
		c.Trace.Printf("In immediate mode\n")
		return op.params[paramIdx] // immediate mode
	} else if val == 2 {
		relativeOffset := c.relative + op.params[paramIdx]
		c.Trace.Printf("In relative mode\n")
		return c.Program.read(relativeOffset) // relative offset
	} else {
		c.Trace.Printf("In positional mode\n")
		c.Trace.Printf("\t getting value at %v\n", op.params[paramIdx])
		return c.Program.read(op.params[paramIdx]) // positional mode
	}
}

func (c *Computer) inputForParam(param string, op Operation) int {
	switch param {
	case "a":
		return c.inputValueForPosition(op, 2, 0) // maps to
	case "b":
		return c.inputValueForPosition(op, 3, 1)
	case "c":
		retVal := c.inputValueForPosition(op, 4, 2)
		c.Trace.Printf("\t debug program contents: %v\n", c.Program.data)
		c.Trace.Printf("param c value: %v!\n", retVal)
		return retVal
	default:
		return -1
	}
}

func opAdd(op *Operation) func(c *Computer) {
	exec := func(c *Computer) {
		inputs := getTwoInputs(op, c)
		result := inputs[0] + inputs[1]
		output := op.params[2]
		c.Trace.Printf("!! writing to: %v\n", output)
		c.Trace.Printf("!! op params: %v\n", op.params)
		c.Program.store(result, output)
	}
	return exec
}

func opMult(op *Operation) func(c *Computer) {
	exec := func(c *Computer) {
		inputs := getTwoInputs(op, c)
		result := inputs[0] * inputs[1]
		output := op.params[2]
		c.Program.store(result, output)
	}
	return exec
}

func opSave(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		c.Program.store(c.getInput(), op.params[0])
	}
}

func opJit(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		inputs := getTwoInputs(op, c)
		if inputs[0] != 0 {
			op.nextInstruction = inputs[1]
		}
	}
}

func opJif(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		inputs := getTwoInputs(op, c)
		if inputs[0] == 0 {
			op.nextInstruction = inputs[1]
		}
	}
}

func opLT(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		inputs := getTwoInputs(op, c)
		if inputs[0] < inputs[1] {
			c.Program.store(1, op.params[2])
		} else {
			c.Program.store(0, op.params[2])
		}
	}
}

func opEq(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		inputs := getTwoInputs(op, c)
		if inputs[0] == inputs[1] {
			c.Program.store(1, op.params[2])
		} else {
			c.Program.store(0, op.params[2])
		}
	}
}

func opOutput(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		c.output = c.inputForParam("a", *op)
		c.UserInputStreams.Write(c.output)
	}
}

func getTwoInputs(op *Operation, c *Computer) []int {
	var inputs []int
	inputs = append(inputs, c.inputForParam("a", *op))
	inputs = append(inputs, c.inputForParam("b", *op))
	return inputs
}

func (c *Computer) getCurrentOperation() *Operation {
	return ParseOperation(c)
}

func ParseOperation(c *Computer) *Operation {
	opcodes := c.Program.data
	address := *c.functionPointer
	c.Trace.Printf("\tcurrently at address: %v\n", address)
	instructionIdx := int(address)
	encodedOp := opcodes[instructionIdx]
	op := &Operation{opcode: Decode(encodedOp), encoded: encodedOp}
	switch op.opcode {
	case OpcodeAdd:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opAdd(op)
	case OpcodeMultiply:
		op.output = opcodes[instructionIdx+3]
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opMult(op)
	case OpcodeOutput:
		op.output = -1
		op.params = opcodes[instructionIdx+1 : instructionIdx+2]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opOutput(op)
	case OpcodeSave:
		op.params = opcodes[instructionIdx+1 : instructionIdx+2]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opSave(op)
	case OpcodeJIT:
		op.params = opcodes[instructionIdx+1 : instructionIdx+3]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opJit(op)
	case OpcodeJIF:
		op.params = opcodes[instructionIdx+1 : instructionIdx+3]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opJif(op)
	case OpcodeLT:
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opLT(op)
	case OpcodeEq:
		op.params = opcodes[instructionIdx+1 : instructionIdx+4]
		op.nextInstruction = instructionIdx + OpLengths[op.opcode]
		op.exec = opEq(op)
	case OpcodeErr:
		c.Trace.Printf("\t !! received halt code\n")
	default:
		log.Fatal("unknown opcode", op.opcode)
	}
	return op
}

func (c *Computer) performOperation(op *Operation) {
	op.exec(c)
	*c.functionPointer = memoryAddress(op.nextInstruction)
}

func (c *Computer) DumpMemory() string {
	return fmt.Sprint(c.Program.data)
}
