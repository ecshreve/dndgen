package schema

import (
	"entgo.io/contrib/entgql"
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
		field.String("full_name"),
		field.Strings("desc"),
	}
}

// Edges of the AbilityScore.
func (AbilityScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classes", Class.Type).
			Ref("saving_throws").
			Annotations(
				entgql.Skip(),
			),
		edge.From("skills", Skill.Type).
			Ref("ability_score"),
	}
}

// Annotations of the AbilityScore.
func (AbilityScore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
