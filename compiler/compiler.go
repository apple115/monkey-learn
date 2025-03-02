package compiler

import (
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

type Compiler struct {
	//操作码
	instructions code.Instructions
	//常量池
	constants []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

// Compile ...
func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

// Bytecode ...
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
