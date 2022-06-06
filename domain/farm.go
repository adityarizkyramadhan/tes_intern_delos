package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FarmModel struct {
	gorm.Model
	Username        string `gorm:"uniqueIndex"`
	Name            string
	Leader          string
	Supervisor      string
	NumberEmployees int
	Location        string
	Contact         string
	PondModels      []PondModel `gorm:"foreignkey:FarmModelId;references:ID"`
}

type FarmInput struct {
	Name            string `json:"name" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Leader          string `json:"leader" binding:"required"`
	Supervisor      string `json:"supervisor" binding:"required"`
	NumberEmployees int    `json:"number_employees" binding:"required"`
	Location        string `json:"location" binding:"required"`
	Contact         string `json:"contact" binding:"required"`
}

type FarmUseCase interface {
	Create(*FarmModel) (uint, error)
	Read(uint) (FarmModel, error)
	Delete(uint) error
	Update(*FarmModel, uint) error
	ReadAll() ([]FarmModel, error)
}

type FarmHandler interface {
	Create(*gin.Context)
	Read(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	ReadAll(*gin.Context)
}
