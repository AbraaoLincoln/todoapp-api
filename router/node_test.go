package router

import "testing"

func TestAddChild(t *testing.T) {
	root := NewNode()
	child := NewNode()
	child.value = "node1"

	root.addChild(child.value, &child)

	if root.children["node1"].value == "" {
		t.Error("Expecting to child to be add to children map")
	}
}

func TestHasNode(t *testing.T) {
	node := NewNode()
	node.value = "root"
	node.children["node1"] = &Node{value: "node1"}

	if !node.hasChild("node1") {
		t.Error("Expecting hasChildren to return true")
	}

	if node.hasChild("node2") {
		t.Error("Expecting hasChildren to return false")
	}
}

func TestGetChildren(t *testing.T) {
	root := NewNode()
	root.value = "root"
	child := NewNode()
	child.value = "child"
	root.children["child"] = &child

	childResult := root.getChild("child")

	if childResult.value != child.value {
		t.Error("Expecting to get the child with the same name, but didn't")
	}
}

func TestHasPathVariable(t *testing.T) {
	node := NewNode()

	if node.hasPathVariable() {
		t.Error("Expecting node not to have path variable set")
	}

	node.pathVariable = &Node{}

	if !node.hasPathVariable() {
		t.Error("Expecting node to have a path variable")
	}
}
