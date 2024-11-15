package main_test

import (
	"bytes"
	"encoding/json"
	"golang-gorm-gin/internal/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignInRoute(t *testing.T) {
	router := router.InitializeRouter()
	payload := `{"username": "testuser", "password": "password123"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/signin", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginRoute(t *testing.T) {
	router := router.InitializeRouter()
	payload := `{"username": "testuser", "password": "password123"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}


func TestListLayanan(t *testing.T) {
	r := router.InitializeRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/laundry-services", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3MjkwMjgsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.BuNq4V-ZfT_FTdPBGWghsWtgQQSKTlXfX0E4DT36E9E"})

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetLayanan(t *testing.T)  {
	r := router.InitializeRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/laundry-services/1", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3MjkwMjgsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.BuNq4V-ZfT_FTdPBGWghsWtgQQSKTlXfX0E4DT36E9E"})
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostLayanan(t *testing.T){
	r := router.InitializeRouter()
	payload := map[string]interface{}{
		"name" : "tes iki",
		"price_per_kg" : 23.09,
		"description" : "Laundry test",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/laundry-services", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3MjkwMjgsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.BuNq4V-ZfT_FTdPBGWghsWtgQQSKTlXfX0E4DT36E9E"})

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateLayanan(t *testing.T){
	r := router.InitializeRouter()
	payload := map[string]interface{}{
		"name" : "tes iki",
		"price" : 90.23,
		"description" : "Update tes iki",
	}
	body, _ := json.Marshal(payload)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/laundry-services/1",bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3MjkwMjgsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.BuNq4V-ZfT_FTdPBGWghsWtgQQSKTlXfX0E4DT36E9E"})

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteLayanan(t *testing.T){
	r := router.InitializeRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/laundry-services/8", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE3MjkwMjgsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.BuNq4V-ZfT_FTdPBGWghsWtgQQSKTlXfX0E4DT36E9E"})

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}