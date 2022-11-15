package card

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
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
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
		Email:       strconv.Itoa(rand.Int()) + "doe@doe.doe",
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

func TestAdd(t *testing.T) {
	type pair struct {
		body interface{}
		code int
	}

	pairs := []pair{
		{nil, http.StatusBadRequest},
		{ent.Card{
			Number:         "",
			CardholderName: "",
			ValidUntil:     "",
		}, http.StatusBadRequest}, {ent.Card{
			Number:         "aaaaaa",
			CardholderName: "aa",
			ValidUntil:     "aa",
		}, http.StatusBadRequest}, {ent.Card{
			Number:         "213123123",
			CardholderName: "aa",
			ValidUntil:     "",
		}, http.StatusBadRequest}, {ent.Card{
			Number:         "213123123",
			CardholderName: "aa",
		}, http.StatusBadRequest}, {ent.Card{
			Number:         "123456",
			CardholderName: "abc",
			ValidUntil:     "11/11",
		}, http.StatusCreated}, {ent.Card{
			Number:         "1234567",
			CardholderName: "abcd",
			ValidUntil:     "11/11",
		}, http.StatusCreated},
	}

	for _, p := range pairs {
		req, _ := http.NewRequest(http.MethodPost, "/", utils.ToJson(p.body))
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		log.Println(w.Body)
		assert.Equal(t, p.code, w.Code)
	}

	{
		// second account
		register()
		req, _ := http.NewRequest(http.MethodPost, "/", utils.ToJson(ent.Card{
			Number:         "123456788888",
			CardholderName: "abcd",
			ValidUntil:     "11/11",
		}))
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		log.Println(w.Body)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

}

func TestGet(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenContent)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := []*ent.Card{}
	log.Println(w.Body)
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "123456788888", result[0].Number)

}

func TestEdit(t *testing.T) {
	{
		req, _ := http.NewRequest(http.MethodPost, "/3", utils.ToJson(ent.Card{
			Number:         "999",
			CardholderName: "abcd",
			ValidUntil:     "11/11",
		}))
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		result := ent.Card{}
		log.Println(w.Body)
		json.Unmarshal(w.Body.Bytes(), &result)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "999", result.Number)
	}
	{
		req, _ := http.NewRequest(http.MethodPost, "/2", utils.ToJson(ent.Card{
			Number:         "999",
			CardholderName: "abcd",
			ValidUntil:     "11/11",
		}))
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)

	}
}

func TestDelete(t *testing.T) {
	{
		req, _ := http.NewRequest(http.MethodDelete, "/2", nil)
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
	}
	{
		req, _ := http.NewRequest(http.MethodDelete, "/3", nil)
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	}
	{
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+tokenContent)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		result := []*ent.Card{}
		log.Println(w.Body)
		json.Unmarshal(w.Body.Bytes(), &result)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, 0, len(result))
	}
}
