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

func (Race) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Race.
func (Race) Fields() []ent.Field {
	return []ent.Field{
		field.Int("speed"),
	}
}

// Edges of the Race.
func (Race) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("languages", Language.Type),
		// edge.To("ability_bonuses", AbilityBonus.Type),
		edge.To("proficiencies", Proficiency.Type),
		edge.To("subrace", Subrace.Type).Unique(),
		edge.To("traits", Trait.Type),
	}
}

// Annotations of the Race.
func (Race) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

// Subrace holds the schema definition for the Subrace entity.
type Subrace struct {
	ent.Schema
}

func (Subrace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Subrace.
func (Subrace) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
	}
}

// Edges of the Subrace.
func (Subrace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("race", Race.Type).
			Ref("subrace").Unique(),
		edge.To("proficiencies", Proficiency.Type).
			StructTag(`json:"starting_proficiencies,omitempty"`),
		edge.To("traits", Trait.Type),
	}
}

// Annotations of the Subrace.
func (Subrace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
