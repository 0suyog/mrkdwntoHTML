package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type TextParser struct{}

func (tp TextParser) Parse(tokens tokenizer.TokenList) nodes.INode {
	if tokens.Peek(tokenizer.TEXT) {
		return nodes.NewNode("TEXT", tokens.Tokens()[0].Value(), 1)
	}
	return nodes.NullNode{}
}
