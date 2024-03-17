package server

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramonmedeiros/key_value_store/internal/keystore"
)

func (s *Server) getKey(c *gin.Context) {
	key := getKey(c)

	value, err := s.keyStore.GetKey(key)
	if errors.Is(err, keystore.ErrNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "text/plain", value)
}

func (s *Server) addKey(c *gin.Context) {
	key := getKey(c)

	defer c.Request.Body.Close()
	content, err := io.ReadAll(c.Request.Body)
	if err != nil {
		s.logger.Error("could not read body: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = s.keyStore.AddKey(key, content)
	if err != nil {
		s.logger.Error("could not add key: %s, error %s", key, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func getKey(c *gin.Context) string {
	key := c.Param(keyParam)
	if key == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	return key
}
