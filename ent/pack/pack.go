// Code generated by ent, DO NOT EDIT.

package pack

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the pack type in the database.
	Label = "pack"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContents holds the string denoting the contents field in the database.
	FieldContents = "contents"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the pack in the database.
	Table = "packs"
	// EquipmentTable is the table that holds the equipment relation/edge. The primary key declared below.
	EquipmentTable = "equipment_pack"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
)

// Columns holds all SQL columns for pack fields.
var Columns = []string{
	FieldID,
	FieldContents,
}

var (
	// EquipmentPrimaryKey and EquipmentColumn2 are the table columns denoting the
	// primary key for the equipment relation (M2M).
	EquipmentPrimaryKey = []string{"equipment_id", "pack_id"}
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

// OrderOption defines the ordering options for the Pack queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContents orders the results by the contents field.
func ByContents(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContents, opts...).ToFunc()
}

// ByEquipmentCount orders the results by equipment count.
func ByEquipmentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEquipmentStep(), opts...)
	}
}

// ByEquipment orders the results by equipment terms.
func ByEquipment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEquipmentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, EquipmentTable, EquipmentPrimaryKey...),
	)
}