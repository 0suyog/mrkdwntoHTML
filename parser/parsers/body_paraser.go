package parsers

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

type BodyParser struct{}

type IParser interface {
	Parse(tokens tokenizer.TokenList) nodes.INode
}

func (bp BodyParser) Parse(tokens tokenizer.TokenList) nodes.INode {

	matchedNodes, consumed := match_star(tokens, ParagraphParser{})
	var paragraphNodes = make([]nodes.ParagraphNode, 0, len(matchedNodes))

	for _, node := range matchedNodes {
		if p, ok := node.(nodes.ParagraphNode); ok {
			paragraphNodes = append(paragraphNodes, p)
		}
	}

	return nodes.NewBodyNode(matchedNodes, consumed)
}
