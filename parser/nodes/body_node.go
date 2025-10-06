package nodes

import "fmt"

type BodyNode struct {
	paragraphs []INode
	consumed   int
}

func NewBodyNode(paragraphs []INode, consumed int) BodyNode {
	return BodyNode{paragraphs, consumed}
}

func (bn BodyNode) IsNull() bool {
	return false
}

func (bn BodyNode) IsPresent() bool {
	return true
}

func (bn BodyNode) NodeType() string {
	return "PARAGRAPH"
}

func (bn BodyNode) Consumed() int {
	return bn.consumed
}

func (bn BodyNode) Value() string {

	joinedValue := ""

	for _, node := range bn.paragraphs {
		joinedValue += fmt.Sprintf("\t%s,\n", node.String())
	}

	return fmt.Sprintf("Nodes: [\n%s]", joinedValue)
}

func (bn BodyNode) String() string {

	joinedValue := ""

	for _, node := range bn.paragraphs {
		joinedValue += fmt.Sprintf("\t%s,\n", node.String())
	}
	return fmt.Sprintf("Type: Body Node,\nNodes: [\n%s],Consumed: %d", joinedValue, bn.Consumed())
}
