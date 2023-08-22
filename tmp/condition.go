package schema

import "entgo.io/ent"

// Condition holds the schema definition for the Condition entity.
type Condition struct {
	ent.Schema
}

func (Condition) Mixin() []ent.Mixin {
    return []ent.Mixin{
        CommonMixin{},
    }
}

// Fields of the Condition.
func (Condition) Fields() []ent.Field {
	return nil
}

// Edges of the Condition.
func (Condition) Edges() []ent.Edge {
	return nil
}

