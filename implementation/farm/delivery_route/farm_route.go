package delivery_route

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/gin-gonic/gin"
)

func RouteFarm(farm *gin.RouterGroup, farmHandler domain.FarmHandler) {
	farm.POST("/", farmHandler.Create)
	farm.PUT("/", middleware.ValidateJWToken(), farmHandler.Update)
	farm.DELETE("/", middleware.ValidateJWToken(), farmHandler.Delete)
	farm.GET("/", middleware.ValidateJWToken(), farmHandler.Read)
	farm.GET("/all", farmHandler.ReadAll)
}
