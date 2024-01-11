package token

type TokenType string;

type Token struct{
    Type TokenType;
    Literal string;
}


const (
    EOF = "EOF"
    ILLEGAL = "ILLEGAL"

    IDENT = "IDENT"
    INT = "INT"

    ASSIGN = "="
    PLUS = "+"

    COMMA = ","
    SEMICOLON = ";"

    LPARA = "("
    RPARA = ")"
    LCURL = "{"
    RCURL = "}"

    FUNCTION = "FUNCTION"
    VAR = "VAR"
)
