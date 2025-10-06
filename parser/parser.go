package parser

import (
	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/parser/parsers"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

func Parse(tokens tokenizer.TokenList) nodes.INode {
	bp := parsers.BodyParser{}
	node := bp.Parse(tokens)
	return node
}
