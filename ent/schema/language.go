package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Mixin of the Language.
func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").StructTag(`json:"desc,omitempty"`),
		field.Enum("language_type").
			Values("STANDARD", "EXOTIC").Default("STANDARD").
			StructTag(`json:"type"`),
		field.Enum("script").Optional().
			Values("Common", "Dwarvish", "Elvish", "Infernal", "Draconic", "Celestial", "Abyssal", "Giant", "Gnomish", "Goblin", "Halfling", "Orc", "Other").Default("Common"),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
}

// Annotations of the Language.
func (Language) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
