package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// RuleSection holds the schema definition for the RuleSection entity.
type RuleSection struct {
	ent.Schema
}

// Mixin implements the ent.Mixin for the RuleSection.
func (RuleSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Edges of the RuleSection.
func (RuleSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rule", Rule.Type).
			Ref("sections").
			Unique(),
	}
}
