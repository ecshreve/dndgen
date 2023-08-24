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

// Mixin of the Proficiency.
func (Proficiency) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Proficiency.
func (Proficiency) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("proficiency_category").
			Values("weapons",
				"armor",
				"artisans_tools",
				"vehicles",
				"gaming_sets",
				"musical_instruments",
				"saving_throws",
				"skills",
				"other").Default("other"),
	}
}

// Edges of the Proficiency.
func (Proficiency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("classes", Class.Type).Ref("proficiencies"),
		edge.From("races", Race.Type).Ref("proficiencies"),
		edge.To("skill", Skill.Type),
		edge.To("equipment", Equipment.Type),
	}
}

// Annotations of the Proficiency.
func (Proficiency) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
