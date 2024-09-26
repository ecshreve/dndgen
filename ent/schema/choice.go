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
		edge.From("class", Class.Type).
			Ref("proficiency_options").
			Unique(),
	}
}

type AbilityBonusChoice struct {
	ent.Schema
}

// Fields of the Choice.
func (AbilityBonusChoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("choose").
			Positive(),
	}
}

// Edges of the Choice.
func (AbilityBonusChoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ability_bonuses", AbilityBonus.Type),
		edge.From("race", Race.Type).
			Ref("ability_bonus_options"),
	}
}

type LanguageChoice struct {
	ent.Schema
}

// Fields of the Choice.
func (LanguageChoice) Fields() []ent.Field {
	return []ent.Field{
		field.Int("choose").
			Positive(),
	}
}

// Edges of the Choice.
func (LanguageChoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("languages", Language.Type),
		edge.From("race", Race.Type).
			Ref("language_options").
			Unique(),
		edge.From("subrace", Subrace.Type).
			Ref("language_options").
			Unique(),
	}
}
