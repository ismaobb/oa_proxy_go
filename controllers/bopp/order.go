package bopp

import (
	"ims_oa/proxy/models"
	"ims_oa/proxy/services"
	"log"

	"github.com/gin-gonic/gin"
)

//	@Summary	更新
//	@Tags		BOPP
//	@Accept		json
//	@Produce	json
//
//	@Param		job_order	path		string					true	"通知单号"
//	@Param		update_dto	body		models.UpdateBoppOrder	true	"更新数据体"
//
//	@Success	200			{boolean}	true
//	@Failure	400			{object}	models.Response
//	@Router		/bopp/{job_order} [patch]
func PatchOrder(c *gin.Context) {
	jobOrder := c.Param("job_order")
	if len(jobOrder) == 0 {
		models.Fail("job_order not exists", c)
		return
	}
	update := models.UpdateBoppOrder{}
	if err := c.ShouldBindJSON(&update); err != nil {
		models.Fail(err.Error(), c)
		return
	}

	log.Println("jobOrder=", jobOrder, "  state=", update.State)
	err := services.BoppPatchOrder(jobOrder, update.State)
	if err != nil {
		models.Fail(err.Error(), c)
		return
	}

	models.Ok(true, c)
}
