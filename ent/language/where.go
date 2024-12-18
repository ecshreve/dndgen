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

// DescIsNil applies the IsNil predicate on the "desc" field.
func DescIsNil() predicate.Language {
	return predicate.Language(sql.FieldIsNull(FieldDesc))
}

// DescNotNil applies the NotNil predicate on the "desc" field.
func DescNotNil() predicate.Language {
	return predicate.Language(sql.FieldNotNull(FieldDesc))
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

// HasRace applies the HasEdge predicate on the "race" edge.
func HasRace() predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RaceTable, RacePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRaceWith applies the HasEdge predicate on the "race" edge with a given conditions (other predicates).
func HasRaceWith(preds ...predicate.Race) predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		step := newRaceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOptions applies the HasEdge predicate on the "options" edge.
func HasOptions() predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OptionsTable, OptionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOptionsWith applies the HasEdge predicate on the "options" edge with a given conditions (other predicates).
func HasOptionsWith(preds ...predicate.LanguageChoice) predicate.Language {
	return predicate.Language(func(s *sql.Selector) {
		step := newOptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Language) predicate.Language {
	return predicate.Language(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Language) predicate.Language {
	return predicate.Language(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Language) predicate.Language {
	return predicate.Language(sql.NotPredicates(p))
}
