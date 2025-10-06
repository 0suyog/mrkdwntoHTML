package nodes

import "fmt"

type Node struct {
	nodeType string
	consumed int
	value    string
}

type NullNode struct {
}

type INode interface {
	IsNull() bool
	IsPresent() bool
	NodeType() string
	Value() string
	Consumed() int
	String() string
}

func NewNode(nodeType string, value string, consumed int) Node {
	node := Node{
		nodeType: nodeType,
		value:    value,
		consumed: consumed,
	}
	return node
}

func (n Node) IsNull() bool {
	return false
}

func (n Node) IsPresent() bool {
	return true
}

func (n Node) NodeType() string {
	return n.nodeType
}

func (n Node) Value() string {
	return n.value
}

func (n Node) Consumed() int {
	return n.consumed
}

func (n Node) String() string {
	return fmt.Sprintf("Node: %s, Value: %s, Consumed: %d", n.NodeType(), n.Value(), n.Consumed())
}

func (n NullNode) IsNull() bool {
	return true
}

func (n NullNode) IsPresent() bool {
	return false
}

func (n NullNode) NodeType() string {
	return ""
}

func (n NullNode) Value() string {
	return ""
}

func (n NullNode) Consumed() int {
	return 0
}

func (n NullNode) String() string {
	return "Node: Null Node"
}
