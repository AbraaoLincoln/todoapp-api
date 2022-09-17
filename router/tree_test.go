package router

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsPathVariable(t *testing.T) {
	pathVariable1 := ":id"
	pathVariable2 := "id"

	if !isPathVariable(pathVariable1) {
		t.Errorf("Expecting %v to be a path variable", pathVariable1)
	}

	if isPathVariable(pathVariable2) {
		t.Errorf("Expecting %v not to be a path variable", pathVariable2)
	}
}

func TestInsert(t *testing.T) {
	path := "/users/123/posts"
	handle := func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler1") }
	root := NewNode()
	tree := Tree{Root: &root}

	tree.Insert(path, handle)

	if root.children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children["users"].children["123"] == nil {
		t.Error("Expecting /users/123 not to be nil")
	}

	if root.children["users"].children["123"].children["posts"] == nil {
		t.Error("Expecting /users/123/posts not to be nil")
	}

	path = "/users/profile"
	handle = func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler2") }

	tree.Insert(path, handle)

	if root.children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children["users"].children["profile"] == nil {
		t.Error("Expecting /users/profile not to be nil")
	}
}

func TestInsertWithPathVariable(t *testing.T) {
	path := "/users/:id/posts"
	handle := func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler1") }
	root := NewNode()
	tree := Tree{Root: &root}

	tree.Insert(path, handle)

	path = "/users/new/posts"
	handle = func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler2") }

	tree.Insert(path, handle)

	if root.children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children["users"].pathVariable == nil {
		t.Error("Expecting /users/:id path variable not to be nil")
	}

	if root.children["users"].pathVariable.children["posts"] == nil {
		t.Error("Expecting /users/:id/posts not to be nil")
	}

	if root.children["users"].children["new"] == nil {
		t.Error("Expecting /users/new/posts not to be nil")
	}

	if root.children["users"].children["new"].children["posts"] == nil {
		t.Error("Expecting /users/new/posts not to be nil")
	}
}

func TestWith2PathVariable(t *testing.T) {
	path := "/users/:id/:region"
	root := NewNode()
	root.value = "/"
	tree := Tree{Root: &root}

	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) {})

	if root.children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children["users"].pathVariable == nil {
		t.Error("Expecting :id path variable not to be nil")
	}

	if root.children["users"].pathVariable.pathVariable == nil {
		t.Error("Expecting :region path variable not to be nil")
	}

	path = "/users/:id/test/:region"

	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) {})

	if root.children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children["users"].pathVariable == nil {
		t.Error("Expecting :id path variable not to be nil")
	}

	if root.children["users"].pathVariable.children["test"] == nil {
		t.Error("Expecting /test path variable not to be nil")
	}

	if root.children["users"].pathVariable.children["test"].pathVariable == nil {
		t.Error("Expecting :region path variable not to be nil")
	}
}

func TestWith3PathVariable(t *testing.T) {
	path := "/users/:id/:region/:country"
	root := NewNode()
	root.value = "/"
	tree := Tree{Root: &root}

	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) {})

	if root.children["users"] == nil {
		t.Error("Expecting /users not to be nil")
	}

	if root.children["users"].pathVariable == nil {
		t.Error("Expecting :id path variable not to be nil")
	}

	if root.children["users"].pathVariable.pathVariable == nil {
		t.Error("Expecting :region path variable not to be nil")
	}

	if root.children["users"].pathVariable.pathVariable.pathVariable == nil {
		t.Error("Expecting :country path variable not to be nil")
	}
}

func TestFindNodeWithoutPathVariable(t *testing.T) {
	path := "/users/region/country"
	root := NewNodeWith("/", nil)
	tree := Tree{Root: root}
	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler 1") })

	node, pathVariables, err := tree.Find(path)

	if err != nil {
		t.Error("Expecting error to be nil")
	}

	if node.value != "country" {
		t.Error("Expecting node value to be equal to country, but got ", node.value)
	}

	if len(pathVariables) != 0 {
		t.Error("Expecting 0 path variables, but got ", len(pathVariables))
	}

}

func TestFindNodeWithPathVariable(t *testing.T) {
	path := "/users/:id/test/:region/:country"
	root := NewNodeWith("/", nil)
	tree := Tree{Root: root}
	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler 1") })

	node, pathVariables, err := tree.Find("/users/123/test/sa/br")

	if err != nil {
		t.Error("Expecting error to be nil")
	}

	if node.value != ":country" {
		t.Error("Expecting node value to be equal to country, but got ", node.value)
	}

	if len(pathVariables) != 3 {
		t.Error("Expecting 3 path variables, but got ", len(pathVariables))
	}

	if pathVariables[":id"] == "" {
		t.Error("Expecting the :id variable to be on the result map")
	}

	if pathVariables[":id"] != "123" {
		t.Error("Expecting value of :id variable to be 123 but got, ", pathVariables[":id"])
	}

	if pathVariables[":region"] == "" {
		t.Error("Expecting the :region variable to be on the result map")
	}

	if pathVariables[":region"] != "sa" {
		t.Error("Expecting value of :region variable to be sa but got, ", pathVariables[":region"])
	}

	if pathVariables[":country"] == "" {
		t.Error("Expecting the :country variable to be on the result map")
	}

	if pathVariables[":country"] != "br" {
		t.Error("Expecting value of :country variable to be br but got, ", pathVariables[":country"])
	}
}

func TestFindNodeWith2PathVariable(t *testing.T) {
	path := "/users/:id/:region"
	root := NewNodeWith("/", nil)
	tree := Tree{Root: root}
	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler 1") })

	node, pathVariables, err := tree.Find("/users/123/sa")

	if err != nil {
		t.Error("Expecting error to be nil")
	}

	if node.value != ":region" {
		t.Error("Expecting node value to be equal to region, but got ", node.value)
	}

	if len(pathVariables) != 2 {
		t.Error("Expecting 2 path variables, but got ", len(pathVariables))
	}

	if pathVariables[":id"] == "" {
		t.Error("Expecting the :id variable to be on the result map")
	}

	if pathVariables[":id"] != "123" {
		t.Error("Expecting value of :id variable to be 123 but got, ", pathVariables[":id"])
	}

	if pathVariables[":region"] == "" {
		t.Error("Expecting the :region variable to be on the result map")
	}

	if pathVariables[":region"] != "sa" {
		t.Error("Expecting value of :region variable to be sa but got, ", pathVariables[":region"])
	}
}

func TestFindNodeWithOnlyPathVariable(t *testing.T) {
	path := "/:id/:region"
	root := NewNodeWith("/", nil)
	tree := Tree{Root: root}
	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler 1") })

	node, pathVariables, err := tree.Find("/123/sa")

	if err != nil {
		t.Error("Expecting error to be nil")
	}

	if node.value != ":region" {
		t.Error("Expecting node value to be equal to region, but got ", node.value)
	}

	if len(pathVariables) != 2 {
		t.Error("Expecting 2 path variables, but got ", len(pathVariables))
	}

	if pathVariables[":id"] == "" {
		t.Error("Expecting the :id variable to be on the result map")
	}

	if pathVariables[":id"] != "123" {
		t.Error("Expecting value of :id variable to be 123 but got, ", pathVariables[":id"])
	}

	if pathVariables[":region"] == "" {
		t.Error("Expecting the :region variable to be on the result map")
	}

	if pathVariables[":region"] != "sa" {
		t.Error("Expecting value of :region variable to be sa but got, ", pathVariables[":region"])
	}
}

func TestFindNodeWith1PathVariable1Fixed(t *testing.T) {
	path := "/:id/region"
	root := NewNodeWith("/", nil)
	tree := Tree{Root: root}
	tree.Insert(path, func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo) { fmt.Println("handler 1") })

	node, pathVariables, err := tree.Find("/123/region")

	if err != nil {
		t.Error("Expecting error to be nil")
	}

	if node.value != "region" {
		t.Error("Expecting node value to be equal to region, but got ", node.value)
	}

	if len(pathVariables) != 1 {
		t.Error("Expecting 1 path variables, but got ", len(pathVariables))
	}

	if pathVariables[":id"] == "" {
		t.Error("Expecting the :id variable to be on the result map")
	}

	if pathVariables[":id"] != "123" {
		t.Error("Expecting value of :id variable to be 123 but got, ", pathVariables[":id"])
	}
}
