package lexer

import "monkey/token"

type Lexer struct {
	input             string
	currentPosition   int
	readAheadPosition int
	currentChar       byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(currentChar) + string(l.currentChar)}
		} else {
			tok = newToken(token.ASSIGN, l.currentChar)
		}
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '-':
		tok = newToken(token.MINUS, l.currentChar)
	case '!':
		if l.peekChar() == '=' {
			currentChar := l.currentChar
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(currentChar) + string(l.currentChar)}
		} else {
			tok = newToken(token.BANG, l.currentChar)
		}
	case '/':
		tok = newToken(token.SLASH, l.currentChar)
	case '*':
		tok = newToken(token.ASTERISK, l.currentChar)
	case '<':
		tok = newToken(token.LT, l.currentChar)
	case '>':
		tok = newToken(token.GT, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.currentChar)
	case ']':
		tok = newToken(token.RBRACKET, l.currentChar)
	case ':':
		tok = newToken(token.COLON, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)};
}

func (l *Lexer) readChar() {
	if l.readAheadPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.readAheadPosition]
	}
	l.currentPosition = l.readAheadPosition
	l.readAheadPosition += 1
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) peekChar() byte {
	if l.readAheadPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readAheadPosition]
	}
}

func (l *Lexer) readString() string {
	position := l.currentPosition + 1
	for {
		l.readChar()
		if l.currentChar == '"' || l.currentChar == 0 {
			break
		}
	}
	return l.input[position:l.currentPosition]
}
