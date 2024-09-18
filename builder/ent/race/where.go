// Code generated by ent, DO NOT EDIT.

package race

import (
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/builder/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldName, v))
}

// Speed applies equality check predicate on the "speed" field. It's identical to SpeedEQ.
func Speed(v int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSpeed, v))
}

// SizeDescription applies equality check predicate on the "size_description" field. It's identical to SizeDescriptionEQ.
func SizeDescription(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSizeDescription, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldAge, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Race {
	return predicate.Race(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Race {
	return predicate.Race(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Race {
	return predicate.Race(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Race {
	return predicate.Race(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Race {
	return predicate.Race(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Race {
	return predicate.Race(sql.FieldContainsFold(FieldName, v))
}

// SpeedEQ applies the EQ predicate on the "speed" field.
func SpeedEQ(v int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSpeed, v))
}

// SpeedNEQ applies the NEQ predicate on the "speed" field.
func SpeedNEQ(v int) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldSpeed, v))
}

// SpeedIn applies the In predicate on the "speed" field.
func SpeedIn(vs ...int) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldSpeed, vs...))
}

// SpeedNotIn applies the NotIn predicate on the "speed" field.
func SpeedNotIn(vs ...int) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldSpeed, vs...))
}

// SpeedGT applies the GT predicate on the "speed" field.
func SpeedGT(v int) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldSpeed, v))
}

// SpeedGTE applies the GTE predicate on the "speed" field.
func SpeedGTE(v int) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldSpeed, v))
}

// SpeedLT applies the LT predicate on the "speed" field.
func SpeedLT(v int) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldSpeed, v))
}

// SpeedLTE applies the LTE predicate on the "speed" field.
func SpeedLTE(v int) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldSpeed, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v Size) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v Size) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...Size) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...Size) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldSize, vs...))
}

// SizeDescriptionEQ applies the EQ predicate on the "size_description" field.
func SizeDescriptionEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSizeDescription, v))
}

// SizeDescriptionNEQ applies the NEQ predicate on the "size_description" field.
func SizeDescriptionNEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldSizeDescription, v))
}

// SizeDescriptionIn applies the In predicate on the "size_description" field.
func SizeDescriptionIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldSizeDescription, vs...))
}

// SizeDescriptionNotIn applies the NotIn predicate on the "size_description" field.
func SizeDescriptionNotIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldSizeDescription, vs...))
}

// SizeDescriptionGT applies the GT predicate on the "size_description" field.
func SizeDescriptionGT(v string) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldSizeDescription, v))
}

// SizeDescriptionGTE applies the GTE predicate on the "size_description" field.
func SizeDescriptionGTE(v string) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldSizeDescription, v))
}

// SizeDescriptionLT applies the LT predicate on the "size_description" field.
func SizeDescriptionLT(v string) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldSizeDescription, v))
}

// SizeDescriptionLTE applies the LTE predicate on the "size_description" field.
func SizeDescriptionLTE(v string) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldSizeDescription, v))
}

// SizeDescriptionContains applies the Contains predicate on the "size_description" field.
func SizeDescriptionContains(v string) predicate.Race {
	return predicate.Race(sql.FieldContains(FieldSizeDescription, v))
}

// SizeDescriptionHasPrefix applies the HasPrefix predicate on the "size_description" field.
func SizeDescriptionHasPrefix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasPrefix(FieldSizeDescription, v))
}

// SizeDescriptionHasSuffix applies the HasSuffix predicate on the "size_description" field.
func SizeDescriptionHasSuffix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasSuffix(FieldSizeDescription, v))
}

// SizeDescriptionEqualFold applies the EqualFold predicate on the "size_description" field.
func SizeDescriptionEqualFold(v string) predicate.Race {
	return predicate.Race(sql.FieldEqualFold(FieldSizeDescription, v))
}

// SizeDescriptionContainsFold applies the ContainsFold predicate on the "size_description" field.
func SizeDescriptionContainsFold(v string) predicate.Race {
	return predicate.Race(sql.FieldContainsFold(FieldSizeDescription, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v string) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v string) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v string) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v string) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldAge, v))
}

// AgeContains applies the Contains predicate on the "age" field.
func AgeContains(v string) predicate.Race {
	return predicate.Race(sql.FieldContains(FieldAge, v))
}

// AgeHasPrefix applies the HasPrefix predicate on the "age" field.
func AgeHasPrefix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasPrefix(FieldAge, v))
}

// AgeHasSuffix applies the HasSuffix predicate on the "age" field.
func AgeHasSuffix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasSuffix(FieldAge, v))
}

// AgeEqualFold applies the EqualFold predicate on the "age" field.
func AgeEqualFold(v string) predicate.Race {
	return predicate.Race(sql.FieldEqualFold(FieldAge, v))
}

// AgeContainsFold applies the ContainsFold predicate on the "age" field.
func AgeContainsFold(v string) predicate.Race {
	return predicate.Race(sql.FieldContainsFold(FieldAge, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Race) predicate.Race {
	return predicate.Race(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Race) predicate.Race {
	return predicate.Race(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Race) predicate.Race {
	return predicate.Race(sql.NotPredicates(p))
}
