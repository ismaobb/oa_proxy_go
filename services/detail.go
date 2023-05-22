package services

import (
	"ims_oa/proxy/bootstrap"
	"ims_oa/proxy/repositories"
)

func BoppFindDetail() ([]string, error) {
	details := []string{}
	err := bootstrap.App.Client.KL.Model(&repositories.AccessoryColumn{}).Distinct().Pluck("Accessory", &details).Error
	if err != nil {
		return nil, err
	}
	return details, nil
}
