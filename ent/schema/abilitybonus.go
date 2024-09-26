package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AbilityBonus holds the schema definition for the AbilityBonus entity.
type AbilityBonus struct {
	ent.Schema
}

// Annotations of the AbilityBonus.
func (AbilityBonus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("race_id", "ability_score_id"),
	}
}

// Fields of the AbilityBonus.
func (AbilityBonus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("bonus").Positive(),
		field.Int("race_id"),
		field.Int("ability_score_id"),
	}
}

// Edges of the AbilityBonus.
func (AbilityBonus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("race", Race.Type).
			Unique().
			Required().
			Field("race_id"),
		edge.To("ability_score", AbilityScore.Type).
			Unique().
			Required().
			Field("ability_score_id"),
	}
}
