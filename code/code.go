package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Instructions 定义了一个字符切片类型，用于表示一系列的字节码指令
type Instructions []byte

// Opcode 定义了操作码类型，本质是一个字符类型，用于表示不同的操作
type Opcode byte

const (
	OpConstant Opcode = iota
	//OpAdd 是一个操作码，用于表示加法操作
	OpAdd
)

// 定义：名字 操作符占用字符数
type Definition struct {
	Name          string
	OperandWidths []int //
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpAdd:      {"OpAdd", []int{}},
}

// Lookup ...
func Lookup(op byte) (*Definition, error) {
	//尝试definitions 映射 Definition
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

// Make 是一个函数，用于生成字节码指令
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1

	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			// 如果操作数宽度为 2 字节，使用大端字节序将操作数转换为 16 位无符号整数并存储到字节切片中
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	return instruction
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstructions(def, operands))

		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstructions(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}
	return operands, offset
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
