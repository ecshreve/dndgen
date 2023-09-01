package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Choice holds the schema definition for the Choice entity.
type ProficiencyChoice struct {
	ent.Schema
}

// Fields of the Choice.
func (ProficiencyChoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").Optional(),
		field.Int("choose"),
	}
}

// Edges of the Choice.
func (ProficiencyChoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("options", Proficiency.Type),
		edge.From("class", Class.Type).
			Ref("proficiency_choices"),
		edge.From("race", Race.Type).
			Ref("starting_proficiency_option"),
	}
}
