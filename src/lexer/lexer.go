package lexer

import "devscript/src/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// return Lexer instance
func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

/*
reads next character in the input string
and increments the [position] & [readPosition]
*/
func (lexer *Lexer) readChar() {
	// check for EOF
	if lexer.readPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

/*
returns next token from the input string
*/
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.char {
	case '=':
		// check for "==" (equality operator)
		if lexer.peekChar() == '=' {
			ch := lexer.char
			lexer.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(lexer.char)}
		} else {
			tok = newToken(token.ASSIGN, lexer.char)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.char)
	case '(':
		tok = newToken(token.LPAREN, lexer.char)
	case ')':
		tok = newToken(token.RPAREN, lexer.char)
	case ',':
		tok = newToken(token.COMMA, lexer.char)
	case '+':
		tok = newToken(token.PLUS, lexer.char)
	case '-':
		tok = newToken(token.MINUS, lexer.char)
	case '*':
		tok = newToken(token.ASTERISK, lexer.char)
	case '/':
		// check for "//" (comment)
		if lexer.peekChar() == '/' {
			// read the second '/'
			lexer.readChar()

			// skip until end of line
			lexer.skipLine()

			tok = lexer.NextToken()
		} else {
			tok = newToken(token.SLASH, lexer.char)
		}
	case '!':
		// check for "!=" (NOT_EQ)
		if lexer.peekChar() == '=' {
			ch := lexer.char
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(lexer.char)}
		} else {
			tok = newToken(token.BANG, lexer.char)
		}
	case '<':
		tok = newToken(token.LT, lexer.char)
	case '>':
		tok = newToken(token.GT, lexer.char)
	case '{':
		tok = newToken(token.LBRACE, lexer.char)
	case '}':
		tok = newToken(token.RBRACE, lexer.char)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if validStartIdentifier(lexer.char) {
			tok.Literal = lexer.readIdentifier()
			// check if the identifier is a token
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.char) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.char)
		}
	}

	lexer.readChar()
	return tok
}

/*
newToken return new instance of type Token struct
{Type: tokenType, Literal: string}
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
function to return an identifier name
using maximal munch rule (longest common prefix)
*/
func (lexer *Lexer) readIdentifier() string {
	start := lexer.position

	// readChar until char is not a letter
	for validIdentifier(lexer.char) {
		lexer.readChar()
	}

	// end position = lexer.position
	return lexer.input[start:lexer.position]
}

/*
function to return number
using maximal munch rule (longest common prefix)
*/
func (lexer *Lexer) readNumber() string {
	start := lexer.position

	// readChar until char is not a number
	for isDigit(lexer.char) {
		lexer.readChar()
	}

	// end position = lexer.position
	return lexer.input[start:lexer.position]
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

/*
returns true only if the identifier starts with a letter or _ else false
*/
func validStartIdentifier(ch byte) bool {
	return isLetter(ch) || ch == '_'
}

func validIdentifier(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_' || ch == '-'
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// skips the current line
func (lexer *Lexer) skipLine() {
	for lexer.char != '\n' {
		lexer.readChar()
	}
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}
