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
		field.Enum("equipment_category").
			Values("weapon", "armor", "adventuring_gear", "tools", "mounts_and_vehicles", "trade_goods", "other").Default("other"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("weapon", Weapon.Type).Unique(),
		edge.To("armor", Armor.Type).Unique(),
		edge.To("gear", Gear.Type).Unique(),
		edge.To("tool", Tool.Type).Unique(),
	}
}

// Annotations of the Equipment.
func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
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

type Gear struct {
	ent.Schema
}

// Mixin of the Armor.
func (Gear) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Gear) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("gear_category").
			Values("ammunition", "standard_gear", "kits", "equipment_packs", "arcane_foci", "druidic_foci", "holy_symbols", "other").Default("other"),
		field.Strings("desc"),
		field.Int("quantity").Optional(),
	}
}

func (Gear) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("gear"),
	}
}

func (Gear) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

type Tool struct {
	ent.Schema
}

// Mixin of the Tool.
func (Tool) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Tool) Fields() []ent.Field {
	return []ent.Field{
		field.String("tool_category"),
	}
}

func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).Ref("tool"),
	}
}

func (Tool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
