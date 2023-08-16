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
	// FieldProperties holds the string denoting the properties field in the database.
	FieldProperties = "properties"
	// EdgeRange holds the string denoting the range edge name in mutations.
	EdgeRange = "range"
	// EdgeDamage holds the string denoting the damage edge name in mutations.
	EdgeDamage = "damage"
	// EdgeTwoHandedDamage holds the string denoting the two_handed_damage edge name in mutations.
	EdgeTwoHandedDamage = "two_handed_damage"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the weapon in the database.
	Table = "weapons"
	// RangeTable is the table that holds the range relation/edge.
	RangeTable = "weapon_ranges"
	// RangeInverseTable is the table name for the WeaponRange entity.
	// It exists in this package in order to avoid circular dependency with the "weaponrange" package.
	RangeInverseTable = "weapon_ranges"
	// RangeColumn is the table column denoting the range relation/edge.
	RangeColumn = "weapon_range"
	// DamageTable is the table that holds the damage relation/edge.
	DamageTable = "weapon_damages"
	// DamageInverseTable is the table name for the WeaponDamage entity.
	// It exists in this package in order to avoid circular dependency with the "weapondamage" package.
	DamageInverseTable = "weapon_damages"
	// DamageColumn is the table column denoting the damage relation/edge.
	DamageColumn = "weapon_damage"
	// TwoHandedDamageTable is the table that holds the two_handed_damage relation/edge.
	TwoHandedDamageTable = "weapon_damages"
	// TwoHandedDamageInverseTable is the table name for the WeaponDamage entity.
	// It exists in this package in order to avoid circular dependency with the "weapondamage" package.
	TwoHandedDamageInverseTable = "weapon_damages"
	// TwoHandedDamageColumn is the table column denoting the two_handed_damage relation/edge.
	TwoHandedDamageColumn = "weapon_two_handed_damage"
	// EquipmentTable is the table that holds the equipment relation/edge. The primary key declared below.
	EquipmentTable = "equipment_weapon"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
)

// Columns holds all SQL columns for weapon fields.
var Columns = []string{
	FieldID,
	FieldProperties,
}

var (
	// EquipmentPrimaryKey and EquipmentColumn2 are the table columns denoting the
	// primary key for the equipment relation (M2M).
	EquipmentPrimaryKey = []string{"equipment_id", "weapon_id"}
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

// OrderOption defines the ordering options for the Weapon queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProperties orders the results by the properties field.
func ByProperties(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProperties, opts...).ToFunc()
}

// ByRangeCount orders the results by range count.
func ByRangeCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRangeStep(), opts...)
	}
}

// ByRange orders the results by range terms.
func ByRange(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRangeStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDamageCount orders the results by damage count.
func ByDamageCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDamageStep(), opts...)
	}
}

// ByDamage orders the results by damage terms.
func ByDamage(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDamageStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTwoHandedDamageCount orders the results by two_handed_damage count.
func ByTwoHandedDamageCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTwoHandedDamageStep(), opts...)
	}
}

// ByTwoHandedDamage orders the results by two_handed_damage terms.
func ByTwoHandedDamage(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTwoHandedDamageStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
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
func newRangeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RangeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RangeTable, RangeColumn),
	)
}
func newDamageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DamageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DamageTable, DamageColumn),
	)
}
func newTwoHandedDamageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TwoHandedDamageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TwoHandedDamageTable, TwoHandedDamageColumn),
	)
}
func newEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, EquipmentTable, EquipmentPrimaryKey...),
	)
}