// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipment"
)

// Equipment is the model entity for the Equipment schema.
type Equipment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost string `json:"cost,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight string `json:"weight,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentQuery when eager-loading is set.
	Edges                    EquipmentEdges `json:"edges"`
	class_starting_equipment *int
	selectValues             sql.SelectValues
}

// EquipmentEdges holds the relations/edges for other nodes in the graph.
type EquipmentEdges struct {
	// Weapon holds the value of the weapon edge.
	Weapon []*Weapon `json:"weapon,omitempty"`
	// Armor holds the value of the armor edge.
	Armor []*Armor `json:"armor,omitempty"`
	// Gear holds the value of the gear edge.
	Gear []*Gear `json:"gear,omitempty"`
	// Pack holds the value of the pack edge.
	Pack []*Pack `json:"pack,omitempty"`
	// Ammunition holds the value of the ammunition edge.
	Ammunition []*Ammunition `json:"ammunition,omitempty"`
	// Vehicle holds the value of the vehicle edge.
	Vehicle []*Vehicle `json:"vehicle,omitempty"`
	// MagicItem holds the value of the magic_item edge.
	MagicItem []*MagicItem `json:"magic_item,omitempty"`
	// Category holds the value of the category edge.
	Category []*EquipmentCategory `json:"category,omitempty"`
	// Subcategory holds the value of the subcategory edge.
	Subcategory []*EquipmentCategory `json:"subcategory,omitempty"`
	// Proficiencies holds the value of the proficiencies edge.
	Proficiencies []*Proficiency `json:"proficiencies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [10]bool
}

// WeaponOrErr returns the Weapon value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) WeaponOrErr() ([]*Weapon, error) {
	if e.loadedTypes[0] {
		return e.Weapon, nil
	}
	return nil, &NotLoadedError{edge: "weapon"}
}

// ArmorOrErr returns the Armor value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) ArmorOrErr() ([]*Armor, error) {
	if e.loadedTypes[1] {
		return e.Armor, nil
	}
	return nil, &NotLoadedError{edge: "armor"}
}

// GearOrErr returns the Gear value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) GearOrErr() ([]*Gear, error) {
	if e.loadedTypes[2] {
		return e.Gear, nil
	}
	return nil, &NotLoadedError{edge: "gear"}
}

// PackOrErr returns the Pack value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) PackOrErr() ([]*Pack, error) {
	if e.loadedTypes[3] {
		return e.Pack, nil
	}
	return nil, &NotLoadedError{edge: "pack"}
}

// AmmunitionOrErr returns the Ammunition value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) AmmunitionOrErr() ([]*Ammunition, error) {
	if e.loadedTypes[4] {
		return e.Ammunition, nil
	}
	return nil, &NotLoadedError{edge: "ammunition"}
}

// VehicleOrErr returns the Vehicle value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) VehicleOrErr() ([]*Vehicle, error) {
	if e.loadedTypes[5] {
		return e.Vehicle, nil
	}
	return nil, &NotLoadedError{edge: "vehicle"}
}

// MagicItemOrErr returns the MagicItem value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) MagicItemOrErr() ([]*MagicItem, error) {
	if e.loadedTypes[6] {
		return e.MagicItem, nil
	}
	return nil, &NotLoadedError{edge: "magic_item"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) CategoryOrErr() ([]*EquipmentCategory, error) {
	if e.loadedTypes[7] {
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// SubcategoryOrErr returns the Subcategory value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) SubcategoryOrErr() ([]*EquipmentCategory, error) {
	if e.loadedTypes[8] {
		return e.Subcategory, nil
	}
	return nil, &NotLoadedError{edge: "subcategory"}
}

// ProficienciesOrErr returns the Proficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) ProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[9] {
		return e.Proficiencies, nil
	}
	return nil, &NotLoadedError{edge: "proficiencies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Equipment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipment.FieldID:
			values[i] = new(sql.NullInt64)
		case equipment.FieldIndx, equipment.FieldName, equipment.FieldDesc, equipment.FieldCost, equipment.FieldWeight:
			values[i] = new(sql.NullString)
		case equipment.ForeignKeys[0]: // class_starting_equipment
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Equipment fields.
func (e *Equipment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case equipment.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				e.Indx = value.String
			}
		case equipment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case equipment.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				e.Desc = value.String
			}
		case equipment.FieldCost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cost", values[i])
			} else if value.Valid {
				e.Cost = value.String
			}
		case equipment.FieldWeight:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				e.Weight = value.String
			}
		case equipment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field class_starting_equipment", value)
			} else if value.Valid {
				e.class_starting_equipment = new(int)
				*e.class_starting_equipment = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Equipment.
// This includes values selected through modifiers, order, etc.
func (e *Equipment) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryWeapon queries the "weapon" edge of the Equipment entity.
func (e *Equipment) QueryWeapon() *WeaponQuery {
	return NewEquipmentClient(e.config).QueryWeapon(e)
}

// QueryArmor queries the "armor" edge of the Equipment entity.
func (e *Equipment) QueryArmor() *ArmorQuery {
	return NewEquipmentClient(e.config).QueryArmor(e)
}

// QueryGear queries the "gear" edge of the Equipment entity.
func (e *Equipment) QueryGear() *GearQuery {
	return NewEquipmentClient(e.config).QueryGear(e)
}

// QueryPack queries the "pack" edge of the Equipment entity.
func (e *Equipment) QueryPack() *PackQuery {
	return NewEquipmentClient(e.config).QueryPack(e)
}

// QueryAmmunition queries the "ammunition" edge of the Equipment entity.
func (e *Equipment) QueryAmmunition() *AmmunitionQuery {
	return NewEquipmentClient(e.config).QueryAmmunition(e)
}

// QueryVehicle queries the "vehicle" edge of the Equipment entity.
func (e *Equipment) QueryVehicle() *VehicleQuery {
	return NewEquipmentClient(e.config).QueryVehicle(e)
}

// QueryMagicItem queries the "magic_item" edge of the Equipment entity.
func (e *Equipment) QueryMagicItem() *MagicItemQuery {
	return NewEquipmentClient(e.config).QueryMagicItem(e)
}

// QueryCategory queries the "category" edge of the Equipment entity.
func (e *Equipment) QueryCategory() *EquipmentCategoryQuery {
	return NewEquipmentClient(e.config).QueryCategory(e)
}

// QuerySubcategory queries the "subcategory" edge of the Equipment entity.
func (e *Equipment) QuerySubcategory() *EquipmentCategoryQuery {
	return NewEquipmentClient(e.config).QuerySubcategory(e)
}

// QueryProficiencies queries the "proficiencies" edge of the Equipment entity.
func (e *Equipment) QueryProficiencies() *ProficiencyQuery {
	return NewEquipmentClient(e.config).QueryProficiencies(e)
}

// Update returns a builder for updating this Equipment.
// Note that you need to call Equipment.Unwrap() before calling this method if this Equipment
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Equipment) Update() *EquipmentUpdateOne {
	return NewEquipmentClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Equipment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Equipment) Unwrap() *Equipment {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Equipment is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Equipment) String() string {
	var builder strings.Builder
	builder.WriteString("Equipment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("indx=")
	builder.WriteString(e.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(e.Desc)
	builder.WriteString(", ")
	builder.WriteString("cost=")
	builder.WriteString(e.Cost)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(e.Weight)
	builder.WriteByte(')')
	return builder.String()
}

// EquipmentSlice is a parsable slice of Equipment.
type EquipmentSlice []*Equipment
