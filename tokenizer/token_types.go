package tokenizer

type TokenType string

func newTokenType(tokenType string) TokenType {
	return TokenType(tokenType)
}

var UNDERSCORE = newTokenType("UNDERSCORE")
var ASTERIK = newTokenType("ASTERIK")
var TAB = newTokenType("TAB")
var NEWLINE = newTokenType("NEWLINE")
var TEXT = newTokenType("TEXT")
var EOF = newTokenType("EOF")
var NULLTOKEN = newTokenType("NULL")
