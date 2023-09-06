package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ClassEquipment struct {
	ent.Schema
}

func (ClassEquipment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("class_id", "equipment_id"),
	}
}

func (ClassEquipment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("class_id"),
		field.Int("equipment_id"),
		field.Int("quantity"),
	}
}

// Edges of the StartingEquipment.
func (ClassEquipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class", Class.Type).
			Required().Unique().
			Field("class_id"),
		edge.To("equipment", Equipment.Type).
			Required().Unique().
			Field("equipment_id"),
	}
}

type EquipmentCategory struct {
	ent.Schema
}

func (EquipmentCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_category_id").Optional(),
		field.String("name"),
		// .
		// 	Values(
		// 		"weapon",
		// 		"simple",
		// 		"martial",
		// 		"melee",
		// 		"ranged",
		// 		"armor",
		// 		"light",
		// 		"medium",
		// 		"heavy",
		// 		"shield",
		// 		"adventuring_gear",
		// 		"tools",
		// 		"mounts_and_vehicles",
		// 		"other").Default("other"),
	}
}

func (EquipmentCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", EquipmentCategory.Type).
			From("parent").
			Unique().
			Field("parent_category_id"),
		edge.To("equipment", Equipment.Type),
	}
}

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
		field.Int("weight").Optional(),
	}
}

// Edges of the Equipment.
func (Equipment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment_category", EquipmentCategory.Type).
			Ref("equipment"),
		edge.To("cost", EquipmentCost.Type).Unique(),
		edge.To("weapon", Weapon.Type).Unique(),
		edge.To("armor", Armor.Type).Unique(),
		edge.To("gear", Gear.Type).Unique(),
		edge.To("tool", Tool.Type).Unique(),
		edge.To("vehicle", Vehicle.Type).Unique(),
		edge.From("class", Class.Type).
			Ref("equipment").
			Through("class_equipment", ClassEquipment.Type),
		edge.From("choice", EquipmentChoice.Type).
			Ref("equipment"),
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
		// field.Enum("armor_category").
		// 	Values("light", "medium", "heavy", "shield", "other").Default("other"),
		field.String("armor_category"),
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
