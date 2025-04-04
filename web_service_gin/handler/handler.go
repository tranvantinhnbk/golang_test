package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Users
// @Tags users
// @Produce json
// @Success 200 {array} string
// @Router /api/users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, []string{"Alice", "Bob", "Charlie"})
}
