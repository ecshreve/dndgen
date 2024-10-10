package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type BaseGQLMixin struct {
	mixin.Schema
}

func (BaseGQLMixin) Fields() []ent.Field {
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
func (BaseGQLMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
