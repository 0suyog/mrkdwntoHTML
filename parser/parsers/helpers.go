package parsers

import (
	"fmt"
	"reflect"

	"github.com/0suyog/mrkdwntoHTML/parser/nodes"
	"github.com/0suyog/mrkdwntoHTML/tokenizer"
)

func match_first(tokens tokenizer.TokenList, parsers ...IParser) (nodes.INode, int) {
	parserNames := ""
	for _, parser := range parsers {
		parserNames += fmt.Sprintf("%s, ", reflect.TypeOf(parser).Name())
	}
	for _, parser := range parsers {
		if node := parser.Parse(tokens); !node.IsNull() {
			return node, node.Consumed()
		}
	}
	return nodes.NullNode{}, 0
}

func match_star(tokens tokenizer.TokenList, p IParser) ([]nodes.INode, int) {
	var matchedNodes = make([]nodes.INode, 0)
	consumed := 0

	for {
		node := p.Parse(*tokens.Offset(consumed))
		if node.IsNull() {
			return matchedNodes, consumed
		}
		matchedNodes = append(matchedNodes, node)
		consumed += node.Consumed()
	}

}

func match_plus(tokens tokenizer.TokenList, p IParser) ([]nodes.INode, int, error) {
	matchedNodes, consumed := match_star(tokens, p)
	if consumed == 0 {
		return matchedNodes, 0, fmt.Errorf("No tokens matched")
	}
	return matchedNodes, consumed, nil
}
