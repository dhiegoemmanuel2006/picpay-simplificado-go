package controller

import (
	"github.com/dhiegoemmanuel2006/picpay-simplificado-go/services"
	"gorm.io/gorm"
)

type Api struct {
	Us services.UserService
}

func NewApi(db *gorm.DB) *Api {
	return &Api{
		Us: services.NewUserService(db),
	}
}
