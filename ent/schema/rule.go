package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Rule holds the schema definition for the Rule entity.
type Rule struct {
	ent.Schema
}

// Mixin implements the ent.Mixin for the Rule.
func (Rule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Edges of the Rule.
func (Rule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sections", RuleSection.Type).
			StorageKey(edge.Column("rule_id")),
	}
}
