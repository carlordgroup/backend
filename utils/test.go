package utils

import (
	"bytes"
	"carlord/ent"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

func ToJson(v interface{}) io.Reader {
	if v == nil {
		return nil
	}
	data, _ := json.Marshal(v)
	println(string(data))
	return bytes.NewReader(data)
}

func Request(router *gin.Engine, token string, method string, url string, data any) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, ToJson(data))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func DBInit() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return nil
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.Fatalf("failed creation: %v", err)
		return nil
	}
	return client
}
