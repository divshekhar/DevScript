package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	DATATYPE = "DATATYPE"

	// Identifiers
	IDENT = "IDENT"

	// Literals
	INT    = "INT"
	STRING = "STRING"
	BOOL   = "BOOL"
	NULL   = "NULL"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,

	// Data Types
	"int":      DATATYPE,
	"string":   DATATYPE,
	"bool":     DATATYPE,
	"function": DATATYPE,
}

/*
returns Keyword Specific token if the identifier is a keyword
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
