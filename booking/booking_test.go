package booking

import (
	"carlord/auth"
	"carlord/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
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
	client.Card.Create().SetNumber("9").SetValidUntil("").SetCardholderName("").SetOwnerID(1).Save(context.Background())
	client.Location.Create().SetName("1").SetLatitude(1).SetLongitude(2).Save(context.Background())
	client.Car.Create().SetBrand("vw").SetCarType("suv").SetUnitPrice(2).SetDeposit(2).SetColor("1").SetYear(1).SetMileage(1).SetPrice(1).SetPlateNumber("").SetPlateCountry("").SetPlateNumber("").SetModel("").SetLocationID(1).Save(context.Background())

	json.Unmarshal(w.Body.Bytes(), &token)

}

func TestBook(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodPost, "/", bookStruct{
		CarID:     1,
		CardID:    1,
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Hour),
	})
	assert.Equal(t, 201, w.Code)
}
