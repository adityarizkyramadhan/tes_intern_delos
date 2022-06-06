package delivery_route

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"github.com/gin-gonic/gin"
)

func RoutePonds(pond *gin.RouterGroup, pondHandler domain.PondHandler) {
	pond.POST("/", pondHandler.Create)
	pond.GET("/", pondHandler.Read)
	pond.PUT("/:id", pondHandler.Update)
	pond.DELETE("/:id", pondHandler.Delete)
	pond.GET("/:id", pondHandler.ReadById)
}
