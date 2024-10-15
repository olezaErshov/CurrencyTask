package handler

import (
	"CurrencyTask/services/currency/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Servicer
}

func NewHandler(service service.Servicer) Handler {
	return Handler{service: service}
}

func InitRoutes(h *Handler) *gin.Engine {
	router := gin.New()
	api := router.Group("/api/v1")
	{
		trips := api.Group("/rate")
		{
			trips.POST("/date", h.RateByDay)
		}

	}
	return router
}
