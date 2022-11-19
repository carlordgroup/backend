package booking

import (
	"carlord/ent"
	"carlord/ent/billing"
	"carlord/ent/booking"
	"carlord/ent/car"
	"carlord/ent/card"
	"carlord/ent/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type bookStruct struct {
	CarID     int       `json:"car_id" binding:"required"`
	CardID    int       `json:"card_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
}

const (
	StatusPlan   = "plan"
	StatusCancel = "cancel"
	StatusInUse  = "in_use"
	StatusFinish = "finish"
)

func (s *service) carNotAvailable(ctx *gin.Context, carID int, startAt, endAt time.Time) (bool, error) {
	return s.client.Booking.Query().Where(
		booking.And(
			booking.HasCarWith(car.ID(carID)),
			booking.And(
				booking.StartAtLTE(endAt),
				booking.EndAtGTE(startAt),
				booking.BookingStatusNEQ(StatusCancel),
				booking.BookingStatusNEQ(StatusFinish),
			),
		)).Exist(ctx)
}

func pay(card *ent.Card) bool {
	time.Sleep(time.Second * 5)
	return !strings.HasSuffix(card.Number, "00")
}

// Add booking godoc
// @Tags booking
// @Param request body bookStruct true "booking info"
// @Summary add a booking
// @Accept json
// @Produce json
// @Success 200 {object} ent.Booking
// @Router /booking/ [post]
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
	notOK, err := s.carNotAvailable(ctx, c.ID, book.StartTime, book.EndTime)
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

	b, err := s.client.Booking.Create().
		SetStartAt(book.StartTime).
		SetEndAt(book.EndTime).
		SetCarID(book.CarID).
		SetUserID(userId).
		SetBookingStatus(StatusPlan).
		Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}

	userCard, err := s.client.Card.Query().Where(card.ID(book.CardID), card.HasOwnerWith(user.ID(userId))).Only(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}

	if !pay(userCard) {
		_, err = b.Update().SetBookingStatus(StatusCancel).Save(ctx)
		if err != nil {
			return http.StatusBadRequest, err
		}
		return http.StatusForbidden, errors.New("payment is rejected")
	}

	return http.StatusCreated, b
}

// Cancel booking godoc
// @Tags booking
// @Summary remove a booking
// @Produce json
// @Success 204
// @Router /booking/:id [delete]
func (s *service) cancelBookingCar(ctx *gin.Context, id int) (int, any) {
	userID := ctx.GetInt("id")
	_, err := s.client.Booking.Update().Where(booking.And(booking.ID(id), booking.HasUserWith(user.ID(userID)))).SetBookingStatus(StatusCancel).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}

// List booking godoc
// @Tags booking
// @Summary list booking
// @Produce json
// @Success 200 {object} []ent.Booking
// @Router /booking/ [get]
func (s *service) listBooking(ctx *gin.Context) (int, any) {
	userID := ctx.GetInt("id")
	data, err := s.client.Booking.Query().Where(booking.HasUserWith(user.ID(userID))).All(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, data
}
