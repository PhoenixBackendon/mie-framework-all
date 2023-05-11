package mie

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	HTTPStatus int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

func (c *Context) PostForm(k string) string {
	return c.Req.FormValue(k)
}

func (c *Context) Query(k string) string {
	return c.Req.URL.Query().Get(k)
}

func (c *Context) Status(s int) error {
	if s < 100 || s >= 600 {
		return errors.New("invalid HTTP status code")
	}
	c.HTTPStatus = s
	c.Writer.WriteHeader(s)
	return nil
}

func (c *Context) SetHeader(k string, v string) {
	c.Writer.Header().Set(k, v)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	_ = c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	_ = c.Status(code)
	encode := json.NewEncoder(c.Writer)
	if err := encode.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), code)
	}
}

func (c *Context) Data(code int, data []byte) {
	_ = c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	_ = c.Status(code)
	c.Writer.Write([]byte(html))
}
