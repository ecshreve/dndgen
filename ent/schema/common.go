package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type EquipmentMixin struct {
	mixin.Schema
}

func (EquipmentMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("INDX"),
			),
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Float("weight"),
	}
}

// Edges of the EquipmentMixin
func (EquipmentMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cost", Cost.Type).
			Unique(),
	}
}

type CommonMixin struct {
	mixin.Schema
}

func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("INDX"),
			),
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Strings("desc").
			Optional(),
	}
}

// Annotations returns the annotations for the schema.
func (CommonMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
