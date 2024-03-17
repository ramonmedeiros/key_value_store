package server

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const keyParam = "key"

type Server struct {
	port   string
	logger *slog.Logger
	router *gin.Engine
	cache  map[string][]byte
}

type API interface {
	Serve()
}

func New(port string, logger *slog.Logger) *Server {
	return &Server{
		port:   port,
		logger: logger,
		cache:  map[string][]byte{},
	}
}

func (s *Server) Serve() {
	router := gin.Default()
	s.router = router

	s.setupConfig(router)

	s.setupKeyValueStoreEndpoints()

	router.Run("0.0.0.0:" + s.port)
}

func (s *Server) setupConfig(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
}

func (s *Server) setupKeyValueStoreEndpoints() {
	keyStoreEndpoints := s.router.Group("/")

	keyStoreEndpoints.GET(":"+keyParam, s.getKey)
	keyStoreEndpoints.POST(":"+keyParam, s.addKey)
}
