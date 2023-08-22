// Code generated by ent, DO NOT EDIT.

package armorclass

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the armorclass type in the database.
	Label = "armor_class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBase holds the string denoting the base field in the database.
	FieldBase = "base"
	// FieldDexBonus holds the string denoting the dex_bonus field in the database.
	FieldDexBonus = "dex_bonus"
	// FieldMaxBonus holds the string denoting the max_bonus field in the database.
	FieldMaxBonus = "max_bonus"
	// Table holds the table name of the armorclass in the database.
	Table = "armor_classes"
)

// Columns holds all SQL columns for armorclass fields.
var Columns = []string{
	FieldID,
	FieldBase,
	FieldDexBonus,
	FieldMaxBonus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "armor_classes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"armor_armor_class",
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

// OrderOption defines the ordering options for the ArmorClass queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBase orders the results by the base field.
func ByBase(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBase, opts...).ToFunc()
}

// ByDexBonus orders the results by the dex_bonus field.
func ByDexBonus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDexBonus, opts...).ToFunc()
}

// ByMaxBonus orders the results by the max_bonus field.
func ByMaxBonus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxBonus, opts...).ToFunc()
}