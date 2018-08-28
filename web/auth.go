package web

import (
	"github.com/gocraft/web"
	"fmt"
	"net/http"
)

func (c *Context) Auth(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	if req.Header.Get("Authorization") == fmt.Sprintf("pre-shared: %s", c.AuthKey) {
			next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}