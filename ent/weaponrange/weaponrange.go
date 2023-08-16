// Code generated by ent, DO NOT EDIT.

package weaponrange

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// EdgeWeapon holds the string denoting the weapon edge name in mutations.
	EdgeWeapon = "weapon"
	// Table holds the table name of the weaponrange in the database.
	Table = "weapon_ranges"
	// WeaponTable is the table that holds the weapon relation/edge. The primary key declared below.
	WeaponTable = "weapon_range"
	// WeaponInverseTable is the table name for the Weapon entity.
	// It exists in this package in order to avoid circular dependency with the "weapon" package.
	WeaponInverseTable = "weapons"
)

// Columns holds all SQL columns for weaponrange fields.
var Columns = []string{
	FieldID,
	FieldDesc,
	FieldNormal,
	FieldLong,
}

var (
	// WeaponPrimaryKey and WeaponColumn2 are the table columns denoting the
	// primary key for the weapon relation (M2M).
	WeaponPrimaryKey = []string{"weapon_id", "weapon_range_id"}
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

// ByWeaponCount orders the results by weapon count.
func ByWeaponCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWeaponStep(), opts...)
	}
}

// ByWeapon orders the results by weapon terms.
func ByWeapon(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeaponStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWeaponStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeaponInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, WeaponTable, WeaponPrimaryKey...),
	)
}
