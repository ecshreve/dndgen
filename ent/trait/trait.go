// Code generated by ent, DO NOT EDIT.

package trait

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the trait type in the database.
	Label = "trait"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeRace holds the string denoting the race edge name in mutations.
	EdgeRace = "race"
	// EdgeSubrace holds the string denoting the subrace edge name in mutations.
	EdgeSubrace = "subrace"
	// Table holds the table name of the trait in the database.
	Table = "traits"
	// RaceTable is the table that holds the race relation/edge. The primary key declared below.
	RaceTable = "race_traits"
	// RaceInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	RaceInverseTable = "races"
	// SubraceTable is the table that holds the subrace relation/edge. The primary key declared below.
	SubraceTable = "subrace_traits"
	// SubraceInverseTable is the table name for the Subrace entity.
	// It exists in this package in order to avoid circular dependency with the "subrace" package.
	SubraceInverseTable = "subraces"
)

// Columns holds all SQL columns for trait fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldDesc,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "traits"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"character_traits",
}

var (
	// RacePrimaryKey and RaceColumn2 are the table columns denoting the
	// primary key for the race relation (M2M).
	RacePrimaryKey = []string{"race_id", "trait_id"}
	// SubracePrimaryKey and SubraceColumn2 are the table columns denoting the
	// primary key for the subrace relation (M2M).
	SubracePrimaryKey = []string{"subrace_id", "trait_id"}
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
	// IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	IndxValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Trait queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIndx orders the results by the indx field.
func ByIndx(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndx, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
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

// BySubraceCount orders the results by subrace count.
func BySubraceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubraceStep(), opts...)
	}
}

// BySubrace orders the results by subrace terms.
func BySubrace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubraceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RaceTable, RacePrimaryKey...),
	)
}
func newSubraceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubraceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SubraceTable, SubracePrimaryKey...),
	)
}
