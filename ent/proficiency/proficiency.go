// Code generated by ent, DO NOT EDIT.

package proficiency

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the proficiency type in the database.
	Label = "proficiency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldReference holds the string denoting the reference field in the database.
	FieldReference = "reference"
	// Table holds the table name of the proficiency in the database.
	Table = "proficiencies"
)

// Columns holds all SQL columns for proficiency fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldReference,
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
	// ReferenceValidator is a validator for the "reference" field. It is called by the builders before save.
	ReferenceValidator func(string) error
)

// OrderOption defines the ordering options for the Proficiency queries.
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

// ByReference orders the results by the reference field.
func ByReference(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReference, opts...).ToFunc()
}
