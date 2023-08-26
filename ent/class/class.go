// Code generated by ent, DO NOT EDIT.

package class

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the class type in the database.
	Label = "class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHitDie holds the string denoting the hit_die field in the database.
	FieldHitDie = "hit_die"
	// Table holds the table name of the class in the database.
	Table = "classes"
)

// Columns holds all SQL columns for class fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldHitDie,
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
)

// OrderOption defines the ordering options for the Class queries.
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

// ByHitDie orders the results by the hit_die field.
func ByHitDie(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHitDie, opts...).ToFunc()
}
