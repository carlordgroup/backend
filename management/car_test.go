package management

import (
	"carlord/ent"
	"carlord/utils"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)
import _ "github.com/mattn/go-sqlite3"

var carObj = carData{ent.Car{
	Color:        "red",
	Brand:        "vw",
	Model:        "golf",
	Year:         2021,
	CarType:      "normal",
	PlateNumber:  "A2292",
	PlateCountry: "CA",
	UnitPrice:    10,
	Price:        20000,
	Mileage:      55,
	Deposit:      5000,
}, 0}

func addALocation() {
	w := utils.Request(router, token.Token, http.MethodPost, "/location/", map[string]any{"name": "1", "latitude": 1.1, "longitude": 2.2})
	utils.Request(router, token.Token, http.MethodPost, "/location/", map[string]any{"name": "1", "latitude": 1.1, "longitude": 2.2})
	var data map[string]any
	json.Unmarshal(w.Body.Bytes(), &data)
	carObj.LocationID = int(data["id"].(float64))
}

func TestAddCar(t *testing.T) {
	addALocation()
	w := utils.Request(router, token.Token, http.MethodPost, "/car/", carObj)
	log.Println(w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
	var data map[string]any
	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, "red", data["color"])
	assert.Equal(t, 10.0, data["unit_price"])
	assert.Equal(t, "A2292", data["plate_number"])

}

func TestEditCar(t *testing.T) {
	car2 := carObj
	car2.Color = "white"
	w := utils.Request(router, token.Token, http.MethodPost, "/car/1", car2)
	var data map[string]any
	json.Unmarshal(w.Body.Bytes(), &data)
	log.Println(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "white", data["color"])

}

func TestGetCar(t *testing.T) {
	r, _ := http.NewRequest("GET", "/car/", nil)
	q := r.URL.Query()
	q.Add("color", "white")
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	var data []any

	json.Unmarshal(w.Body.Bytes(), &data)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, len(data))
}

func TestDeleteCar(t *testing.T) {
	w := utils.Request(router, token.Token, http.MethodDelete, "/car/1", nil)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
