package web

import (
	"github.com/cthit/gotify/pkg/mail"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port         string
	preSharedKey string
	debug        bool
	mailService  mail.MailService
}

func NewServer(port, preSharedKey string, debug bool, mailService mail.MailService) (*Server, error) {
	return &Server{
		port:         port,
		preSharedKey: preSharedKey,
		debug:        debug,
		mailService:  mailService,
	}, nil
}

func (s *Server) Start() error {
	if s.debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	authorized := router.Group("/")
	authorized.Use(s.MustAuthorize)
	{
		authorized.POST("/mail", s.MailHandler)
	}
	return router.Run(":" + s.port)
}
