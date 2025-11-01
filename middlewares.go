package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) userMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"message": "missing Authorization header"})
		c.Abort()
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		c.JSON(401, gin.H{"message": "invalid Authorization format"})
		c.Abort()
		return
	}

	token := parts[1]
	session := s.validateSession(token)
	if session == nil {
		c.JSON(401, gin.H{"message": "invalid or expired session"})
		c.Abort()
		return
	}

	c.Set("user_id", session.UserID)
	c.Next()
}
