package data

import (
	"carlord/ent"
	"carlord/ent/billing"
	"carlord/ent/booking"
	"context"
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

func (u *User) Book(client *ent.Client, car *Car, startAt int64, endAt int64, cardID int) (*Booking, error) {
	b, err := client.Booking.Create().
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetCarID(car.ID).
		SetUserID(u.ID).
		SetBookingStatus(BookingStatusPlan).
		Save(u.ctx)
	_, err = client.Billing.Create().SetUserID(u.ID).SetCardID(cardID).SetBookingID(b.ID).Save(u.ctx)
	if err != nil {
		return nil, err
	}
	b, err = client.Booking.Query().Where(booking.ID(b.ID)).WithCar().WithBilling().Only(u.ctx)
	if err != nil {
		return nil, err
	}

	return NewBooking(u.ctx, b), err
}

func (u *User) HasUnpaidBill() (bool, error) {
	return u.QueryBill().Where(billing.Status(BillUnpaid)).Exist(u.ctx)
}

func (u *User) Bookings() ([]*Booking, error) {
	books, err := u.QueryBooking().WithBilling().WithCar().All(u.ctx)
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
