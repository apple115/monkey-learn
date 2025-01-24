package compiler

import (
	"go/constant"
	"monkey/code"
)

type Compiler struct{
	instructions code.Instructions
	constants []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants: []object.Object{},
	}
}

// Compile ...
func (c *Compiler) Compile(node ast.Node)error  {
	return nil
}

// Bytecode ...
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instruction: c.instructions,
		Constants: c.constants
	}
}

type Bytecode struct{
	Instructions code.Instructions
	Constants []object.Object
}
