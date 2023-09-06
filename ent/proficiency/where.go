// Code generated by ent, DO NOT EDIT.

package proficiency

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldName, v))
}

// ProficiencyCategory applies equality check predicate on the "proficiency_category" field. It's identical to ProficiencyCategoryEQ.
func ProficiencyCategory(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldProficiencyCategory, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldContainsFold(FieldName, v))
}

// ProficiencyCategoryEQ applies the EQ predicate on the "proficiency_category" field.
func ProficiencyCategoryEQ(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEQ(FieldProficiencyCategory, v))
}

// ProficiencyCategoryNEQ applies the NEQ predicate on the "proficiency_category" field.
func ProficiencyCategoryNEQ(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNEQ(FieldProficiencyCategory, v))
}

// ProficiencyCategoryIn applies the In predicate on the "proficiency_category" field.
func ProficiencyCategoryIn(vs ...string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldIn(FieldProficiencyCategory, vs...))
}

// ProficiencyCategoryNotIn applies the NotIn predicate on the "proficiency_category" field.
func ProficiencyCategoryNotIn(vs ...string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldNotIn(FieldProficiencyCategory, vs...))
}

// ProficiencyCategoryGT applies the GT predicate on the "proficiency_category" field.
func ProficiencyCategoryGT(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGT(FieldProficiencyCategory, v))
}

// ProficiencyCategoryGTE applies the GTE predicate on the "proficiency_category" field.
func ProficiencyCategoryGTE(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldGTE(FieldProficiencyCategory, v))
}

// ProficiencyCategoryLT applies the LT predicate on the "proficiency_category" field.
func ProficiencyCategoryLT(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLT(FieldProficiencyCategory, v))
}

// ProficiencyCategoryLTE applies the LTE predicate on the "proficiency_category" field.
func ProficiencyCategoryLTE(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldLTE(FieldProficiencyCategory, v))
}

// ProficiencyCategoryContains applies the Contains predicate on the "proficiency_category" field.
func ProficiencyCategoryContains(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldContains(FieldProficiencyCategory, v))
}

// ProficiencyCategoryHasPrefix applies the HasPrefix predicate on the "proficiency_category" field.
func ProficiencyCategoryHasPrefix(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldHasPrefix(FieldProficiencyCategory, v))
}

// ProficiencyCategoryHasSuffix applies the HasSuffix predicate on the "proficiency_category" field.
func ProficiencyCategoryHasSuffix(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldHasSuffix(FieldProficiencyCategory, v))
}

// ProficiencyCategoryEqualFold applies the EqualFold predicate on the "proficiency_category" field.
func ProficiencyCategoryEqualFold(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldEqualFold(FieldProficiencyCategory, v))
}

// ProficiencyCategoryContainsFold applies the ContainsFold predicate on the "proficiency_category" field.
func ProficiencyCategoryContainsFold(v string) predicate.Proficiency {
	return predicate.Proficiency(sql.FieldContainsFold(FieldProficiencyCategory, v))
}

// HasClasses applies the HasEdge predicate on the "classes" edge.
func HasClasses() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ClassesTable, ClassesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassesWith applies the HasEdge predicate on the "classes" edge with a given conditions (other predicates).
func HasClassesWith(preds ...predicate.Class) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newClassesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRaces applies the HasEdge predicate on the "races" edge.
func HasRaces() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RacesTable, RacesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRacesWith applies the HasEdge predicate on the "races" edge with a given conditions (other predicates).
func HasRacesWith(preds ...predicate.Race) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newRacesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubraces applies the HasEdge predicate on the "subraces" edge.
func HasSubraces() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubracesTable, SubracesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubracesWith applies the HasEdge predicate on the "subraces" edge with a given conditions (other predicates).
func HasSubracesWith(preds ...predicate.Subrace) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newSubracesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChoice applies the HasEdge predicate on the "choice" edge.
func HasChoice() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ChoiceTable, ChoicePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChoiceWith applies the HasEdge predicate on the "choice" edge with a given conditions (other predicates).
func HasChoiceWith(preds ...predicate.ProficiencyChoice) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newChoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSkill applies the HasEdge predicate on the "skill" edge.
func HasSkill() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SkillTable, SkillColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSkillWith applies the HasEdge predicate on the "skill" edge with a given conditions (other predicates).
func HasSkillWith(preds ...predicate.Skill) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newSkillStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEquipment applies the HasEdge predicate on the "equipment" edge.
func HasEquipment() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, EquipmentTable, EquipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentWith applies the HasEdge predicate on the "equipment" edge with a given conditions (other predicates).
func HasEquipmentWith(preds ...predicate.Equipment) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newEquipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSavingThrow applies the HasEdge predicate on the "saving_throw" edge.
func HasSavingThrow() predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SavingThrowTable, SavingThrowColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSavingThrowWith applies the HasEdge predicate on the "saving_throw" edge with a given conditions (other predicates).
func HasSavingThrowWith(preds ...predicate.AbilityScore) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		step := newSavingThrowStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Proficiency) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Proficiency) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
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
func Not(p predicate.Proficiency) predicate.Proficiency {
	return predicate.Proficiency(func(s *sql.Selector) {
		p(s.Not())
	})
}
