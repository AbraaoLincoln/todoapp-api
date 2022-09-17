package router

import (
	"errors"
	"net/http"
	"strings"
)

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(path string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) {
	splitPath := strings.Split(cleanPath(path), "/")
	currentNode := t.Root

	for index, value := range splitPath {
		if isPathVariable(value) {
			updateNodePathVariable(value, handler, currentNode)
			currentNode = currentNode.pathVariable
			continue
		}

		if isEndpoint(index, len(splitPath)) && hasTheSameValueAsTheCurrentNode(value, currentNode.value) {
			updateNodeHandler(handler, currentNode)
			break
		}

		if currentNode.hasChild(value) {
			currentNode = currentNode.getChild(value)
			continue
		}

		newNode := NewNodeWith(value, handler)
		currentNode.addChild(value, newNode)
		currentNode = newNode
	}
}

func cleanPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	if strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}

	return path
}

func isPathVariable(value string) bool {
	return strings.HasPrefix(value, ":")
}

func updateNodePathVariable(value string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo), node *Node) {

	if node.pathVariable != nil {
		node.pathVariable.Handle = handler
		return
	}

	newNode := NewNodeWith(value, handler)
	node.pathVariable = newNode
}

func isEndpoint(currentIndex int, qtyOfPoints int) bool {
	return currentIndex == (qtyOfPoints - 1)
}

func updateNodeHandler(handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo), node *Node) {
	node.Handle = handler
}

func hasTheSameValueAsTheCurrentNode(value string, currentNodeValue string) bool {
	return value == currentNodeValue
}

func updateCurrentNode(currentNode *Node, newNode *Node) {
	currentNode = newNode
}

func addNewChildAndUpdateCurrentNode(currentNode *Node, childNode *Node) {
	currentNode.addChild(childNode.value, childNode)
	currentNode = childNode
}

func (t *Tree) Find(path string) (Node, map[string]string, error) {
	splitPath := strings.Split(cleanPath(path), "/")
	qtyNodesOnPath := len(splitPath)
	currentNode := t.Root
	pathVariables := make(map[string]string)

	for index, value := range splitPath {

		if isEndpoint(index, qtyNodesOnPath) {
			if currentNode.hasChild(value) {
				return currentNode.getChild(value).getCopy(), getPathVariables(pathVariables), nil
			}

			if currentNode.hasPathVariable() {
				pathVariables[currentNode.pathVariable.value] = value
				return currentNode.pathVariable.getCopy(), getPathVariables(pathVariables), nil
			}
		}

		if currentNode.hasChild(value) {
			currentNode = currentNode.getChild(value)
			continue
		}

		if currentNode.hasPathVariable() {
			pathVariables[currentNode.pathVariable.value] = value
			currentNode = currentNode.pathVariable
			continue
		}
	}

	return Node{}, nil, errors.New("Node not found for the path: " + path)
}

func getPathVariables(pathVariables map[string]string) map[string]string {
	if len(pathVariables) > 0 {
		return pathVariables
	}
	return nil
}
