package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;autoIncrement:false"`
	FullName string    `json:"fullName"`
	Document string    `json:"document" gorm:"unique"`
	Email    string    `json:"email" gorm:"unique"`
	Password string    `json:"-"`
	Role     string    `json:"role"`
	Balance  float64   `json:"balance"`
}

const (
	LOJISTA = "LOJISTA"
	USUARIO = "USUARIO"
)
