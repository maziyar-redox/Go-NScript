package lexer

import (
	"github.com/maziyar-redox/Go-NScript/token"
)

type Lexer struct {
	input			string // source code input
	position		int // current position in input (points to current char)
	readPosition	int // current reading position to input (after current char)
	ch				byte // current char under examination
}

// ====================== //
// New funcion for initiating input and lexing analysis
// ====================== //

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// ====================== //
// This is just a little helper function that returns a type of Token for
// Clean code purposes
// ====================== //

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// ====================== //
// This helper function will help us to read peekChar and now what is coming after
// Current character
// ====================== //

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// ====================== //
// Helper function for skiping white spaces
// ====================== //

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// ====================== //
// each time the NextToken or New function is called, This method(readChar) will be called too
// And it will advence our position in characters, changing position and readPosition in input
// ====================== //

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// ====================== //
// Reading JUST identifiers with help of is letter
// ====================== //

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// ====================== //
// Reading number with help of a helper function
// ====================== //

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// ====================== //
// Main Next Token function
// ====================== //

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace()
	switch l.ch {
		case '=':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.ASSIGN, l.ch)
			}
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '!':
			if l.peekChar() == '=' {
				ch := l.ch
				l.readChar()
				tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
			} else {
				tok = newToken(token.EXCLAMATION, l.ch)
			}
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case ';':
			tok = newToken(token.SEMICOLON, l.ch)
		case '(':
			tok = newToken(token.LPAREN, l.ch)
		case ')':
			tok = newToken(token.RPAREN, l.ch)
		case ',':
			tok = newToken(token.COMMA, l.ch)
		case '+':
			tok = newToken(token.PLUS, l.ch)
		case '{':
			tok = newToken(token.LBRACE, l.ch)
		case '}':
			tok = newToken(token.RBRACE, l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}
	l.readChar()
	return tok
}