package tokenizer

import (
	"fmt"
	"strings"

	"github.com/0suyog/mrkdwntoHTML/helpers"
)

func SingleLetterScanner(plain_markdown string) Token {
	firstCharacter, err := helpers.AtPosition(plain_markdown, 0)
	if err != nil {
		panic("Something bad happened")
	}
	token, ok := SingleLetterTokens[firstCharacter]

	if !ok {
		return NullToken()
	}

	return NewToken(token, firstCharacter)
}

func TextScanner(plain_markdown string) Token {
	value := make([]string, 0)
	fmt.Println(plain_markdown)
	for i := range len([]rune(plain_markdown)) {
		charAtPos, _ := helpers.AtPosition(plain_markdown, i)
		if !SingleLetterScanner(charAtPos).isNull {
			break
		}
		value = append(value, charAtPos)
	}
	return NewToken("TEXT", strings.Join(value, ""))
}
