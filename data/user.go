package data

import (
	"carlord/ent"
	"carlord/ent/billing"
	"context"
	"time"
)

type User struct {
	*ent.User
	ctx context.Context
}

func NewUser(ctx context.Context, e *ent.User) *User {
	return &User{User: e, ctx: ctx}
}

func (u *User) Cards() ([]*Card, error) {
	cards, err := u.QueryCard().All(u.ctx)
	if err != nil {
		return nil, err
	}
	return NewCards(u.ctx, cards), nil
}

func (u *User) Book(client *ent.Client, car *Car, startAt time.Time, endAt time.Time, cardID int) (*Booking, error) {
	billCreate := client.Billing.Create().SetUserID(u.ID).SetCardID(cardID)
	b, err := client.Booking.Create().
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetCarID(car.ID).
		SetUserID(u.ID).
		SetBookingStatus(BookingStatusPlan).
		Save(u.ctx)
	billCreate.SetBooking(b)
	_, err = billCreate.Save(u.ctx)
	if err != nil {
		return nil, err
	}
	return NewBooking(u.ctx, b), err
}

func (u *User) HasUnpaidBill() (bool, error) {
	return u.QueryBill().Where(billing.Status(BillUnpaid)).Exist(u.ctx)
}

func (u *User) Bookings() ([]*Booking, error) {
	books, err := u.QueryBooking().All(u.ctx)
	if err != nil {
		return nil, err
	}
	return NewBookings(u.ctx, books), nil
}

func (u *User) Bills() ([]*Bill, error) {
	data, err := u.QueryBill().Where(billing.Status(BillUnpaid)).All(u.ctx)
	if err != nil {
		return nil, err
	}
	return NewBills(u.ctx, data), nil
}
