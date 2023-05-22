package routes

import (
	"ims_oa/proxy/controllers/bopp"

	"github.com/gin-gonic/gin"
)

func setupBopp(r *gin.Engine) {
	rb := r.Group("/bopp")
	{
		rb.GET("", bopp.FindAll)
		rb.GET("/accessory/category", bopp.FindCategory)
		rb.GET("/accessory/detail", bopp.FindDetail)
		rb.PATCH(":job_order", bopp.PatchOrder)
	}
}
