package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
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
