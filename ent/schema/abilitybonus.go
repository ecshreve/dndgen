package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Int("bonus").Annotations(
			entgql.OrderField("BONUS"),
		),
	}
}

// Edges of the AbilityBonus.
func (AbilityBonus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability_score", AbilityScore.Type),
		edge.From("race", Race.Type).
			Ref("ability_bonuses"),
	}
}

// Annotations of the AbilityBonus.
func (AbilityBonus) Annotations() []schema.Annotation {
	return nil
}
