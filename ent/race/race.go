// Code generated by ent, DO NOT EDIT.

package race

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the race type in the database.
	Label = "race"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSpeed holds the string denoting the speed field in the database.
	FieldSpeed = "speed"
	// EdgeLanguages holds the string denoting the languages edge name in mutations.
	EdgeLanguages = "languages"
	// EdgeProficiencies holds the string denoting the proficiencies edge name in mutations.
	EdgeProficiencies = "proficiencies"
	// Table holds the table name of the race in the database.
	Table = "races"
	// LanguagesTable is the table that holds the languages relation/edge. The primary key declared below.
	LanguagesTable = "race_languages"
	// LanguagesInverseTable is the table name for the Language entity.
	// It exists in this package in order to avoid circular dependency with the "language" package.
	LanguagesInverseTable = "languages"
	// ProficienciesTable is the table that holds the proficiencies relation/edge. The primary key declared below.
	ProficienciesTable = "race_proficiencies"
	// ProficienciesInverseTable is the table name for the Proficiency entity.
	// It exists in this package in order to avoid circular dependency with the "proficiency" package.
	ProficienciesInverseTable = "proficiencies"
)

// Columns holds all SQL columns for race fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldSpeed,
}

var (
	// LanguagesPrimaryKey and LanguagesColumn2 are the table columns denoting the
	// primary key for the languages relation (M2M).
	LanguagesPrimaryKey = []string{"race_id", "language_id"}
	// ProficienciesPrimaryKey and ProficienciesColumn2 are the table columns denoting the
	// primary key for the proficiencies relation (M2M).
	ProficienciesPrimaryKey = []string{"race_id", "proficiency_id"}
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

var (
	// IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	IndxValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Race queries.
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

// BySpeed orders the results by the speed field.
func BySpeed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpeed, opts...).ToFunc()
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
func newLanguagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LanguagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LanguagesTable, LanguagesPrimaryKey...),
	)
}
func newProficienciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProficienciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProficienciesTable, ProficienciesPrimaryKey...),
	)
}
