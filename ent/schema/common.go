package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CommonMixin struct {
	mixin.Schema
}

func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).NotEmpty().Unique().
			Annotations(
				entgql.OrderField("INDX"),
				entproto.Field(2),
			),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
			entproto.Field(3),
		),
	}
}

func (CommonMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}

type EquipmentMixin struct {
	mixin.Schema
}

func (EquipmentMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).NotEmpty().Unique().Annotations(
			entgql.OrderField("INDX"),
		),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

type ArmorClass struct {
	ent.Schema
}

func (ArmorClass) Fields() []ent.Field {
	return []ent.Field{
		field.Int("base"),
		field.Bool("dex_bonus"),
		field.Int("max_bonus").Optional(),
	}
}
