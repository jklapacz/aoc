package computer

import (
	"fmt"
	"log"
	"math"
)

// Opcode determines which operation gets performed
type Opcode = int

func (o Operation) ToString() string {
	details := fmt.Sprintf("(%v) %v", o.encoded, o.params)
	action := ""
	switch o.opcode {
	case OpcodeAdd:
		action = "ADD"
	case OpcodeMultiply:
		action = "MULT"
	case OpcodeSave:
		action = "SAVE"
	case OpcodeOutput:
		action = "OUT"
	case OpcodeJIT:
		action = "JIT"
	case OpcodeJIF:
		action = "JIF"
	case OpcodeLT:
		action = "LT"
	case OpcodeEq:
		action = "EQ"
	case OpcodeRel:
		action = "REL"
	case OpcodeUnknown:
		action = "???"
	case OpcodeErr:
		action = "ERR"
	}
	return fmt.Sprintf("%s\t%s", action, details)
}

const (
	// OpcodeAdd (paramA + paramB) => paramC
	OpcodeAdd Opcode = iota + 1
	// OpcodeMultiply (paramA * paramB) => paramC
	OpcodeMultiply
	// OpcodeSave (userInput) => paramA
	OpcodeSave
	// OpcodeOutput (paramA) => userOutput
	OpcodeOutput
	// OpcodeJIT (paramA != 0) => returnAddress = paramB
	OpcodeJIT
	// OpcodeJIF (paramA == 0) => returnAddress = paramB
	OpcodeJIF
	// OpcodeLT (paramA < paramB) => 1 -> paramC : 0 -> paramC
	OpcodeLT
	// OpcodeEq (paramA == paramB) => 1 -> paramC : 0 -> paramC
	OpcodeEq
	// OpcodeRel (paramA) => relativeOffset += paramA
	OpcodeRel
	// OpcodeUnknown unsupported operation (due to parsing usually)
	OpcodeUnknown Opcode = 98
	// OpcodeErr halts the program
	OpcodeErr Opcode = 99
)

// OpLengths is a map of opcode to the number of instructions it encompasses
var OpLengths = map[Opcode]int{
	OpcodeAdd:      4,
	OpcodeMultiply: 4,
	OpcodeSave:     2,
	OpcodeOutput:   2,
	OpcodeJIT:      3,
	OpcodeJIF:      3,
	OpcodeLT:       4,
	OpcodeEq:       4,
	OpcodeRel:      2,
	OpcodeUnknown:  0,
	OpcodeErr:      0,
}

// Decode takes an int and returns the opcode
func Decode(encodedOp int) Opcode {
	tens := nthdigit(encodedOp, 1)
	ones := nthdigit(encodedOp, 0)
	if tens != 0 && tens != 9 {
		return OpcodeUnknown
	}
	if tens == 9 && ones == 9 {
		return OpcodeErr
	}
	if ones < 1 || ones > 9 {
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
	opParams        []opParam
}

type paramMode = int

const (
	modePositional paramMode = iota
	modeImmediate
	modeRelative
)

type opParam struct {
	mode  paramMode
	value int
}

type operable interface {
	exec()
}

func (c *Computer) inputValueForPosition(op Operation, pos, paramIdx int, output bool) int {
	c.Playback.Printf("at instruction: %v", *c.functionPointer)
	startAddress := int(*c.functionPointer) + 1 // first parameter address
	memA, _ := c.ReadFromMemory(startAddress)
	memB, _ := c.ReadFromMemory(startAddress + 1)
	memC, _ := c.ReadFromMemory(startAddress + 2)
	c.Playback.Printf("memory layout: [%v,%v,%v]", memA, memB, memC)
	switch nthdigit(op.encoded, pos) {
	case modePositional:
		c.Playback.Printf("looking for param: %v at %v", paramIdx, startAddress+paramIdx)
		desiredAddress, _ := c.ReadFromMemory(startAddress + paramIdx)
		c.Playback.Printf("desired param address: %v", desiredAddress)
		if op.opcode == OpcodeSave || output {
			return desiredAddress
		}
		value, _ := c.ReadFromMemory(desiredAddress)
		c.Playback.Printf("found value: %v", value)
		return value
	case modeImmediate:
		value, _ := c.ReadFromMemory(startAddress + paramIdx)
		return value
	case modeRelative:
		relativeOffset := int(*c.relativeOffset)
		c.Playback.Printf("looking for offset delta: %v at %v", paramIdx, startAddress+paramIdx)
		delta, _ := c.ReadFromMemory(startAddress + paramIdx)
		c.Playback.Printf("delta: %v", delta)
		desiredAddress := relativeOffset + delta
		c.Playback.Printf("desired param address: %v", desiredAddress)
		if op.opcode == OpcodeSave || output {
			return desiredAddress
		}
		value, _ := c.ReadFromMemory(desiredAddress)
		c.Playback.Printf("found value: %v", value)
		return value
	}
	return math.MinInt64
}

func (c *Computer) inputForParam(param string, op Operation) int {
	switch param {
	case "a":
		return c.inputValueForPosition(op, 2, 0, false) // maps to
	case "b":
		return c.inputValueForPosition(op, 3, 1, false)
	case "c":
		retVal := c.inputValueForPosition(op, 4, 2, true)
		return retVal
	default:
		return -1
	}
}

func opAdd(op *Operation) func(c *Computer) {
	exec := func(c *Computer) {
		inputs := getTwoInputs(op, c)
		result := inputs[0] + inputs[1]
		output := c.inputForParam("c", *op)
		before, _ := c.ReadFromMemory(output)
		c.WriteToMemory(output, result)
		after, _ := c.ReadFromMemory(output)
		c.Playback.Printf("|--------------> address[%v]: %v -> %v", output, before, after)
	}
	return exec
}

func opMult(op *Operation) func(c *Computer) {
	exec := func(c *Computer) {
		inputs := getTwoInputs(op, c)
		result := inputs[0] * inputs[1]
		output := c.inputForParam("c", *op)
		before, _ := c.ReadFromMemory(output)
		c.WriteToMemory(output, result)
		after, _ := c.ReadFromMemory(output)
		c.Playback.Printf("|--------------> address[%v]: %v -> %v", output, before, after)
	}
	return exec
}

func opSave(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		output := c.inputForParam("a", *op)
		inputVal := c.getInput()
		before, _ := c.ReadFromMemory(output)
		c.Trace.Printf("\t in save operation. storing %v at index %v\n", inputVal, output)
		c.WriteToMemory(output, inputVal)
		after, _ := c.ReadFromMemory(output)
		c.Playback.Printf("|--------------> address[%v]: %v -> %v", output, before, after)
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
		output := c.inputForParam("c", *op)
		before, _ := c.ReadFromMemory(output)
		if inputs[0] < inputs[1] {
			c.WriteToMemory(output, 1)
			c.Playback.Printf("|--------------> %v < %v; address[%v]: %v -> 1", inputs[0], inputs[1], output, before)
		} else {
			c.WriteToMemory(output, 0)
			c.Playback.Printf("|--------------> %v >= %v; address[%v]: %v -> 0", inputs[0], inputs[1], output, before)
		}
	}
}

func opEq(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		inputs := getTwoInputs(op, c)
		output := c.inputForParam("c", *op)
		before, _ := c.ReadFromMemory(output)
		if inputs[0] == inputs[1] {
			c.WriteToMemory(output, 1)
			c.Playback.Printf("|--------------> %v == %v; address[%v]: %v -> 1", inputs[0], inputs[1], output, before)
		} else {
			c.WriteToMemory(output, 0)
			c.Playback.Printf("|--------------> %v != %v; address[%v]: %v -> 0", inputs[0], inputs[1], output, before)
		}
	}
}

func opOutput(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		output := c.inputForParam("a", *op)
		c.Trace.Printf("\t\t======= current operation: %v\n", *op)
		c.Trace.Printf("\t\t======= relative offset: %v\n", *c.relativeOffset)
		c.UserInputStreams.Write(output)
		c.Playback.Printf("|--------------> output is %v", output)
	}
}

func opRel(op *Operation) func(c *Computer) {
	return func(c *Computer) {
		offset := c.inputForParam("a", *op)
		before := *c.relativeOffset
		*c.relativeOffset += memoryAddress(offset)
		after := *c.relativeOffset
		c.Playback.Printf("|--------------> relative offset: %v -> %v", before, after)
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
	address := *c.functionPointer
	instructionIdx := int(address)
	encodedOp, _ := c.ReadFromMemory(instructionIdx)
	op := &Operation{
		opcode:  Decode(encodedOp),
		encoded: encodedOp,
	}

	if op.opcode != OpcodeErr && op.opcode != OpcodeUnknown {
		op.opParams = make([]opParam, OpLengths[op.opcode]-1)
		//		for paramAddress := instructionIdx + 1; address < instructionIdx + OpLengths[op.opcode] ; address++ {
		//			op.params = append(op.params, c.ReadFromMemory(paramAddress
		//		}
	}
	op.nextInstruction = instructionIdx + OpLengths[op.opcode]
	switch op.opcode {
	case OpcodeAdd:
		op.exec = opAdd(op)
	case OpcodeMultiply:
		op.exec = opMult(op)
	case OpcodeOutput:
		op.exec = opOutput(op)
	case OpcodeSave:
		op.exec = opSave(op)
	case OpcodeJIT:
		op.exec = opJit(op)
	case OpcodeJIF:
		op.exec = opJif(op)
	case OpcodeLT:
		op.exec = opLT(op)
	case OpcodeEq:
		op.exec = opEq(op)
	case OpcodeRel:
		op.exec = opRel(op)
	case OpcodeErr:
		c.Trace.Printf("\t !! received halt code\n")
	default:
		log.Fatal("unknown opcode", op.opcode)
	}
	return op
}

func (c *Computer) performOperation(op *Operation) {
	c.Playback.Printf("%s", op.ToString())
	c.Trace.Printf("\n\n __PERFORMING OPERATION: %s__ (address: %v)\n", op.ToString(), *c.functionPointer)
	c.Trace.Printf("OPCODE: \t %s (%v) \t PARAMS: %v\n", op.ToString(), op.encoded, op.params)
	//c.Trace.Printf("\tcurrently at address: %v\n", *c.functionPointer)
	op.exec(c)
	c.Playback.Printf("next instruction at: %v", op.nextInstruction)
	*c.functionPointer = memoryAddress(op.nextInstruction)
	c.Trace.Printf("__JUMPING TO: %v__ \n\n", *c.functionPointer)
}

func (c *Computer) DumpMemory() string {
	return fmt.Sprint(c.Program.data)
}
