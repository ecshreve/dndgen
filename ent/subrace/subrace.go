// Code generated by ent, DO NOT EDIT.

package subrace

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subrace type in the database.
	Label = "subrace"
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
	// EdgeProficiencies holds the string denoting the proficiencies edge name in mutations.
	EdgeProficiencies = "proficiencies"
	// EdgeTraits holds the string denoting the traits edge name in mutations.
	EdgeTraits = "traits"
	// Table holds the table name of the subrace in the database.
	Table = "subraces"
	// RaceTable is the table that holds the race relation/edge.
	RaceTable = "subraces"
	// RaceInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	RaceInverseTable = "races"
	// RaceColumn is the table column denoting the race relation/edge.
	RaceColumn = "race_subrace"
	// ProficienciesTable is the table that holds the proficiencies relation/edge. The primary key declared below.
	ProficienciesTable = "subrace_proficiencies"
	// ProficienciesInverseTable is the table name for the Proficiency entity.
	// It exists in this package in order to avoid circular dependency with the "proficiency" package.
	ProficienciesInverseTable = "proficiencies"
	// TraitsTable is the table that holds the traits relation/edge. The primary key declared below.
	TraitsTable = "subrace_traits"
	// TraitsInverseTable is the table name for the Trait entity.
	// It exists in this package in order to avoid circular dependency with the "trait" package.
	TraitsInverseTable = "traits"
)

// Columns holds all SQL columns for subrace fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldDesc,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "subraces"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"race_subrace",
}

var (
	// ProficienciesPrimaryKey and ProficienciesColumn2 are the table columns denoting the
	// primary key for the proficiencies relation (M2M).
	ProficienciesPrimaryKey = []string{"subrace_id", "proficiency_id"}
	// TraitsPrimaryKey and TraitsColumn2 are the table columns denoting the
	// primary key for the traits relation (M2M).
	TraitsPrimaryKey = []string{"subrace_id", "trait_id"}
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

// OrderOption defines the ordering options for the Subrace queries.
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

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByRaceField orders the results by race field.
func ByRaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRaceStep(), sql.OrderByField(field, opts...))
	}
}

// ByProficienciesCount orders the results by proficiencies count.
func ByProficienciesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProficienciesStep(), opts...)
	}
}

// ByProficiencies orders the results by proficiencies terms.
func ByProficiencies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProficienciesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTraitsCount orders the results by traits count.
func ByTraitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTraitsStep(), opts...)
	}
}

// ByTraits orders the results by traits terms.
func ByTraits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTraitsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RaceTable, RaceColumn),
	)
}
func newProficienciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProficienciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProficienciesTable, ProficienciesPrimaryKey...),
	)
}
func newTraitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TraitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TraitsTable, TraitsPrimaryKey...),
	)
}
