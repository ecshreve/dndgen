// Code generated by ent, DO NOT EDIT.

package damage

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the damage type in the database.
	Label = "damage"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDamageDice holds the string denoting the damage_dice field in the database.
	FieldDamageDice = "damage_dice"
	// EdgeDamageType holds the string denoting the damage_type edge name in mutations.
	EdgeDamageType = "damage_type"
	// Table holds the table name of the damage in the database.
	Table = "damages"
	// DamageTypeTable is the table that holds the damage_type relation/edge.
	DamageTypeTable = "damages"
	// DamageTypeInverseTable is the table name for the DamageType entity.
	// It exists in this package in order to avoid circular dependency with the "damagetype" package.
	DamageTypeInverseTable = "damage_types"
	// DamageTypeColumn is the table column denoting the damage_type relation/edge.
	DamageTypeColumn = "damage_damage_type"
)

// Columns holds all SQL columns for damage fields.
var Columns = []string{
	FieldID,
	FieldDamageDice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "damages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"damage_damage_type",
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

// OrderOption defines the ordering options for the Damage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDamageDice orders the results by the damage_dice field.
func ByDamageDice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDamageDice, opts...).ToFunc()
}

// ByDamageTypeField orders the results by damage_type field.
func ByDamageTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDamageTypeStep(), sql.OrderByField(field, opts...))
	}
}
func newDamageTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DamageTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DamageTypeTable, DamageTypeColumn),
	)
}