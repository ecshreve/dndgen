// Code generated by ent, DO NOT EDIT.

package proficiencychoice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the proficiencychoice type in the database.
	Label = "proficiency_choice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldChoose holds the string denoting the choose field in the database.
	FieldChoose = "choose"
	// EdgeOptions holds the string denoting the options edge name in mutations.
	EdgeOptions = "options"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeRace holds the string denoting the race edge name in mutations.
	EdgeRace = "race"
	// Table holds the table name of the proficiencychoice in the database.
	Table = "proficiency_choices"
	// OptionsTable is the table that holds the options relation/edge. The primary key declared below.
	OptionsTable = "proficiency_choice_options"
	// OptionsInverseTable is the table name for the Proficiency entity.
	// It exists in this package in order to avoid circular dependency with the "proficiency" package.
	OptionsInverseTable = "proficiencies"
	// ClassTable is the table that holds the class relation/edge. The primary key declared below.
	ClassTable = "class_proficiency_choices"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// RaceTable is the table that holds the race relation/edge.
	RaceTable = "races"
	// RaceInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	RaceInverseTable = "races"
	// RaceColumn is the table column denoting the race relation/edge.
	RaceColumn = "race_starting_proficiency_option"
)

// Columns holds all SQL columns for proficiencychoice fields.
var Columns = []string{
	FieldID,
	FieldDesc,
	FieldChoose,
}

var (
	// OptionsPrimaryKey and OptionsColumn2 are the table columns denoting the
	// primary key for the options relation (M2M).
	OptionsPrimaryKey = []string{"proficiency_choice_id", "proficiency_id"}
	// ClassPrimaryKey and ClassColumn2 are the table columns denoting the
	// primary key for the class relation (M2M).
	ClassPrimaryKey = []string{"class_id", "proficiency_choice_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the ProficiencyChoice queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByChoose orders the results by the choose field.
func ByChoose(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChoose, opts...).ToFunc()
}

// ByOptionsCount orders the results by options count.
func ByOptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOptionsStep(), opts...)
	}
}

// ByOptions orders the results by options terms.
func ByOptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByClassCount orders the results by class count.
func ByClassCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClassStep(), opts...)
	}
}

// ByClass orders the results by class terms.
func ByClass(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRaceCount orders the results by race count.
func ByRaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRaceStep(), opts...)
	}
}

// ByRace orders the results by race terms.
func ByRace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, OptionsTable, OptionsPrimaryKey...),
	)
}
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ClassTable, ClassPrimaryKey...),
	)
}
func newRaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, RaceTable, RaceColumn),
	)
}
