package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type BoldParser struct{}

func (bp BoldParser) Parse(tokens tokenizer.TokenList) nodes.INode {
	UnderscoreVariant := []tokenizer.TokenType{tokenizer.UNDERSCORE, tokenizer.UNDERSCORE, tokenizer.TEXT, tokenizer.UNDERSCORE, tokenizer.UNDERSCORE}

	AsterikVariant := []tokenizer.TokenType{tokenizer.ASTERIK, tokenizer.ASTERIK, tokenizer.TEXT, tokenizer.ASTERIK, tokenizer.ASTERIK}
	if !tokens.PeekOr(UnderscoreVariant, AsterikVariant) {
		return nodes.NullNode{}
	}

	return nodes.NewNode("BOLD", tokens.Third().Value(), 5)
}
