package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Rule holds the schema definition for the Rule entity.
type Rule struct {
	ent.Schema
}

func (Rule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Rule.
func (Rule) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").Annotations(),
	}
}

// Edges of the Rule.
func (Rule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rule_sections", RuleSection.Type).
			StructTag(`json:"subsections,omitempty"`),
	}
}

// Annotations of the Rule.
func (Rule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

// RuleSection holds the schema definition for the RuleSection entity.
type RuleSection struct {
	ent.Schema
}

func (RuleSection) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the RuleSection.
func (RuleSection) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc").Annotations(),
	}
}

// Edges of the RuleSection.
func (RuleSection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rules", Rule.Type).
			Ref("rule_sections"),
	}
}

// Annotations of the RuleSection.
func (RuleSection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
