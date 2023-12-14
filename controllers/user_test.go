package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignUpValidParams(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := []byte(`{
		"name": "Test",
		"email": "test@gmail.com",
		"password": "Aa@123456"
	}`)
	req, _ := http.NewRequest("POST", "/api/v1/signup", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "New user account registered", w.Body.String())
}
