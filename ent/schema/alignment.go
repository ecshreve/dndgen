package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Alignment holds the schema definition for the Alignment entity.
type Alignment struct {
	ent.Schema
}

// Fields of the Alignment.
func (Alignment) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).Unique(),
		field.String("name"),
		field.String("abbr"),
		field.String("desc"),
	}
}

// Edges of the Alignment.
func (Alignment) Edges() []ent.Edge {
	return nil
}
