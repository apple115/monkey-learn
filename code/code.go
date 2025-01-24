package code

import (
	"encoding/binary"
	"fmt"
)

// Instructions 定义了一个字符切片类型，用于表示一系列的字节码指令
type Instructions []byte

// Opcode 定义了操作码类型，本质是一个字符类型，用于表示不同的操作
type Opcode byte

const (
	OpConstant Opcode = iota
)

// 定义：名字 操作符占用字符数
type Definition struct {
	Name          string
	OperandWidths []int //
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
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

// Make ...
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
