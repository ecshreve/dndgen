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
		edge.From("ability_bonuses", AbilityBonus.Type).
			Ref("race"),
		edge.To("languages", Language.Type).
			StorageKey(edge.Table("race_languages"), edge.Columns("race_id", "language_id")),
		edge.To("proficiencies", Proficiency.Type),
	}
}

// Annotations of the Race.
func (Race) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
