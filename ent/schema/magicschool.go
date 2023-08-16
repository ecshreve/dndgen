package schema

import "entgo.io/ent"

// MagicSchool holds the schema definition for the MagicSchool entity.
type MagicSchool struct {
	ent.Schema
}

func (MagicSchool) Mixin() []ent.Mixin {
    return []ent.Mixin{
        CommonMixin{},
    }
}

// Fields of the MagicSchool.
func (MagicSchool) Fields() []ent.Field {
	return nil
}

// Edges of the MagicSchool.
func (MagicSchool) Edges() []ent.Edge {
	return nil
}

