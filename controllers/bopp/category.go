package bopp

import (
	"ims_oa/proxy/models"
	"ims_oa/proxy/services"

	"github.com/gin-gonic/gin"
)

//	@Summary	查询辅材大类
//	@Tags		BOPP
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]string
//	@Failure	400	{object}	models.Response
//	@Router		/bopp/accessory/category [get]
func FindCategory(ctx *gin.Context) {
	res, err := services.BoppFindCategory()
	if err != nil {
		models.Fail(err.Error(), ctx)
		return
	}
	models.Ok(res, ctx)
}
