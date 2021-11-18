package server

import (
	"log"

	"github.com/emersonluiz/go-user/logger"
	"github.com/emersonluiz/go-user/route"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (server *Server) Run() {
	router := route.ConfigRoute(server.server)
	logger.SetLog("Server running at port: " + server.port)
	log.Fatal(router.Run(":" + server.port))
}
