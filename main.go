package main

import (
	"fmt"
	"ims_oa/proxy/bootstrap"
	"ims_oa/proxy/routes"

	"github.com/gin-gonic/gin"
)

//	@title			Go-OA-Proxy
//	@version		1.0
//	@description	This is a Go-OA-Proxy server.

// @BasePath	/
func main() {
	gin.SetMode(bootstrap.App.Env.Mode)
	r := routes.SetupRouter()
	fmt.Println(fmt.Sprint("[Swagger]  ", "http://localhost:", bootstrap.App.Env.Port, "/swagger/index.html"))
	r.Run(fmt.Sprint(":", bootstrap.App.Env.Port))
}
