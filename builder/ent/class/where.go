// Code generated by ent, DO NOT EDIT.

package class

import (
	"builder/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldName, v))
}

// HitDie applies equality check predicate on the "hit_die" field. It's identical to HitDieEQ.
func HitDie(v int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldHitDie, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Class {
	return predicate.Class(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Class {
	return predicate.Class(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Class {
	return predicate.Class(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Class {
	return predicate.Class(sql.FieldContainsFold(FieldName, v))
}

// HitDieEQ applies the EQ predicate on the "hit_die" field.
func HitDieEQ(v int) predicate.Class {
	return predicate.Class(sql.FieldEQ(FieldHitDie, v))
}

// HitDieNEQ applies the NEQ predicate on the "hit_die" field.
func HitDieNEQ(v int) predicate.Class {
	return predicate.Class(sql.FieldNEQ(FieldHitDie, v))
}

// HitDieIn applies the In predicate on the "hit_die" field.
func HitDieIn(vs ...int) predicate.Class {
	return predicate.Class(sql.FieldIn(FieldHitDie, vs...))
}

// HitDieNotIn applies the NotIn predicate on the "hit_die" field.
func HitDieNotIn(vs ...int) predicate.Class {
	return predicate.Class(sql.FieldNotIn(FieldHitDie, vs...))
}

// HitDieGT applies the GT predicate on the "hit_die" field.
func HitDieGT(v int) predicate.Class {
	return predicate.Class(sql.FieldGT(FieldHitDie, v))
}

// HitDieGTE applies the GTE predicate on the "hit_die" field.
func HitDieGTE(v int) predicate.Class {
	return predicate.Class(sql.FieldGTE(FieldHitDie, v))
}

// HitDieLT applies the LT predicate on the "hit_die" field.
func HitDieLT(v int) predicate.Class {
	return predicate.Class(sql.FieldLT(FieldHitDie, v))
}

// HitDieLTE applies the LTE predicate on the "hit_die" field.
func HitDieLTE(v int) predicate.Class {
	return predicate.Class(sql.FieldLTE(FieldHitDie, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Class) predicate.Class {
	return predicate.Class(sql.NotPredicates(p))
}
