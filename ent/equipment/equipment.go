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
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// EdgeWeapon holds the string denoting the weapon edge name in mutations.
	EdgeWeapon = "weapon"
	// EdgeArmor holds the string denoting the armor edge name in mutations.
	EdgeArmor = "armor"
	// EdgeGear holds the string denoting the gear edge name in mutations.
	EdgeGear = "gear"
	// EdgePack holds the string denoting the pack edge name in mutations.
	EdgePack = "pack"
	// EdgeAmmunition holds the string denoting the ammunition edge name in mutations.
	EdgeAmmunition = "ammunition"
	// EdgeVehicle holds the string denoting the vehicle edge name in mutations.
	EdgeVehicle = "vehicle"
	// EdgeMagicItem holds the string denoting the magic_item edge name in mutations.
	EdgeMagicItem = "magic_item"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeSubcategory holds the string denoting the subcategory edge name in mutations.
	EdgeSubcategory = "subcategory"
	// Table holds the table name of the equipment in the database.
	Table = "equipment"
	// WeaponTable is the table that holds the weapon relation/edge. The primary key declared below.
	WeaponTable = "equipment_weapon"
	// WeaponInverseTable is the table name for the Weapon entity.
	// It exists in this package in order to avoid circular dependency with the "weapon" package.
	WeaponInverseTable = "weapons"
	// ArmorTable is the table that holds the armor relation/edge. The primary key declared below.
	ArmorTable = "equipment_armor"
	// ArmorInverseTable is the table name for the Armor entity.
	// It exists in this package in order to avoid circular dependency with the "armor" package.
	ArmorInverseTable = "armors"
	// GearTable is the table that holds the gear relation/edge. The primary key declared below.
	GearTable = "equipment_gear"
	// GearInverseTable is the table name for the Gear entity.
	// It exists in this package in order to avoid circular dependency with the "gear" package.
	GearInverseTable = "gears"
	// PackTable is the table that holds the pack relation/edge. The primary key declared below.
	PackTable = "equipment_pack"
	// PackInverseTable is the table name for the Pack entity.
	// It exists in this package in order to avoid circular dependency with the "pack" package.
	PackInverseTable = "packs"
	// AmmunitionTable is the table that holds the ammunition relation/edge. The primary key declared below.
	AmmunitionTable = "equipment_ammunition"
	// AmmunitionInverseTable is the table name for the Ammunition entity.
	// It exists in this package in order to avoid circular dependency with the "ammunition" package.
	AmmunitionInverseTable = "ammunitions"
	// VehicleTable is the table that holds the vehicle relation/edge. The primary key declared below.
	VehicleTable = "equipment_vehicle"
	// VehicleInverseTable is the table name for the Vehicle entity.
	// It exists in this package in order to avoid circular dependency with the "vehicle" package.
	VehicleInverseTable = "vehicles"
	// MagicItemTable is the table that holds the magic_item relation/edge. The primary key declared below.
	MagicItemTable = "equipment_magic_item"
	// MagicItemInverseTable is the table name for the MagicItem entity.
	// It exists in this package in order to avoid circular dependency with the "magicitem" package.
	MagicItemInverseTable = "magic_items"
	// CategoryTable is the table that holds the category relation/edge. The primary key declared below.
	CategoryTable = "equipment_category"
	// CategoryInverseTable is the table name for the EquipmentCategory entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentcategory" package.
	CategoryInverseTable = "equipment_categories"
	// SubcategoryTable is the table that holds the subcategory relation/edge.
	SubcategoryTable = "equipment_categories"
	// SubcategoryInverseTable is the table name for the EquipmentCategory entity.
	// It exists in this package in order to avoid circular dependency with the "equipmentcategory" package.
	SubcategoryInverseTable = "equipment_categories"
	// SubcategoryColumn is the table column denoting the subcategory relation/edge.
	SubcategoryColumn = "equipment_subcategory"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldDesc,
	FieldCost,
	FieldWeight,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "equipment"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"class_starting_equipment",
}

var (
	// WeaponPrimaryKey and WeaponColumn2 are the table columns denoting the
	// primary key for the weapon relation (M2M).
	WeaponPrimaryKey = []string{"equipment_id", "weapon_id"}
	// ArmorPrimaryKey and ArmorColumn2 are the table columns denoting the
	// primary key for the armor relation (M2M).
	ArmorPrimaryKey = []string{"equipment_id", "armor_id"}
	// GearPrimaryKey and GearColumn2 are the table columns denoting the
	// primary key for the gear relation (M2M).
	GearPrimaryKey = []string{"equipment_id", "gear_id"}
	// PackPrimaryKey and PackColumn2 are the table columns denoting the
	// primary key for the pack relation (M2M).
	PackPrimaryKey = []string{"equipment_id", "pack_id"}
	// AmmunitionPrimaryKey and AmmunitionColumn2 are the table columns denoting the
	// primary key for the ammunition relation (M2M).
	AmmunitionPrimaryKey = []string{"equipment_id", "ammunition_id"}
	// VehiclePrimaryKey and VehicleColumn2 are the table columns denoting the
	// primary key for the vehicle relation (M2M).
	VehiclePrimaryKey = []string{"equipment_id", "vehicle_id"}
	// MagicItemPrimaryKey and MagicItemColumn2 are the table columns denoting the
	// primary key for the magic_item relation (M2M).
	MagicItemPrimaryKey = []string{"equipment_id", "magic_item_id"}
	// CategoryPrimaryKey and CategoryColumn2 are the table columns denoting the
	// primary key for the category relation (M2M).
	CategoryPrimaryKey = []string{"equipment_id", "equipment_category_id"}
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

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByCost orders the results by the cost field.
func ByCost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCost, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
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

// ByArmorCount orders the results by armor count.
func ByArmorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newArmorStep(), opts...)
	}
}

// ByArmor orders the results by armor terms.
func ByArmor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArmorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGearCount orders the results by gear count.
func ByGearCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGearStep(), opts...)
	}
}

// ByGear orders the results by gear terms.
func ByGear(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGearStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPackCount orders the results by pack count.
func ByPackCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPackStep(), opts...)
	}
}

// ByPack orders the results by pack terms.
func ByPack(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPackStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAmmunitionCount orders the results by ammunition count.
func ByAmmunitionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAmmunitionStep(), opts...)
	}
}

// ByAmmunition orders the results by ammunition terms.
func ByAmmunition(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAmmunitionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVehicleCount orders the results by vehicle count.
func ByVehicleCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVehicleStep(), opts...)
	}
}

// ByVehicle orders the results by vehicle terms.
func ByVehicle(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVehicleStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMagicItemCount orders the results by magic_item count.
func ByMagicItemCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMagicItemStep(), opts...)
	}
}

// ByMagicItem orders the results by magic_item terms.
func ByMagicItem(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMagicItemStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoryCount orders the results by category count.
func ByCategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoryStep(), opts...)
	}
}

// ByCategory orders the results by category terms.
func ByCategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySubcategoryCount orders the results by subcategory count.
func BySubcategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubcategoryStep(), opts...)
	}
}

// BySubcategory orders the results by subcategory terms.
func BySubcategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubcategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newWeaponStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeaponInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, WeaponTable, WeaponPrimaryKey...),
	)
}
func newArmorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArmorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ArmorTable, ArmorPrimaryKey...),
	)
}
func newGearStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GearInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, GearTable, GearPrimaryKey...),
	)
}
func newPackStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PackInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PackTable, PackPrimaryKey...),
	)
}
func newAmmunitionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AmmunitionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AmmunitionTable, AmmunitionPrimaryKey...),
	)
}
func newVehicleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VehicleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, VehicleTable, VehiclePrimaryKey...),
	)
}
func newMagicItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MagicItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, MagicItemTable, MagicItemPrimaryKey...),
	)
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CategoryTable, CategoryPrimaryKey...),
	)
}
func newSubcategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubcategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubcategoryTable, SubcategoryColumn),
	)
}
