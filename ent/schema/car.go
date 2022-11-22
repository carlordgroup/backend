package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("color"),
		field.String("brand"),
		field.String("model"),
		field.Int("year"),
		field.String("status").Default("idle"),
		field.String("car_type"),
		field.String("plate_number"),
		field.String("plate_country"),
		field.Float32("unit_price"),
		field.Float32("price"),
		field.Int("mileage"),
		field.Float32("deposit"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("location", Location.Type).Unique(),
		edge.From("booking", Booking.Type).Ref("car"),
	}
}
