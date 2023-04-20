package lexer

import (
	"strings"

	"github.com/hculpan/kmonkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	linePos      int
}

var keywords = map[string]token.TokenType{
	"fn":     token.FUNCTION,
	"let":    token.LET,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
}

func NewLexer(input []string) *Lexer {
	// Join the slice of strings into a single string with newline characters
	joinedInput := strings.Join(input, "\n")

	return NewLexerForString(joinedInput)
}

func NewLexerForString(input string) *Lexer {
	l := &Lexer{input: input, line: 1, linePos: 0}
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
	l.linePos++
	if l.ch == '\n' {
		l.line++
		l.linePos = 0
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) LookupIdent(ident string) token.TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return token.IDENT
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Line: l.line, Pos: l.linePos}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.line, l.linePos)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch), Line: l.line, Pos: l.linePos}
		} else {
			tok = newToken(token.NOT, l.ch, l.line, l.linePos)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GT_EQ, Literal: string(ch) + string(l.ch), Line: l.line, Pos: l.linePos}
		} else {
			tok = newToken(token.GT, l.ch, l.line, l.linePos)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LT_EQ, Literal: string(ch) + string(l.ch), Line: l.line, Pos: l.linePos}
		} else {
			tok = newToken(token.LT, l.ch, l.line, l.linePos)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line, l.linePos)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line, l.linePos)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line, l.linePos)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line, l.linePos)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.line, l.linePos)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.line, l.linePos)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, l.line, l.linePos)
	case '/':
		if l.peekChar() == '/' {
			l.readChar()
			tok.Literal = l.readSingleLineComment()
			tok.Type = token.SINGLE_LINE_COMMENT
			tok.Line = l.line
			tok.Pos = l.linePos - len(tok.Literal) + 1
		} else if l.peekChar() == '*' {
			l.readChar()
			tok.Literal, tok.Line, tok.Pos = l.readMultiLineComment()
			tok.Type = token.MULTI_LINE_COMMENT
		} else {
			tok = newToken(token.SLASH, l.ch, l.line, l.linePos)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line, l.linePos)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line, l.linePos)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.LookupIdent(tok.Literal)
			tok.Line = l.line
			tok.Pos = l.linePos - len(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Line = l.line
			tok.Pos = l.linePos - len(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line, l.linePos)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readSingleLineComment() string {
	position := l.readPosition
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) readMultiLineComment() (string, int, int) {
	startLine := l.line
	startPos := l.linePos - 1
	for !(l.ch == '*' && l.peekChar() == '/') && l.ch != 0 {
		if l.ch == '\n' {
			l.line++
			l.linePos = 0
		} else {
			l.linePos++
		}
		l.readChar()
	}
	l.readChar() // Consume the final '/'
	return l.input[startPos:l.readPosition], startLine, startPos
}

func newToken(tokenType token.TokenType, ch byte, line int, pos int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line, Pos: pos}
}

func (l *Lexer) readIdentifier() string {
	position := l.readPosition - 1
	for isIdentifierChar(l.ch) {
		l.readChar()
	}
	return l.input[position : l.readPosition-1]
}

func isIdentifierChar(ch byte) bool {
	return isLetter(ch) || isDigit(ch) || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
