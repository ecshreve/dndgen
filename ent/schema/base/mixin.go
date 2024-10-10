package base

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BaseMixin is a mixin for common fields and annotations.
type BaseMixin struct {
	ent.Schema
}

// Fields returns the fields for the schema.
func (BaseMixin) Fields() []ent.Field {
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
