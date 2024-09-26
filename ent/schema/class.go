package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
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
		field.Int("hit_die").
			Positive(),
	}
}

func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proficiencies", Proficiency.Type),
		edge.To("proficiency_options", ProficiencyChoice.Type),
		edge.To("starting_equipment", EquipmentEntry.Type),
		edge.To("saving_throws", AbilityScore.Type),
	}
}

// Annotations of the Class.
func (Class) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}
