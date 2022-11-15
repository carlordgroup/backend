package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name"),
		field.String("last_name"),
		field.String("address"),
		field.String("postal_code"),
		field.String("tel"),
		field.String("driver_license_id"),
		field.String("driver_license_country"),
		field.Time("birthday"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).StructTag(`json:"cards"`),
		edge.To("note_flaws", Flaw.Type).StructTag(`json:"flaws"`),
		edge.From("account", Account.Type).Ref("user").Required(),
	}
}
