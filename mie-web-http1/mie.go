package mie

import (
	"fmt"
	"net/http"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	k := routeKey(method, pattern)
	engine.router[k] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.AddRoute(http.MethodGet, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.AddRoute(http.MethodPost, pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	k := routeKey(r.Method, r.URL.Path)
	if h := engine.router[k]; h != nil {
		h(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func routeKey(method string, pattern string) string {
	s := []string{
		method,
		"-",
		pattern,
	}
	k := strings.Join(s, "")
	return k
}
