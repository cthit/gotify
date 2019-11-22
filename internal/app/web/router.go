package web

import (
	"github.com/cthit/gotify/pkg/mail"
	"github.com/gocraft/web"
	"net/http"
)

type Context struct {
	MailService mail.MailService
	AuthKey     string
	Debug       bool
}

func Router(authKey string, mailService mail.MailService, debug bool) http.Handler {

	router := web.NewWithPrefix(
		Context{},
		"")

	router.Middleware(web.LoggerMiddleware)
	if debug {
		router.Middleware(web.ShowErrorsMiddleware)
	}

	router.Middleware(setDebugMode(debug))
	router.Middleware(setMailServiceProvider(mailService))
	router.Middleware(setAuthKey(authKey))
	router.Middleware((*Context).Auth)

	router.Post("/mail", (*Context).SendMail)
	return router
}

func setMailServiceProvider(mailService mail.MailService) func(*Context, web.ResponseWriter, *web.Request, web.NextMiddlewareFunc) {
	return func(c *Context, rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
		c.MailService = mailService
		next(rw, req)
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
