package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type EquipmentChoice struct {
	ent.Schema
}

func (EquipmentChoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("class_id"),
		field.Int("choose"),
		field.String("desc").Optional(),
	}
}

func (EquipmentChoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).
			Ref("equipment_choice").
			Unique().Required().
			Field("class_id"),
		edge.From("equipment", Equipment.Type).
			Ref("choice"),
	}
}

// Choice holds the schema definition for the Choice entity.
type Choice struct {
	ent.Schema
}

// Fields of the Choice.
func (Choice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("choose"),
		field.String("desc").Optional(),
	}
}

// Edges of the Choice.
func (Choice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("choices", Choice.Type).
			From("parent_choice").
			Unique(),
		edge.To("proficiency_options", Proficiency.Type),
		edge.From("class", Class.Type).
			Ref("proficiency_choices"),
		edge.From("race", Race.Type).
			Ref("starting_proficiency_options"),
	}
}
