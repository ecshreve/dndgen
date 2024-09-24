package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AbilityBonus holds the schema definition for the AbilityBonus entity.
type AbilityBonus struct {
	ent.Schema
}

// Fields of the AbilityBonus.
func (AbilityBonus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("bonus").Positive(),
	}
}

// Edges of the AbilityBonus.
func (AbilityBonus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability_score", AbilityScore.Type).
			Required().
			Unique(),
		edge.From("race", Race.Type).
			Ref("ability_bonuses"),
		edge.From("options", AbilityBonusChoice.Type).
			Ref("ability_bonuses"),
		edge.From("subrace", Subrace.Type).
			Ref("ability_bonuses"),
	}
}
