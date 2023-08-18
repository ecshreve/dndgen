package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

func (Skill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return nil
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability_score", AbilityScore.Type).Unique(),
		edge.From("proficiencies", Proficiency.Type).
			Ref("skill"),
	}
}

// Annotations of the Skill.
func (Skill) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
