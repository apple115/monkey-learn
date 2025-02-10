package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

// Inspect ...
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type ...
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

// Inspect ...
func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}

// Type ...
func (i *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type Null struct{}

// Inspect ...
func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Error struct {
	Message string
}

// Type ...
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect ...
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type ...
func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

// Inspect ...
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

// Type ...
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect ...
func (s *String) Inspect() string {
	return s.Value
}
