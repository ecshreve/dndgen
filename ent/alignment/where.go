// Code generated by ent, DO NOT EDIT.

package alignment

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Alignment {
	return predicate.Alignment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Alignment {
	return predicate.Alignment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Alignment {
	return predicate.Alignment(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldName, v))
}

// Abbr applies equality check predicate on the "abbr" field. It's identical to AbbrEQ.
func Abbr(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldAbbr, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Alignment {
	return predicate.Alignment(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Alignment {
	return predicate.Alignment(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Alignment {
	return predicate.Alignment(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Alignment {
	return predicate.Alignment(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldContainsFold(FieldName, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.Alignment {
	return predicate.Alignment(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.Alignment {
	return predicate.Alignment(sql.FieldNotNull(FieldDesc))
}

// AbbrEQ applies the EQ predicate on the "abbr" field.
func AbbrEQ(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEQ(FieldAbbr, v))
}

// AbbrNEQ applies the NEQ predicate on the "abbr" field.
func AbbrNEQ(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldNEQ(FieldAbbr, v))
}

// AbbrIn applies the In predicate on the "abbr" field.
func AbbrIn(vs ...string) predicate.Alignment {
	return predicate.Alignment(sql.FieldIn(FieldAbbr, vs...))
}

// AbbrNotIn applies the NotIn predicate on the "abbr" field.
func AbbrNotIn(vs ...string) predicate.Alignment {
	return predicate.Alignment(sql.FieldNotIn(FieldAbbr, vs...))
}

// AbbrGT applies the GT predicate on the "abbr" field.
func AbbrGT(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldGT(FieldAbbr, v))
}

// AbbrGTE applies the GTE predicate on the "abbr" field.
func AbbrGTE(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldGTE(FieldAbbr, v))
}

// AbbrLT applies the LT predicate on the "abbr" field.
func AbbrLT(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldLT(FieldAbbr, v))
}

// AbbrLTE applies the LTE predicate on the "abbr" field.
func AbbrLTE(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldLTE(FieldAbbr, v))
}

// AbbrContains applies the Contains predicate on the "abbr" field.
func AbbrContains(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldContains(FieldAbbr, v))
}

// AbbrHasPrefix applies the HasPrefix predicate on the "abbr" field.
func AbbrHasPrefix(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldHasPrefix(FieldAbbr, v))
}

// AbbrHasSuffix applies the HasSuffix predicate on the "abbr" field.
func AbbrHasSuffix(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldHasSuffix(FieldAbbr, v))
}

// AbbrEqualFold applies the EqualFold predicate on the "abbr" field.
func AbbrEqualFold(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldEqualFold(FieldAbbr, v))
}

// AbbrContainsFold applies the ContainsFold predicate on the "abbr" field.
func AbbrContainsFold(v string) predicate.Alignment {
	return predicate.Alignment(sql.FieldContainsFold(FieldAbbr, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Alignment) predicate.Alignment {
	return predicate.Alignment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Alignment) predicate.Alignment {
	return predicate.Alignment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Alignment) predicate.Alignment {
	return predicate.Alignment(sql.NotPredicates(p))
}
