package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id"),
		field.Float("weight"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return nil
}

type Weapon struct {
	ent.Schema
}

// Mixin of the Weapon.
func (Weapon) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		EquipmentMixin{},
	}
}

func (Weapon) Fields() []ent.Field {
	return []ent.Field{}
}

func (Weapon) Edges() []ent.Edge {
	return nil
}

type Armor struct {
	ent.Schema
}

// Mixin of the Armor.
func (Armor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
		EquipmentMixin{},
	}
}

func (Armor) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("stealth_disadvantage"),
		field.String("armor_class"),
		field.Int("min_strength"),
	}
}

func (Armor) Edges() []ent.Edge {
	return nil
}
