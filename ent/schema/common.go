package schema

import (
	"entgo.io/contrib/entgql"
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
			),
		field.String("name").NotEmpty().Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

func (CommonMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
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

type Cost struct {
	ent.Schema
}

func (Cost) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity"),
		field.String("unit"),
	}
}
