package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type SentenceWithEOFParser struct{}

func (sp SentenceWithEOFParser) Parse(tokens tokenizer.TokenList) nodes.INode {
	sentenceParser := SentenceParser{}

	matchedNodes, consumed, err := match_plus(tokens, sentenceParser)

	if err != nil {
		return nodes.NullNode{}
	}

	if tokens.PeekAt(consumed, tokenizer.NEWLINE, tokenizer.EOF) {
		consumed += 1
		return nodes.NewParagraphNode(matchedNodes, consumed)
	}
	if tokens.PeekAt(consumed, tokenizer.EOF) {
		consumed += 1
		return nodes.NewParagraphNode(matchedNodes, consumed)
	}
	return nodes.NullNode{}

}
