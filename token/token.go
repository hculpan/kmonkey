package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Pos     int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and keywords
	IDENT = "IDENT"
	INT   = "INT"

	// Keywords
	LET      = "LET"
	FUNCTION = "FN"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	NOT      = "!"

	// Comparison
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="
	GT_EQ  TokenType = ">="
	LT_EQ  TokenType = "<="
	GT     TokenType = ">"
	LT     TokenType = "<"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Comments
	SINGLE_LINE_COMMENT TokenType = "//"
	MULTI_LINE_COMMENT  TokenType = "/*"
)
