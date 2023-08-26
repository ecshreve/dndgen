package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CommonMixin struct {
	mixin.Schema
}

func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").StructTag(`json:"index"`).NotEmpty().Unique(),
		field.String("name").NotEmpty(),
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
