package handler

import "github.com/gin-gonic/gin"

type ApiHandler struct {
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{}
}

func (api *ApiHandler) GoldFoilImage(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (api *ApiHandler) GoldFoilSvg(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func (api *ApiHandler) GoldFoilImageFile(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
