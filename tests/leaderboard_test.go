package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/sedat/leaderboard-api/handlers"
)

func TestLeaderboardLimit(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/leaderboard/limit/10", nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{"limit", "10"})
	handlers.GetLeaderboardWithLimit(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be
}

func TestLeaderboardByCountry(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8000/leaderboard/country/TR", nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	handlers.GetLeaderboardByCountry(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be
}

func TestLeaderboardByCountryWithLimit(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8000/leaderboard/country/TR", nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{"limit", "10"})
	handlers.GetLeaderboardByCountryWithLimit(c)
	assert.Equal(t, 200, w.Code)
}

func TestLeaderboardInRange(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/leaderboard/range/0/10", nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err.Error())
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = append(c.Params, gin.Param{"start", "0"})
	c.Params = append(c.Params, gin.Param{"end", "10"})
	handlers.GetLeaderboardInRange(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be
}
