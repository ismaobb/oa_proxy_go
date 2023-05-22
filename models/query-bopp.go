package models

import (
	"ims_oa/proxy/repositories"
)

type ReqQueryBopp struct {
	JobOrder string `form:"job_order"`
	Take     int    `form:"take" binding:"required"`
	Factory  string `form:"factory"`
	From     string `form:"from"`
	To       string `form:"to"`
}

type ResQueryBopp struct {
	repositories.Order
	OrderDetails     []repositories.OrderDetail    `json:"order_details"`
	OrderAccessories []repositories.OrderAccessory `json:"order_accessories"`
}
