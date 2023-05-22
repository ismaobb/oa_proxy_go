package bopp

import (
	"ims_oa/proxy/models"
	"ims_oa/proxy/services"
	"log"

	"github.com/gin-gonic/gin"
)

//	@Tags		BOPP
//	@Summary	查询
//	@Accept		json
//	@Produce	json
//	@Param		request	query		models.ReqQueryBopp	true	"查询"
//	@Success	200		{object}	models.ResQueryBopp
//	@Failure	400		{object}	models.Response
//	@Router		/bopp [get]
func FindAll(ctx *gin.Context) {
	q := models.ReqQueryBopp{}
	if err := ctx.ShouldBindQuery(&q); err != nil {
		models.Fail(err.Error(), ctx)
		return
	}

	log.Printf("%+v", q)
	res, err := services.BoppFindAll(&q)
	if err != nil {
		models.Fail(err.Error(), ctx)
		return
	}

	models.Ok(res, ctx)
}
