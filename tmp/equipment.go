package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Equipment holds the schema definition for the Equipment entity.
type Equipment struct {
	ent.Schema
}

func (Equipment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Equipment.
func (Equipment) Fields() []ent.Field {
	return []ent.Field{
		field.String("cost"),
		field.String("weight"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("weapon", Weapon.Type),
		edge.To("armor", Armor.Type),
		edge.To("gear", Gear.Type),
		edge.To("pack", Pack.Type),
		edge.To("ammunition", Ammunition.Type),
		edge.To("vehicle", Vehicle.Type),
		edge.To("magic_item", MagicItem.Type),
		edge.To("category", EquipmentCategory.Type),
		edge.To("subcategory", EquipmentCategory.Type),
		edge.From("proficiencies", Proficiency.Type).Ref("equipment"),
	}
}

type WeaponRange struct {
	ent.Schema
}

func (WeaponRange) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
		field.Int("normal"),
		field.Int("long"),
	}
}

func (WeaponRange) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("weapon", Weapon.Type).Ref("range"),
	}
}

type Weapon struct {
	ent.Schema
}

func (Weapon) Fields() []ent.Field {
	return []ent.Field{
		field.String("properties"),
	}
}

func (Weapon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("range", WeaponRange.Type),
		edge.To("damage", WeaponDamage.Type),
		edge.To("two_handed_damage", WeaponDamage.Type),
		edge.From("equipment", Equipment.Type).Ref("weapon"),
	}
}

type Armor struct {
	ent.Schema
}

func (Armor) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("stealth_disadvantage"),
		field.String("armor_class"),
		field.Int("min_strength"),
	}
}

func (Armor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("armor"),
	}
}

type Gear struct {
	ent.Schema
}

func (Gear) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("gear"),
	}
}

type Pack struct {
	ent.Schema
}

func (Pack) Fields() []ent.Field {
	return []ent.Field{
		field.String("contents"),
	}
}

func (Pack) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("pack"),
	}
}

type Ammunition struct {
	ent.Schema
}

func (Ammunition) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
	}
}

func (Ammunition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("ammunition"),
	}
}

type Vehicle struct {
	ent.Schema
}

func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.String("speed"),
		field.String("capacity"),
	}
}

func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("vehicle"),
	}
}

type MagicItem struct {
	ent.Schema
}

func (MagicItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("rarity"),
	}
}

func (MagicItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("magic_item"),
	}
}
