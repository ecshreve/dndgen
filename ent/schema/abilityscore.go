package schema

import "entgo.io/ent"

// AbilityScore holds the schema definition for the AbilityScore entity.
type AbilityScore struct {
	ent.Schema
}

func (AbilityScore) Mixin() []ent.Mixin {
    return []ent.Mixin{
        CommonMixin{},
    }
}

// Fields of the AbilityScore.
func (AbilityScore) Fields() []ent.Field {
	return nil
}

// Edges of the AbilityScore.
func (AbilityScore) Edges() []ent.Edge {
	return nil
}

