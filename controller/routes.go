package controller

import (
	"github.com/gin-gonic/gin"
)

func (api *Api) NewHandler() *gin.Engine {
	h := gin.New()

	h.Use(gin.Recovery())
	h.Use(gin.Logger())

	h.POST("/tranfer", api.HandleTranfer)
	h.POST("/create-user", api.HandleCreateUser)

	return h
}
