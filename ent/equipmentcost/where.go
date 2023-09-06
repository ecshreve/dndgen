// Code generated by ent, DO NOT EDIT.

package equipmentcost

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldLTE(FieldID, id))
}

// EquipmentID applies equality check predicate on the "equipment_id" field. It's identical to EquipmentIDEQ.
func EquipmentID(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldEquipmentID, v))
}

// CoinID applies equality check predicate on the "coin_id" field. It's identical to CoinIDEQ.
func CoinID(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldCoinID, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldQuantity, v))
}

// GpValue applies equality check predicate on the "gp_value" field. It's identical to GpValueEQ.
func GpValue(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldGpValue, v))
}

// EquipmentIDEQ applies the EQ predicate on the "equipment_id" field.
func EquipmentIDEQ(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldEquipmentID, v))
}

// EquipmentIDNEQ applies the NEQ predicate on the "equipment_id" field.
func EquipmentIDNEQ(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNEQ(FieldEquipmentID, v))
}

// EquipmentIDIn applies the In predicate on the "equipment_id" field.
func EquipmentIDIn(vs ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldIn(FieldEquipmentID, vs...))
}

// EquipmentIDNotIn applies the NotIn predicate on the "equipment_id" field.
func EquipmentIDNotIn(vs ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNotIn(FieldEquipmentID, vs...))
}

// CoinIDEQ applies the EQ predicate on the "coin_id" field.
func CoinIDEQ(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldCoinID, v))
}

// CoinIDNEQ applies the NEQ predicate on the "coin_id" field.
func CoinIDNEQ(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNEQ(FieldCoinID, v))
}

// CoinIDIn applies the In predicate on the "coin_id" field.
func CoinIDIn(vs ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldIn(FieldCoinID, vs...))
}

// CoinIDNotIn applies the NotIn predicate on the "coin_id" field.
func CoinIDNotIn(vs ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNotIn(FieldCoinID, vs...))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldLTE(FieldQuantity, v))
}

// GpValueEQ applies the EQ predicate on the "gp_value" field.
func GpValueEQ(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldEQ(FieldGpValue, v))
}

// GpValueNEQ applies the NEQ predicate on the "gp_value" field.
func GpValueNEQ(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNEQ(FieldGpValue, v))
}

// GpValueIn applies the In predicate on the "gp_value" field.
func GpValueIn(vs ...float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldIn(FieldGpValue, vs...))
}

// GpValueNotIn applies the NotIn predicate on the "gp_value" field.
func GpValueNotIn(vs ...float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldNotIn(FieldGpValue, vs...))
}

// GpValueGT applies the GT predicate on the "gp_value" field.
func GpValueGT(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldGT(FieldGpValue, v))
}

// GpValueGTE applies the GTE predicate on the "gp_value" field.
func GpValueGTE(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldGTE(FieldGpValue, v))
}

// GpValueLT applies the LT predicate on the "gp_value" field.
func GpValueLT(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldLT(FieldGpValue, v))
}

// GpValueLTE applies the LTE predicate on the "gp_value" field.
func GpValueLTE(v float64) predicate.EquipmentCost {
	return predicate.EquipmentCost(sql.FieldLTE(FieldGpValue, v))
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		step := newEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCoin applies the HasEdge predicate on the "coin" edge.
func HasCoin() predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CoinTable, CoinColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoinWith applies the HasEdge predicate on the "coin" edge with a given conditions (other predicates).
func HasCoinWith(preds ...predicate.Coin) predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		step := newCoinStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EquipmentCost) predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EquipmentCost) predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EquipmentCost) predicate.EquipmentCost {
	return predicate.EquipmentCost(func(s *sql.Selector) {
		p(s.Not())
	})
}