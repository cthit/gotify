package web

import (
	"github.com/gocraft/web"
	"../../gotify"
	"net/http"
)

var DEBUG = false

type Context struct {
	MailService gotify.MailService
	AuthKey string
}

func Router(authKey string, mailServiceCreator func()  gotify.MailService) http.Handler {

	router := web.NewWithPrefix(
		Context{},
		"")

	router.Middleware(web.LoggerMiddleware)
	if DEBUG {
		router.Middleware(web.ShowErrorsMiddleware)
	}

	router.Middleware(setMailServiceProvider(mailServiceCreator))
	router.Middleware(setAuthKey(authKey))
	router.Middleware((*Context).Auth)

	router.Post("/mail", (*Context).SendMail)
	return router
}

func setMailServiceProvider(mailServiceProvider func() gotify.MailService) func(*Context, web.ResponseWriter, *web.Request, web.NextMiddlewareFunc) {
	return func(c *Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.MailService = mailServiceProvider()
		next(rw, req)
		c.MailService.Destroy()
	}
}

func setAuthKey(authKey string) func(*Context, web.ResponseWriter, *web.Request, web.NextMiddlewareFunc) {
	return func(c *Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.AuthKey = authKey
		next(rw, req)
	}
}