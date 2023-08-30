package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AbilityScore holds the schema definition for the AbilityScore entity.
type AbilityScore struct {
	ent.Schema
}

func (AbilityScore) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the AbilityScore.
func (AbilityScore) Fields() []ent.Field {
	return []ent.Field{
		field.String("full_name").Annotations(
			entgql.OrderField("FULL_NAME"),
			entproto.Field(4),
		),
		field.Strings("desc").Annotations(
			entproto.Field(5),
		),
	}
}

// Edges of the AbilityScore.
func (AbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("skills", Skill.Type).
			Ref("ability_score").Annotations(
			entproto.Field(6),
		),
		edge.To("ability_bonuses", AbilityBonus.Type).Annotations(
			entproto.Field(7),
		),
	}
}

// Annotations of the AbilityScore.
func (AbilityScore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

// AbilityBonus holds the schema definition for the AbilityBonus entity.
type AbilityBonus struct {
	ent.Schema
}

// Fields of the AbilityBonus.
func (AbilityBonus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ability_score_id"),
		field.Int("bonus"),
	}
}

// Edges of the AbilityBonus.
func (AbilityBonus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ability_score", AbilityScore.Type).
			Ref("ability_bonuses").
			Unique().
			Required().
			Field("ability_score_id"),
		edge.From("race", Race.Type).
			Ref("ability_bonuses").Unique(),
		edge.From("subrace", Subrace.Type).
			Ref("ability_bonuses").Unique(),
	}
}
