package schema

import "entgo.io/ent"

// Skill holds the schema definition for the Skill entity.
type Skill struct {
	ent.Schema
}

func (Skill) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Skill.
func (Skill) Fields() []ent.Field {
	return nil
}

// Edges of the Skill.
func (Skill) Edges() []ent.Edge {
	return nil
}
