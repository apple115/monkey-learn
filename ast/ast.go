package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	//返回与其关联的词法单元的字面量
	TokenLiteral() string
	//
	String() string
}

type Statement interface {
	Node
	//
	StatementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// TokenLiteral ...
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String ...
func (p Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// statementNode ...
func (ls *LetStatement) StatementNode() {
}

// TokenLiteral ...
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String ...
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

// expressionNode ...
func (i *Identifier) expressionNode() {
}

// TokenLiter ...
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String ...
func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// statementNode ...
func (rs *ReturnStatement) StatementNode() {}

// TokenLiteral ...
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String ...
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// statementNode ...
func (es *ExpressionStatement) StatementNode() {
}

// TokenLiteral ...
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String ...
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral ...
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// name ...
func (il *IntegerLiteral) expressionNode() {

}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// expressionNode ...
func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral ...
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String ...
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// expressionNode ...
func (ie *InfixExpression) expressionNode() {
}

// TokenLiteral ...
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String ...
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type BlockStatement struct {
	Token      token.Token // '{'词法单元
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type IfExpression struct {
	Token       token.Token // if
	Condition   Expression
	Consequence *BlockStatement //
	Alternative *BlockStatement // else
}

// expressionNode ...
func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
