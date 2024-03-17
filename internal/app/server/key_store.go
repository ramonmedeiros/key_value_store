package server

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getKey(c *gin.Context) {
	key := c.Param(keyParam)
	if key == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	value, err := s.keyStore.GetKey(key)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Data(http.StatusOK, "text/plain", value)
}

func (s *Server) addKey(c *gin.Context) {
	key := c.Param(keyParam)
	if key == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	defer c.Request.Body.Close()
	content, err := io.ReadAll(c.Request.Body)
	if err != nil {
		s.logger.Error("could not read body: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	err = s.keyStore.AddKey(key, content)
	if err != nil {
		s.logger.Error("could not add key: %s, error %s", key, err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Status(http.StatusCreated)
}
