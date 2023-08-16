package schema

import (
	"entgo.io/ent"
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
		edge.From("ability_scores", AbilityScore.Type).
			Ref("skills"),
	}
}
