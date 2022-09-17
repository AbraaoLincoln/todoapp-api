package router

import "net/http"

type Node struct {
	value        string
	children     map[string]*Node
	Handle       func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)
	pathVariable *Node
}

func NewNode() Node {
	return Node{children: make(map[string]*Node)}
}

func NewNodeWith(value string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) *Node {
	return &Node{
		value:    value,
		children: make(map[string]*Node),
		Handle:   handler,
	}
}

func (n *Node) addChild(name string, newChild *Node) {
	n.children[name] = newChild
}

func (n *Node) hasChild(childrenName string) bool {
	return n.children[childrenName] != nil
}

func (n *Node) getChild(childrenName string) *Node {
	return n.children[childrenName]
}

func (n *Node) hasPathVariable() bool {
	return n.pathVariable != nil
}

func (n *Node) getCopy() Node {
	return Node{
		value:        n.value,
		children:     n.children,
		Handle:       n.Handle,
		pathVariable: n.pathVariable,
	}
}
