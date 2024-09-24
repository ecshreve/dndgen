// Code generated by ent, DO NOT EDIT.

package armor

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Armor {
	return predicate.Armor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Armor {
	return predicate.Armor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Armor {
	return predicate.Armor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Armor {
	return predicate.Armor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Armor {
	return predicate.Armor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Armor {
	return predicate.Armor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Armor {
	return predicate.Armor(sql.FieldLTE(FieldID, id))
}

// StrMinimum applies equality check predicate on the "str_minimum" field. It's identical to StrMinimumEQ.
func StrMinimum(v int) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldStrMinimum, v))
}

// StealthDisadvantage applies equality check predicate on the "stealth_disadvantage" field. It's identical to StealthDisadvantageEQ.
func StealthDisadvantage(v bool) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldStealthDisadvantage, v))
}

// ArmorCategoryEQ applies the EQ predicate on the "armor_category" field.
func ArmorCategoryEQ(v ArmorCategory) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldArmorCategory, v))
}

// ArmorCategoryNEQ applies the NEQ predicate on the "armor_category" field.
func ArmorCategoryNEQ(v ArmorCategory) predicate.Armor {
	return predicate.Armor(sql.FieldNEQ(FieldArmorCategory, v))
}

// ArmorCategoryIn applies the In predicate on the "armor_category" field.
func ArmorCategoryIn(vs ...ArmorCategory) predicate.Armor {
	return predicate.Armor(sql.FieldIn(FieldArmorCategory, vs...))
}

// ArmorCategoryNotIn applies the NotIn predicate on the "armor_category" field.
func ArmorCategoryNotIn(vs ...ArmorCategory) predicate.Armor {
	return predicate.Armor(sql.FieldNotIn(FieldArmorCategory, vs...))
}

// StrMinimumEQ applies the EQ predicate on the "str_minimum" field.
func StrMinimumEQ(v int) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldStrMinimum, v))
}

// StrMinimumNEQ applies the NEQ predicate on the "str_minimum" field.
func StrMinimumNEQ(v int) predicate.Armor {
	return predicate.Armor(sql.FieldNEQ(FieldStrMinimum, v))
}

// StrMinimumIn applies the In predicate on the "str_minimum" field.
func StrMinimumIn(vs ...int) predicate.Armor {
	return predicate.Armor(sql.FieldIn(FieldStrMinimum, vs...))
}

// StrMinimumNotIn applies the NotIn predicate on the "str_minimum" field.
func StrMinimumNotIn(vs ...int) predicate.Armor {
	return predicate.Armor(sql.FieldNotIn(FieldStrMinimum, vs...))
}

// StrMinimumGT applies the GT predicate on the "str_minimum" field.
func StrMinimumGT(v int) predicate.Armor {
	return predicate.Armor(sql.FieldGT(FieldStrMinimum, v))
}

// StrMinimumGTE applies the GTE predicate on the "str_minimum" field.
func StrMinimumGTE(v int) predicate.Armor {
	return predicate.Armor(sql.FieldGTE(FieldStrMinimum, v))
}

// StrMinimumLT applies the LT predicate on the "str_minimum" field.
func StrMinimumLT(v int) predicate.Armor {
	return predicate.Armor(sql.FieldLT(FieldStrMinimum, v))
}

// StrMinimumLTE applies the LTE predicate on the "str_minimum" field.
func StrMinimumLTE(v int) predicate.Armor {
	return predicate.Armor(sql.FieldLTE(FieldStrMinimum, v))
}

// StealthDisadvantageEQ applies the EQ predicate on the "stealth_disadvantage" field.
func StealthDisadvantageEQ(v bool) predicate.Armor {
	return predicate.Armor(sql.FieldEQ(FieldStealthDisadvantage, v))
}

// StealthDisadvantageNEQ applies the NEQ predicate on the "stealth_disadvantage" field.
func StealthDisadvantageNEQ(v bool) predicate.Armor {
	return predicate.Armor(sql.FieldNEQ(FieldStealthDisadvantage, v))
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.Armor {
	return predicate.Armor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.Armor {
	return predicate.Armor(func(s *sql.Selector) {
		step := newEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArmorClass applies the HasEdge predicate on the "armor_class" edge.
func HasArmorClass() predicate.Armor {
	return predicate.Armor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ArmorClassTable, ArmorClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArmorClassWith applies the HasEdge predicate on the "armor_class" edge with a given conditions (other predicates).
func HasArmorClassWith(preds ...predicate.ArmorClass) predicate.Armor {
	return predicate.Armor(func(s *sql.Selector) {
		step := newArmorClassStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Armor) predicate.Armor {
	return predicate.Armor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Armor) predicate.Armor {
	return predicate.Armor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Armor) predicate.Armor {
	return predicate.Armor(sql.NotPredicates(p))
}
