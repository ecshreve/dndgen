// Code generated by ent, DO NOT EDIT.

package equipment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the equipment type in the database.
	Label = "equipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldEquipmentCategoryID holds the string denoting the equipment_category_id field in the database.
	FieldEquipmentCategoryID = "equipment_category_id"
	// EdgeEquipmentCategory holds the string denoting the equipment_category edge name in mutations.
	EdgeEquipmentCategory = "equipment_category"
	// EdgeCost holds the string denoting the cost edge name in mutations.
	EdgeCost = "cost"
	// EdgeWeapon holds the string denoting the weapon edge name in mutations.
	EdgeWeapon = "weapon"
	// EdgeArmor holds the string denoting the armor edge name in mutations.
	EdgeArmor = "armor"
	// EdgeGear holds the string denoting the gear edge name in mutations.
	EdgeGear = "gear"
	// EdgeTool holds the string denoting the tool edge name in mutations.
	EdgeTool = "tool"
	// EdgeVehicle holds the string denoting the vehicle edge name in mutations.
	EdgeVehicle = "vehicle"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeChoice holds the string denoting the choice edge name in mutations.
	EdgeChoice = "choice"
	// EdgeClassEquipment holds the string denoting the class_equipment edge name in mutations.
	EdgeClassEquipment = "class_equipment"
	// Table holds the table name of the equipment in the database.
	Table = "equipment"
	// EquipmentCategoryTable is the table that holds the equipment_category relation/edge.
	EquipmentCategoryTable = "equipment"
	// EquipmentCategoryInverseTable is the table name for the EquipmentCategory entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentcategory" package.
	EquipmentCategoryInverseTable = "equipment_categories"
	// EquipmentCategoryColumn is the table column denoting the equipment_category relation/edge.
	EquipmentCategoryColumn = "equipment_category_id"
	// CostTable is the table that holds the cost relation/edge.
	CostTable = "equipment_costs"
	// CostInverseTable is the table name for the EquipmentCost entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentcost" package.
	CostInverseTable = "equipment_costs"
	// CostColumn is the table column denoting the cost relation/edge.
	CostColumn = "equipment_id"
	// WeaponTable is the table that holds the weapon relation/edge.
	WeaponTable = "weapons"
	// WeaponInverseTable is the table name for the Weapon entity.
	// It exists in this package in order to avoid circular dependency with the "weapon" package.
	WeaponInverseTable = "weapons"
	// WeaponColumn is the table column denoting the weapon relation/edge.
	WeaponColumn = "equipment_id"
	// ArmorTable is the table that holds the armor relation/edge.
	ArmorTable = "armors"
	// ArmorInverseTable is the table name for the Armor entity.
	// It exists in this package in order to avoid circular dependency with the "armor" package.
	ArmorInverseTable = "armors"
	// ArmorColumn is the table column denoting the armor relation/edge.
	ArmorColumn = "equipment_id"
	// GearTable is the table that holds the gear relation/edge.
	GearTable = "gears"
	// GearInverseTable is the table name for the Gear entity.
	// It exists in this package in order to avoid circular dependency with the "gear" package.
	GearInverseTable = "gears"
	// GearColumn is the table column denoting the gear relation/edge.
	GearColumn = "equipment_id"
	// ToolTable is the table that holds the tool relation/edge.
	ToolTable = "tools"
	// ToolInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolInverseTable = "tools"
	// ToolColumn is the table column denoting the tool relation/edge.
	ToolColumn = "equipment_id"
	// VehicleTable is the table that holds the vehicle relation/edge.
	VehicleTable = "vehicles"
	// VehicleInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	VehicleInverseTable = "vehicles"
	// VehicleColumn is the table column denoting the vehicle relation/edge.
	VehicleColumn = "equipment_id"
	// ClassTable is the table that holds the class relation/edge. The primary key declared below.
	ClassTable = "class_equipments"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ChoiceTable is the table that holds the choice relation/edge. The primary key declared below.
	ChoiceTable = "equipment_choice_equipment"
	// ChoiceInverseTable is the table name for the EquipmentChoice entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentchoice" package.
	ChoiceInverseTable = "equipment_choices"
	// ClassEquipmentTable is the table that holds the class_equipment relation/edge.
	ClassEquipmentTable = "class_equipments"
	// ClassEquipmentInverseTable is the table name for the ClassEquipment entity.
	// It exists in this package in order to avoid circular dependency with the "classequipment" package.
	ClassEquipmentInverseTable = "class_equipments"
	// ClassEquipmentColumn is the table column denoting the class_equipment relation/edge.
	ClassEquipmentColumn = "equipment_id"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldWeight,
	FieldEquipmentCategoryID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "equipment"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"proficiency_equipment",
}

var (
	// ClassPrimaryKey and ClassColumn2 are the table columns denoting the
	// primary key for the class relation (M2M).
	ClassPrimaryKey = []string{"class_id", "equipment_id"}
	// ChoicePrimaryKey and ChoiceColumn2 are the table columns denoting the
	// primary key for the choice relation (M2M).
	ChoicePrimaryKey = []string{"equipment_choice_id", "equipment_id"}
)

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

var (
	// IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	IndxValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Equipment queries.
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

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByEquipmentCategoryID orders the results by the equipment_category_id field.
func ByEquipmentCategoryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEquipmentCategoryID, opts...).ToFunc()
}

// ByEquipmentCategoryField orders the results by equipment_category field.
func ByEquipmentCategoryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEquipmentCategoryStep(), sql.OrderByField(field, opts...))
	}
}

// ByCostField orders the results by cost field.
func ByCostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCostStep(), sql.OrderByField(field, opts...))
	}
}

// ByWeaponField orders the results by weapon field.
func ByWeaponField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeaponStep(), sql.OrderByField(field, opts...))
	}
}

// ByArmorField orders the results by armor field.
func ByArmorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArmorStep(), sql.OrderByField(field, opts...))
	}
}

// ByGearField orders the results by gear field.
func ByGearField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGearStep(), sql.OrderByField(field, opts...))
	}
}

// ByToolField orders the results by tool field.
func ByToolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToolStep(), sql.OrderByField(field, opts...))
	}
}

// ByVehicleField orders the results by vehicle field.
func ByVehicleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVehicleStep(), sql.OrderByField(field, opts...))
	}
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

// ByChoiceCount orders the results by choice count.
func ByChoiceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChoiceStep(), opts...)
	}
}

// ByChoice orders the results by choice terms.
func ByChoice(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChoiceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByClassEquipmentCount orders the results by class_equipment count.
func ByClassEquipmentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClassEquipmentStep(), opts...)
	}
}

// ByClassEquipment orders the results by class_equipment terms.
func ByClassEquipment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassEquipmentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEquipmentCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, EquipmentCategoryTable, EquipmentCategoryColumn),
	)
}
func newCostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CostTable, CostColumn),
	)
}
func newWeaponStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeaponInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, WeaponTable, WeaponColumn),
	)
}
func newArmorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArmorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ArmorTable, ArmorColumn),
	)
}
func newGearStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GearInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, GearTable, GearColumn),
	)
}
func newToolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ToolTable, ToolColumn),
	)
}
func newVehicleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VehicleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, VehicleTable, VehicleColumn),
	)
}
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ClassTable, ClassPrimaryKey...),
	)
}
func newChoiceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChoiceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ChoiceTable, ChoicePrimaryKey...),
	)
}
func newClassEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassEquipmentInverseTable, ClassEquipmentColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, ClassEquipmentTable, ClassEquipmentColumn),
	)
}
