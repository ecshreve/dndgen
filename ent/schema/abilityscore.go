package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AbilityScore holds the schema definition for the AbilityScore entity.
type AbilityScore struct {
	ent.Schema
}

func (AbilityScore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Fields of the AbilityScore.
func (AbilityScore) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Annotations(
			entgql.OrderField("FULL_NAME"),
		),
	}
}

// Edges of the AbilityScore.
func (AbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("skills", Skill.Type),
		edge.From("character_ability_score", CharacterAbilityScore.Type).
			Ref("ability_score").
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		edge.From("classes", Class.Type).
			Ref("saving_throws").
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		edge.From("race", Race.Type).
			Ref("ability_bonuses").
			Through("race_ability_bonuses", AbilityBonus.Type).
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
	}
}
