// Code generated by ent, DO NOT EDIT.

package character

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldName, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldAge, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldLevel, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Character {
	return predicate.Character(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Character {
	return predicate.Character(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Character {
	return predicate.Character(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Character {
	return predicate.Character(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Character {
	return predicate.Character(sql.FieldContainsFold(FieldName, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldAge, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.Character {
	return predicate.Character(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.Character {
	return predicate.Character(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.Character {
	return predicate.Character(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.Character {
	return predicate.Character(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.Character {
	return predicate.Character(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.Character {
	return predicate.Character(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.Character {
	return predicate.Character(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.Character {
	return predicate.Character(sql.FieldLTE(FieldLevel, v))
}

// HasRace applies the HasEdge predicate on the "race" edge.
func HasRace() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RaceTable, RaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRaceWith applies the HasEdge predicate on the "race" edge with a given conditions (other predicates).
func HasRaceWith(preds ...predicate.Race) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newRaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ClassTable, ClassColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newClassStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAlignment applies the HasEdge predicate on the "alignment" edge.
func HasAlignment() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AlignmentTable, AlignmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAlignmentWith applies the HasEdge predicate on the "alignment" edge with a given conditions (other predicates).
func HasAlignmentWith(preds ...predicate.Alignment) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newAlignmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTraits applies the HasEdge predicate on the "traits" edge.
func HasTraits() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TraitsTable, TraitsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTraitsWith applies the HasEdge predicate on the "traits" edge with a given conditions (other predicates).
func HasTraitsWith(preds ...predicate.Trait) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newTraitsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLanguages applies the HasEdge predicate on the "languages" edge.
func HasLanguages() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LanguagesTable, LanguagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLanguagesWith applies the HasEdge predicate on the "languages" edge with a given conditions (other predicates).
func HasLanguagesWith(preds ...predicate.Language) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newLanguagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProficiencies applies the HasEdge predicate on the "proficiencies" edge.
func HasProficiencies() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProficienciesTable, ProficienciesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProficienciesWith applies the HasEdge predicate on the "proficiencies" edge with a given conditions (other predicates).
func HasProficienciesWith(preds ...predicate.Proficiency) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newProficienciesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAbilityScores applies the HasEdge predicate on the "ability_scores" edge.
func HasAbilityScores() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AbilityScoresTable, AbilityScoresPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAbilityScoresWith applies the HasEdge predicate on the "ability_scores" edge with a given conditions (other predicates).
func HasAbilityScoresWith(preds ...predicate.AbilityScore) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newAbilityScoresStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCharacterAbilityScores applies the HasEdge predicate on the "character_ability_scores" edge.
func HasCharacterAbilityScores() predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CharacterAbilityScoresTable, CharacterAbilityScoresColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCharacterAbilityScoresWith applies the HasEdge predicate on the "character_ability_scores" edge with a given conditions (other predicates).
func HasCharacterAbilityScoresWith(preds ...predicate.CharacterAbilityScore) predicate.Character {
	return predicate.Character(func(s *sql.Selector) {
		step := newCharacterAbilityScoresStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Character) predicate.Character {
	return predicate.Character(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Character) predicate.Character {
	return predicate.Character(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Character) predicate.Character {
	return predicate.Character(sql.NotPredicates(p))
}