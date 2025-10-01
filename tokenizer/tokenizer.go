package tokenizer

import (
	"fmt"
	"log"

	"github.com/0suyog/mrkdwntoHTML/helpers"
	"golang.org/x/tools/go/callgraph/rta"
)

var TokenScanners = []func(string) Token{
	SingleLetterScanner,
	TextScanner,
}

func scan_one_token(plain_markdown string) (Token, error) {
	for _, scanner := range TokenScanners {
		token := scanner(plain_markdown)
		if !token.IsNull() {
			return token, nil
		}
	}
	return NullToken(), fmt.Errorf("No scanner matched given input: %s", plain_markdown)
}

func tokens_as_array(plain_markdown string) []Token {
	tokens := make([]Token, 0)

	if plain_markdown == "" {
		return []Token{EOFToken()}
	}

	token, err := scan_one_token(plain_markdown)
	if err != nil {
		panic(err)
	}
	tokens = append(tokens, token)
	tokens = append(tokens)

	for i := range len([]rune(plain_markdown)) {
		stringAtPos, err := helpers.AtPosition(plain_markdown, i)
		if err != nil {
			log.Fatal(err)
		}
		token, err := scan_one_token(stringAtPos)
		if err != nil {
			return tokens
		}
		tokens = append(tokens, token)
	}
	tokens = append(tokens, EOFToken())
	return tokens
}

func Tokenize(plain_markdown string) *TokenList {
	tokens := tokens_as_array(plain_markdown)
	return NewTokenList(tokens)
}
