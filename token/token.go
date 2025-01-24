package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //illegal 非法的
	EOF = "EOF"

	IDENT = "IDENT" // add,foobar,x,y
	INT="INT" //132131

	//运算符
	ASSIGN ="="
	PLUS="+"
	MINUS = '-'
	BANG = "!"
	ASTERISK="*"
	SLASH = "/"
	LT = "<"
	GT = ">"


	//分隔符
	COMMA = ","
	SEMICOLON=";"

	LPAREN="("
	RPAREN=")"
	LBRACE="{"
	RBRACE="}"

	//关键字
	FUNCTION="FUNCTION"
	LET = "LET"
)


var keywords = map[string]TokenType{
	"fn":FUNCTION,
	"let":LET,
}

// LookupIdent ...
func LookupIdent(ident string) TokenType {
	if tok,ok := keywords[ident];ok{
		return tok
	}
	return IDENT
}
