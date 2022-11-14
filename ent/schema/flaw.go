package schema

import "entgo.io/ent"

// Flaw holds the schema definition for the Flaw entity.
type Flaw struct {
	ent.Schema
}

// Fields of the Flaw.
func (Flaw) Fields() []ent.Field {
	return nil
}

// Edges of the Flaw.
func (Flaw) Edges() []ent.Edge {
	return nil
}
