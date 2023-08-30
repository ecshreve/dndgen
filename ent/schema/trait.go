package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Trait holds the schema definition for the Trait entity.
type Trait struct {
	ent.Schema
}

// Mixin of the Trait.
func (Trait) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Trait.
func (Trait) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("desc"),
	}
}

// Edges of the Trait.
func (Trait) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("races", Race.Type).Ref("traits"),
		edge.From("subraces", Subrace.Type).Ref("traits"),
	}
}

// Annotations of the Trait.
func (Trait) Annotations() []schema.Annotation {
	return nil
}
