// Code generated by ent, DO NOT EDIT.

package tool

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldID, id))
}

// ToolCategory applies equality check predicate on the "tool_category" field. It's identical to ToolCategoryEQ.
func ToolCategory(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldToolCategory, v))
}

// ToolCategoryEQ applies the EQ predicate on the "tool_category" field.
func ToolCategoryEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEQ(FieldToolCategory, v))
}

// ToolCategoryNEQ applies the NEQ predicate on the "tool_category" field.
func ToolCategoryNEQ(v string) predicate.Tool {
	return predicate.Tool(sql.FieldNEQ(FieldToolCategory, v))
}

// ToolCategoryIn applies the In predicate on the "tool_category" field.
func ToolCategoryIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldIn(FieldToolCategory, vs...))
}

// ToolCategoryNotIn applies the NotIn predicate on the "tool_category" field.
func ToolCategoryNotIn(vs ...string) predicate.Tool {
	return predicate.Tool(sql.FieldNotIn(FieldToolCategory, vs...))
}

// ToolCategoryGT applies the GT predicate on the "tool_category" field.
func ToolCategoryGT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGT(FieldToolCategory, v))
}

// ToolCategoryGTE applies the GTE predicate on the "tool_category" field.
func ToolCategoryGTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldGTE(FieldToolCategory, v))
}

// ToolCategoryLT applies the LT predicate on the "tool_category" field.
func ToolCategoryLT(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLT(FieldToolCategory, v))
}

// ToolCategoryLTE applies the LTE predicate on the "tool_category" field.
func ToolCategoryLTE(v string) predicate.Tool {
	return predicate.Tool(sql.FieldLTE(FieldToolCategory, v))
}

// ToolCategoryContains applies the Contains predicate on the "tool_category" field.
func ToolCategoryContains(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContains(FieldToolCategory, v))
}

// ToolCategoryHasPrefix applies the HasPrefix predicate on the "tool_category" field.
func ToolCategoryHasPrefix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasPrefix(FieldToolCategory, v))
}

// ToolCategoryHasSuffix applies the HasSuffix predicate on the "tool_category" field.
func ToolCategoryHasSuffix(v string) predicate.Tool {
	return predicate.Tool(sql.FieldHasSuffix(FieldToolCategory, v))
}

// ToolCategoryEqualFold applies the EqualFold predicate on the "tool_category" field.
func ToolCategoryEqualFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldEqualFold(FieldToolCategory, v))
}

// ToolCategoryContainsFold applies the ContainsFold predicate on the "tool_category" field.
func ToolCategoryContainsFold(v string) predicate.Tool {
	return predicate.Tool(sql.FieldContainsFold(FieldToolCategory, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.Tool {
	return predicate.Tool(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.Tool {
	return predicate.Tool(sql.FieldNotNull(FieldDesc))
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.Tool {
	return predicate.Tool(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.Tool {
	return predicate.Tool(func(s *sql.Selector) {
		step := newEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tool) predicate.Tool {
	return predicate.Tool(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tool) predicate.Tool {
	return predicate.Tool(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tool) predicate.Tool {
	return predicate.Tool(sql.NotPredicates(p))
}