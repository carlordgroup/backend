package order

import (
	"carlord/ent/billing"
	"carlord/ent/booking"
	"carlord/ent/car"
	"carlord/ent/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type bookStruct struct {
	CarID     int       `json:"car_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

func (s *service) carAvailable(ctx *gin.Context, carID int, startAt, endAt time.Time) (bool, error) {
	return s.client.Booking.Query().Where(
		booking.And(
			booking.HasCarWith(car.ID(carID)),
			booking.StartAtLTE(endAt),
			booking.EndAtGTE(startAt),
		)).Exist(ctx)
}

func (s *service) bookCar(ctx *gin.Context) (int, any) {
	var book bookStruct
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		return http.StatusBadRequest, err
	}
	userId := ctx.GetInt("id")
	c, err := s.client.Car.Get(ctx, book.CarID)
	if err != nil {
		return http.StatusBadRequest, err
	}
	notOK, err := s.carAvailable(ctx, c.ID, book.StartTime, book.EndTime)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if notOK {
		return http.StatusForbidden, errors.New("the time is occupied")
	}
	unpaid, err := s.client.Billing.Query().Where(billing.HasUserWith(user.ID(userId)), billing.Status("unpaid")).Exist(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if unpaid {
		return http.StatusForbidden, errors.New("cannot start new booking with unpaid bill")
	}
	b, err := s.client.Booking.Create().SetStartAt(book.StartTime).SetEndAt(book.EndTime).SetCarID(book.CarID).SetUserID(userId).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusCreated, b
}

func (s *service) cancelBookingCar(ctx *gin.Context, id int) (int, any) {
	userID := ctx.GetInt("id")
	_, err := s.client.Booking.Delete().Where(booking.And(booking.ID(id), booking.HasUserWith(user.ID(userID)))).Exec(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}
