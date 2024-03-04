package api

import "github.com/gin-gonic/gin"

type Server struct {
	Router *gin.Engine
}

func SpawnServer() *Server {
	engine := gin.Default()
	return &Server{
		Router: engine,
	}
}

func (s *Server) MountHandlers() {
	api := s.Router.Group("/api/queue")
	api.POST("/join", JoinQueue)
}
