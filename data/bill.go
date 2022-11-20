package data

import (
	"carlord/ent"
	"context"
)

const (
	BillUnpaid = "unpaid"
	BillPaid   = "paid"
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

func (b *Bill) Charge() (bool, error) {
	c, err := b.QueryCard().Only(b.ctx)
	if err != nil {
		return false, err
	}
	return NewCard(b.ctx, c).Pay(), nil
}
