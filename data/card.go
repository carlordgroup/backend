package data

import (
	"carlord/ent"
	"context"
	"strings"
)

type Card struct {
	*ent.Card
	ctx context.Context
}

func NewCard(ctx context.Context, card *ent.Card) *Card {
	return &Card{
		Card: card,
		ctx:  ctx,
	}
}

func NewCards(ctx context.Context, data []*ent.Card) []*Card {
	result := make([]*Card, len(data))
	for i, d := range data {
		result[i] = NewCard(ctx, d)
	}
	return result
}

// Pay a fake payment function
func (c *Card) Pay() bool {
	return !strings.HasSuffix(c.Number, "00")
}

func (c *Card) Verify() bool {
	return !strings.HasSuffix(c.Number, "01")
}
