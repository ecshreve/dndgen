package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Race holds the schema definition for the Race entity.
type Race struct {
	ent.Schema
}

func (Race) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Race.
func (Race) Fields() []ent.Field {
	return []ent.Field{
		field.Int("speed"),
	}
}

// Edges of the Race.
func (Race) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("languages", Language.Type),
		edge.To("ability_bonuses", AbilityBonus.Type),
		edge.To("starting_proficiencies", Proficiency.Type),
	}
}