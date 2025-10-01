package tokenizer

import "fmt"

type Token struct {
	tokenType string
	value     string
	isNull    bool
}

func NewToken(tokenType string, value string) Token {
	var t Token
	if tokenType == "" || value == "" {
		t.isNull = true
		return t
	}
	t.tokenType = tokenType
	t.value = value
	return t
}

func NullToken() Token {
	t := NewToken("", "")
	return t

}

func EOFToken() Token {
	return NewToken("EOF", "eof")
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

func (t Token) TokenType() string {
	return t.tokenType
}

func (t Token) Present() bool {
	return !t.isNull
}

func (t Token) String() string {
	return fmt.Sprintf("{\n\ttype: %s,\n\tvalue: %s\n}\n", t.tokenType, t.value)
}
