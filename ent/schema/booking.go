package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_at"),
		field.Time("end_at"),
		field.Time("return_car_at"),
		field.Float32("fuel_level_at_begin"),
		field.Float32("fuel_level_at_end"),
		field.Int("mileage_begin"),
		field.Int("mileage_end"),
		field.String("booking_status"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required(),
		edge.To("car", Car.Type).Required(),
		edge.From("billing", Billing.Type).Ref("booking"),
	}
}
