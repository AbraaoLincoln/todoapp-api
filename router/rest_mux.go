package router

import (
	"errors"
	"log"
	"net/http"
)

type RestMux struct {
	getTree    *Tree
	postTree   *Tree
	putTree    *Tree
	patchTree  *Tree
	deleteTree *Tree
}

func (restMux *RestMux) Get(path string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) {
	restMux.getTree.Insert(path, handler)
}

func (restMux *RestMux) Post(path string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) {
	restMux.postTree.Insert(path, handler)
}

func (restMux *RestMux) Put(path string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) {
	restMux.putTree.Insert(path, handler)
}

func (restMux *RestMux) Patch(path string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) {
	restMux.patchTree.Insert(path, handler)
}

func (restMux *RestMux) Delete(path string, handler func(w http.ResponseWriter, r *http.Request, extraInfo *ExtraInfo)) {
	restMux.deleteTree.Insert(path, handler)
}

func (restMux *RestMux) Find(httpMethod string, path string) (Node, map[string]string, error) {
	switch httpMethod {
	case http.MethodGet:
		return restMux.getTree.Find(path)
	case http.MethodPost:
		return restMux.postTree.Find(path)
	case http.MethodPut:
		return restMux.putTree.Find(path)
	case http.MethodPatch:
		return restMux.patchTree.Find(path)
	case http.MethodDelete:
		return restMux.deleteTree.Find(path)
	default:
		return Node{}, nil, errors.New("Node not found for path: " + path)
	}
}

func (restMux *RestMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Looking for handle")
	log.Println("http method:", r.Method)
	log.Println("path:", r.URL.Path)

	extraInfo := ExtraInfo{}
	node, pathVariables, err := restMux.Find(r.Method, r.URL.Path)

	if err != nil {
		log.Println("Handler not found for path", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if pathVariables != nil {
		log.Println("Some path variables was found")
		extraInfo = ExtraInfo{
			PathVariables: pathVariables,
		}
	}

	node.Handle(w, r, &extraInfo)
}

func NewRestMux() *RestMux {
	return &RestMux{
		getTree:    &Tree{Root: NewNodeWith("/", nil)},
		postTree:   &Tree{Root: NewNodeWith("/", nil)},
		putTree:    &Tree{Root: NewNodeWith("/", nil)},
		patchTree:  &Tree{Root: NewNodeWith("/", nil)},
		deleteTree: &Tree{Root: NewNodeWith("/", nil)},
	}
}
