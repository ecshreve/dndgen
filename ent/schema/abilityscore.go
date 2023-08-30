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
		edge.From("ability_bonus", AbilityBonus.Type).
			Ref("ability_score").Annotations(
			entproto.Field(7),
		),
	}
}

// Annotations of the AbilityScore.
func (AbilityScore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entproto.Message(),
		entproto.Service(),
	}
}

// AbilityBonus holds the schema definition for the AbilityBonus entity.
type AbilityBonus struct {
	ent.Schema
}

// Fields of the AbilityBonus.
func (AbilityBonus) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value").Annotations(
			entgql.OrderField("VALUE"),
			entproto.Field(2),
		),
	}
}

// Edges of the AbilityBonus.
func (AbilityBonus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability_score", AbilityScore.Type).
			Annotations(
				entproto.Field(3),
			),
	}
}

// Annotations of the AbilityBonus.
func (AbilityBonus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
