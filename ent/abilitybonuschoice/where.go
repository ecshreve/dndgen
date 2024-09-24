// Code generated by ent, DO NOT EDIT.

package abilitybonuschoice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldLTE(FieldID, id))
}

// Choose applies equality check predicate on the "choose" field. It's identical to ChooseEQ.
func Choose(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldEQ(FieldChoose, v))
}

// ChooseEQ applies the EQ predicate on the "choose" field.
func ChooseEQ(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldEQ(FieldChoose, v))
}

// ChooseNEQ applies the NEQ predicate on the "choose" field.
func ChooseNEQ(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldNEQ(FieldChoose, v))
}

// ChooseIn applies the In predicate on the "choose" field.
func ChooseIn(vs ...int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldIn(FieldChoose, vs...))
}

// ChooseNotIn applies the NotIn predicate on the "choose" field.
func ChooseNotIn(vs ...int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldNotIn(FieldChoose, vs...))
}

// ChooseGT applies the GT predicate on the "choose" field.
func ChooseGT(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldGT(FieldChoose, v))
}

// ChooseGTE applies the GTE predicate on the "choose" field.
func ChooseGTE(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldGTE(FieldChoose, v))
}

// ChooseLT applies the LT predicate on the "choose" field.
func ChooseLT(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldLT(FieldChoose, v))
}

// ChooseLTE applies the LTE predicate on the "choose" field.
func ChooseLTE(v int) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.FieldLTE(FieldChoose, v))
}

// HasAbilityBonuses applies the HasEdge predicate on the "ability_bonuses" edge.
func HasAbilityBonuses() predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AbilityBonusesTable, AbilityBonusesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAbilityBonusesWith applies the HasEdge predicate on the "ability_bonuses" edge with a given conditions (other predicates).
func HasAbilityBonusesWith(preds ...predicate.AbilityBonus) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(func(s *sql.Selector) {
		step := newAbilityBonusesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRace applies the HasEdge predicate on the "race" edge.
func HasRace() predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RaceTable, RaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRaceWith applies the HasEdge predicate on the "race" edge with a given conditions (other predicates).
func HasRaceWith(preds ...predicate.Race) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(func(s *sql.Selector) {
		step := newRaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AbilityBonusChoice) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AbilityBonusChoice) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AbilityBonusChoice) predicate.AbilityBonusChoice {
	return predicate.AbilityBonusChoice(sql.NotPredicates(p))
}