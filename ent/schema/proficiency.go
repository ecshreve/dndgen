package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			Ref("proficiencies"),
		edge.From("class", Class.Type).
			Ref("proficiencies"),
		edge.From("character", Character.Type).
			Ref("proficiencies").
			Through("character_proficiencies", CharacterProficiency.Type),
	}
}

type CharacterProficiency struct {
	ent.Schema
}

// Annotatioins
func (CharacterProficiency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("character_id", "proficiency_id"),
	}
}

func (CharacterProficiency) Fields() []ent.Field {
	return []ent.Field{
		field.Int("character_id"),
		field.Int("proficiency_id"),
	}
}

func (CharacterProficiency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("character", Character.Type).
			Unique().
			Required().
			Field("character_id"),
		edge.To("proficiency", Proficiency.Type).
			Unique().
			Required().
			Field("proficiency_id"),
	}
}
