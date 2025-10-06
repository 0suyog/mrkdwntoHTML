package tokenizer

import (
	"fmt"
)

type TokenList struct {
	tokens []IToken
}

func NewTokenList(tokens []IToken) *TokenList {
	tokenList := TokenList{
		tokens: tokens,
	}
	return &tokenList
}

func (tl *TokenList) Tokens() []IToken {
	return tl.tokens
}

func (tl *TokenList) Peek(tokenTypes ...TokenType) bool {
	if length := len(tl.Tokens()); length == 0 {
		return false
	}
	for i, tokenType := range tokenTypes {
		if indexedToken := tl.Tokens()[i]; indexedToken.TokenType() != tokenType {
			return false
		}
	}
	return true
}

func (tl *TokenList) PeekOr(tokenTypes ...[]TokenType) bool {
	for _, tokenType := range tokenTypes {
		if tl.Peek(tokenType...) {
			return true
		}
	}
	return false
}

func (tl *TokenList) Offset(index int) *TokenList {
	if index == 0 {
		return tl
	}
	offsetedTokens := tl.Tokens()[index:]
	return NewTokenList(offsetedTokens)
}

func (tl *TokenList) PeekAt(index int, tokenTypes ...TokenType) bool {
	if index == 0 {
		return tl.Peek(tokenTypes...)
	}
	offsetedTokenList := tl.Offset(index)

	return offsetedTokenList.Peek(tokenTypes...)
}

// Grab modifies the Tokens inside TokenList
// it left shifts the list by provided amount
// and returns the shifted elements
func (tl *TokenList) Grab(amount int) (returnTokens []IToken, e error) {
	if amount == 0 {
		return
	}
	if amount == len(tl.Tokens()) {
		return returnTokens, fmt.Errorf("Invalid Amount, There are only %d elements and you are trying to select %d elements", len(tl.Tokens()), amount)
	}
	shiftedList := tl.Tokens()[amount:]
	returnTokens = tl.Tokens()[0:amount]
	tl.tokens = shiftedList
	return returnTokens, nil
}

func (tl *TokenList) First() IToken {
	return tl.Tokens()[0]
}

func (tl *TokenList) Second() (returnToken IToken) {
	return tl.Tokens()[1]
}

func (tl *TokenList) Third() (returnToken IToken) {
	return tl.Tokens()[2]
}

func (tl *TokenList) String() string {
	var joinedString string

	for _, token := range tl.Tokens() {
		joinedString += fmt.Sprintf("%s", token.String())
	}

	return fmt.Sprintf("[\n%s\n]", joinedString)

}
