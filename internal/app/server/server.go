package server

import (
	"log/slog"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ramonmedeiros/key_value_store/internal/keystore"
)

const keyParam = "key"

type Server struct {
	port     string
	logger   *slog.Logger
	router   *gin.Engine
	keyStore keystore.KeyStorer
}

type API interface {
	Serve()
}

func New(port string, logger *slog.Logger, keyStore keystore.KeyStorer) *Server {
	server := &Server{
		port:     port,
		logger:   logger,
		keyStore: keyStore,
	}

	server.router = gin.Default()
	server.setupConfig()
	server.setupKeyValueStoreEndpoints()
	return server
}

func (s *Server) Serve() {
	s.router.Run("0.0.0.0:" + s.port)
}

func (s *Server) setupConfig() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	s.router.Use(cors.New(config))
}

func (s *Server) setupKeyValueStoreEndpoints() {
	keyStoreEndpoints := s.router.Group("/")

	keyStoreEndpoints.GET(":"+keyParam, s.getKey)
	keyStoreEndpoints.POST(":"+keyParam, s.addKey)
}
