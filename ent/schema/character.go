package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Character holds the schema definition for the Character entity.
type Character struct {
	ent.Schema
}

// Fields of the Character.
func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("age").Positive().Default(25),
		field.Int("level").Positive().Default(1),
	}
}

// Edges of the Character.
func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("race", Race.Type).Unique(),
		edge.To("class", Class.Type).Unique(),
		edge.To("alignment", Alignment.Type).Unique(),
		edge.To("traits", Trait.Type),
		edge.To("languages", Language.Type),
		edge.To("proficiencies", Proficiency.Type),
		edge.From("ability_scores", CharacterAbilityScore.Type).
			Ref("character"),
	}
}

// Annotations returns the annotations for the Character.
func (Character) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
