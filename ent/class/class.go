// Code generated by ent, DO NOT EDIT.

package class

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the class type in the database.
	Label = "class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHitDie holds the string denoting the hit_die field in the database.
	FieldHitDie = "hit_die"
	// EdgeSavingThrows holds the string denoting the saving_throws edge name in mutations.
	EdgeSavingThrows = "saving_throws"
	// Table holds the table name of the class in the database.
	Table = "classes"
	// SavingThrowsTable is the table that holds the saving_throws relation/edge. The primary key declared below.
	SavingThrowsTable = "class_saving_throws"
	// SavingThrowsInverseTable is the table name for the AbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "abilityscore" package.
	SavingThrowsInverseTable = "ability_scores"
)

// Columns holds all SQL columns for class fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldHitDie,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "classes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"character_class",
}

var (
	// SavingThrowsPrimaryKey and SavingThrowsColumn2 are the table columns denoting the
	// primary key for the saving_throws relation (M2M).
	SavingThrowsPrimaryKey = []string{"class_id", "ability_score_id"}
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
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Class queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByHitDie orders the results by the hit_die field.
func ByHitDie(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHitDie, opts...).ToFunc()
}

// BySavingThrowsCount orders the results by saving_throws count.
func BySavingThrowsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSavingThrowsStep(), opts...)
	}
}

// BySavingThrows orders the results by saving_throws terms.
func BySavingThrows(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSavingThrowsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSavingThrowsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SavingThrowsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SavingThrowsTable, SavingThrowsPrimaryKey...),
	)
}
