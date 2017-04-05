package controllers

import (
	"github.com/showntop/tripper/models"
)

type Categories struct {
	application
}

func (c *Categories) List() ([]byte, *HttpError) {


	data, err := models.GetCategories()
	if err != nil {
		return nil, DBErr
	}
	return WrapResp(data)
}
