// Code generated by ent, DO NOT EDIT.

package characterskill

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldLTE(FieldID, id))
}

// Proficient applies equality check predicate on the "proficient" field. It's identical to ProficientEQ.
func Proficient(v bool) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldProficient, v))
}

// CharacterID applies equality check predicate on the "character_id" field. It's identical to CharacterIDEQ.
func CharacterID(v int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldCharacterID, v))
}

// SkillID applies equality check predicate on the "skill_id" field. It's identical to SkillIDEQ.
func SkillID(v int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldSkillID, v))
}

// ProficientEQ applies the EQ predicate on the "proficient" field.
func ProficientEQ(v bool) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldProficient, v))
}

// ProficientNEQ applies the NEQ predicate on the "proficient" field.
func ProficientNEQ(v bool) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNEQ(FieldProficient, v))
}

// CharacterIDEQ applies the EQ predicate on the "character_id" field.
func CharacterIDEQ(v int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldCharacterID, v))
}

// CharacterIDNEQ applies the NEQ predicate on the "character_id" field.
func CharacterIDNEQ(v int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNEQ(FieldCharacterID, v))
}

// CharacterIDIn applies the In predicate on the "character_id" field.
func CharacterIDIn(vs ...int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldIn(FieldCharacterID, vs...))
}

// CharacterIDNotIn applies the NotIn predicate on the "character_id" field.
func CharacterIDNotIn(vs ...int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNotIn(FieldCharacterID, vs...))
}

// SkillIDEQ applies the EQ predicate on the "skill_id" field.
func SkillIDEQ(v int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldEQ(FieldSkillID, v))
}

// SkillIDNEQ applies the NEQ predicate on the "skill_id" field.
func SkillIDNEQ(v int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNEQ(FieldSkillID, v))
}

// SkillIDIn applies the In predicate on the "skill_id" field.
func SkillIDIn(vs ...int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldIn(FieldSkillID, vs...))
}

// SkillIDNotIn applies the NotIn predicate on the "skill_id" field.
func SkillIDNotIn(vs ...int) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.FieldNotIn(FieldSkillID, vs...))
}

// HasCharacter applies the HasEdge predicate on the "character" edge.
func HasCharacter() predicate.CharacterSkill {
	return predicate.CharacterSkill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CharacterTable, CharacterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCharacterWith applies the HasEdge predicate on the "character" edge with a given conditions (other predicates).
func HasCharacterWith(preds ...predicate.Character) predicate.CharacterSkill {
	return predicate.CharacterSkill(func(s *sql.Selector) {
		step := newCharacterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkill applies the HasEdge predicate on the "skill" edge.
func HasSkill() predicate.CharacterSkill {
	return predicate.CharacterSkill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SkillTable, SkillColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkillWith applies the HasEdge predicate on the "skill" edge with a given conditions (other predicates).
func HasSkillWith(preds ...predicate.Skill) predicate.CharacterSkill {
	return predicate.CharacterSkill(func(s *sql.Selector) {
		step := newSkillStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCharacterAbilityScore applies the HasEdge predicate on the "character_ability_score" edge.
func HasCharacterAbilityScore() predicate.CharacterSkill {
	return predicate.CharacterSkill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CharacterAbilityScoreTable, CharacterAbilityScoreColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCharacterAbilityScoreWith applies the HasEdge predicate on the "character_ability_score" edge with a given conditions (other predicates).
func HasCharacterAbilityScoreWith(preds ...predicate.CharacterAbilityScore) predicate.CharacterSkill {
	return predicate.CharacterSkill(func(s *sql.Selector) {
		step := newCharacterAbilityScoreStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CharacterSkill) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CharacterSkill) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CharacterSkill) predicate.CharacterSkill {
	return predicate.CharacterSkill(sql.NotPredicates(p))
}
