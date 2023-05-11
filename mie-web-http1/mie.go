package mie

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.AddRoute(http.MethodGet, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.AddRoute(http.MethodPost, pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
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
