package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type SentenceParser struct{}

func (sp SentenceParser) Parse(tokens tokenizer.TokenList) nodes.INode {

	parserOrder := []IParser{EmphasisParser{}, BoldParser{}, TextParser{}}
	node, _ := match_first(tokens, parserOrder...)
	return node
}
