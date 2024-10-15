package handler

import (
	"CurrencyTask/services/gateway/service"
	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
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
		auth := api.Group("/auth")
		{
			auth.POST("/sign-in", h.signIn)
		}
		currency := api.Group("/currency")
		{
			currency.GET("/rate", h.authMiddleware(), h.GetCurrency)
			currency.GET("/history", h.authMiddleware(), h.GetRateHistory)
		}
	}
	return router
}
