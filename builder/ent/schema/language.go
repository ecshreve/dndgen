package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("STANDARD", "EXOTIC"),
		field.Enum("script").Values("Common", "Dwarvish", "Elvish", "Giant", "Gnomish", "Goblin", "Halfling", "Orc", "Abyssal", "Celestial", "Draconic", "Drow Sign Language", "Infernal", "Primordial", "Sylvan", "Undercommon", "Other"),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
}
