package handler

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
}

type authHandler struct{}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (h *authHandler) Register(c *gin.Context) {

}

func (h *authHandler) Login(c *gin.Context) {

}

func (h *authHandler) GetProfile(c *gin.Context) {

}
