package computer

import (
	"fmt"
)

type registerContents int

type register struct {
	address  memoryAddress
	contents *registerContents
}

type computerMemory struct {
	start, end memoryAddress
	registers  []register
}

func (m *computerMemory) readFromAddress(a memoryAddress) (*register, error) {
	if a < m.start || a > m.end {
		return nil, fmt.Errorf("invalid write address: %v", a)
	}
	return &m.registers[a], nil
}

func (m *computerMemory) writeToAddress(a memoryAddress, value registerContents) error {
	if a < m.start || a > m.end {
		return fmt.Errorf("invalid write address: %v", a)
	}
	if m.registers[a].contents == nil {
		m.registers[a].contents = &value
	} else {
		*m.registers[a].contents = value
	}
	return nil
}

func initializeMemory(inputValues []int) *computerMemory {
	registers := make([]register, 10000)
	//registers := make([]register, int(math.Max(float64(len(inputValues)*5), float64(500))))
	registerIdx := 0
	for _, value := range inputValues {
		r := registerContents(value)
		registers[memoryAddress(registerIdx)] = register{memoryAddress(registerIdx), &r}
		registerIdx++
	}
	for ; registerIdx < len(registers); registerIdx++ {
		registers[memoryAddress(registerIdx)] = register{memoryAddress(registerIdx), nil}
	}
	return &computerMemory{
		memoryAddress(0),
		memoryAddress(len(registers) - 1),
		registers,
	}
}

func (m *computerMemory) Dump() string {
	s := fmt.Sprintf("\n==== DUMPING MEMORY CONTENTS ====\n")
	for _, register := range m.registers {
		if register.contents != nil {
			s += fmt.Sprintf("[%v]: %v \t(%v)\n", register.address, *register.contents, register.contents)
		} else {
			s += fmt.Sprintf("[%v]: %v \n", register.address, register.contents)
		}
	}
	//fmt.Println(s)
	return s
}

// ReadFromMemory returns the register at a given memory address
func (c *Computer) ReadFromMemory(address int) (int, error) {
	register, err := c.Memory.readFromAddress(memoryAddress(address))
	if err != nil {
		return 0, err
	}
	if register.contents != nil {
		return int(*register.contents), nil
	}
	return 0, fmt.Errorf("address %v has empty contents", address)
}

func (c *Computer) WriteToMemory(address, value int) error {
	if err := c.Memory.writeToAddress(memoryAddress(address), registerContents(value)); err != nil {
		return err
	}
	return nil
}
