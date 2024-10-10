package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Property holds the schema definition for the Property entity.
type Property struct {
	ent.Schema
}

// Mixin of the Property.
func (Property) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseGQLMixin{},
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapons", Weapon.Type).
			Ref("properties"),
	}
}
