package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/ecshreve/dndgen/ent/schema/base"
)

// Trait holds the schema definition for the Trait entity.
type Trait struct {
	ent.Schema
}

// Mixin of the Trait.
func (Trait) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.BaseMixin{},
	}
}

// Edges of the Trait.
func (Trait) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("race", Race.Type).
			Ref("traits"),
		edge.From("subrace", Subrace.Type).
			Ref("traits"),
	}
}
