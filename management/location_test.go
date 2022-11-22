package management

import (
	"carlord/auth"
	"carlord/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)
import _ "github.com/mattn/go-sqlite3"

var router *gin.Engine
var token struct {
	Token string `json:"token"`
}

func init() {
	client := utils.DBInit()
	router = gin.Default()
	a := auth.New(client)
	a.RegisterRouter(router.Group("/account"))
	New(client).RegisterRouter(router, a)

	utils.Request(router, "", http.MethodPost, "/account/register", map[string]any{"email": "a@b.cc", "password": "112233"})
	w := utils.Request(router, "", http.MethodPost, "/account/login", map[string]any{"email": "a@b.cc", "password": "112233"})
	client.Account.UpdateOneID(1).SetIsAdmin(true).Save(context.Background())

	json.Unmarshal(w.Body.Bytes(), &token)

}

func TestAdd(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodPost, "/location/", map[string]any{"name": "1", "latitude": 1.1, "longitude": 2.2})
	utils.Request(router, token.Token, http.MethodPost, "/location/", map[string]any{"name": "1", "latitude": 1.1, "longitude": 2.2})
	log.Println(w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
	var data map[string]any
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, "1", data["name"])
	assert.Equal(t, 1.1, data["latitude"])
	assert.Equal(t, 2.2, data["longitude"])

}
func TestEdit(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodPost, "/location/1", map[string]any{"name": "2", "latitude": 1.1, "longitude": 2.2})
	var data map[string]any
	json.Unmarshal(w.Body.Bytes(), &data)
	log.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", data["name"])

}

func TestGet(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodGet, "/location/", nil)
	var data []any
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(data))
}

func TestDelete(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodDelete, "/location/1", nil)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
