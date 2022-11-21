package pickup

import (
	"carlord/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

type bookingData struct {
	FuelLevel float32 `json:"fuel_level"`
	Mileage   int     `json:"mileage"`
}

// realizeBooking godoc
// @Tags pickup
// @Param request body bookingData true "updated info"
// @Summary run the booking
// @Accept json
// @Produce json
// @Success 200 {object} ent.Booking
// @Router /pickup/start/:id [post]
func (s *service) realizeBooking(ctx *gin.Context, id int) (int, any) {
	var d bookingData
	err := ctx.ShouldBindJSON(&d)
	if err != nil {
		return http.StatusBadRequest, err
	}
	book, err := s.client.Booking.Get(ctx, id)
	if err != nil {
		return http.StatusBadRequest, err
	}

	err = data.NewAdmin(ctx).RealizeBooking(data.NewBooking(ctx, book), d.FuelLevel, d.Mileage)
	if err != nil {
		return http.StatusBadRequest, nil
	}
	return http.StatusOK, book
}

// returnCar godoc
// @Tags pickup
// @Summary end the booking
// @Param request body bookingData true "updated info"
// @Accept json
// @Success 204
// @Router /pickup/finish/:id [post]
func (s *service) returnCar(ctx *gin.Context, id int) (int, any) {
	var d bookingData
	err := ctx.ShouldBindJSON(&d)
	if err != nil {
		return http.StatusBadRequest, err
	}

	book, err := s.client.Booking.Get(ctx, id)
	if err != nil {
		return http.StatusBadRequest, err
	}
	bookingObj := data.NewBooking(ctx, book)
	err = data.NewAdmin(ctx).ReturnACar(bookingObj, d.FuelLevel, d.Mileage)
	if err != nil {
		return http.StatusBadRequest, nil
	}
	bill, err := bookingObj.Bill()
	if err != nil {
		return http.StatusBadRequest, err
	}
	err = bill.Calculate()
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}
