package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Alignment holds the schema definition for the Alignment entity.
type Alignment struct {
	ent.Schema
}

func (Alignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Alignment.
func (Alignment) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
		field.String("abbr").StructTag(`json:"abbreviation"`),
	}
}

// Edges of the Alignment.
func (Alignment) Edges() []ent.Edge {
	return nil
}
