package auth

import (
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

// init the database
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
	s := New(client)
	s.RegisterRouter(router)

}

type testPack struct {
	login  LoginCredential
	result int
}

func TestRegister(t *testing.T) {
	payload := []testPack{
		{
			LoginCredential{
				Email:       "doe@doe.doe",
				RawPassword: "aaaaaddddd",
			}, http.StatusCreated,
		}, {
			LoginCredential{
				Email:       "doe@doe.doe",
				RawPassword: "asdaaasd",
			}, http.StatusBadRequest,
		}, {
			LoginCredential{
				Email: "doe2@doe.doe",
			}, http.StatusBadRequest,
		},
	}

	for _, pack := range payload {
		req, _ := http.NewRequest(http.MethodPost, "/register", utils.ToJson(pack.login))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, pack.result, w.Code)
	}

}

func TestLogin(t *testing.T) {
	payload := []testPack{
		{
			LoginCredential{
				Email:       "doe@doe.doe",
				RawPassword: "aaaaaddddd",
			}, http.StatusOK,
		}, {
			LoginCredential{
				Email:       "doe@doe.doe",
				RawPassword: "asdaaasd",
			}, http.StatusUnauthorized,
		}, {
			LoginCredential{
				Email:       "doe2@doe.doe",
				RawPassword: "asdaaasd",
			}, http.StatusUnauthorized,
		},
	}

	for _, pack := range payload {

		req, _ := http.NewRequest(http.MethodPost, "/login", utils.ToJson(pack.login))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, pack.result, w.Code)
	}
}

func TestSelf(t *testing.T) {
	type token struct {
		Token string `json:"token"`
	}

	req, _ := http.NewRequest(http.MethodPost, "/login", utils.ToJson(LoginCredential{
		Email:       "doe@doe.doe",
		RawPassword: "aaaaaddddd",
	}))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	var tok token
	json.Unmarshal(w.Body.Bytes(), &tok)
	req, _ = http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tok.Token)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	log.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)

}
