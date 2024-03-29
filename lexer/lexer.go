package lexer

import "john/token"

type Lexer struct {
	input        string
	position     int  //current position in input (current char)
	readPosition int  //current reading position (after current char)
	ch           byte //current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch{
        case '=':
            tok = newToken(token.ASSIGN, l.ch)
        case '+':
            tok = newToken(token.PLUS, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '(':
            tok = newToken(token.LPARA, l.ch)
        case ')':
            tok = newToken(token.RPARA, l.ch)
        case '{':
            tok = newToken(token.LCURL, l.ch)
        case '}':
            tok = newToken(token.RCURL, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
    }

    l.readChar()
    return tok
}


func newToken(tokenType token.TokenType, ch byte)token.Token {
    return token.Token{Type:tokenType, Literal: string(ch)}
}
