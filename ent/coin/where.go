// Code generated by ent, DO NOT EDIT.

package coin

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldName, v))
}

// GoldConversionRate applies equality check predicate on the "gold_conversion_rate" field. It's identical to GoldConversionRateEQ.
func GoldConversionRate(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldGoldConversionRate, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContainsFold(FieldName, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.Coin {
	return predicate.Coin(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.Coin {
	return predicate.Coin(sql.FieldNotNull(FieldDesc))
}

// GoldConversionRateEQ applies the EQ predicate on the "gold_conversion_rate" field.
func GoldConversionRateEQ(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldGoldConversionRate, v))
}

// GoldConversionRateNEQ applies the NEQ predicate on the "gold_conversion_rate" field.
func GoldConversionRateNEQ(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldGoldConversionRate, v))
}

// GoldConversionRateIn applies the In predicate on the "gold_conversion_rate" field.
func GoldConversionRateIn(vs ...float64) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldGoldConversionRate, vs...))
}

// GoldConversionRateNotIn applies the NotIn predicate on the "gold_conversion_rate" field.
func GoldConversionRateNotIn(vs ...float64) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldGoldConversionRate, vs...))
}

// GoldConversionRateGT applies the GT predicate on the "gold_conversion_rate" field.
func GoldConversionRateGT(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldGoldConversionRate, v))
}

// GoldConversionRateGTE applies the GTE predicate on the "gold_conversion_rate" field.
func GoldConversionRateGTE(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldGoldConversionRate, v))
}

// GoldConversionRateLT applies the LT predicate on the "gold_conversion_rate" field.
func GoldConversionRateLT(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldGoldConversionRate, v))
}

// GoldConversionRateLTE applies the LTE predicate on the "gold_conversion_rate" field.
func GoldConversionRateLTE(v float64) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldGoldConversionRate, v))
}

// HasCosts applies the HasEdge predicate on the "costs" edge.
func HasCosts() predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CostsTable, CostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCostsWith applies the HasEdge predicate on the "costs" edge with a given conditions (other predicates).
func HasCostsWith(preds ...predicate.Cost) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		step := newCostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Coin) predicate.Coin {
	return predicate.Coin(sql.NotPredicates(p))
}
