package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Billing holds the schema definition for the Billing entity.
type Billing struct {
	ent.Schema
}

// Fields of the Billing.
func (Billing) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Billing.
func (Billing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("booking", Booking.Type).Required(),
		edge.To("card", Card.Type),
	}
}
