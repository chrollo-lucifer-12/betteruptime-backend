package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const defaultPort = ":8080"

type ServerOpts struct {
	Port string
	DB   *gorm.DB
}

type Server struct {
	r    *gin.Engine
	port string
	db   *gorm.DB
}

func NewServer(opts ServerOpts) *Server {
	r := gin.Default()
	return &Server{
		r:    r,
		port: opts.Port,
		db:   opts.DB,
	}
}

func (s *Server) Start() error {
	if len(s.port) == 0 {
		s.port = defaultPort
	}
	s.RegisterRoutes(s.r)
	err := s.r.Run(s.port)
	if err != nil {
		return err
	}
	return nil
}
