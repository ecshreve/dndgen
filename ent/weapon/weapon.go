// Code generated by ent, DO NOT EDIT.

package weapon

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the weapon type in the database.
	Label = "weapon"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWeaponCategory holds the string denoting the weapon_category field in the database.
	FieldWeaponCategory = "weapon_category"
	// FieldWeaponSubcategory holds the string denoting the weapon_subcategory field in the database.
	FieldWeaponSubcategory = "weapon_subcategory"
	// EdgeDamage holds the string denoting the damage edge name in mutations.
	EdgeDamage = "damage"
	// EdgeProperties holds the string denoting the properties edge name in mutations.
	EdgeProperties = "properties"
	// EdgeEquipment holds the string denoting the equipment edge name in mutations.
	EdgeEquipment = "equipment"
	// EdgeWeaponRange holds the string denoting the weapon_range edge name in mutations.
	EdgeWeaponRange = "weapon_range"
	// Table holds the table name of the weapon in the database.
	Table = "weapons"
	// DamageTable is the table that holds the damage relation/edge.
	DamageTable = "weapons"
	// DamageInverseTable is the table name for the Damage entity.
	// It exists in this package in order to avoid circular dependency with the "damage" package.
	DamageInverseTable = "damages"
	// DamageColumn is the table column denoting the damage relation/edge.
	DamageColumn = "weapon_damage"
	// PropertiesTable is the table that holds the properties relation/edge. The primary key declared below.
	PropertiesTable = "weapon_properties"
	// PropertiesInverseTable is the table name for the Property entity.
	// It exists in this package in order to avoid circular dependency with the "property" package.
	PropertiesInverseTable = "properties"
	// EquipmentTable is the table that holds the equipment relation/edge.
	EquipmentTable = "weapons"
	// EquipmentInverseTable is the table name for the Equipment entity.
	// It exists in this package in order to avoid circular dependency with the "equipment" package.
	EquipmentInverseTable = "equipment"
	// EquipmentColumn is the table column denoting the equipment relation/edge.
	EquipmentColumn = "weapon_equipment"
	// WeaponRangeTable is the table that holds the weapon_range relation/edge.
	WeaponRangeTable = "weapons"
	// WeaponRangeInverseTable is the table name for the WeaponRange entity.
	// It exists in this package in order to avoid circular dependency with the "weaponrange" package.
	WeaponRangeInverseTable = "weapon_ranges"
	// WeaponRangeColumn is the table column denoting the weapon_range relation/edge.
	WeaponRangeColumn = "weapon_weapon_range"
)

// Columns holds all SQL columns for weapon fields.
var Columns = []string{
	FieldID,
	FieldWeaponCategory,
	FieldWeaponSubcategory,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "weapons"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"weapon_damage",
	"weapon_equipment",
	"weapon_weapon_range",
}

var (
	// PropertiesPrimaryKey and PropertiesColumn2 are the table columns denoting the
	// primary key for the properties relation (M2M).
	PropertiesPrimaryKey = []string{"weapon_id", "property_id"}
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

// WeaponCategory defines the type for the "weapon_category" enum field.
type WeaponCategory string

// WeaponCategory values.
const (
	WeaponCategorySimple  WeaponCategory = "simple"
	WeaponCategoryMartial WeaponCategory = "martial"
	WeaponCategoryExotic  WeaponCategory = "exotic"
	WeaponCategoryOther   WeaponCategory = "other"
)

func (wc WeaponCategory) String() string {
	return string(wc)
}

// WeaponCategoryValidator is a validator for the "weapon_category" field enum values. It is called by the builders before save.
func WeaponCategoryValidator(wc WeaponCategory) error {
	switch wc {
	case WeaponCategorySimple, WeaponCategoryMartial, WeaponCategoryExotic, WeaponCategoryOther:
		return nil
	default:
		return fmt.Errorf("weapon: invalid enum value for weapon_category field: %q", wc)
	}
}

// WeaponSubcategory defines the type for the "weapon_subcategory" enum field.
type WeaponSubcategory string

// WeaponSubcategory values.
const (
	WeaponSubcategoryMelee  WeaponSubcategory = "melee"
	WeaponSubcategoryRanged WeaponSubcategory = "ranged"
	WeaponSubcategoryOther  WeaponSubcategory = "other"
)

func (ws WeaponSubcategory) String() string {
	return string(ws)
}

// WeaponSubcategoryValidator is a validator for the "weapon_subcategory" field enum values. It is called by the builders before save.
func WeaponSubcategoryValidator(ws WeaponSubcategory) error {
	switch ws {
	case WeaponSubcategoryMelee, WeaponSubcategoryRanged, WeaponSubcategoryOther:
		return nil
	default:
		return fmt.Errorf("weapon: invalid enum value for weapon_subcategory field: %q", ws)
	}
}

// OrderOption defines the ordering options for the Weapon queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByWeaponCategory orders the results by the weapon_category field.
func ByWeaponCategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeaponCategory, opts...).ToFunc()
}

// ByWeaponSubcategory orders the results by the weapon_subcategory field.
func ByWeaponSubcategory(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeaponSubcategory, opts...).ToFunc()
}

// ByDamageField orders the results by damage field.
func ByDamageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDamageStep(), sql.OrderByField(field, opts...))
	}
}

// ByPropertiesCount orders the results by properties count.
func ByPropertiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPropertiesStep(), opts...)
	}
}

// ByProperties orders the results by properties terms.
func ByProperties(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPropertiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEquipmentField orders the results by equipment field.
func ByEquipmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEquipmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByWeaponRangeField orders the results by weapon_range field.
func ByWeaponRangeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWeaponRangeStep(), sql.OrderByField(field, opts...))
	}
}
func newDamageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DamageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DamageTable, DamageColumn),
	)
}
func newPropertiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PropertiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PropertiesTable, PropertiesPrimaryKey...),
	)
}
func newEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, EquipmentTable, EquipmentColumn),
	)
}
func newWeaponRangeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WeaponRangeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, WeaponRangeTable, WeaponRangeColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e WeaponCategory) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *WeaponCategory) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = WeaponCategory(str)
	if err := WeaponCategoryValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid WeaponCategory", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e WeaponSubcategory) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *WeaponSubcategory) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = WeaponSubcategory(str)
	if err := WeaponSubcategoryValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid WeaponSubcategory", str)
	}
	return nil
}
