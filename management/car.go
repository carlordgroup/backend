package management

import (
	"carlord/ent"
	"carlord/ent/car"
	"carlord/ent/location"
	"github.com/gin-gonic/gin"
	"net/http"
)

type carData struct {
	ent.Car
	LocationID int `json:"location_id"`
}

// Add Car godoc
// @Tags car
// @Param request body ent.Car true "car info"
// @Summary add a car
// @Accept json
// @Produce json
// @Success 201 {object} ent.Car
// @Router /management/car [post]
func (s *service) addCar(ctx *gin.Context) (int, any) {
	var data carData
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	newData, err := s.client.Car.Create().
		SetColor(data.Color).
		SetBrand(data.Brand).
		SetModel(data.Model).
		SetYear(data.Year).
		SetStatus(data.Status).
		SetCarType(data.CarType).
		SetDeposit(data.Deposit).
		SetMileage(data.Mileage).
		SetPrice(data.Price).
		SetUnitPrice(data.UnitPrice).
		SetPlateNumber(data.PlateNumber).
		SetPlateCountry(data.PlateCountry).
		SetLocationID(data.LocationID).
		Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, newData

}

// Update Car godoc
// @Tags car
// @Param request body ent.Car true "car info"
// @Summary update a car
// @Accept json
// @Produce json
// @Success 200 {object} ent.Car
// @Router /management/car/:id [post]
func (s *service) updateCar(ctx *gin.Context, id int) (int, any) {
	var data carData
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	newData, err := s.client.Car.UpdateOneID(id).
		SetColor(data.Color).
		SetBrand(data.Brand).
		SetModel(data.Model).
		SetYear(data.Year).
		SetStatus(data.Status).
		SetCarType(data.CarType).
		SetDeposit(data.Deposit).
		SetMileage(data.Mileage).
		SetPrice(data.Price).
		SetUnitPrice(data.UnitPrice).
		SetPlateNumber(data.PlateNumber).
		SetPlateCountry(data.PlateCountry).
		SetLocationID(data.LocationID).
		Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, newData

}

// Update Car godoc
// @Tags car
// @Summary update a car
// @Accept json
// @Success 204
// @Router /management/car/:id [delete]
func (s *service) deleteCar(ctx *gin.Context, id int) (int, any) {

	err := s.client.Car.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}

type queryCar struct {
	CarType  *string `query:"car_type"`
	Color    *string `query:"color"`
	Brand    *string `query:"brand"`
	Model    *string `query:"model"`
	Location *string `query:"location"`
}

// Query Car godoc
// @Tags car
// @Param request query queryCar true "car info"
// @Summary query cars
// @Accept json
// @Produce json
// @Success 204
// @Router /management/car/ [get]
func (s *service) filterCar(ctx *gin.Context) (int, any) {

	var query queryCar
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		return http.StatusBadRequest, err
	}

	q := s.client.Car.Query()
	if query.CarType != nil {
		q = q.Where(car.CarType(*query.CarType))
	}
	if query.Color != nil {
		q = q.Where(car.Color(*query.Color))
	}
	if query.Brand != nil {
		q = q.Where(car.Brand(*query.Brand))
	}
	if query.Model != nil {
		q = q.Where(car.Model(*query.Model))
	}
	if query.Location != nil {
		q = q.Where(car.HasLocationWith(location.Name(*query.Location)))
	}

	cars, err := q.All(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, cars
}
