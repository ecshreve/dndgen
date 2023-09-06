// Code generated by ent, DO NOT EDIT.

package coin

import (
	"entgo.io/ent/dialect/sql"
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

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldDesc, v))
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

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Coin {
	return predicate.Coin(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Coin {
	return predicate.Coin(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Coin {
	return predicate.Coin(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Coin {
	return predicate.Coin(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Coin {
	return predicate.Coin(sql.FieldContainsFold(FieldDesc, v))
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
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
func Not(p predicate.Coin) predicate.Coin {
	return predicate.Coin(func(s *sql.Selector) {
		p(s.Not())
	})
}
