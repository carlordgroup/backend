package booking

import (
	"carlord/data"
	"carlord/ent/account"
	"carlord/ent/booking"
	"carlord/ent/card"
	"carlord/ent/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type bookStruct struct {
	CarID     int       `json:"car_id" binding:"required"`
	CardID    int       `json:"card_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
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
	// data binding
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		return http.StatusBadRequest, err
	}

	accountId := ctx.GetInt("id")

	c, err := s.client.Car.Get(ctx, book.CarID)
	if err != nil {
		return http.StatusBadRequest, err
	}
	carObj := data.NewCar(ctx, c)
	// see if car is available at that time
	available, err := carObj.Available(book.StartTime, book.EndTime)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if !available {
		return http.StatusForbidden, errors.New("the time is occupied")
	}
	accountObj := ctx.MustGet("account").(*data.Account)
	userObj, err := accountObj.User()
	if err != nil {
		return http.StatusBadRequest, nil
	}

	// reject if user has any unpaid bill
	unpaid, err := userObj.HasUnpaidBill()
	if err != nil {
		return http.StatusBadRequest, err
	}
	if unpaid {
		return http.StatusForbidden, errors.New("cannot start new booking with unpaid bill")
	}

	b, err := userObj.Book(s.client, carObj, book.StartTime, book.EndTime, book.CardID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// check if the user do have that card
	userCard, err := s.client.Card.Query().Where(card.ID(book.CardID), card.HasOwnerWith(user.HasAccountWith(account.ID(accountId)))).Only(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	cardObj := data.NewCard(ctx, userCard)

	// try to pay
	if !cardObj.Pay() {
		err = b.UpdateStatus(data.BookingStatusCancel)
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
	accountID := ctx.GetInt("id")
	// query the booking under the user and cancel it
	_, err := s.client.Booking.Update().
		Where(booking.And(booking.ID(id), booking.HasUserWith(user.HasAccountWith(account.ID(accountID))))).
		SetBookingStatus(data.BookingStatusCancel).Save(ctx)
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
	acc := ctx.MustGet("account").(*data.Account)
	u, err := acc.User()
	if err != nil {
		return http.StatusBadRequest, err
	}
	b, err := u.Bookings()
	return http.StatusOK, b
}
