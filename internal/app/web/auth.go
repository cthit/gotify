package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) MustAuthorize(c *gin.Context) {
	if c.GetHeader("Authorization") != fmt.Sprintf("pre-shared: %s", s.preSharedKey) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
