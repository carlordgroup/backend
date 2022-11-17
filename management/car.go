package management

import (
	"carlord/ent"
	"carlord/ent/car"
	"carlord/ent/location"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *service) addCar(ctx *gin.Context) (int, any) {
	var data ent.Car
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
		Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusCreated, newData

}

func (s *service) updateCar(ctx *gin.Context, id int) (int, any) {
	var data ent.Car
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
		Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, newData

}

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
