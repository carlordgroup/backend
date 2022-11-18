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
		field.Time("return_car_at").Nillable(),
		field.Float32("fuel_level_at_begin").Nillable(),
		field.Float32("fuel_level_at_end").Nillable(),
		field.Int("mileage_begin").Nillable(),
		field.Int("mileage_end").Nillable(),
		field.String("booking_status"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required().Unique(),
		edge.To("car", Car.Type).Required().Unique(),
		edge.From("billing", Billing.Type).Ref("booking").Unique(),
	}
}
