package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DamageType holds the schema definition for the DamageType entity.
type DamageType struct {
	ent.Schema
}

func (DamageType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Damage holds the schema definition for the Damage entity.
type Damage struct {
	ent.Schema
}

// Fields of the Damage.
func (Damage) Fields() []ent.Field {
	return []ent.Field{
		field.String("damage_dice"),
	}
}

// Edges of the Damage.
func (Damage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("damage_type", DamageType.Type).
			Unique(),
	}
}
