package data

import (
	"carlord/ent"
	"carlord/ent/booking"
	"context"
)

type Car struct {
	*ent.Car
	ctx context.Context
}

const (
	CarInUse   = "in_use"
	CarIdle    = "idle"
	CarDamaged = "damaged"
)

func NewCar(ctx context.Context, car *ent.Car) *Car {
	return &Car{
		Car: car,
		ctx: ctx,
	}
}

func (c *Car) Available(startAt, endAt int64) (bool, error) {
	result, err := c.QueryBooking().Where(
		booking.And(
			booking.StartAtLTE(endAt),
			booking.EndAtGTE(startAt),
			booking.BookingStatusNEQ(BookingStatusCancel),
			booking.BookingStatusNEQ(BookingStatusFinish),
		)).Exist(c.ctx)
	return !result, err
}

func (c *Car) InUse() error {
	_, err := c.Update().SetStatus(CarInUse).Save(c.ctx)
	return err
}

func (c *Car) Returned() error {
	_, err := c.Update().SetStatus(CarIdle).Save(c.ctx)
	return err
}
func (c *Car) SeriouslyDamaged() error {
	_, err := c.Update().SetStatus(CarDamaged).Save(c.ctx)
	return err
}
