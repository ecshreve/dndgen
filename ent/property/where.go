// Code generated by ent, DO NOT EDIT.

package property

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldName, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Property {
	return predicate.Property(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Property {
	return predicate.Property(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Property {
	return predicate.Property(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Property {
	return predicate.Property(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Property {
	return predicate.Property(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Property {
	return predicate.Property(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Property {
	return predicate.Property(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Property {
	return predicate.Property(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Property {
	return predicate.Property(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Property {
	return predicate.Property(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Property {
	return predicate.Property(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Property {
	return predicate.Property(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Property {
	return predicate.Property(sql.FieldContainsFold(FieldName, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.Property {
	return predicate.Property(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.Property {
	return predicate.Property(sql.FieldNotNull(FieldDesc))
}

// HasWeapons applies the HasEdge predicate on the "weapons" edge.
func HasWeapons() predicate.Property {
	return predicate.Property(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, WeaponsTable, WeaponsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeaponsWith applies the HasEdge predicate on the "weapons" edge with a given conditions (other predicates).
func HasWeaponsWith(preds ...predicate.Weapon) predicate.Property {
	return predicate.Property(func(s *sql.Selector) {
		step := newWeaponsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Property) predicate.Property {
	return predicate.Property(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Property) predicate.Property {
	return predicate.Property(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Property) predicate.Property {
	return predicate.Property(sql.NotPredicates(p))
}