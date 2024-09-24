// Code generated by ent, DO NOT EDIT.

package tool

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tool type in the database.
	Label = "tool"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToolCategory holds the string denoting the tool_category field in the database.
	FieldToolCategory = "tool_category"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the tool in the database.
	Table = "tools"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "tools"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "tool_equipment"
)

// Columns holds all SQL columns for tool fields.
var Columns = []string{
	FieldID,
	FieldToolCategory,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tools"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"tool_equipment",
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

// OrderOption defines the ordering options for the Tool queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByToolCategory orders the results by the tool_category field.
func ByToolCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToolCategory, opts...).ToFunc()
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
