package api

import (
	"net/http"

	"github.com/TutorialEdge/notification-service/internal/list"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateList(c *gin.Context) {
	h.log.Info(c.Request.Context(), "creating a new list")
	var newList list.List
	if err := c.BindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request to list object",
			"error":   err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	createdList, err := h.listService.CreateList(c.Request.Context(), newList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create new list",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "list created",
		"list":    createdList,
	})
}
