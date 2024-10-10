package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Race holds the schema definition for the Race entity.
type Race struct {
	ent.Schema
}

// Fields of the Race.
func (Race) Fields() []ent.Field {
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
		field.Int("speed").Positive(),
		field.Enum("size").Values(
			"Small",
			"Medium",
			"Large",
		).Default("Medium"),
		field.String("size_desc").StructTag(`json:"size_description"`),
		field.String("alignment_desc").StructTag(`json:"alignment"`),
		field.String("age_desc").StructTag(`json:"age"`),
		field.String("language_desc"),
	}
}

// Edges of the Race.
func (Race) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("traits", Trait.Type),
		edge.To("starting_proficiencies", Proficiency.Type),
		edge.To("starting_proficiency_options", ProficiencyChoice.Type),
		edge.To("ability_bonuses", AbilityScore.Type).
			Through("race_ability_bonuses", AbilityBonus.Type),
		// edge.To("ability_bonus_options", AbilityBonusChoice.Type).
		// 	Unique(),
		edge.To("languages", Language.Type),
		edge.To("language_options", LanguageChoice.Type).
			Unique(),
		// edge.To("subraces", Subrace.Type),
		edge.From("characters", Character.Type).
			Ref("race"),
	}
}

// Annotations of the Race.
func (Race) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

// type Subrace struct {
// 	ent.Schema
// }

// func (Subrace) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.String("indx").StructTag(`json:"index"`).
// 			NotEmpty().
// 			Unique().
// 			Annotations(
// 				entgql.OrderField("INDX"),
// 			),
// 		field.String("name").
// 			NotEmpty().
// 			Annotations(
// 				entgql.OrderField("NAME"),
// 			),
// 		field.Strings("desc"),
// 	}

// }

// func (Subrace) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("race", Race.Type).
// 			Ref("subraces").
// 			Unique(),
// 		edge.To("ability_bonuses", AbilityBonus.Type),
// 		edge.To("proficiencies", Proficiency.Type),
// 		edge.To("traits", Trait.Type),
// 		edge.To("language_options", LanguageChoice.Type),
// 	}
// }
