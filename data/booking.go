package data

import (
	"carlord/ent"
	"context"
)

const (
	BookingStatusPlan   = "plan"
	BookingStatusCancel = "cancel"
	BookingStatusInUse  = "in_use"
	BookingStatusFinish = "finish"
)

type Booking struct {
	*ent.Booking
	ctx context.Context
}

func NewBooking(ctx context.Context, booking *ent.Booking) *Booking {
	return &Booking{
		Booking: booking,
		ctx:     ctx,
	}
}

func NewBookings(ctx context.Context, data []*ent.Booking) []*Booking {
	result := make([]*Booking, len(data))
	for i, d := range data {
		result[i] = NewBooking(ctx, d)
	}
	return result
}

func (b *Booking) Bill() (*Bill, error) {
	bill, err := b.QueryBilling().Only(b.ctx)
	return NewBill(b.ctx, bill), err
}

func (b *Booking) UpdateStatus(status string) (err error) {
	_, err = b.Update().SetBookingStatus(status).Save(b.ctx)
	return err
}
