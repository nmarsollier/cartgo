package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"gopkg.in/go-playground/assert.v1"
)

// Obtiene Router server con el contexto de testing adecuado
// mockeando interfaces a serivcios externos
func TestRouter(deps ...interface{}) *gin.Engine {
	server = nil
	Router(deps...)
	if len(deps) > 0 {
		server.Use(func(c *gin.Context) {
			c.Set("mock_deps", deps)
			c.Next()
		})
	}
	return server
}

// Requests Test functions

func TestGetRequest(url string, tokenString string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if len(tokenString) > 0 {
		req.Header.Add("Authorization", "Bearer "+tokenString)
	}
	w := httptest.NewRecorder()
	return req, w
}

func TestDeleteRequest(url string, tokenString string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte{}))
	if len(tokenString) > 0 {
		req.Header.Add("Authorization", "Bearer "+tokenString)
	}
	w := httptest.NewRecorder()
	return req, w
}

func TestPostRequest(url string, body interface{}, tokenString string) (*http.Request, *httptest.ResponseRecorder) {
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if len(tokenString) > 0 {
		req.Header.Add("Authorization", "Bearer "+tokenString)
	}
	w := httptest.NewRecorder()
	return req, w
}

// Assertion Functions
func AssertUnauthorized(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	var result map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal(t, result["error"], "Unauthorized")
}

func AssertDocumentNotFound(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusNotFound, w.Code)

	var result map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, "Document not found", result["error"])
}

func AssertInternalServerError(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func AssertBadRequestError(t *testing.T, w *httptest.ResponseRecorder) {
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
