// Code generated by ent, DO NOT EDIT.

package rulesection

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldName, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldDesc, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldContainsFold(FieldName, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.RuleSection {
	return predicate.RuleSection(sql.FieldContainsFold(FieldDesc, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RuleSection) predicate.RuleSection {
	return predicate.RuleSection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RuleSection) predicate.RuleSection {
	return predicate.RuleSection(func(s *sql.Selector) {
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
func Not(p predicate.RuleSection) predicate.RuleSection {
	return predicate.RuleSection(func(s *sql.Selector) {
		p(s.Not())
	})
}
