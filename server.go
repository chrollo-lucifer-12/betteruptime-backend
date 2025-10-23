package main

import "github.com/gin-gonic/gin"

const defaultPort = ":8080"

type ServerOpts struct {
	Port string
}

type Server struct {
	r    *gin.Engine
	port string
}

func NewServer(opts ServerOpts) *Server {
	r := gin.Default()
	return &Server{
		r:    r,
		port: opts.Port,
	}
}

func (s *Server) Start() error {
	if len(s.port) == 0 {
		s.port = defaultPort
	}
	err := s.r.Run(s.port)
	if err != nil {
		return err
	}
	return nil
}
