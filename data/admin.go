package data

import "context"

type Admin struct {
	ctx context.Context
}

func NewAdmin(ctx context.Context) *Admin {
	return &Admin{ctx}
}

func (a *Admin) RealizeBooking(booking *Booking, fuelLevelAtBegin float32, mileageBegin int) (err error) {
	_, err = booking.Update().SetBookingStatus(BookingStatusInUse).SetFuelLevelAtBegin(fuelLevelAtBegin).SetMileageBegin(mileageBegin).Save(a.ctx)
	return err
}

func (a *Admin) ReturnACar(booking *Booking, fuelLevelAtEnd float32, mileageEnd int) (err error) {
	_, err = booking.Update().SetBookingStatus(BookingStatusFinish).SetFuelLevelAtBegin(fuelLevelAtEnd).SetMileageBegin(mileageEnd).Save(a.ctx)
	if err != nil {
		return err
	}
	b, err := booking.Bill()
	if err != nil {
		return err
	}
	car, err := booking.QueryCar().Only(a.ctx)
	if err != nil {
		return err
	}
	_, err = car.Update().SetStatus(CarIdle).Save(a.ctx)
	if err != nil {
		return err
	}
	return b.Calculate()

}
