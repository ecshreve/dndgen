// Code generated by ent, DO NOT EDIT.

package weapondamage

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the weapondamage type in the database.
	Label = "weapon_damage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWeaponID holds the string denoting the weapon_id field in the database.
	FieldWeaponID = "weapon_id"
	// FieldDamageTypeID holds the string denoting the damage_type_id field in the database.
	FieldDamageTypeID = "damage_type_id"
	// FieldDice holds the string denoting the dice field in the database.
	FieldDice = "dice"
	// EdgeWeapon holds the string denoting the weapon edge name in mutations.
	EdgeWeapon = "weapon"
	// EdgeDamageType holds the string denoting the damage_type edge name in mutations.
	EdgeDamageType = "damage_type"
	// Table holds the table name of the weapondamage in the database.
	Table = "weapon_damages"
	// WeaponTable is the table that holds the weapon relation/edge.
	WeaponTable = "weapon_damages"
	// WeaponInverseTable is the table name for the Weapon entity.
	// It exists in this package in order to avoid circular dependency with the "weapon" package.
	WeaponInverseTable = "weapons"
	// WeaponColumn is the table column denoting the weapon relation/edge.
	WeaponColumn = "weapon_id"
	// DamageTypeTable is the table that holds the damage_type relation/edge.
	DamageTypeTable = "weapon_damages"
	// DamageTypeInverseTable is the table name for the DamageType entity.
	// It exists in this package in order to avoid circular dependency with the "damagetype" package.
	DamageTypeInverseTable = "damage_types"
	// DamageTypeColumn is the table column denoting the damage_type relation/edge.
	DamageTypeColumn = "damage_type_id"
)

// Columns holds all SQL columns for weapondamage fields.
var Columns = []string{
	FieldID,
	FieldWeaponID,
	FieldDamageTypeID,
	FieldDice,
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

// OrderOption defines the ordering options for the WeaponDamage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByWeaponID orders the results by the weapon_id field.
func ByWeaponID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeaponID, opts...).ToFunc()
}

// ByDamageTypeID orders the results by the damage_type_id field.
func ByDamageTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDamageTypeID, opts...).ToFunc()
}

// ByDice orders the results by the dice field.
func ByDice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDice, opts...).ToFunc()
}

// ByWeaponField orders the results by weapon field.
func ByWeaponField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeaponStep(), sql.OrderByField(field, opts...))
	}
}

// ByDamageTypeField orders the results by damage_type field.
func ByDamageTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDamageTypeStep(), sql.OrderByField(field, opts...))
	}
}
func newWeaponStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeaponInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WeaponTable, WeaponColumn),
	)
}
func newDamageTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DamageTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DamageTypeTable, DamageTypeColumn),
	)
}
