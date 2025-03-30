package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"monkey/code"
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
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	QUOTE_OBJ        = "QUOTE"
	MACRO_OBJ        = "MACRO"

	COMPILED_FUNCTION_OBJ = "COMPILED_FUNCTION_OBJ"

	CLOSURE_OBJ = "CLOSURE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value     int64
	HashValue *HashKey
}

// Inspect ...
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type ...
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// HashKey ...
func (i *Integer) HashKey() HashKey {
	if i.HashValue == nil {
		i.HashValue = &HashKey{Type: i.Type(), Value: uint64(i.Value)}
	}
	return *i.HashValue
}

type Boolean struct {
	Value     bool
	HashValue *HashKey
}

// Inspect ...
func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}

// Type ...
func (i *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// HashKey ...
func (b *Boolean) HashKey() HashKey {
	if b.HashValue == nil {
		var value uint64
		if b.Value {
			value = 1
		} else {
			value = 0
		}
		b.HashValue = &HashKey{Type: b.Type(), Value: value}
	}
	return *b.HashValue
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
	Value     string
	HashValue *HashKey
}

// Type ...
func (s *String) Type() ObjectType {
	return STRING_OBJ
}

// Inspect ...
func (s *String) Inspect() string {
	return s.Value
}

// HashKey ...
func (s *String) HashKey() HashKey {
	if s.HashValue == nil {
		h := fnv.New64a()
		h.Write([]byte(s.Value))
		s.HashValue = &HashKey{Type: s.Type(), Value: h.Sum64()}
	}
	return *s.HashValue
}

// 函数需要接受零个或多个object.Object作为参数并能返回一个object.Object
type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

// Type ...
func (ao *Array) Type() ObjectType {
	return ARRAY_OBJ
}
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type ...
func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

// Inspect ...
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ","))
	out.WriteString("}")

	return out.String()
}

type Quote struct {
	Node ast.Node
}

// Type ...
func (q *Quote) Type() ObjectType {
	return QUOTE_OBJ
}

// Inspect ...
func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}

type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (m *Macro) Type() ObjectType { return MACRO_OBJ }
func (m *Macro) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type CompiledFunction struct {
	Instructions  code.Instructions
	NumLocals     int
	NumParameters int
}

func (cf *CompiledFunction) Type() ObjectType {
	return COMPILED_FUNCTION_OBJ
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}

type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) Type() ObjectType {
	return CLOSURE_OBJ
}

func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
