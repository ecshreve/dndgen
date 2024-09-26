package schema

import (
	"entgo.io/ent"
)

// MagicSchool holds the schema definition for the MagicSchool entity.
type MagicSchool struct {
	ent.Schema
}

func (MagicSchool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}
