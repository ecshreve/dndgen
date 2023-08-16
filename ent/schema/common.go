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
		field.String("indx").StructTag(`json:"index"`).Annotations(entgql.OrderField("INDX")),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.String("desc"),
	}
}

func (CommonMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
