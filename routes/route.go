package routes

import (
	"ims_oa/proxy/controllers"
	docs "ims_oa/proxy/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/"
	r := gin.Default()
	r.GET("/", controllers.Index)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	setupBopp(r)
	return r
}
