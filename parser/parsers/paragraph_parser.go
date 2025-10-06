package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type ParagraphParser struct {
}

func (pp ParagraphParser) Parse(tokens tokenizer.TokenList) nodes.INode {
	matchedNode, _ := match_first(tokens, SentenceWithNewLineParser{}, SentenceWithEOFParser{})
	return matchedNode
}
