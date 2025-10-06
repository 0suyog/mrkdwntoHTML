package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type SentenceWithNewLineParser struct{}

func (sp SentenceWithNewLineParser) Parse(tokens tokenizer.TokenList) nodes.INode {

	sentenceParser := SentenceParser{}
	matchedNodes, consumed, err := match_plus(tokens, sentenceParser)

	if err != nil {
		return nodes.NullNode{}
	}

	if !tokens.PeekAt(consumed, tokenizer.NEWLINE, tokenizer.NEWLINE) {
		return nodes.NullNode{}
	}

	consumed += 2

	return nodes.NewParagraphNode(matchedNodes, consumed)

}
