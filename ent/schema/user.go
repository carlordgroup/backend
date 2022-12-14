package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").Default(""),
		field.String("last_name").Default(""),
		field.String("address").Default(""),
		field.String("postal_code").Default(""),
		field.String("tel").Default(""),
		field.String("driver_license_id").Default(""),
		field.String("driver_license_country").Default(""),
		field.Int64("birthday").DefaultFunc(func() int64 { return time.Now().Unix() }),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).StructTag(`json:"cards"`),
		edge.From("account", Account.Type).Ref("user").Unique(),
		edge.From("booking", Booking.Type).Ref("user"),
		edge.From("bill", Billing.Type).Ref("user"),
	}
}
