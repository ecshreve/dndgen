// Code generated by ent, DO NOT EDIT.

package weaponrange

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the weaponrange type in the database.
	Label = "weapon_range"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldNormal holds the string denoting the normal field in the database.
	FieldNormal = "normal"
	// FieldLong holds the string denoting the long field in the database.
	FieldLong = "long"
	// Table holds the table name of the weaponrange in the database.
	Table = "weapon_ranges"
)

// Columns holds all SQL columns for weaponrange fields.
var Columns = []string{
	FieldID,
	FieldDesc,
	FieldNormal,
	FieldLong,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "weapon_ranges"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"weapon_range",
}

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

// OrderOption defines the ordering options for the WeaponRange queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByNormal orders the results by the normal field.
func ByNormal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNormal, opts...).ToFunc()
}

// ByLong orders the results by the long field.
func ByLong(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLong, opts...).ToFunc()
}
