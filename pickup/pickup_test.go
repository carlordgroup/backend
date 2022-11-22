package pickup

import (
	"carlord/auth"
	"carlord/booking"
	"carlord/ent"
	"carlord/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)
import _ "github.com/mattn/go-sqlite3"

var router *gin.Engine
var client *ent.Client
var token struct {
	Token string `json:"token"`
}

func init() {
	client = utils.DBInit()
	router = gin.Default()
	a := auth.New(client)
	a.RegisterRouter(router.Group("/account"))
	New(client).RegisterRouter(router, a)
	booking.New(client).RegisterRouter(router.Group("/booking"), a)

	utils.Request(router, "", http.MethodPost, "/account/register", map[string]any{"email": "a@b.cc", "password": "112233"})
	w := utils.Request(router, "", http.MethodPost, "/account/login", map[string]any{"email": "a@b.cc", "password": "112233"})
	json.Unmarshal(w.Body.Bytes(), &token)
	client.Account.UpdateOneID(1).SetIsAdmin(true).Save(context.Background())
	client.Card.Create().SetNumber("9").SetValidUntil("").SetCardholderName("").SetOwnerID(1).Save(context.Background())
	client.Location.Create().SetName("1").SetLatitude(1).SetLongitude(2).Save(context.Background())
	client.Car.Create().SetBrand("vw").SetCarType("suv").SetUnitPrice(2).SetDeposit(2).SetColor("1").SetYear(1).SetMileage(1).SetPrice(1).SetPlateNumber("").SetPlateCountry("").SetPlateNumber("").SetModel("").SetLocationID(1).Save(context.Background())
	w = utils.Request(router, token.Token, http.MethodPost, "/booking/", map[string]any{"car_id": 1, "card_id": 1, "start_time": time.Now(), "end_time": time.Now().Add(time.Hour)})
	log.Println(w.Body)

}

func TestPickup(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodPost, "/start/1", map[string]any{"fuel_level": 1, "mileage": 1})
	log.Println(w.Body)
	assert.Equal(t, 200, w.Code)
}
func TestReturn(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodPost, "/finish/1", map[string]any{"fuel_level": 1, "mileage": 2})
	log.Println(client.Billing.Query().AllX(context.Background()))
	assert.Equal(t, 204, w.Code)

}
