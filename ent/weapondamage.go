// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// WeaponDamage is the model entity for the WeaponDamage schema.
type WeaponDamage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WeaponID holds the value of the "weapon_id" field.
	WeaponID int `json:"weapon_id,omitempty"`
	// DamageTypeID holds the value of the "damage_type_id" field.
	DamageTypeID int `json:"damage_type_id,omitempty"`
	// Dice holds the value of the "dice" field.
	Dice string `json:"dice,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WeaponDamageQuery when eager-loading is set.
	Edges        WeaponDamageEdges `json:"-"`
	selectValues sql.SelectValues
}

// WeaponDamageEdges holds the relations/edges for other nodes in the graph.
type WeaponDamageEdges struct {
	// Weapon holds the value of the weapon edge.
	Weapon *Weapon `json:"weapon,omitempty"`
	// DamageType holds the value of the damage_type edge.
	DamageType *DamageType `json:"damage_type,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// WeaponOrErr returns the Weapon value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WeaponDamageEdges) WeaponOrErr() (*Weapon, error) {
	if e.loadedTypes[0] {
		if e.Weapon == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: weapon.Label}
		}
		return e.Weapon, nil
	}
	return nil, &NotLoadedError{edge: "weapon"}
}

// DamageTypeOrErr returns the DamageType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WeaponDamageEdges) DamageTypeOrErr() (*DamageType, error) {
	if e.loadedTypes[1] {
		if e.DamageType == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: damagetype.Label}
		}
		return e.DamageType, nil
	}
	return nil, &NotLoadedError{edge: "damage_type"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WeaponDamage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case weapondamage.FieldID, weapondamage.FieldWeaponID, weapondamage.FieldDamageTypeID:
			values[i] = new(sql.NullInt64)
		case weapondamage.FieldDice:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WeaponDamage fields.
func (wd *WeaponDamage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case weapondamage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wd.ID = int(value.Int64)
		case weapondamage.FieldWeaponID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weapon_id", values[i])
			} else if value.Valid {
				wd.WeaponID = int(value.Int64)
			}
		case weapondamage.FieldDamageTypeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field damage_type_id", values[i])
			} else if value.Valid {
				wd.DamageTypeID = int(value.Int64)
			}
		case weapondamage.FieldDice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dice", values[i])
			} else if value.Valid {
				wd.Dice = value.String
			}
		default:
			wd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WeaponDamage.
// This includes values selected through modifiers, order, etc.
func (wd *WeaponDamage) Value(name string) (ent.Value, error) {
	return wd.selectValues.Get(name)
}

// QueryWeapon queries the "weapon" edge of the WeaponDamage entity.
func (wd *WeaponDamage) QueryWeapon() *WeaponQuery {
	return NewWeaponDamageClient(wd.config).QueryWeapon(wd)
}

// QueryDamageType queries the "damage_type" edge of the WeaponDamage entity.
func (wd *WeaponDamage) QueryDamageType() *DamageTypeQuery {
	return NewWeaponDamageClient(wd.config).QueryDamageType(wd)
}

// Update returns a builder for updating this WeaponDamage.
// Note that you need to call WeaponDamage.Unwrap() before calling this method if this WeaponDamage
// was returned from a transaction, and the transaction was committed or rolled back.
func (wd *WeaponDamage) Update() *WeaponDamageUpdateOne {
	return NewWeaponDamageClient(wd.config).UpdateOne(wd)
}

// Unwrap unwraps the WeaponDamage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wd *WeaponDamage) Unwrap() *WeaponDamage {
	_tx, ok := wd.config.driver.(*txDriver)
	if !ok {
		panic("ent: WeaponDamage is not a transactional entity")
	}
	wd.config.driver = _tx.drv
	return wd
}

// String implements the fmt.Stringer.
func (wd *WeaponDamage) String() string {
	var builder strings.Builder
	builder.WriteString("WeaponDamage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wd.ID))
	builder.WriteString("weapon_id=")
	builder.WriteString(fmt.Sprintf("%v", wd.WeaponID))
	builder.WriteString(", ")
	builder.WriteString("damage_type_id=")
	builder.WriteString(fmt.Sprintf("%v", wd.DamageTypeID))
	builder.WriteString(", ")
	builder.WriteString("dice=")
	builder.WriteString(wd.Dice)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (wd *WeaponDamage) MarshalJSON() ([]byte, error) {
	type Alias WeaponDamage
	return json.Marshal(&struct {
		*Alias
		WeaponDamageEdges
	}{
		Alias:             (*Alias)(wd),
		WeaponDamageEdges: wd.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (wd *WeaponDamage) UnmarshalJSON(data []byte) error {
	type Alias WeaponDamage
	aux := &struct {
		*Alias
		WeaponDamageEdges
	}{
		Alias: (*Alias)(wd),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	wd.Edges = aux.WeaponDamageEdges
	return nil
}

func (wdc *WeaponDamageCreate) SetWeaponDamage(input *WeaponDamage) *WeaponDamageCreate {
	wdc.SetWeaponID(input.WeaponID)
	wdc.SetDamageTypeID(input.DamageTypeID)
	wdc.SetDice(input.Dice)
	return wdc
}

// WeaponDamages is a parsable slice of WeaponDamage.
type WeaponDamages []*WeaponDamage
