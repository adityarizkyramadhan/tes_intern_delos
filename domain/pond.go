package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PondModel struct {
	gorm.Model
	FarmModelId uint
	Commodity   string
	Capacity    int
	Price       int
}

type PondInput struct {
	Commodity string `json:"commodity" binding:"required"`
	Capacity  int    `json:"capacity" binding:"required"`
	Price     int    `json:"price" binding:"required"`
}

type PondUseCase interface {
	Create(*PondModel) error
	Read(uint) ([]PondModel, error)
	Update(uint, *PondModel) error
	Delete(uint) error
	ReadById(uint) (PondModel, error)
}

type PondHandler interface {
	Create(*gin.Context)
	Read(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	ReadById(*gin.Context)
}
