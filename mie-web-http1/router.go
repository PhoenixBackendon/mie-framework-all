package mie

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	k := routeKey(method, pattern)
	r.handlers[k] = handler
}

func (r *Router) handle(c *Context) {
	k := routeKey(c.Method, c.Path)
	if h := r.handlers[k]; h != nil {
		h(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
