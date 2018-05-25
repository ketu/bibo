package routing

import (
	"net/http"
	"fmt"
)

type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func (r *Route) Match(req *http.Request) bool {
	if req.Method != r.Method{
		return false
	}
	path := r.Path
	fmt.Println(path)
	requestURI := req.RequestURI

	if requestURI != path {
		return false
	}
	fmt.Println(requestURI)

	return true
}
