package schema

import "entgo.io/ent"

// Choice holds the schema definition for the Choice entity.
type Choice struct {
	ent.Schema
}

// Fields of the Choice.
func (Choice) Fields() []ent.Field {
	return nil
}

// Edges of the Choice.
func (Choice) Edges() []ent.Edge {
	return nil
}
