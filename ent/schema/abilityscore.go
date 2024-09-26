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
		edge.From("ability_bonuses", AbilityBonus.Type).
			Ref("ability_score"),
		edge.From("classes", Class.Type).
			Ref("saving_throws"),
	}
}

// CharacterAbilityScore holds the schema definition for the CharacterAbilityScore entity.
type CharacterAbilityScore struct {
	ent.Schema
}

func (CharacterAbilityScore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("score").Positive(),
	}
}

func (CharacterAbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("character", Character.Type).
			Unique(),
		edge.To("ability_score", AbilityScore.Type).
			Required().
			Unique(),
	}
}
