package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Coin struct {
	ent.Schema
}

func (Coin) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").Unique(), //Values("cp", "sp", "ep", "gp", "pp", "other").Default("other"),
		field.String("desc"),
		field.Float("gold_conversion_rate"),
	}
}

// type Cost struct {
// 	ent.Schema
// }

// func (Cost) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.Int("coin_id"),
// 		field.Int("quantity"),
// 	}
// }

// func (Cost) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.To("coin", Coin.Type).
// 			Unique().
// 			Required().
// 			Field("coin_id"),
// 	}
// }

type EquipmentCost struct {
	ent.Schema
}

func (EquipmentCost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("equipment_id"),
		field.Int("coin_id"),
		field.Int("quantity"),
		field.Float("gp_value"),
	}
}

func (EquipmentCost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("equipment", Equipment.Type).
			Ref("cost").
			Unique().
			Required().
			Field("equipment_id"),
		edge.To("coin", Coin.Type).
			Unique().
			Required().
			Field("coin_id"),
	}
}
