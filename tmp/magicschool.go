package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MagicSchool holds the schema definition for the MagicSchool entity.
type MagicSchool struct {
	ent.Schema
}

// Fields of the MagicSchool.
func (MagicSchool) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).Unique(),
		field.String("name"),
		field.String("desc"),
	}
}

// Edges of the MagicSchool.
func (MagicSchool) Edges() []ent.Edge {
	return nil
}
