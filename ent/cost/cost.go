// Code generated by ent, DO NOT EDIT.

package cost

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cost type in the database.
	Label = "cost"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// Table holds the table name of the cost in the database.
	Table = "costs"
)

// Columns holds all SQL columns for cost fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
	FieldUnit,
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

// OrderOption defines the ordering options for the Cost queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByUnit orders the results by the unit field.
func ByUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnit, opts...).ToFunc()
}
