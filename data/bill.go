package data

import (
	"carlord/ent"
	"context"
	"time"
)

const (
	BillUnpaid = "unpaid"
	BillPaid   = "paid"
)

const (
	Rate = 5
)

type Bill struct {
	*ent.Billing
	ctx context.Context
}

func NewBill(ctx context.Context, bill *ent.Billing) *Bill {
	return &Bill{
		Billing: bill,
		ctx:     ctx,
	}
}

func NewBills(ctx context.Context, data []*ent.Billing) []*Bill {
	result := make([]*Bill, len(data))
	for i, d := range data {
		result[i] = NewBill(ctx, d)
	}
	return result
}

// Charge from user card
func (b *Bill) Charge() (bool, error) {
	c, err := b.QueryCard().Only(b.ctx)
	if err != nil {
		return false, err
	}
	result := NewCard(b.ctx, c).Pay()
	if !result {
		return false, nil
	}
	b.Billing, err = b.Update().SetStatus(BillPaid).Save(b.ctx)
	return true, nil
}

// Calculate the bill calculation
func (b *Bill) Calculate() error {
	book, err := b.QueryBooking().Only(b.ctx)
	if err != nil {
		return err
	}
	basic := float32(time.Unix(book.EndAt-book.StartAt, 0).Hour()) * book.Rate
	var compensation float32
	if time.Unix(book.ReturnCarAt-book.EndAt, 0).Hour() > 0 {
		compensation = float32(time.Unix(book.ReturnCarAt-book.EndAt, 0).Hour()) * book.ExceedRate
	}
	_, err = b.Update().SetBasicCost(basic).SetCompensation(compensation).SetStatus(BillUnpaid).SetDeposit(book.Deposit).Save(b.ctx)
	return err
}
