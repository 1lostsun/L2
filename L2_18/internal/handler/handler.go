package handler

import (
	"github.com/1lostsun/L2/tree/main/L2_18/internal/middleware"
	"github.com/1lostsun/L2/tree/main/L2_18/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Handler : Структура ручки
type Handler struct {
	uc *usecase.UseCase
}

// New : функция конструктор, возвращающая структуру Handler
func New(uc *usecase.UseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

// InitRoutes : Инициализация маршрутов
func (h *Handler) InitRoutes(r *gin.Engine) {
	r.Use(middleware.GinLogger(), middleware.Force400Middleware())
	v1 := r.Group("/api/v1")
	{
		v1.POST("/create_event", h.CreateEvent)
		v1.POST("/update_event", h.UpdateEvent)
		v1.POST("/delete_event", h.DeleteEvent)

		v1.GET("/events_for_day", h.EventsForDay)
		v1.GET("/events_for_week", h.EventsForWeek)
		v1.GET("/events_for_month", h.EventsForMonth)
	}
}
