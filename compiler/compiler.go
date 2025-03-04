package compiler

import (
	"fmt"
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
// 递归遍历AST、找到*ast.IntegerLiterals、对其进行求值并将其转换为*object.Integers、将它们添加到常量字段、最后将OpConstant指令添加到内部的Instructions切片
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return nil
			}
		}
	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
	case *ast.InfixExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}
		err = c.Compile(node.Right)
		if err != nil {
			return err
		}
		switch node.Operator{
		case "+":
		c.emit(code.OpAdd)
		default:
		return fmt.Errorf("unknown operator %s",node.Operator)
		}
	case *ast.IntegerLiteral:
		//TOOD
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant,c.addConstant(integer))
	}

	return nil
}

// Bytecode ...
// 返回一个包含编译器内部指令和常量的*Bytecode结构体指针
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


//生成指令并将其添加到最终结果
func (c *Compiler) emit(op code.Opcode,operands ...int) int  {
	ins := code.Make(op,operands...)
	pos := c.addInstruction(ins)
	return pos
}

func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

// 添加
func (c *Compiler) addInstruction(ins []byte) int {
	c.instructions = append(c.instructions, ins...)
	posNewInstruction := len(c.instructions)
	return posNewInstruction
}
