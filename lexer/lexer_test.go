package lexer_test

import (
	"john/lexer"
	"john/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+{}(),;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LCURL, "{"},
		{token.RCURL, "}"},
		{token.LPARA, "("},
		{token.RPARA, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected %q. Got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected %q. Got %q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
