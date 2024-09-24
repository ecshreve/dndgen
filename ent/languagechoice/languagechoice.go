// Code generated by ent, DO NOT EDIT.

package languagechoice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the languagechoice type in the database.
	Label = "language_choice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChoose holds the string denoting the choose field in the database.
	FieldChoose = "choose"
	// EdgeLanguages holds the string denoting the languages edge name in mutations.
	EdgeLanguages = "languages"
	// EdgeRace holds the string denoting the race edge name in mutations.
	EdgeRace = "race"
	// EdgeSubrace holds the string denoting the subrace edge name in mutations.
	EdgeSubrace = "subrace"
	// Table holds the table name of the languagechoice in the database.
	Table = "language_choices"
	// LanguagesTable is the table that holds the languages relation/edge. The primary key declared below.
	LanguagesTable = "language_choice_languages"
	// LanguagesInverseTable is the table name for the Language entity.
	// It exists in this package in order to avoid circular dependency with the "language" package.
	LanguagesInverseTable = "languages"
	// RaceTable is the table that holds the race relation/edge.
	RaceTable = "language_choices"
	// RaceInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	RaceInverseTable = "races"
	// RaceColumn is the table column denoting the race relation/edge.
	RaceColumn = "race_language_options"
	// SubraceTable is the table that holds the subrace relation/edge.
	SubraceTable = "language_choices"
	// SubraceInverseTable is the table name for the Subrace entity.
	// It exists in this package in order to avoid circular dependency with the "subrace" package.
	SubraceInverseTable = "subraces"
	// SubraceColumn is the table column denoting the subrace relation/edge.
	SubraceColumn = "subrace_language_options"
)

// Columns holds all SQL columns for languagechoice fields.
var Columns = []string{
	FieldID,
	FieldChoose,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "language_choices"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"race_language_options",
	"subrace_language_options",
}

var (
	// LanguagesPrimaryKey and LanguagesColumn2 are the table columns denoting the
	// primary key for the languages relation (M2M).
	LanguagesPrimaryKey = []string{"language_choice_id", "language_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ChooseValidator is a validator for the "choose" field. It is called by the builders before save.
	ChooseValidator func(int) error
)

// OrderOption defines the ordering options for the LanguageChoice queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChoose orders the results by the choose field.
func ByChoose(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChoose, opts...).ToFunc()
}

// ByLanguagesCount orders the results by languages count.
func ByLanguagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLanguagesStep(), opts...)
	}
}

// ByLanguages orders the results by languages terms.
func ByLanguages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLanguagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRaceField orders the results by race field.
func ByRaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRaceStep(), sql.OrderByField(field, opts...))
	}
}

// BySubraceField orders the results by subrace field.
func BySubraceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubraceStep(), sql.OrderByField(field, opts...))
	}
}
func newLanguagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LanguagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LanguagesTable, LanguagesPrimaryKey...),
	)
}
func newRaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RaceTable, RaceColumn),
	)
}
func newSubraceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubraceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SubraceTable, SubraceColumn),
	)
}
