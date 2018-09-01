package web

import (
	"github.com/cthit/gotify"
	"github.com/gocraft/web"
	"net/http"
)

type Context struct {
	MailService gotify.MailService
	AuthKey     string
	Debug       bool
}

func Router(authKey string, mailServiceCreator func() gotify.MailService, debug bool) http.Handler {

	router := web.NewWithPrefix(
		Context{},
		"")

	router.Middleware(web.LoggerMiddleware)
	if debug {
		router.Middleware(web.ShowErrorsMiddleware)
	}

	router.Middleware(setDebugMode(debug))
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

func setDebugMode(debug bool) func(*Context, web.ResponseWriter, *web.Request, web.NextMiddlewareFunc) {
	return func(c *Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.Debug = debug
		next(rw, req)
	}
}
