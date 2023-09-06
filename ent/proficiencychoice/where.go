// Code generated by ent, DO NOT EDIT.

package proficiencychoice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldLTE(FieldID, id))
}

// Choose applies equality check predicate on the "choose" field. It's identical to ChooseEQ.
func Choose(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEQ(FieldChoose, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEQ(FieldDesc, v))
}

// ChooseEQ applies the EQ predicate on the "choose" field.
func ChooseEQ(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEQ(FieldChoose, v))
}

// ChooseNEQ applies the NEQ predicate on the "choose" field.
func ChooseNEQ(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNEQ(FieldChoose, v))
}

// ChooseIn applies the In predicate on the "choose" field.
func ChooseIn(vs ...int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldIn(FieldChoose, vs...))
}

// ChooseNotIn applies the NotIn predicate on the "choose" field.
func ChooseNotIn(vs ...int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNotIn(FieldChoose, vs...))
}

// ChooseGT applies the GT predicate on the "choose" field.
func ChooseGT(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldGT(FieldChoose, v))
}

// ChooseGTE applies the GTE predicate on the "choose" field.
func ChooseGTE(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldGTE(FieldChoose, v))
}

// ChooseLT applies the LT predicate on the "choose" field.
func ChooseLT(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldLT(FieldChoose, v))
}

// ChooseLTE applies the LTE predicate on the "choose" field.
func ChooseLTE(v int) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldLTE(FieldChoose, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldHasSuffix(FieldDesc, v))
}

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldNotNull(FieldDesc))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(sql.FieldContainsFold(FieldDesc, v))
}

// HasProficiency applies the HasEdge predicate on the "proficiency" edge.
func HasProficiency() predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProficiencyTable, ProficiencyPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProficiencyWith applies the HasEdge predicate on the "proficiency" edge with a given conditions (other predicates).
func HasProficiencyWith(preds ...predicate.Proficiency) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := newProficiencyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParentChoice applies the HasEdge predicate on the "parent_choice" edge.
func HasParentChoice() predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentChoiceTable, ParentChoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentChoiceWith applies the HasEdge predicate on the "parent_choice" edge with a given conditions (other predicates).
func HasParentChoiceWith(preds ...predicate.ProficiencyChoice) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := newParentChoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubChoice applies the HasEdge predicate on the "sub_choice" edge.
func HasSubChoice() predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SubChoiceTable, SubChoiceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubChoiceWith applies the HasEdge predicate on the "sub_choice" edge with a given conditions (other predicates).
func HasSubChoiceWith(preds ...predicate.ProficiencyChoice) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := newSubChoiceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasClass applies the HasEdge predicate on the "class" edge.
func HasClass() predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ClassTable, ClassPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClassWith applies the HasEdge predicate on the "class" edge with a given conditions (other predicates).
func HasClassWith(preds ...predicate.Class) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		step := newClassStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProficiencyChoice) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProficiencyChoice) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
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
func Not(p predicate.ProficiencyChoice) predicate.ProficiencyChoice {
	return predicate.ProficiencyChoice(func(s *sql.Selector) {
		p(s.Not())
	})
}
