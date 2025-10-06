package tokenizer

import (
	"strings"

	"github.com/0suyog/mrkdwntoHTML/helpers"
)

func SingleLetterScanner(plain_markdown string) IToken {
	firstCharacter, err := helpers.AtPosition(plain_markdown, 0)
	if err != nil {
		panic("Something bad happened")
	}
	token, ok := SingleLetterTokens[firstCharacter]

	if !ok {
		return NewNullToken()
	}

	return NewToken(token, firstCharacter)
}

func TextScanner(plain_markdown string) IToken {
	value := make([]string, 0)
	for i := range len([]rune(plain_markdown)) {
		charAtPos, _ := helpers.AtPosition(plain_markdown, i)
		if !SingleLetterScanner(charAtPos).IsNull() {
			break
		}
		value = append(value, charAtPos)
	}
	return NewToken(TEXT, strings.Join(value, ""))
}
