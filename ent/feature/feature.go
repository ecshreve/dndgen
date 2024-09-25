// Code generated by ent, DO NOT EDIT.

package feature

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the feature type in the database.
	Label = "feature"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// EdgePrerequisites holds the string denoting the prerequisites edge name in mutations.
	EdgePrerequisites = "prerequisites"
	// Table holds the table name of the feature in the database.
	Table = "features"
	// PrerequisitesTable is the table that holds the prerequisites relation/edge.
	PrerequisitesTable = "prerequisites"
	// PrerequisitesInverseTable is the table name for the Prerequisite entity.
	// It exists in this package in order to avoid circular dependency with the "prerequisite" package.
	PrerequisitesInverseTable = "prerequisites"
	// PrerequisitesColumn is the table column denoting the prerequisites relation/edge.
	PrerequisitesColumn = "feature_prerequisites"
)

// Columns holds all SQL columns for feature fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldDesc,
	FieldLevel,
}

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
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int) error
)

// OrderOption defines the ordering options for the Feature queries.
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

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByPrerequisitesCount orders the results by prerequisites count.
func ByPrerequisitesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPrerequisitesStep(), opts...)
	}
}

// ByPrerequisites orders the results by prerequisites terms.
func ByPrerequisites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrerequisitesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPrerequisitesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrerequisitesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PrerequisitesTable, PrerequisitesColumn),
	)
}
