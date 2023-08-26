// Code generated by ent, DO NOT EDIT.

package race

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldName, v))
}

// Speed applies equality check predicate on the "speed" field. It's identical to SpeedEQ.
func Speed(v int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSpeed, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Race {
	return predicate.Race(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Race {
	return predicate.Race(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Race {
	return predicate.Race(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Race {
	return predicate.Race(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Race {
	return predicate.Race(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Race {
	return predicate.Race(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Race {
	return predicate.Race(sql.FieldContainsFold(FieldName, v))
}

// SpeedEQ applies the EQ predicate on the "speed" field.
func SpeedEQ(v int) predicate.Race {
	return predicate.Race(sql.FieldEQ(FieldSpeed, v))
}

// SpeedNEQ applies the NEQ predicate on the "speed" field.
func SpeedNEQ(v int) predicate.Race {
	return predicate.Race(sql.FieldNEQ(FieldSpeed, v))
}

// SpeedIn applies the In predicate on the "speed" field.
func SpeedIn(vs ...int) predicate.Race {
	return predicate.Race(sql.FieldIn(FieldSpeed, vs...))
}

// SpeedNotIn applies the NotIn predicate on the "speed" field.
func SpeedNotIn(vs ...int) predicate.Race {
	return predicate.Race(sql.FieldNotIn(FieldSpeed, vs...))
}

// SpeedGT applies the GT predicate on the "speed" field.
func SpeedGT(v int) predicate.Race {
	return predicate.Race(sql.FieldGT(FieldSpeed, v))
}

// SpeedGTE applies the GTE predicate on the "speed" field.
func SpeedGTE(v int) predicate.Race {
	return predicate.Race(sql.FieldGTE(FieldSpeed, v))
}

// SpeedLT applies the LT predicate on the "speed" field.
func SpeedLT(v int) predicate.Race {
	return predicate.Race(sql.FieldLT(FieldSpeed, v))
}

// SpeedLTE applies the LTE predicate on the "speed" field.
func SpeedLTE(v int) predicate.Race {
	return predicate.Race(sql.FieldLTE(FieldSpeed, v))
}

// HasLanguages applies the HasEdge predicate on the "languages" edge.
func HasLanguages() predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, LanguagesTable, LanguagesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLanguagesWith applies the HasEdge predicate on the "languages" edge with a given conditions (other predicates).
func HasLanguagesWith(preds ...predicate.Language) predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
		step := newLanguagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProficiencies applies the HasEdge predicate on the "proficiencies" edge.
func HasProficiencies() predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ProficienciesTable, ProficienciesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProficienciesWith applies the HasEdge predicate on the "proficiencies" edge with a given conditions (other predicates).
func HasProficienciesWith(preds ...predicate.Proficiency) predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
		step := newProficienciesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Race) predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Race) predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
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
func Not(p predicate.Race) predicate.Race {
	return predicate.Race(func(s *sql.Selector) {
		p(s.Not())
	})
}
