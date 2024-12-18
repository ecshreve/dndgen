package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Proficiency holds the schema definition for the Proficiency entity.
type Proficiency struct {
	ent.Schema
}

// Fields of the Proficiency.
func (Proficiency) Fields() []ent.Field {
	return []ent.Field{
		field.String("indx").
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
		field.String("reference").
			NotEmpty(),
	}
}

// Edges of the Proficiency.
func (Proficiency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("race", Race.Type).
			Ref("starting_proficiencies"),
		edge.From("options", ProficiencyChoice.Type).
			Ref("proficiencies").
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		edge.From("class", Class.Type).
			Ref("proficiencies"),
		edge.From("character_proficiencies", CharacterProficiency.Type).
			Ref("proficiency").
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
	}
}
