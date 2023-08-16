package schema

import "entgo.io/ent"

// DamageType holds the schema definition for the DamageType entity.
type DamageType struct {
	ent.Schema
}

func (DamageType) Mixin() []ent.Mixin {
    return []ent.Mixin{
        CommonMixin{},
    }
}

// Fields of the DamageType.
func (DamageType) Fields() []ent.Field {
	return nil
}

// Edges of the DamageType.
func (DamageType) Edges() []ent.Edge {
	return nil
}

