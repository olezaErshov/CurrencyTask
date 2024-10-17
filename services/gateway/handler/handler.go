package handler

import (
	"CurrencyTask/services/gateway/config"
	"CurrencyTask/services/gateway/service"

	"github.com/gin-gonic/gin"
)

const (
	authHeader              = "Authorization"
	requestExpiredInSeconds = 2
)

type Handler struct {
	service service.Servicer
	cfg     config.UrlsConfig
}

func NewHandler(service service.Servicer, cfg config.UrlsConfig) Handler {
	return Handler{service: service, cfg: cfg}
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
			currency.GET("/rate", h.authMiddleware(), h.GetRate)
			currency.GET("/history", h.authMiddleware(), h.GetRateHistory)
		}
	}
	return router
}
