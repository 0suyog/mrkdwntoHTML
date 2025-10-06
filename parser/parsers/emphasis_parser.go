package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type EmphasisParser struct{}

func (ep EmphasisParser) Parse(t tokenizer.TokenList) nodes.INode {
	UndersocreVriant := []tokenizer.TokenType{tokenizer.UNDERSCORE, tokenizer.TEXT, tokenizer.UNDERSCORE}

	AsterikVariant := []tokenizer.TokenType{tokenizer.ASTERIK, tokenizer.TEXT, tokenizer.ASTERIK}

	if !t.PeekOr(UndersocreVriant, AsterikVariant) {
		return nodes.NullNode{}
	}

	return nodes.NewNode("EMPHASIS", t.Second().Value(), 3)
}
