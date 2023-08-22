package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Mixin of the Equipment.
func (Equipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("category").
			Values("weapon", "armor", "adventuring_gear", "tools", "mounts_and_vehicles", "trade_goods", "other"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("weapon", Weapon.Type).Unique(),
		edge.To("armor", Armor.Type).Unique(),
	}
}

//==========================================================
// Weapon and Armor
//==========================================================

type Weapon struct {
	ent.Schema
}

// Mixin of the Weapon.
func (Weapon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Weapon) Fields() []ent.Field {
	return []ent.Field{
		field.String("weapon_range"),
	}
}

func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("weapon"),
	}
}

func (Weapon) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

type Armor struct {
	ent.Schema
}

// Mixin of the Armor.
func (Armor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Armor) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("stealth_disadvantage"),
		field.Int("min_strength"),
	}
}

func (Armor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("armor"),
		edge.To("armor_class", ArmorClass.Type),
	}
}

func (Armor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
