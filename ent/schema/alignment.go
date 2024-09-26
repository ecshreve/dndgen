package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Alignment holds the schema definition for the Alignment entity.
type Alignment struct {
	ent.Schema
}

// Mixin returns the Alignment mixed-in schema.
func (Alignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Fields of the Alignment.
func (Alignment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("abbr").Values(
			"LG", // Lawful Good
			"NG", // Neutral Good
			"CG", // Chaotic Good
			"LN", // Lawful Neutral
			"N",  // Neutral
			"CN", // Chaotic Neutral
			"LE", // Lawful Evil
			"NE", // Neutral Evil
			"CE", // Chaotic Evil
		).StructTag(`json:"abbreviation"`),
	}
}
