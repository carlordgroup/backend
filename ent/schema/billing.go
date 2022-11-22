package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Billing holds the schema definition for the Billing entity.
type Billing struct {
	ent.Schema
}

// Fields of the Billing.
func (Billing) Fields() []ent.Field {
	return []ent.Field{
		field.String("status").Default("unpaid"),
		field.Float32("basic_cost").Default(0),
		field.Float32("fuel_cost").Default(0),
		field.Float32("compensation").Default(0),
		field.Float32("deposit").Default(0),
	}
}

// Edges of the Billing.
func (Billing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("booking", Booking.Type).Unique(),
		edge.To("card", Card.Type).Unique(),
		edge.To("user", User.Type).Unique(),
	}
}
