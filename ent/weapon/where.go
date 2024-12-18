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

// RangeNormal applies equality check predicate on the "range_normal" field. It's identical to RangeNormalEQ.
func RangeNormal(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldRangeNormal, v))
}

// RangeLong applies equality check predicate on the "range_long" field. It's identical to RangeLongEQ.
func RangeLong(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldRangeLong, v))
}

// ThrowRangeNormal applies equality check predicate on the "throw_range_normal" field. It's identical to ThrowRangeNormalEQ.
func ThrowRangeNormal(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldThrowRangeNormal, v))
}

// ThrowRangeLong applies equality check predicate on the "throw_range_long" field. It's identical to ThrowRangeLongEQ.
func ThrowRangeLong(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldThrowRangeLong, v))
}

// DamageDice applies equality check predicate on the "damage_dice" field. It's identical to DamageDiceEQ.
func DamageDice(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldDamageDice, v))
}

// WeaponCategoryEQ applies the EQ predicate on the "weapon_category" field.
func WeaponCategoryEQ(v WeaponCategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldWeaponCategory, v))
}

// WeaponCategoryNEQ applies the NEQ predicate on the "weapon_category" field.
func WeaponCategoryNEQ(v WeaponCategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldWeaponCategory, v))
}

// WeaponCategoryIn applies the In predicate on the "weapon_category" field.
func WeaponCategoryIn(vs ...WeaponCategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldWeaponCategory, vs...))
}

// WeaponCategoryNotIn applies the NotIn predicate on the "weapon_category" field.
func WeaponCategoryNotIn(vs ...WeaponCategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldWeaponCategory, vs...))
}

// WeaponSubcategoryEQ applies the EQ predicate on the "weapon_subcategory" field.
func WeaponSubcategoryEQ(v WeaponSubcategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldWeaponSubcategory, v))
}

// WeaponSubcategoryNEQ applies the NEQ predicate on the "weapon_subcategory" field.
func WeaponSubcategoryNEQ(v WeaponSubcategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldWeaponSubcategory, v))
}

// WeaponSubcategoryIn applies the In predicate on the "weapon_subcategory" field.
func WeaponSubcategoryIn(vs ...WeaponSubcategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldWeaponSubcategory, vs...))
}

// WeaponSubcategoryNotIn applies the NotIn predicate on the "weapon_subcategory" field.
func WeaponSubcategoryNotIn(vs ...WeaponSubcategory) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldWeaponSubcategory, vs...))
}

// RangeNormalEQ applies the EQ predicate on the "range_normal" field.
func RangeNormalEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldRangeNormal, v))
}

// RangeNormalNEQ applies the NEQ predicate on the "range_normal" field.
func RangeNormalNEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldRangeNormal, v))
}

// RangeNormalIn applies the In predicate on the "range_normal" field.
func RangeNormalIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldRangeNormal, vs...))
}

// RangeNormalNotIn applies the NotIn predicate on the "range_normal" field.
func RangeNormalNotIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldRangeNormal, vs...))
}

// RangeNormalGT applies the GT predicate on the "range_normal" field.
func RangeNormalGT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldRangeNormal, v))
}

// RangeNormalGTE applies the GTE predicate on the "range_normal" field.
func RangeNormalGTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldRangeNormal, v))
}

// RangeNormalLT applies the LT predicate on the "range_normal" field.
func RangeNormalLT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldRangeNormal, v))
}

// RangeNormalLTE applies the LTE predicate on the "range_normal" field.
func RangeNormalLTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldRangeNormal, v))
}

// RangeNormalIsNil applies the IsNil predicate on the "range_normal" field.
func RangeNormalIsNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldIsNull(FieldRangeNormal))
}

// RangeNormalNotNil applies the NotNil predicate on the "range_normal" field.
func RangeNormalNotNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldNotNull(FieldRangeNormal))
}

// RangeLongEQ applies the EQ predicate on the "range_long" field.
func RangeLongEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldRangeLong, v))
}

// RangeLongNEQ applies the NEQ predicate on the "range_long" field.
func RangeLongNEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldRangeLong, v))
}

// RangeLongIn applies the In predicate on the "range_long" field.
func RangeLongIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldRangeLong, vs...))
}

// RangeLongNotIn applies the NotIn predicate on the "range_long" field.
func RangeLongNotIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldRangeLong, vs...))
}

// RangeLongGT applies the GT predicate on the "range_long" field.
func RangeLongGT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldRangeLong, v))
}

// RangeLongGTE applies the GTE predicate on the "range_long" field.
func RangeLongGTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldRangeLong, v))
}

// RangeLongLT applies the LT predicate on the "range_long" field.
func RangeLongLT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldRangeLong, v))
}

// RangeLongLTE applies the LTE predicate on the "range_long" field.
func RangeLongLTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldRangeLong, v))
}

// RangeLongIsNil applies the IsNil predicate on the "range_long" field.
func RangeLongIsNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldIsNull(FieldRangeLong))
}

// RangeLongNotNil applies the NotNil predicate on the "range_long" field.
func RangeLongNotNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldNotNull(FieldRangeLong))
}

// ThrowRangeNormalEQ applies the EQ predicate on the "throw_range_normal" field.
func ThrowRangeNormalEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldThrowRangeNormal, v))
}

// ThrowRangeNormalNEQ applies the NEQ predicate on the "throw_range_normal" field.
func ThrowRangeNormalNEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldThrowRangeNormal, v))
}

// ThrowRangeNormalIn applies the In predicate on the "throw_range_normal" field.
func ThrowRangeNormalIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldThrowRangeNormal, vs...))
}

// ThrowRangeNormalNotIn applies the NotIn predicate on the "throw_range_normal" field.
func ThrowRangeNormalNotIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldThrowRangeNormal, vs...))
}

// ThrowRangeNormalGT applies the GT predicate on the "throw_range_normal" field.
func ThrowRangeNormalGT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldThrowRangeNormal, v))
}

// ThrowRangeNormalGTE applies the GTE predicate on the "throw_range_normal" field.
func ThrowRangeNormalGTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldThrowRangeNormal, v))
}

// ThrowRangeNormalLT applies the LT predicate on the "throw_range_normal" field.
func ThrowRangeNormalLT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldThrowRangeNormal, v))
}

// ThrowRangeNormalLTE applies the LTE predicate on the "throw_range_normal" field.
func ThrowRangeNormalLTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldThrowRangeNormal, v))
}

// ThrowRangeNormalIsNil applies the IsNil predicate on the "throw_range_normal" field.
func ThrowRangeNormalIsNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldIsNull(FieldThrowRangeNormal))
}

// ThrowRangeNormalNotNil applies the NotNil predicate on the "throw_range_normal" field.
func ThrowRangeNormalNotNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldNotNull(FieldThrowRangeNormal))
}

// ThrowRangeLongEQ applies the EQ predicate on the "throw_range_long" field.
func ThrowRangeLongEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldThrowRangeLong, v))
}

// ThrowRangeLongNEQ applies the NEQ predicate on the "throw_range_long" field.
func ThrowRangeLongNEQ(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldThrowRangeLong, v))
}

// ThrowRangeLongIn applies the In predicate on the "throw_range_long" field.
func ThrowRangeLongIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldThrowRangeLong, vs...))
}

// ThrowRangeLongNotIn applies the NotIn predicate on the "throw_range_long" field.
func ThrowRangeLongNotIn(vs ...int) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldThrowRangeLong, vs...))
}

// ThrowRangeLongGT applies the GT predicate on the "throw_range_long" field.
func ThrowRangeLongGT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldThrowRangeLong, v))
}

// ThrowRangeLongGTE applies the GTE predicate on the "throw_range_long" field.
func ThrowRangeLongGTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldThrowRangeLong, v))
}

// ThrowRangeLongLT applies the LT predicate on the "throw_range_long" field.
func ThrowRangeLongLT(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldThrowRangeLong, v))
}

// ThrowRangeLongLTE applies the LTE predicate on the "throw_range_long" field.
func ThrowRangeLongLTE(v int) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldThrowRangeLong, v))
}

// ThrowRangeLongIsNil applies the IsNil predicate on the "throw_range_long" field.
func ThrowRangeLongIsNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldIsNull(FieldThrowRangeLong))
}

// ThrowRangeLongNotNil applies the NotNil predicate on the "throw_range_long" field.
func ThrowRangeLongNotNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldNotNull(FieldThrowRangeLong))
}

// DamageDiceEQ applies the EQ predicate on the "damage_dice" field.
func DamageDiceEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEQ(FieldDamageDice, v))
}

// DamageDiceNEQ applies the NEQ predicate on the "damage_dice" field.
func DamageDiceNEQ(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNEQ(FieldDamageDice, v))
}

// DamageDiceIn applies the In predicate on the "damage_dice" field.
func DamageDiceIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldIn(FieldDamageDice, vs...))
}

// DamageDiceNotIn applies the NotIn predicate on the "damage_dice" field.
func DamageDiceNotIn(vs ...string) predicate.Weapon {
	return predicate.Weapon(sql.FieldNotIn(FieldDamageDice, vs...))
}

// DamageDiceGT applies the GT predicate on the "damage_dice" field.
func DamageDiceGT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGT(FieldDamageDice, v))
}

// DamageDiceGTE applies the GTE predicate on the "damage_dice" field.
func DamageDiceGTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldGTE(FieldDamageDice, v))
}

// DamageDiceLT applies the LT predicate on the "damage_dice" field.
func DamageDiceLT(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLT(FieldDamageDice, v))
}

// DamageDiceLTE applies the LTE predicate on the "damage_dice" field.
func DamageDiceLTE(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldLTE(FieldDamageDice, v))
}

// DamageDiceContains applies the Contains predicate on the "damage_dice" field.
func DamageDiceContains(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContains(FieldDamageDice, v))
}

// DamageDiceHasPrefix applies the HasPrefix predicate on the "damage_dice" field.
func DamageDiceHasPrefix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasPrefix(FieldDamageDice, v))
}

// DamageDiceHasSuffix applies the HasSuffix predicate on the "damage_dice" field.
func DamageDiceHasSuffix(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldHasSuffix(FieldDamageDice, v))
}

// DamageDiceIsNil applies the IsNil predicate on the "damage_dice" field.
func DamageDiceIsNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldIsNull(FieldDamageDice))
}

// DamageDiceNotNil applies the NotNil predicate on the "damage_dice" field.
func DamageDiceNotNil() predicate.Weapon {
	return predicate.Weapon(sql.FieldNotNull(FieldDamageDice))
}

// DamageDiceEqualFold applies the EqualFold predicate on the "damage_dice" field.
func DamageDiceEqualFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldEqualFold(FieldDamageDice, v))
}

// DamageDiceContainsFold applies the ContainsFold predicate on the "damage_dice" field.
func DamageDiceContainsFold(v string) predicate.Weapon {
	return predicate.Weapon(sql.FieldContainsFold(FieldDamageDice, v))
}

// HasProperties applies the HasEdge predicate on the "properties" edge.
func HasProperties() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, PropertiesTable, PropertiesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPropertiesWith applies the HasEdge predicate on the "properties" edge with a given conditions (other predicates).
func HasPropertiesWith(preds ...predicate.Property) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newPropertiesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDamageType applies the HasEdge predicate on the "damage_type" edge.
func HasDamageType() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DamageTypeTable, DamageTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDamageTypeWith applies the HasEdge predicate on the "damage_type" edge with a given conditions (other predicates).
func HasDamageTypeWith(preds ...predicate.DamageType) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newDamageTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Weapon) predicate.Weapon {
	return predicate.Weapon(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Weapon) predicate.Weapon {
	return predicate.Weapon(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Weapon) predicate.Weapon {
	return predicate.Weapon(sql.NotPredicates(p))
}
