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
		field.Time("return_car_at").Optional(),
		field.Float32("rate").Default(5),
		field.Float32("exceed_rate").Default(10),
		field.Float32("deposit").Default(0),
		field.Float32("fuel_level_at_begin").Optional(),
		field.Float32("fuel_level_at_end").Optional(),
		field.Int("mileage_begin").Optional(),
		field.Int("mileage_end").Optional(),
		field.String("booking_status").Default("plan"),
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
