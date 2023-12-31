// Code generated by ent, DO NOT EDIT.

package language

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Language {
	return predicate.Language(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Language {
	return predicate.Language(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Language {
	return predicate.Language(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Language {
	return predicate.Language(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Language {
	return predicate.Language(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Language {
	return predicate.Language(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Language {
	return predicate.Language(sql.FieldLTE(FieldID, id))
}

// Indx applies equality check predicate on the "indx" field. It's identical to IndxEQ.
func Indx(v string) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldIndx, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldName, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldDesc, v))
}

// IndxEQ applies the EQ predicate on the "indx" field.
func IndxEQ(v string) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldIndx, v))
}

// IndxNEQ applies the NEQ predicate on the "indx" field.
func IndxNEQ(v string) predicate.Language {
	return predicate.Language(sql.FieldNEQ(FieldIndx, v))
}

// IndxIn applies the In predicate on the "indx" field.
func IndxIn(vs ...string) predicate.Language {
	return predicate.Language(sql.FieldIn(FieldIndx, vs...))
}

// IndxNotIn applies the NotIn predicate on the "indx" field.
func IndxNotIn(vs ...string) predicate.Language {
	return predicate.Language(sql.FieldNotIn(FieldIndx, vs...))
}

// IndxGT applies the GT predicate on the "indx" field.
func IndxGT(v string) predicate.Language {
	return predicate.Language(sql.FieldGT(FieldIndx, v))
}

// IndxGTE applies the GTE predicate on the "indx" field.
func IndxGTE(v string) predicate.Language {
	return predicate.Language(sql.FieldGTE(FieldIndx, v))
}

// IndxLT applies the LT predicate on the "indx" field.
func IndxLT(v string) predicate.Language {
	return predicate.Language(sql.FieldLT(FieldIndx, v))
}

// IndxLTE applies the LTE predicate on the "indx" field.
func IndxLTE(v string) predicate.Language {
	return predicate.Language(sql.FieldLTE(FieldIndx, v))
}

// IndxContains applies the Contains predicate on the "indx" field.
func IndxContains(v string) predicate.Language {
	return predicate.Language(sql.FieldContains(FieldIndx, v))
}

// IndxHasPrefix applies the HasPrefix predicate on the "indx" field.
func IndxHasPrefix(v string) predicate.Language {
	return predicate.Language(sql.FieldHasPrefix(FieldIndx, v))
}

// IndxHasSuffix applies the HasSuffix predicate on the "indx" field.
func IndxHasSuffix(v string) predicate.Language {
	return predicate.Language(sql.FieldHasSuffix(FieldIndx, v))
}

// IndxEqualFold applies the EqualFold predicate on the "indx" field.
func IndxEqualFold(v string) predicate.Language {
	return predicate.Language(sql.FieldEqualFold(FieldIndx, v))
}

// IndxContainsFold applies the ContainsFold predicate on the "indx" field.
func IndxContainsFold(v string) predicate.Language {
	return predicate.Language(sql.FieldContainsFold(FieldIndx, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Language {
	return predicate.Language(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Language {
	return predicate.Language(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Language {
	return predicate.Language(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Language {
	return predicate.Language(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Language {
	return predicate.Language(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Language {
	return predicate.Language(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Language {
	return predicate.Language(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Language {
	return predicate.Language(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Language {
	return predicate.Language(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Language {
	return predicate.Language(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Language {
	return predicate.Language(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Language {
	return predicate.Language(sql.FieldContainsFold(FieldName, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Language {
	return predicate.Language(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Language {
	return predicate.Language(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Language {
	return predicate.Language(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Language {
	return predicate.Language(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Language {
	return predicate.Language(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Language {
	return predicate.Language(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Language {
	return predicate.Language(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Language {
	return predicate.Language(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Language {
	return predicate.Language(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Language {
	return predicate.Language(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Language {
	return predicate.Language(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Language {
	return predicate.Language(sql.FieldContainsFold(FieldDesc, v))
}

// LanguageTypeEQ applies the EQ predicate on the "language_type" field.
func LanguageTypeEQ(v LanguageType) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldLanguageType, v))
}

// LanguageTypeNEQ applies the NEQ predicate on the "language_type" field.
func LanguageTypeNEQ(v LanguageType) predicate.Language {
	return predicate.Language(sql.FieldNEQ(FieldLanguageType, v))
}

// LanguageTypeIn applies the In predicate on the "language_type" field.
func LanguageTypeIn(vs ...LanguageType) predicate.Language {
	return predicate.Language(sql.FieldIn(FieldLanguageType, vs...))
}

// LanguageTypeNotIn applies the NotIn predicate on the "language_type" field.
func LanguageTypeNotIn(vs ...LanguageType) predicate.Language {
	return predicate.Language(sql.FieldNotIn(FieldLanguageType, vs...))
}

// ScriptEQ applies the EQ predicate on the "script" field.
func ScriptEQ(v Script) predicate.Language {
	return predicate.Language(sql.FieldEQ(FieldScript, v))
}

// ScriptNEQ applies the NEQ predicate on the "script" field.
func ScriptNEQ(v Script) predicate.Language {
	return predicate.Language(sql.FieldNEQ(FieldScript, v))
}

// ScriptIn applies the In predicate on the "script" field.
func ScriptIn(vs ...Script) predicate.Language {
	return predicate.Language(sql.FieldIn(FieldScript, vs...))
}

// ScriptNotIn applies the NotIn predicate on the "script" field.
func ScriptNotIn(vs ...Script) predicate.Language {
	return predicate.Language(sql.FieldNotIn(FieldScript, vs...))
}

// ScriptIsNil applies the IsNil predicate on the "script" field.
func ScriptIsNil() predicate.Language {
	return predicate.Language(sql.FieldIsNull(FieldScript))
}

// ScriptNotNil applies the NotNil predicate on the "script" field.
func ScriptNotNil() predicate.Language {
	return predicate.Language(sql.FieldNotNull(FieldScript))
}

// HasRaceSpeakers applies the HasEdge predicate on the "race_speakers" edge.
func HasRaceSpeakers() predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RaceSpeakersTable, RaceSpeakersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRaceSpeakersWith applies the HasEdge predicate on the "race_speakers" edge with a given conditions (other predicates).
func HasRaceSpeakersWith(preds ...predicate.Race) predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		step := newRaceSpeakersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Language) predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Language) predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
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
func Not(p predicate.Language) predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		p(s.Not())
	})
}
