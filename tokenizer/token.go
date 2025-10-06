package tokenizer

import (
	"fmt"
)

type Token struct {
	tokenType TokenType
	value     string
	isNull    bool
}

type NullToken struct{}

type IToken interface {
	IsNull() bool
	Present() bool
	TokenType() TokenType
	Value() string
	Length() int
	String() string
}

func NewToken(tokenType TokenType, value string) Token {
	var t Token
	if tokenType == "" || value == "" {
		t.isNull = true
		return t
	}
	t.tokenType = tokenType
	t.value = value
	return t
}

func EOFToken() Token {
	return NewToken(EOF, "eof")
}

func (t Token) Length() int {
	return len(t.value)
}

func (t Token) IsNull() bool {
	return t.isNull
}

func (t Token) Value() string {
	return t.value
}

func (t Token) TokenType() TokenType {
	return t.tokenType
}

func (t Token) Present() bool {
	return !t.isNull
}

func (t Token) String() string {
	return fmt.Sprintf("{\n\ttype: %s,\n\tvalue: %s\n}\n", t.tokenType, t.value)
}

func NewNullToken() NullToken {
	return NullToken{}
}

func (t NullToken) Length() int {
	return 0
}

func (t NullToken) IsNull() bool {
	return true
}

func (t NullToken) Value() string {
	return ""
}

func (t NullToken) TokenType() TokenType {
	return NULLTOKEN
}

func (t NullToken) Present() bool {
	return false
}

func (t NullToken) String() string {
	return string(NULLTOKEN)
}
