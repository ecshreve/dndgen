package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CharacterProficiency struct {
	ent.Schema
}

func (CharacterProficiency) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("proficiency_type").
			Values("SKILL", "EQUIPMENT", "EQUIPMENT_CATEGORY", "SAVING_THROW", "OTHER"),
		field.Enum("proficiency_source").
			Values("CLASS_PROFICIENCY", "RACE_PROFICIENCY", "CLASS_CHOICE", "RACE_CHOICE", "OTHER"),
	}
}

func (CharacterProficiency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("character_proficiencies").
			Unique().
			Required(),
		edge.To("proficiency", Proficiency.Type).
			Unique().
			Required(),
		edge.From("character_skill", CharacterSkill.Type).
			Ref("character_proficiency").
			Unique(),
	}
}
