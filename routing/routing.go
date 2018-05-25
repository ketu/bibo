package routing

import (
	"net/http"
	"fmt"
)

const (
	HTTPMethodGET    = "GET"
	HTTPMethodPOST   = "POST"
	HTTPMethodPUT    = "PUT"
	HTTPMethodPATH   = "PATH"
	HTTPMethodDELETE = "DELETE"
)

type Routing struct {
	// See Router.StrictSlash(). This defines the flag for new routes.
	strictSlash bool

	// Configurable Handler to be used when no route matches.
	NotFoundHandler http.Handler

	// Configurable Handler to be used when the request method does not match the route.
	MethodNotAllowedHandler http.Handler

	routes      []*Route
	namedRoutes map[string]*Route
}

func (r *Routing) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("This is a test"))

	for _, route := range (r.routes) {
		fmt.Println(route.Method)
		fmt.Println(route.Name)
		fmt.Println(route.Path)
		if isMatched :=route.Match(req); isMatched{
			route.Handler.ServeHTTP(w, req)
		}
	}

}

func (r *Routing) Map(path string, handler func(w http.ResponseWriter, r *http.Request), methods []string, name string) {
	for _, m := range methods {
		r.add(path, handler, m, name)
	}
}

func (r *Routing) add(path string, handler func(w http.ResponseWriter, r *http.Request), method string, name string) {
	route := &Route{
		Name:    name,
		Path:    path,
		Method:  method,
		Handler: handler,
	}
	r.namedRoutes[route.Name] = route
	r.routes = append(r.routes, route)
}

func (r *Routing) Get(path string, handler func(w http.ResponseWriter, r *http.Request), name string) {
	r.add(path, handler, HTTPMethodGET, name)
}
func (r *Routing) Post(path string, handler func(w http.ResponseWriter, r *http.Request), name string) {
	r.add(path, handler, HTTPMethodPOST, name)
}
func (r *Routing) Put(path string, handler func(w http.ResponseWriter, r *http.Request), name string) {
	r.add(path, handler, HTTPMethodPUT, name)
}
func (r *Routing) Path(path string, handler func(w http.ResponseWriter, r *http.Request), name string) {
	r.add(path, handler, HTTPMethodPATH, name)
}
func (r *Routing) Delete(path string, handler func(w http.ResponseWriter, r *http.Request), name string) {
	r.add(path, handler, HTTPMethodDELETE, name)
}

func NewRouting() *Routing {
	r := &Routing{
		namedRoutes: make(map[string]*Route),
	}
	return r
}
