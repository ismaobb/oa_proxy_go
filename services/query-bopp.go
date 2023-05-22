package services

import (
	"ims_oa/proxy/bootstrap"
	"ims_oa/proxy/models"
	"ims_oa/proxy/repositories"
	"ims_oa/proxy/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BoppFindAll(q *models.ReqQueryBopp) (*[]models.ResQueryBopp, error) {
	if bootstrap.App.Env.Mode == gin.ReleaseMode {
		rKL, _ := dataAdapter(q.Take, bootstrap.App.Client.KL)
		rKW, _ := dataAdapter(q.Take, bootstrap.App.Client.KW)
		rKC, _ := dataAdapter(q.Take, bootstrap.App.Client.KC)
		merged := append(append(*rKL, *rKW...), *rKC...)
		return &merged, nil
	} else {
		return dataAdapter(q.Take, bootstrap.App.Client.KL)
	}
}

func dataAdapter(limit int, db *gorm.DB) (*[]models.ResQueryBopp, error) {
	orders := []repositories.Order{}
	err := db.Limit(limit).Where("SCBDZT IN (?)", []string{utils.Ineffective, utils.NOT_BACKUP}).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	jobOrders := []string{}
	for _, v := range orders {
		jobOrders = append(jobOrders, v.JobOrder)
	}
	orderDetails := []repositories.OrderDetail{}
	db.Where("TZDBH IN ?", jobOrders).Find(&orderDetails)

	orderAccessories := []repositories.OrderAccessory{}
	db.Where("TZDBH IN ?", jobOrders).Find(&orderAccessories)

	res := []models.ResQueryBopp{}
	for _, o := range orders {
		OrderDetails := []repositories.OrderDetail{}
		for _, v := range orderDetails {
			if o.JobOrder == v.JobOrder {
				OrderDetails = append(OrderDetails, v)
			}
		}

		OrderAccessories := []repositories.OrderAccessory{}
		for _, a := range orderAccessories {
			if a.JobOrder == o.JobOrder {
				OrderAccessories = append(OrderAccessories, a)
			}
		}

		ws := models.ResQueryBopp{Order: o, OrderDetails: OrderDetails, OrderAccessories: OrderAccessories}
		res = append(res, ws)
	}

	err = db.Model(&repositories.Order{}).Where("TZDBH IN ?", jobOrders).Update("SCBDZT", utils.BACKUP).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
