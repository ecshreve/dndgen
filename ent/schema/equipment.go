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
			Values("weapon", "armor", "adventuring_gear", "tools", "mounts_and_vehicles", "other").Default("other"),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cost", Cost.Type).Unique(),
		edge.To("weapon", Weapon.Type).Unique(),
		edge.To("armor", Armor.Type).Unique(),
		edge.To("gear", Gear.Type).Unique(),
		edge.To("tool", Tool.Type).Unique(),
		edge.To("vehicle", Vehicle.Type).Unique(),
	}
}

// Annotations of the Equipment.
func (Equipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

//==========================================================
// Armor
//==========================================================

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
		field.Int("equipment_id"),
	}
}

func (Armor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("armor").Unique().
			Required().
			Field("equipment_id"),
		edge.To("armor_class", ArmorClass.Type),
	}
}

func (Armor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

//==========================================================
// Adventuring Gear and Tools
//==========================================================

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
		field.Int("equipment_id"),
	}
}

func (Gear) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("gear").
			Unique().Required().
			Field("equipment_id"),
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
		field.Int("equipment_id"),
	}
}

func (Tool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("tool").
			Unique().Required().
			Field("equipment_id"),
	}
}

func (Tool) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

// ==========================================================
// Vehicles and Mounts
// ==========================================================
type Vehicle struct {
	ent.Schema
}

// Mixin of the Vehicle.
func (Vehicle) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.String("vehicle_category"),
		field.String("capacity"),
		field.Int("equipment_id"),
	}
}

func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("vehicle").
			Unique().Required().
			Field("equipment_id"),
	}
}

func (Vehicle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
