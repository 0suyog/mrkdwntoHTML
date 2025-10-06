package nodes

import "fmt"

type ParagraphNode struct {
	sentences []INode
	consumed  int
}

func NewParagraphNode(sentences []INode, consumed int) ParagraphNode {
	return ParagraphNode{
		sentences: sentences,
		consumed:  consumed,
	}
}

func (pn ParagraphNode) IsNull() bool {
	return false
}

func (pn ParagraphNode) IsPresent() bool {
	return true
}

func (pn ParagraphNode) NodeType() string {
	return "PARAGRAPH"
}

func (pn ParagraphNode) Consumed() int {
	return pn.consumed
}

func (pn ParagraphNode) Value() string {

	joinedValue := ""

	for _, node := range pn.sentences {
		joinedValue += fmt.Sprintf("\t%s,\n", node.String())
	}

	return fmt.Sprintf("Nodes: [\n%s]", joinedValue)
}

func (pn ParagraphNode) String() string {
	joinedValue := ""

	for _, node := range pn.sentences {
		joinedValue += fmt.Sprintf("\t%s,\n", node.String())
	}

	return fmt.Sprintf("Type: Paragraph Node,\nNodes: [\n%s],\nConsumed: %d", joinedValue, pn.Consumed())
}
