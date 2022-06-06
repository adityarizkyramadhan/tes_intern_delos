package route

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/implementation/endpoint_counter/service"
	_farmRoute "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/delivery_route"
	_farmHandler "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/handler"
	_farmService "github.com/adityarizkyramadhan/tes_intern_delos/implementation/farm/service"
	_pondRoute "github.com/adityarizkyramadhan/tes_intern_delos/implementation/pond/delivery_route"
	_pondHandler "github.com/adityarizkyramadhan/tes_intern_delos/implementation/pond/handler"
	_pondService "github.com/adityarizkyramadhan/tes_intern_delos/implementation/pond/service"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Route(e *gin.Engine, db *gorm.DB) {
	tracker := service.NewTracker(db)
	farmUseCase := _farmService.NewServiceFarm(db, tracker)
	farmHandler := _farmHandler.NewHandlerFarm(farmUseCase, tracker)
	farm := e.Group("/farm")
	_farmRoute.RouteFarm(farm, farmHandler)
	pondUseCase := _pondService.NewServicePond(db, tracker)
	pondHandler := _pondHandler.NewHandlerPond(pondUseCase, tracker)
	pond := e.Group("/pond", middleware.ValidateJWToken())
	_pondRoute.RoutePonds(pond, pondHandler)
}
