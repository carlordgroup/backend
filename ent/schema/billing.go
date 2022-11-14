package schema

import "entgo.io/ent"

// Billing holds the schema definition for the Billing entity.
type Billing struct {
	ent.Schema
}

// Fields of the Billing.
func (Billing) Fields() []ent.Field {
	return nil
}

// Edges of the Billing.
func (Billing) Edges() []ent.Edge {
	return nil
}
