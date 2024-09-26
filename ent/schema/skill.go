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
		BaseGQLMixin{},
	}
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ability_score", AbilityScore.Type).
			Ref("skills").
			Unique(),
		edge.From("characters", Character.Type).
			Ref("skills").
			Through("character_skills", CharacterSkill.Type),
	}
}
