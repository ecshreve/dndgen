// Code generated by ent, DO NOT EDIT.

package weapon

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the weapon type in the database.
	Label = "weapon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWeaponRange holds the string denoting the weapon_range field in the database.
	FieldWeaponRange = "weapon_range"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the weapon in the database.
	Table = "weapons"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "equipment_weapon"
)

// Columns holds all SQL columns for weapon fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldWeaponRange,
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

// OrderOption defines the ordering options for the Weapon queries.
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

// ByWeaponRange orders the results by the weapon_range field.
func ByWeaponRange(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeaponRange, opts...).ToFunc()
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
		sqlgraph.Edge(sqlgraph.O2M, true, EquipmentTable, EquipmentColumn),
	)
}
