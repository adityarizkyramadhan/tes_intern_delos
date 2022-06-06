package handler

import (
	"net/http"

	"github.com/adityarizkyramadhan/tes_intern_delos/domain"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/adityarizkyramadhan/tes_intern_delos/utils/response"
	"github.com/gin-gonic/gin"
)

type farmHandler struct {
	FarmUseCase domain.FarmUseCase
	FarmHandler domain.FarmHandler
	Tracker     domain.TrackerUseCase
}

func NewHandlerFarm(useCase domain.FarmUseCase, tracker domain.TrackerUseCase) domain.FarmHandler {
	return &farmHandler{
		FarmUseCase: useCase,
		Tracker:     tracker,
	}
}

func (h farmHandler) Create(c *gin.Context) {
	var input domain.FarmInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Bind JSON fail", err))
		return
	}
	id, err := h.FarmUseCase.Create(&domain.FarmModel{
		Username:        input.Username,
		Contact:         input.Contact,
		Location:        input.Location,
		Leader:          input.Leader,
		Supervisor:      input.Supervisor,
		Name:            input.Name,
		NumberEmployees: input.NumberEmployees,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Create data fail", err))
		return
	}
	token, err := middleware.GenerateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Create token fail", err))
		return
	}
	count, err := h.Tracker.SearchEndpointCalled("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := h.Tracker.SearchUniqueUserAgent("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusCreated, response.ResponseWhenSuccess(count, uniqueCount, "Create data success", gin.H{
		"token": token,
	}))
}

func (h farmHandler) Update(c *gin.Context) {
	var input domain.FarmInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseWhenFail("Bind JSON fail", err))
		return
	}
	idUser := c.MustGet("login").(uint)
	if err := h.FarmUseCase.Update(&domain.FarmModel{
		Username:        input.Username,
		Contact:         input.Contact,
		Location:        input.Location,
		Leader:          input.Leader,
		Supervisor:      input.Supervisor,
		Name:            input.Name,
		NumberEmployees: input.NumberEmployees,
	}, idUser); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Update data fail", err))
		return
	}
	count, err := h.Tracker.SearchEndpointCalled("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := h.Tracker.SearchUniqueUserAgent("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Update data success", nil))
}

func (h farmHandler) Delete(c *gin.Context) {
	idUser := c.MustGet("login").(uint)
	if err := h.FarmUseCase.Delete(idUser); err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Delete data fail", nil))
		return
	}
	count, err := h.Tracker.SearchEndpointCalled("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := h.Tracker.SearchUniqueUserAgent("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Delete data success", nil))
}

func (h farmHandler) Read(c *gin.Context) {
	idUser := c.MustGet("login").(uint)
	data, err := h.FarmUseCase.Read(idUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Read data fail", nil))
		return
	}
	count, err := h.Tracker.SearchEndpointCalled("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := h.Tracker.SearchUniqueUserAgent("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Read data success", data))
}

func (h farmHandler) ReadAll(c *gin.Context) {
	data, err := h.FarmUseCase.ReadAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Read data fail", nil))
		return
	}
	count, err := h.Tracker.SearchEndpointCalled("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Endpoint count fail", err))
		return
	}
	uniqueCount, err := h.Tracker.SearchUniqueUserAgent("/farm")
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ResponseWhenFail("Unique user agent fail", err))
		return
	}
	c.JSON(http.StatusOK, response.ResponseWhenSuccess(count, uniqueCount, "Read data success", data))
}
