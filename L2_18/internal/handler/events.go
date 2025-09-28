package handler

import (
	"github.com/1lostsun/L2/tree/main/L2_18/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateEvent : Ручка создания события
func (h *Handler) CreateEvent(c *gin.Context) {
	var body entity.Request

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.CreateEvent(body); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": body})
}

// UpdateEvent : Ручка обновления события
func (h *Handler) UpdateEvent(c *gin.Context) {
	var body entity.Request

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.UpdateEvent(body); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": body})
}

// DeleteEvent : Ручка удаления события
func (h *Handler) DeleteEvent(c *gin.Context) {
	var body entity.Request
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.DeleteEvent(body.UserID, body.Title, body.Date); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "event was delete"})
}

// EventsForDay : Ручка получения событий за день
func (h *Handler) EventsForDay(c *gin.Context) {
	var body entity.Request
	var err error

	if userID := c.Query("user_id"); userID != "" {
		body.UserID, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}

		body.Title = c.Query("title")
		body.Description = c.Query("description")
		body.Date = c.Query("date")
	} else {
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.uc.EventsForDay(body.UserID, body.Date)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": resp})
}

// EventsForWeek : Ручка получения событий за неделю
func (h *Handler) EventsForWeek(c *gin.Context) {
	var body entity.Request
	var err error

	if userID := c.Query("user_id"); userID != "" {
		body.UserID, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}

		body.Title = c.Query("title")
		body.Description = c.Query("description")
		body.Date = c.Query("date")
	} else {
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.uc.EventsForWeek(body.UserID, body.Date)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": resp})
}

// EventsForMonth : Ручка получения событий за месяц
func (h *Handler) EventsForMonth(c *gin.Context) {
	var body entity.Request
	var err error

	if userID := c.Query("user_id"); userID != "" {
		body.UserID, err = strconv.ParseUint(userID, 10, 64)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}

		body.Title = c.Query("title")
		body.Description = c.Query("description")
		body.Date = c.Query("date")
	} else {
		if err = c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	resp, err := h.uc.EventsForMonth(body.UserID, body.Date)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": resp})
}
