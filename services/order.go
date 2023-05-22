package services

import (
	"ims_oa/proxy/bootstrap"
	"ims_oa/proxy/repositories"

	"github.com/gin-gonic/gin"
)

func BoppPatchOrder(jobOrder string, state string) (err error) {
	if bootstrap.App.Env.Mode == gin.DebugMode {
		err = bootstrap.App.Client.KL.Model(&repositories.Order{}).Where("TZDBH = ?", jobOrder).Update("SCBDZT", state).Error
		return
	} else {
		bootstrap.App.Client.KL.Model(&repositories.Order{}).Where("TZDBH = ?", jobOrder).Update("SCBDZT", state)
		bootstrap.App.Client.KC.Model(&repositories.Order{}).Where("TZDBH = ?", jobOrder).Update("SCBDZT", state)
		bootstrap.App.Client.KW.Model(&repositories.Order{}).Where("TZDBH = ?", jobOrder).Update("SCBDZT", state)
	}

	return nil
}
