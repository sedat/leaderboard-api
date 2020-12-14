package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/sedat/leaderboard-api/handlers"
)

var id string

func TestUserCreation(t *testing.T) {
	json := []byte(`{"display_name": "sedat", "country": "TR", "password": "1234"}`)
	req, err := http.NewRequest("POST", "localhost:8000/user/create", bytes.NewBuffer(json))
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	handlers.CreateUser(c)
	assert.Equal(t, 201, w.Code) // or what value you need it to be
}

func TestUserLogin(t *testing.T) {
	json := []byte(`{"display_name": "sedat", "password": "1234"}`)
	req, err := http.NewRequest("POST", "localhost:8000/user/login", bytes.NewBuffer(json))
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	handlers.LoginUser(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be
}

func TestSubmitScore(t *testing.T) {
	json := []byte(`{
		"user_id": "", // INSERT ID HERE
		"score_worth": 300,
		"timestamp": "3545234534"
	}`)
	req, err := http.NewRequest("POST", "localhost:8000/user/score/submit", bytes.NewBuffer(json))
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	handlers.SubmitScore(c)
	assert.Equal(t, 201, w.Code) // or what value you need it to be
}

func TestGetProfile(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/user/profile/5fd6844477c053d581836c33", nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, r := gin.CreateTestContext(w)
	handlers.GetProfile(c)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code) // or what value you need it to be
}
