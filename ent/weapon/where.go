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

// HasDamage applies the HasEdge predicate on the "damage" edge.
func HasDamage() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DamageTable, DamageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDamageWith applies the HasEdge predicate on the "damage" edge with a given conditions (other predicates).
func HasDamageWith(preds ...predicate.Damage) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newDamageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
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

// HasWeaponRange applies the HasEdge predicate on the "weapon_range" edge.
func HasWeaponRange() predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WeaponRangeTable, WeaponRangeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeaponRangeWith applies the HasEdge predicate on the "weapon_range" edge with a given conditions (other predicates).
func HasWeaponRangeWith(preds ...predicate.WeaponRange) predicate.Weapon {
	return predicate.Weapon(func(s *sql.Selector) {
		step := newWeaponRangeStep()
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
