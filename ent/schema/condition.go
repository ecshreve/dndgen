package schema

import "entgo.io/ent"

// Condition holds the schema definition for the Condition entity.
type Condition struct {
	ent.Schema
}

// Mixin implements the ent.Mixin for the Condition.
func (Condition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}
