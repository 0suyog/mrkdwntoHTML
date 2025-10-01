package tokenizer

import (
	"fmt"
)

type TokenList struct {
	Tokens []Token
}

func NewTokenList(tokens []Token) *TokenList {
	tokenList := TokenList{
		Tokens: tokens,
	}
	return &tokenList
}

func (tl *TokenList) Peek(tokenTypes ...string) bool {
	if length := len(tl.Tokens); length == 0 {
		return false
	}
	for i, tokenType := range tokenTypes {
		if indexedToken := tl.Tokens[i]; indexedToken.TokenType() != tokenType {
			return false
		}
	}
	return true
}

func (tl *TokenList) PeekOr(tokenTypes ...string) bool {
	for _, tokenType := range tokenTypes {
		if tl.Peek(tokenType) {
			return true
		}
	}
	return false
}

func (tl *TokenList) Offset(index int) *TokenList {
	if index == 0 {
		return tl
	}
	offsetedTokens := tl.Tokens[index:]
	return NewTokenList(offsetedTokens)
}

func (tl *TokenList) PeekAt(index int, tokenTypes ...string) bool {
	if index == 0 {
		return tl.Peek(tokenTypes...)
	}
	offsetedTokenList := tl.Offset(index)

	return offsetedTokenList.Peek(tokenTypes...)
}

// Grab modifies the Tokens inside TokenList
// it left shifts the list by provided amount
// and returns the shifted elements
func (tl *TokenList) Grab(amount int) (returnTokens []Token, e error) {
	if amount == 0 {
		return
	}
	if amount == len(tl.Tokens) {
		return returnTokens, fmt.Errorf("Invalid Amount, There are only %d elements and you are trying to select %d elements", len(tl.Tokens), amount)
	}
	shiftedList := tl.Tokens[amount:]
	returnTokens = tl.Tokens[0:amount]
	tl.Tokens = shiftedList
	return returnTokens, nil
}

func (tl *TokenList) Second() (returnToken Token, e error) {
	if len(tl.Tokens) < 2 {
		return returnToken, fmt.Errorf("List has less than 2 elements ")
	}
	return tl.Tokens[1], nil
}

func (tl *TokenList) Third() (returnToken Token, e error) {
	if len(tl.Tokens) < 2 {
		return returnToken, fmt.Errorf("List has less than 3 elements ")
	}
	return tl.Tokens[2], nil
}

func (tl *TokenList) String() string {
	var joinedString string

	for _, token := range tl.Tokens {
		joinedString += fmt.Sprintf("%s", token.String())
	}

	return fmt.Sprintf("[\n%s\n]", joinedString)

}
