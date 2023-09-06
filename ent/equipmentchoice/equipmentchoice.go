// Code generated by ent, DO NOT EDIT.

package equipmentchoice

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the equipmentchoice type in the database.
	Label = "equipment_choice"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChoose holds the string denoting the choose field in the database.
	FieldChoose = "choose"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// Table holds the table name of the equipmentchoice in the database.
	Table = "equipment_choices"
	// ClassTable is the table that holds the class relation/edge. The primary key declared below.
	ClassTable = "class_equipment_choices"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// EquipmentTable is the table that holds the equipment relation/edge. The primary key declared below.
	EquipmentTable = "equipment_choice_equipment"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
)

// Columns holds all SQL columns for equipmentchoice fields.
var Columns = []string{
	FieldID,
	FieldChoose,
	FieldDesc,
}

var (
	// ClassPrimaryKey and ClassColumn2 are the table columns denoting the
	// primary key for the class relation (M2M).
	ClassPrimaryKey = []string{"class_id", "equipment_choice_id"}
	// EquipmentPrimaryKey and EquipmentColumn2 are the table columns denoting the
	// primary key for the equipment relation (M2M).
	EquipmentPrimaryKey = []string{"equipment_choice_id", "equipment_id"}
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

// OrderOption defines the ordering options for the EquipmentChoice queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChoose orders the results by the choose field.
func ByChoose(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChoose, opts...).ToFunc()
}

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByClassCount orders the results by class count.
func ByClassCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClassStep(), opts...)
	}
}

// ByClass orders the results by class terms.
func ByClass(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ClassTable, ClassPrimaryKey...),
	)
}
func newEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, EquipmentTable, EquipmentPrimaryKey...),
	)
}
