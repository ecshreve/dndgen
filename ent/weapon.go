// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// Weapon is the model entity for the Weapon schema.
type Weapon struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// WeaponRange holds the value of the "weapon_range" field.
	WeaponRange string `json:"weapon_range,omitempty"`
	// EquipmentID holds the value of the "equipment_id" field.
	EquipmentID int `json:"equipment_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WeaponQuery when eager-loading is set.
	Edges        WeaponEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WeaponEdges holds the relations/edges for other nodes in the graph.
type WeaponEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// Damage holds the value of the damage edge.
	Damage *WeaponDamage `json:"damage,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WeaponEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// DamageOrErr returns the Damage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WeaponEdges) DamageOrErr() (*WeaponDamage, error) {
	if e.loadedTypes[1] {
		if e.Damage == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: weapondamage.Label}
		}
		return e.Damage, nil
	}
	return nil, &NotLoadedError{edge: "damage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Weapon) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case weapon.FieldID, weapon.FieldEquipmentID:
			values[i] = new(sql.NullInt64)
		case weapon.FieldIndx, weapon.FieldName, weapon.FieldWeaponRange:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Weapon fields.
func (w *Weapon) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case weapon.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case weapon.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				w.Indx = value.String
			}
		case weapon.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case weapon.FieldWeaponRange:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field weapon_range", values[i])
			} else if value.Valid {
				w.WeaponRange = value.String
			}
		case weapon.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				w.EquipmentID = int(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Weapon.
// This includes values selected through modifiers, order, etc.
func (w *Weapon) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryEquipment queries the "equipment" edge of the Weapon entity.
func (w *Weapon) QueryEquipment() *EquipmentQuery {
	return NewWeaponClient(w.config).QueryEquipment(w)
}

// QueryDamage queries the "damage" edge of the Weapon entity.
func (w *Weapon) QueryDamage() *WeaponDamageQuery {
	return NewWeaponClient(w.config).QueryDamage(w)
}

// Update returns a builder for updating this Weapon.
// Note that you need to call Weapon.Unwrap() before calling this method if this Weapon
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Weapon) Update() *WeaponUpdateOne {
	return NewWeaponClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Weapon entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Weapon) Unwrap() *Weapon {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Weapon is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Weapon) String() string {
	var builder strings.Builder
	builder.WriteString("Weapon(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("indx=")
	builder.WriteString(w.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(w.Name)
	builder.WriteString(", ")
	builder.WriteString("weapon_range=")
	builder.WriteString(w.WeaponRange)
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", w.EquipmentID))
	builder.WriteByte(')')
	return builder.String()
}

func (wc *WeaponCreate) SetWeapon(input *Weapon) *WeaponCreate {
	wc.SetIndx(input.Indx)
	wc.SetName(input.Name)
	wc.SetWeaponRange(input.WeaponRange)
	wc.SetEquipmentID(input.EquipmentID)
	return wc
}

// Weapons is a parsable slice of Weapon.
type Weapons []*Weapon
