package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"oanda-mock-api/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	req, err := http.NewRequest("GET", "/accounts", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := gin.Default()
	router.RemoveExtraSlash = true
	handlers.RegisterAccountHandlers(router.Group("/accounts"))
	router.ServeHTTP(rr, req)

	println(rr.Body.String())

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
}
