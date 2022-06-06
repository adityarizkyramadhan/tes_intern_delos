package handler

import (
	"net/http"

	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"github.com/adityarizkyramadhan/tes_intern_delos/utils/response"
	"github.com/gin-gonic/gin"
)

type pondHandler struct {
	UseCase domain.PondUseCase
	Handler domain.FarmHandler
	Tracker domain.TrackerUseCase
}

func NewHandlerPond(useCase domain.PondUseCase, tracker domain.TrackerUseCase) domain.PondHandler {
	return &pondHandler{
		UseCase: useCase,
		Tracker: tracker,
	}
}

func (p pondHandler) Create(c *gin.Context) {
	idUser := c.MustGet("login").(uint)
	var input domain.PondInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind json", err.Error()))
	}
	if err := p.UseCase.Create(&domain.PondModel{
		FarmModelId: idUser,
		Commodity:   input.Commodity,
		Capacity:    input.Capacity,
		Price:       input.Price,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when create pond", err.Error()))
	}
	count, err := p.Tracker.SearchEndpointCalled("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := p.Tracker.SearchUniqueUserAgent("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Success create pond", nil))
}

func (p pondHandler) Read(c *gin.Context) {
	idUser := c.MustGet("login").(uint)
	pond, err := p.UseCase.Read(idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when read pond", err.Error()))
		return
	}
	count, err := p.Tracker.SearchEndpointCalled("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := p.Tracker.SearchUniqueUserAgent("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Success read pond", pond))
}

type uriPond struct {
	Id uint `uri:"id" binding:"required"`
}

func (p pondHandler) Update(c *gin.Context) {
	var idPond uriPond
	idUser := c.MustGet("login").(uint)
	if err := c.BindUri(&idPond); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind uri", err.Error()))
		return
	}
	var input domain.PondInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind json", err.Error()))
	}
	data, err := p.UseCase.ReadById(idPond.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when read pond", err.Error()))
		return
	}
	if data.FarmModelId != idUser {
		c.JSON(http.StatusUnauthorized, response.ResponseWhenFail("Error when read pond", "You are not authorized"))
		return
	}
	if err := p.UseCase.Update(idPond.Id, &domain.PondModel{
		Commodity: input.Commodity,
		Capacity:  input.Capacity,
		Price:     input.Price,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when update pond", err.Error()))
	}
	count, err := p.Tracker.SearchEndpointCalled("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := p.Tracker.SearchUniqueUserAgent("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Success update pond", nil))
}

func (p pondHandler) Delete(c *gin.Context) {
	var idPond uriPond
	idUser := c.MustGet("login").(uint)
	if err := c.BindUri(&idPond); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind uri", err.Error()))
		return
	}
	data, err := p.UseCase.ReadById(idPond.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when read pond", err.Error()))
		return
	}
	if data.FarmModelId != idUser {
		c.JSON(http.StatusUnauthorized, response.ResponseWhenFail("Error when read pond", "You are not authorized"))
		return
	}
	if err := p.UseCase.Delete(idPond.Id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when delete pond", err.Error()))
	}
	count, err := p.Tracker.SearchEndpointCalled("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := p.Tracker.SearchUniqueUserAgent("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Success delete pond", nil))
}

func (p pondHandler) ReadById(c *gin.Context) {
	var idPond uriPond
	idUser := c.MustGet("login").(uint)
	if err := c.BindUri(&idPond); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Error when bind uri", err.Error()))
		return
	}
	data, err := p.UseCase.ReadById(idPond.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Error when read pond", err.Error()))
		return
	}
	if data.FarmModelId != idUser {
		c.JSON(http.StatusUnauthorized, response.ResponseWhenFail("Error when read pond", "You are not authorized"))
		return
	}
	count, err := p.Tracker.SearchEndpointCalled("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := p.Tracker.SearchUniqueUserAgent("/pond")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Success read pond", data))
}
