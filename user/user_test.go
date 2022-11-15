package user

import (
	"carlord/auth"
	"carlord/ent"
	"carlord/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

type token struct {
	Token string `json:"token"`
}

var tokenContent string

func init() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creation: %v", err)
		return
	}
	router = gin.Default()
	a := auth.New(client)
	a.RegisterRouter(router.Group("/account"))
	New(client).RegisterRouter(router, a)
	register()
}

func register() {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	login := auth.LoginCredential{
		Email:       "doe@doe.doe",
		RawPassword: "aaaaaddddd",
	}
	req, _ = http.NewRequest(http.MethodPost, "/account/register", utils.ToJson(login))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	req, _ = http.NewRequest(http.MethodPost, "/account/login", utils.ToJson(login))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var tok token
	json.Unmarshal(w.Body.Bytes(), &tok)
	tokenContent = tok.Token

}

func TestGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenContent)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPost(t *testing.T) {

	{
		req, _ := http.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)

	}

	{
		req, _ := http.NewRequest(http.MethodPost, "/", utils.ToJson(ent.User{}))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}

	{
		type data struct {
			Address string `json:"address"`
		}

		d := data{Address: "aaabbb"}
		req, _ := http.NewRequest(http.MethodPost, "/", utils.ToJson(d))
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		saved := ent.User{}
		log.Println(w.Body)
		json.Unmarshal(w.Body.Bytes(), &saved)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, d.Address, saved.Address)
	}

}
