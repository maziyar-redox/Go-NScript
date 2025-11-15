package lexer

import (
	"testing"

	"github.com/maziyar-redox/Go-NScript/token"
)

type TestTokens struct {
	expectedType			token.TokenType
	expectedLiteral			string
}

func TestNextToken(t *testing.T) {
	input := ``
	tests := []TestTokens{}
	l := New(input)
	for i, TokenTypes := range tests {
		tok := l.NextToken()
		if tok.Type != TokenTypes.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, TokenTypes.expectedType, tok.Type)
		}
		if tok.Literal != TokenTypes.expectedLiteral {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%q, got=%q", i, TokenTypes.expectedLiteral, tok.Literal)
		}
	}
}