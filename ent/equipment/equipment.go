// Code generated by ent, DO NOT EDIT.

package equipment

import (
	"fmt"
	"io"
	"strconv"

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
	// FieldEquipmentCategory holds the string denoting the equipment_category field in the database.
	FieldEquipmentCategory = "equipment_category"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// EdgeCost holds the string denoting the cost edge name in mutations.
	EdgeCost = "cost"
	// EdgeGear holds the string denoting the gear edge name in mutations.
	EdgeGear = "gear"
	// EdgeTool holds the string denoting the tool edge name in mutations.
	EdgeTool = "tool"
	// EdgeWeapon holds the string denoting the weapon edge name in mutations.
	EdgeWeapon = "weapon"
	// EdgeVehicle holds the string denoting the vehicle edge name in mutations.
	EdgeVehicle = "vehicle"
	// EdgeArmor holds the string denoting the armor edge name in mutations.
	EdgeArmor = "armor"
	// EdgeEquipmentEntries holds the string denoting the equipment_entries edge name in mutations.
	EdgeEquipmentEntries = "equipment_entries"
	// Table holds the table name of the equipment in the database.
	Table = "equipment"
	// CostTable is the table that holds the cost relation/edge.
	CostTable = "costs"
	// CostInverseTable is the table name for the Cost entity.
	// It exists in this package in order to avoid circular dependency with the "cost" package.
	CostInverseTable = "costs"
	// CostColumn is the table column denoting the cost relation/edge.
	CostColumn = "equipment_cost"
	// GearTable is the table that holds the gear relation/edge.
	GearTable = "gears"
	// GearInverseTable is the table name for the Gear entity.
	// It exists in this package in order to avoid circular dependency with the "gear" package.
	GearInverseTable = "gears"
	// GearColumn is the table column denoting the gear relation/edge.
	GearColumn = "equipment_gear"
	// ToolTable is the table that holds the tool relation/edge.
	ToolTable = "tools"
	// ToolInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolInverseTable = "tools"
	// ToolColumn is the table column denoting the tool relation/edge.
	ToolColumn = "equipment_tool"
	// WeaponTable is the table that holds the weapon relation/edge.
	WeaponTable = "weapons"
	// WeaponInverseTable is the table name for the Weapon entity.
	// It exists in this package in order to avoid circular dependency with the "weapon" package.
	WeaponInverseTable = "weapons"
	// WeaponColumn is the table column denoting the weapon relation/edge.
	WeaponColumn = "equipment_weapon"
	// VehicleTable is the table that holds the vehicle relation/edge.
	VehicleTable = "vehicles"
	// VehicleInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	VehicleInverseTable = "vehicles"
	// VehicleColumn is the table column denoting the vehicle relation/edge.
	VehicleColumn = "equipment_vehicle"
	// ArmorTable is the table that holds the armor relation/edge.
	ArmorTable = "armors"
	// ArmorInverseTable is the table name for the Armor entity.
	// It exists in this package in order to avoid circular dependency with the "armor" package.
	ArmorInverseTable = "armors"
	// ArmorColumn is the table column denoting the armor relation/edge.
	ArmorColumn = "equipment_armor"
	// EquipmentEntriesTable is the table that holds the equipment_entries relation/edge.
	EquipmentEntriesTable = "equipment_entries"
	// EquipmentEntriesInverseTable is the table name for the EquipmentEntry entity.
	// It exists in this package in order to avoid circular dependency with the "equipmententry" package.
	EquipmentEntriesInverseTable = "equipment_entries"
	// EquipmentEntriesColumn is the table column denoting the equipment_entries relation/edge.
	EquipmentEntriesColumn = "equipment_entry_equipment"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldEquipmentCategory,
	FieldWeight,
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

// EquipmentCategory defines the type for the "equipment_category" enum field.
type EquipmentCategory string

// EquipmentCategory values.
const (
	EquipmentCategoryGear    EquipmentCategory = "GEAR"
	EquipmentCategoryTool    EquipmentCategory = "TOOL"
	EquipmentCategoryWeapon  EquipmentCategory = "WEAPON"
	EquipmentCategoryVehicle EquipmentCategory = "VEHICLE"
	EquipmentCategoryArmor   EquipmentCategory = "ARMOR"
)

func (ec EquipmentCategory) String() string {
	return string(ec)
}

// EquipmentCategoryValidator is a validator for the "equipment_category" field enum values. It is called by the builders before save.
func EquipmentCategoryValidator(ec EquipmentCategory) error {
	switch ec {
	case EquipmentCategoryGear, EquipmentCategoryTool, EquipmentCategoryWeapon, EquipmentCategoryVehicle, EquipmentCategoryArmor:
		return nil
	default:
		return fmt.Errorf("equipment: invalid enum value for equipment_category field: %q", ec)
	}
}

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

// ByEquipmentCategory orders the results by the equipment_category field.
func ByEquipmentCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEquipmentCategory, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByCostField orders the results by cost field.
func ByCostField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCostStep(), sql.OrderByField(field, opts...))
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

// ByWeaponField orders the results by weapon field.
func ByWeaponField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeaponStep(), sql.OrderByField(field, opts...))
	}
}

// ByVehicleField orders the results by vehicle field.
func ByVehicleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVehicleStep(), sql.OrderByField(field, opts...))
	}
}

// ByArmorField orders the results by armor field.
func ByArmorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArmorStep(), sql.OrderByField(field, opts...))
	}
}

// ByEquipmentEntriesCount orders the results by equipment_entries count.
func ByEquipmentEntriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEquipmentEntriesStep(), opts...)
	}
}

// ByEquipmentEntries orders the results by equipment_entries terms.
func ByEquipmentEntries(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEquipmentEntriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCostStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CostInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CostTable, CostColumn),
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
func newWeaponStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeaponInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, WeaponTable, WeaponColumn),
	)
}
func newVehicleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VehicleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, VehicleTable, VehicleColumn),
	)
}
func newArmorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArmorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, ArmorTable, ArmorColumn),
	)
}
func newEquipmentEntriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentEntriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, EquipmentEntriesTable, EquipmentEntriesColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e EquipmentCategory) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *EquipmentCategory) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = EquipmentCategory(str)
	if err := EquipmentCategoryValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid EquipmentCategory", str)
	}
	return nil
}
