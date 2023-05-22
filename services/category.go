package services

import (
	"ims_oa/proxy/bootstrap"
	"ims_oa/proxy/repositories"
)

func BoppFindCategory() ([]string, error) {
	categories := []string{}
	err := bootstrap.App.Client.KL.Model(&repositories.AccessoryColumn{}).Distinct().Pluck("Category", &categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
