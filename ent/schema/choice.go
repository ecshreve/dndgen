package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ProficiencyChoice struct {
	ent.Schema
}

// Fields of the Choice.
func (ProficiencyChoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("choose").
			Positive(),
		field.Strings("desc"),
	}
}

// Edges of the Choice.
func (ProficiencyChoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("proficiencies", Proficiency.Type),
		edge.From("race", Race.Type).
			Ref("starting_proficiency_options").
			Unique(),
	}
}
