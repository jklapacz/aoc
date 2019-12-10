package computer

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
}
