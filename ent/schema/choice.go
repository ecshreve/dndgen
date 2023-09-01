package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// // Choice holds the schema definition for the Choice entity.
// type ProficiencyChoice struct {
// 	ent.Schema
// }

// // Fields of the Choice.
// func (ProficiencyChoice) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.String("desc").Optional(),
// 		field.Int("choose"),
// 	}
// }

// // Edges of the Choice.
// func (ProficiencyChoice) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.To("options", Proficiency.Type),
// 		edge.From("class", Class.Type).
// 			Ref("proficiency_choices"),
// 		edge.From("race", Race.Type).
// 			Ref("starting_proficiency_option"),
// 	}
// }

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
		edge.To("starting_equipment_options", Equipment.Type),
		edge.From("class", Class.Type).
			Ref("proficiency_choices"),
		edge.From("race", Race.Type).
			Ref("starting_proficiency_options"),
	}
}
