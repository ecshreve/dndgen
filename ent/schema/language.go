package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Mixin of the Language.
func (Language) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("language_type").StructTag(`json:"type"`).
			Values(
				"STANDARD",
				"EXOTIC",
			).
			Default("STANDARD"),
		field.Enum("script").
			Values(
				"Common",
				"Dwarvish",
				"Elvish",
				"Infernal",
				"Draconic",
				"Celestial",
				"Abyssal",
				"Giant",
				"Gnomish", "Goblin", "Halfling", "Orc", "Other").
			Default("Common"),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("race", Race.Type).
			Ref("languages"),
		edge.From("options", LanguageChoice.Type).
			Ref("languages"),
	}
}
