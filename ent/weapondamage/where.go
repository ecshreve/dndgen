// Code generated by ent, DO NOT EDIT.

package weapondamage

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldLTE(FieldID, id))
}

// WeaponID applies equality check predicate on the "weapon_id" field. It's identical to WeaponIDEQ.
func WeaponID(v int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldWeaponID, v))
}

// DamageTypeID applies equality check predicate on the "damage_type_id" field. It's identical to DamageTypeIDEQ.
func DamageTypeID(v int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldDamageTypeID, v))
}

// Dice applies equality check predicate on the "dice" field. It's identical to DiceEQ.
func Dice(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldDice, v))
}

// WeaponIDEQ applies the EQ predicate on the "weapon_id" field.
func WeaponIDEQ(v int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldWeaponID, v))
}

// WeaponIDNEQ applies the NEQ predicate on the "weapon_id" field.
func WeaponIDNEQ(v int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNEQ(FieldWeaponID, v))
}

// WeaponIDIn applies the In predicate on the "weapon_id" field.
func WeaponIDIn(vs ...int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldIn(FieldWeaponID, vs...))
}

// WeaponIDNotIn applies the NotIn predicate on the "weapon_id" field.
func WeaponIDNotIn(vs ...int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNotIn(FieldWeaponID, vs...))
}

// DamageTypeIDEQ applies the EQ predicate on the "damage_type_id" field.
func DamageTypeIDEQ(v int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldDamageTypeID, v))
}

// DamageTypeIDNEQ applies the NEQ predicate on the "damage_type_id" field.
func DamageTypeIDNEQ(v int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNEQ(FieldDamageTypeID, v))
}

// DamageTypeIDIn applies the In predicate on the "damage_type_id" field.
func DamageTypeIDIn(vs ...int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldIn(FieldDamageTypeID, vs...))
}

// DamageTypeIDNotIn applies the NotIn predicate on the "damage_type_id" field.
func DamageTypeIDNotIn(vs ...int) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNotIn(FieldDamageTypeID, vs...))
}

// DiceEQ applies the EQ predicate on the "dice" field.
func DiceEQ(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEQ(FieldDice, v))
}

// DiceNEQ applies the NEQ predicate on the "dice" field.
func DiceNEQ(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNEQ(FieldDice, v))
}

// DiceIn applies the In predicate on the "dice" field.
func DiceIn(vs ...string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldIn(FieldDice, vs...))
}

// DiceNotIn applies the NotIn predicate on the "dice" field.
func DiceNotIn(vs ...string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldNotIn(FieldDice, vs...))
}

// DiceGT applies the GT predicate on the "dice" field.
func DiceGT(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldGT(FieldDice, v))
}

// DiceGTE applies the GTE predicate on the "dice" field.
func DiceGTE(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldGTE(FieldDice, v))
}

// DiceLT applies the LT predicate on the "dice" field.
func DiceLT(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldLT(FieldDice, v))
}

// DiceLTE applies the LTE predicate on the "dice" field.
func DiceLTE(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldLTE(FieldDice, v))
}

// DiceContains applies the Contains predicate on the "dice" field.
func DiceContains(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldContains(FieldDice, v))
}

// DiceHasPrefix applies the HasPrefix predicate on the "dice" field.
func DiceHasPrefix(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldHasPrefix(FieldDice, v))
}

// DiceHasSuffix applies the HasSuffix predicate on the "dice" field.
func DiceHasSuffix(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldHasSuffix(FieldDice, v))
}

// DiceEqualFold applies the EqualFold predicate on the "dice" field.
func DiceEqualFold(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldEqualFold(FieldDice, v))
}

// DiceContainsFold applies the ContainsFold predicate on the "dice" field.
func DiceContainsFold(v string) predicate.WeaponDamage {
	return predicate.WeaponDamage(sql.FieldContainsFold(FieldDice, v))
}

// HasWeapon applies the HasEdge predicate on the "weapon" edge.
func HasWeapon() predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WeaponTable, WeaponColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWeaponWith applies the HasEdge predicate on the "weapon" edge with a given conditions (other predicates).
func HasWeaponWith(preds ...predicate.Weapon) predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
		step := newWeaponStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDamageType applies the HasEdge predicate on the "damage_type" edge.
func HasDamageType() predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DamageTypeTable, DamageTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDamageTypeWith applies the HasEdge predicate on the "damage_type" edge with a given conditions (other predicates).
func HasDamageTypeWith(preds ...predicate.DamageType) predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
		step := newDamageTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WeaponDamage) predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WeaponDamage) predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
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
func Not(p predicate.WeaponDamage) predicate.WeaponDamage {
	return predicate.WeaponDamage(func(s *sql.Selector) {
		p(s.Not())
	})
}
