// Code generated by ent, DO NOT EDIT.

package abilityscore

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldName, v))
}

// FullName applies equality check predicate on the "full_name" field. It's identical to FullNameEQ.
func FullName(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldFullName, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldContainsFold(FieldName, v))
}

// FullNameEQ applies the EQ predicate on the "full_name" field.
func FullNameEQ(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEQ(FieldFullName, v))
}

// FullNameNEQ applies the NEQ predicate on the "full_name" field.
func FullNameNEQ(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNEQ(FieldFullName, v))
}

// FullNameIn applies the In predicate on the "full_name" field.
func FullNameIn(vs ...string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldIn(FieldFullName, vs...))
}

// FullNameNotIn applies the NotIn predicate on the "full_name" field.
func FullNameNotIn(vs ...string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldNotIn(FieldFullName, vs...))
}

// FullNameGT applies the GT predicate on the "full_name" field.
func FullNameGT(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGT(FieldFullName, v))
}

// FullNameGTE applies the GTE predicate on the "full_name" field.
func FullNameGTE(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldGTE(FieldFullName, v))
}

// FullNameLT applies the LT predicate on the "full_name" field.
func FullNameLT(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLT(FieldFullName, v))
}

// FullNameLTE applies the LTE predicate on the "full_name" field.
func FullNameLTE(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldLTE(FieldFullName, v))
}

// FullNameContains applies the Contains predicate on the "full_name" field.
func FullNameContains(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldContains(FieldFullName, v))
}

// FullNameHasPrefix applies the HasPrefix predicate on the "full_name" field.
func FullNameHasPrefix(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldHasPrefix(FieldFullName, v))
}

// FullNameHasSuffix applies the HasSuffix predicate on the "full_name" field.
func FullNameHasSuffix(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldHasSuffix(FieldFullName, v))
}

// FullNameEqualFold applies the EqualFold predicate on the "full_name" field.
func FullNameEqualFold(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldEqualFold(FieldFullName, v))
}

// FullNameContainsFold applies the ContainsFold predicate on the "full_name" field.
func FullNameContainsFold(v string) predicate.AbilityScore {
	return predicate.AbilityScore(sql.FieldContainsFold(FieldFullName, v))
}

// HasSkills applies the HasEdge predicate on the "skills" edge.
func HasSkills() predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SkillsTable, SkillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkillsWith applies the HasEdge predicate on the "skills" edge with a given conditions (other predicates).
func HasSkillsWith(preds ...predicate.Skill) predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
		step := newSkillsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAbilityBonuses applies the HasEdge predicate on the "ability_bonuses" edge.
func HasAbilityBonuses() predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AbilityBonusesTable, AbilityBonusesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAbilityBonusesWith applies the HasEdge predicate on the "ability_bonuses" edge with a given conditions (other predicates).
func HasAbilityBonusesWith(preds ...predicate.AbilityBonus) predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
		step := newAbilityBonusesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AbilityScore) predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AbilityScore) predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
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
func Not(p predicate.AbilityScore) predicate.AbilityScore {
	return predicate.AbilityScore(func(s *sql.Selector) {
		p(s.Not())
	})
}
