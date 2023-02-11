package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list created"})
}
