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
	//OpSub 是一个操作码，用于表示减法操作
	OpSub
	//OpMul 是一个操作码，用于表示乘法操作
	OpMul
	//OpDiv 是一个操作码，用于表示除法操作
	OpDiv
	//OpPop 将栈定的抛出
	OpPop
	//OpTrue
	OpTrue
	//OpFalse
	OpFalse
	//OpEqual ==
	OpEqual
	//OpNotEqual !=
	OpNotEqual
	//OpGreaterThan >  <
	OpGreaterThan

	// -
	OpMinus
	// !
	OpBang

	// 有条件跳转
	// OpJumpNotTruthy会让虚拟机在栈顶内容不为真或者为空的时候进行跳转。它的单个操作数是虚拟机应该跳转到的指令的偏移
	OpJumpNotTruthy
	// 直接跳转
	//虚拟机跳转到指令的偏移量处。
	OpJump
	OpNull

	OpGetGlobal
	OpSetGlobal

	OpArray

	OpHash

	OpIndex

	OpCall

	OpReturnValue

	//如果没有返回的OpReturnValue 就添加一个OpReturn
	OpReturn
)

// 定义：名字 操作符占用字符数
type Definition struct {
	Name          string
	OperandWidths []int //
}

var definitions = map[Opcode]*Definition{
	OpConstant:      {"OpConstant", []int{2}},
	OpAdd:           {"OpAdd", []int{}},
	OpPop:           {"OpPop", []int{}},
	OpSub:           {"OpSub", []int{}},
	OpMul:           {"OpMul", []int{}},
	OpDiv:           {"OpDiv", []int{}},
	OpTrue:          {"OpTrue", []int{}},
	OpFalse:         {"OpFalse", []int{}},
	OpEqual:         {"OpEqual", []int{}},
	OpNotEqual:      {"OpNotEqual", []int{}},
	OpGreaterThan:   {"OpGreaterThan", []int{}},
	OpMinus:         {"OpMinus", []int{}},
	OpBang:          {"OpBang", []int{}},
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpJump:          {"OpJump", []int{2}},
	OpNull:          {"OpNull", []int{}},
	OpGetGlobal:     {"OpGetGlobal", []int{2}},
	OpSetGlobal:     {"OpSetGlobal", []int{2}},
	OpArray:         {"OpArrey", []int{2}},
	OpHash:          {"OpHash", []int{2}},
	OpIndex:         {"OpIndex", []int{}},
	OpCall:          {"OpCall", []int{}},
	OpReturnValue:   {"OpReturnValue", []int{}},
	OpReturn:        {"OpReturn", []int{}},
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
// op 是一个操作码，用于表示不同的操作
// operands 是一个整数切片，用于表示操作数
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

// ReadUint16 是一个函数，用于将字节切片转换为 16 位无符号整数
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
