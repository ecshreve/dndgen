package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ArmorClass holds the schema definition for the ArmorClass entity.
type ArmorClass struct {
	ent.Schema
}

// Fields of the ArmorClass.
func (ArmorClass) Fields() []ent.Field {
	return []ent.Field{
		field.Int("base").Positive(),
		field.Bool("dex_bonus").Default(false),
	}
}

// Edges of the ArmorClass.
func (ArmorClass) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("armor", Armor.Type).
			Ref("armor_class").
			Unique(),
	}
}

// Armor holds the schema definition for the Armor entity.
type Armor struct {
	ent.Schema
}

// Fields of the Armor.
func (Armor) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("armor_category").
			Values(
				"light",
				"medium",
				"heavy",
				"shield",
			).
			Annotations(entgql.QueryField("armor_category")),
		field.Int("str_minimum").
			Annotations(entgql.QueryField("str_minimum")),
		field.Bool("stealth_disadvantage").
			Annotations(entgql.QueryField("stealth_disadvantage")),
	}
}

// Edges of the Armor.
func (Armor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("equipment", Equipment.Type).
			Unique(),
		edge.To("armor_class", ArmorClass.Type).
			Unique(),
	}
}
