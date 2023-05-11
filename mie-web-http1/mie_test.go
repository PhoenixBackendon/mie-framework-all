package mie

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

var h = func(c *Context) {}

func TestNew(t *testing.T) {
	e := New()
	assert.Empty(t, e.router.handlers)
}

func TestEngine_AddRoute(t *testing.T) {
	e := New()
	e.router.AddRoute("TEST", "/test", h)
	assert.NotNil(t, e.router.handlers["TEST-/test"])

	e.GET("/test", h)
	assert.NotNil(t, e.router.handlers["GET-/test"])

	e.POST("/test", h)
	assert.NotNil(t, e.router.handlers["POST-/test"])
	assert.Equal(t, 3, len(e.router.handlers))
}

func TestEngine_Run(t *testing.T) {
	e := New()
	e.GET("/test", func(c *Context) {
		fmt.Fprintf(c.Writer, "Hello World!")
	})
	go e.Run("localhost:8080")
	//tr := &http.Transport{}
	client := &http.Client{}
	res, err := client.Get("http://localhost:8080/test")
	if err != nil {
		t.Error(err)
	}
	resp, _ := io.ReadAll(res.Body)
	assert.Equal(t, "Hello World!", string(resp))

	res, err = client.Get("http://localhost:8080/demo")
	if err != nil {
		t.Error(err)
	}
	resp, _ = io.ReadAll(res.Body)
	assert.Equal(t, "404 NOT FOUND: /demo\n", string(resp))
}
