package controllers

import (
	"ims_oa/proxy/docs"
	"ims_oa/proxy/models"
	"runtime"

	"github.com/gin-gonic/gin"
)

// @Tags		/
// @Accept		json
// @Produce	json
// @Success	200	{object}	models.Response
// @Router		/ [get]
func Index(ctx *gin.Context) {
	models.Ok(gin.H{
		"ARCH":    runtime.GOARCH,
		"OS":      runtime.GOOS,
		"SDK":     runtime.Version(),
		"Version": docs.SwaggerInfo.Version,
		"Gin":     gin.Version,
	}, ctx)
}
