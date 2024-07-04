package handlers

import (
	"net/http"

	"oanda-mock-api/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterAccountHandlers(group *gin.RouterGroup) {
	group.GET("", GetAccount)
}

func GetAccount(c *gin.Context) {
	// TODO: Get account ID from auth token

	account, err := services.GetAccount("123")
	if err != nil {
		// Return error as JSON
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return account as JSON
	c.JSON(http.StatusOK, account)
}
