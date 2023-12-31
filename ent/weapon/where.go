// Code generated by ent, DO NOT EDIT.

package weapon

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldName, v))
}

// EquipmentID applies equality check predicate on the "equipment_id" field. It's identical to EquipmentIDEQ.
func EquipmentID(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldEquipmentID, v))
}

// WeaponCategory applies equality check predicate on the "weapon_category" field. It's identical to WeaponCategoryEQ.
func WeaponCategory(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldWeaponCategory, v))
}

// WeaponRange applies equality check predicate on the "weapon_range" field. It's identical to WeaponRangeEQ.
func WeaponRange(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldWeaponRange, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContainsFold(FieldName, v))
}

// EquipmentIDEQ applies the EQ predicate on the "equipment_id" field.
func EquipmentIDEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldEquipmentID, v))
}

// EquipmentIDNEQ applies the NEQ predicate on the "equipment_id" field.
func EquipmentIDNEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldEquipmentID, v))
}

// EquipmentIDIn applies the In predicate on the "equipment_id" field.
func EquipmentIDIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldEquipmentID, vs...))
}

// EquipmentIDNotIn applies the NotIn predicate on the "equipment_id" field.
func EquipmentIDNotIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldEquipmentID, vs...))
}

// WeaponCategoryEQ applies the EQ predicate on the "weapon_category" field.
func WeaponCategoryEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldWeaponCategory, v))
}

// WeaponCategoryNEQ applies the NEQ predicate on the "weapon_category" field.
func WeaponCategoryNEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldWeaponCategory, v))
}

// WeaponCategoryIn applies the In predicate on the "weapon_category" field.
func WeaponCategoryIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldWeaponCategory, vs...))
}

// WeaponCategoryNotIn applies the NotIn predicate on the "weapon_category" field.
func WeaponCategoryNotIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldWeaponCategory, vs...))
}

// WeaponCategoryGT applies the GT predicate on the "weapon_category" field.
func WeaponCategoryGT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldWeaponCategory, v))
}

// WeaponCategoryGTE applies the GTE predicate on the "weapon_category" field.
func WeaponCategoryGTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldWeaponCategory, v))
}

// WeaponCategoryLT applies the LT predicate on the "weapon_category" field.
func WeaponCategoryLT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldWeaponCategory, v))
}

// WeaponCategoryLTE applies the LTE predicate on the "weapon_category" field.
func WeaponCategoryLTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldWeaponCategory, v))
}

// WeaponCategoryContains applies the Contains predicate on the "weapon_category" field.
func WeaponCategoryContains(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContains(FieldWeaponCategory, v))
}

// WeaponCategoryHasPrefix applies the HasPrefix predicate on the "weapon_category" field.
func WeaponCategoryHasPrefix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasPrefix(FieldWeaponCategory, v))
}

// WeaponCategoryHasSuffix applies the HasSuffix predicate on the "weapon_category" field.
func WeaponCategoryHasSuffix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasSuffix(FieldWeaponCategory, v))
}

// WeaponCategoryEqualFold applies the EqualFold predicate on the "weapon_category" field.
func WeaponCategoryEqualFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEqualFold(FieldWeaponCategory, v))
}

// WeaponCategoryContainsFold applies the ContainsFold predicate on the "weapon_category" field.
func WeaponCategoryContainsFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContainsFold(FieldWeaponCategory, v))
}

// WeaponRangeEQ applies the EQ predicate on the "weapon_range" field.
func WeaponRangeEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldWeaponRange, v))
}

// WeaponRangeNEQ applies the NEQ predicate on the "weapon_range" field.
func WeaponRangeNEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldWeaponRange, v))
}

// WeaponRangeIn applies the In predicate on the "weapon_range" field.
func WeaponRangeIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldWeaponRange, vs...))
}

// WeaponRangeNotIn applies the NotIn predicate on the "weapon_range" field.
func WeaponRangeNotIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldWeaponRange, vs...))
}

// WeaponRangeGT applies the GT predicate on the "weapon_range" field.
func WeaponRangeGT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldWeaponRange, v))
}

// WeaponRangeGTE applies the GTE predicate on the "weapon_range" field.
func WeaponRangeGTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldWeaponRange, v))
}

// WeaponRangeLT applies the LT predicate on the "weapon_range" field.
func WeaponRangeLT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldWeaponRange, v))
}

// WeaponRangeLTE applies the LTE predicate on the "weapon_range" field.
func WeaponRangeLTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldWeaponRange, v))
}

// WeaponRangeContains applies the Contains predicate on the "weapon_range" field.
func WeaponRangeContains(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContains(FieldWeaponRange, v))
}

// WeaponRangeHasPrefix applies the HasPrefix predicate on the "weapon_range" field.
func WeaponRangeHasPrefix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasPrefix(FieldWeaponRange, v))
}

// WeaponRangeHasSuffix applies the HasSuffix predicate on the "weapon_range" field.
func WeaponRangeHasSuffix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasSuffix(FieldWeaponRange, v))
}

// WeaponRangeEqualFold applies the EqualFold predicate on the "weapon_range" field.
func WeaponRangeEqualFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEqualFold(FieldWeaponRange, v))
}

// WeaponRangeContainsFold applies the ContainsFold predicate on the "weapon_range" field.
func WeaponRangeContainsFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContainsFold(FieldWeaponRange, v))
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWeaponDamage applies the HasEdge predicate on the "weapon_damage" edge.
func HasWeaponDamage() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WeaponDamageTable, WeaponDamageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeaponDamageWith applies the HasEdge predicate on the "weapon_damage" edge with a given conditions (other predicates).
func HasWeaponDamageWith(preds ...predicate.WeaponDamage) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newWeaponDamageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWeaponProperties applies the HasEdge predicate on the "weapon_properties" edge.
func HasWeaponProperties() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, WeaponPropertiesTable, WeaponPropertiesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeaponPropertiesWith applies the HasEdge predicate on the "weapon_properties" edge with a given conditions (other predicates).
func HasWeaponPropertiesWith(preds ...predicate.WeaponProperty) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newWeaponPropertiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Weapon) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Weapon) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
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
func Not(p predicate.Weapon) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		p(s.Not())
	})
}
