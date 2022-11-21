package management

import (
	"carlord/data"
	"carlord/ent"
	"carlord/ent/account"
	"carlord/ent/booking"
	"carlord/ent/card"
	"carlord/ent/user"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *service) createUser(ctx *gin.Context) (int, any) {
	var u ent.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		return http.StatusBadRequest, err
	}

	user, err := s.client.User.Create().
		SetAddress(u.Address).
		SetBirthday(u.Birthday).
		SetDriverLicenseCountry(u.DriverLicenseCountry).
		SetDriverLicenseID(u.DriverLicenseID).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetPostalCode(u.PostalCode).
		SetTel(u.Tel).
		Save(ctx)
	if err != nil {
		return http.StatusBadRequest, nil
	}
	return http.StatusOK, user
}

type cardBinding struct {
	Number         string `json:"number" binding:"numeric,required"`
	CardholderName string `json:"cardholder_name" binding:"required"`
	ValidUntil     string `json:"valid_until" binding:"required"`
}

func (s *service) createCard(ctx *gin.Context, id int) (int, any) {

	var cardInfo cardBinding
	err := ctx.ShouldBindJSON(&cardInfo)
	if err != nil {
		return http.StatusBadRequest, err
	}

	card, err := s.client.Card.Create().SetNumber(cardInfo.Number).SetCardholderName(cardInfo.CardholderName).SetValidUntil(cardInfo.ValidUntil).SetOwnerID(id).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, card

}

type bookStruct struct {
	CarID     int       `json:"car_id" binding:"required"`
	CardID    int       `json:"card_id" binding:"required"`
	UserID    int       `json:"user_id" binding:"required"`
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
	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		return http.StatusBadRequest, err
	}
	accountID := ctx.GetInt("id")
	c, err := s.client.Car.Get(ctx, book.CarID)
	if err != nil {
		return http.StatusBadRequest, err
	}
	carObj := data.NewCar(ctx, c)
	available, err := carObj.Available(book.StartTime, book.EndTime)
	if err != nil {
		return http.StatusBadRequest, err
	}
	if !available {
		return http.StatusForbidden, errors.New("the time is occupied")
	}
	u, err := s.client.User.Get(ctx, book.UserID)
	if err != nil {
		return http.StatusBadRequest, nil
	}
	userObj := data.NewUser(ctx, u)
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

	userCard, err := s.client.Card.Query().Where(card.ID(book.CardID), card.HasOwnerWith(user.HasAccountWith(account.ID(accountID)))).Only(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	cardObj := data.NewCard(ctx, userCard)
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
	_, err := s.client.Booking.Update().
		Where(booking.And(booking.ID(id))).
		SetBookingStatus(data.BookingStatusCancel).Save(ctx)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusNoContent, nil
}
