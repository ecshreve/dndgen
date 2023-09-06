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

func (Class) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.Int("hit_die"),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proficiencies", Proficiency.Type),
		edge.To("proficiency_choices", ProficiencyChoice.Type),
		edge.To("equipment", Equipment.Type).
			Through("class_equipment", ClassEquipment.Type),
		edge.To("equipment_choices", EquipmentChoice.Type),
	}
}

// Annotations of the Class.
func (Class) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
