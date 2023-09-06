package nodes

import "github.com/smikelson75/parser/tokens"

type Node struct {
	kind     NodeType
	token    tokens.Token
	children []*Node
}

func NewNode(kind NodeType, token tokens.Token) *Node {
	return &Node{
		kind:  kind,
		token: token,
	}
}

func (n *Node) AddChild(child *Node) {
	n.children = append(n.children, child)
}

func (n Node) GetChild(index uint) *Node {
	if int(index) >= len(n.children) {
		return nil
	}

	return n.children[index]
}

func (n Node) NumberOfChildren() int {
	return len(n.children)
}

type NodeType string

const (
	Numeric    NodeType = "numeric"
	String     NodeType = "string"
	Newline    NodeType = "newline"
	Whitespace NodeType = "whitespace"
	Identifier NodeType = "identifier"
)
