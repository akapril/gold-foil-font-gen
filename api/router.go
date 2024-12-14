package api

import (
	"github.com/gin-gonic/gin"
	"gold-foil-font-gen/handler"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	apiHandler := handler.NewApiHandler()
	r := Router.Group("/api")
	r.GET("/gold-foil-image", apiHandler.GoldFoilImage)
	r.GET("/gold-foil-svg", apiHandler.GoldFoilSvg)
	Router.GET("/img/gold-foil-image.png", apiHandler.GoldFoilImageFile)
}
