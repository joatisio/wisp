package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joatisio/wisp/internal/router"
	"net/http"
)

var (
	errLogin = errors.New("login error")
)

type Handlers struct {
	service Auth
}

func (h *Handlers) Login(c *gin.Context) {
	var json LoginRequest
	err := c.ShouldBindJSON(&json)
	if err != nil {
		router.GinJsonError(c, err, http.StatusBadRequest)
		return
	}

	resp, err := h.service.Login(json)
	if err != nil {
		router.GinJsonError(c, errLogin, http.StatusBadRequest)
	}

	router.GinJsonResponse(c, resp, http.StatusOK)
}
