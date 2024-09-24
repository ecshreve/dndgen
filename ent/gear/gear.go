// Code generated by ent, DO NOT EDIT.

package gear

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the gear type in the database.
	Label = "gear"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGearCategory holds the string denoting the gear_category field in the database.
	FieldGearCategory = "gear_category"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the gear in the database.
	Table = "gears"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "gears"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "gear_equipment"
)

// Columns holds all SQL columns for gear fields.
var Columns = []string{
	FieldID,
	FieldGearCategory,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "gears"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"gear_equipment",
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

// OrderOption defines the ordering options for the Gear queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGearCategory orders the results by the gear_category field.
func ByGearCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGearCategory, opts...).ToFunc()
}

// ByEquipmentField orders the results by equipment field.
func ByEquipmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEquipmentStep(), sql.OrderByField(field, opts...))
	}
}
func newEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, EquipmentTable, EquipmentColumn),
	)
}
