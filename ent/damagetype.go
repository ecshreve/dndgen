// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// DamageType is the model entity for the DamageType schema.
type DamageType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DamageTypeQuery when eager-loading is set.
	Edges                     DamageTypeEdges `json:"edges"`
	weapon_damage_damage_type *int
	selectValues              sql.SelectValues
}

// DamageTypeEdges holds the relations/edges for other nodes in the graph.
type DamageTypeEdges struct {
	// WeaponDamage holds the value of the weapon_damage edge.
	WeaponDamage *WeaponDamage `json:"weapon_damage,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// WeaponDamageOrErr returns the WeaponDamage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DamageTypeEdges) WeaponDamageOrErr() (*WeaponDamage, error) {
	if e.loadedTypes[0] {
		if e.WeaponDamage == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: weapondamage.Label}
		}
		return e.WeaponDamage, nil
	}
	return nil, &NotLoadedError{edge: "weapon_damage"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DamageType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case damagetype.FieldID:
			values[i] = new(sql.NullInt64)
		case damagetype.FieldIndx, damagetype.FieldName, damagetype.FieldDesc:
			values[i] = new(sql.NullString)
		case damagetype.ForeignKeys[0]: // weapon_damage_damage_type
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DamageType fields.
func (dt *DamageType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case damagetype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dt.ID = int(value.Int64)
		case damagetype.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				dt.Indx = value.String
			}
		case damagetype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dt.Name = value.String
			}
		case damagetype.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				dt.Desc = value.String
			}
		case damagetype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field weapon_damage_damage_type", value)
			} else if value.Valid {
				dt.weapon_damage_damage_type = new(int)
				*dt.weapon_damage_damage_type = int(value.Int64)
			}
		default:
			dt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DamageType.
// This includes values selected through modifiers, order, etc.
func (dt *DamageType) Value(name string) (ent.Value, error) {
	return dt.selectValues.Get(name)
}

// QueryWeaponDamage queries the "weapon_damage" edge of the DamageType entity.
func (dt *DamageType) QueryWeaponDamage() *WeaponDamageQuery {
	return NewDamageTypeClient(dt.config).QueryWeaponDamage(dt)
}

// Update returns a builder for updating this DamageType.
// Note that you need to call DamageType.Unwrap() before calling this method if this DamageType
// was returned from a transaction, and the transaction was committed or rolled back.
func (dt *DamageType) Update() *DamageTypeUpdateOne {
	return NewDamageTypeClient(dt.config).UpdateOne(dt)
}

// Unwrap unwraps the DamageType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dt *DamageType) Unwrap() *DamageType {
	_tx, ok := dt.config.driver.(*txDriver)
	if !ok {
		panic("ent: DamageType is not a transactional entity")
	}
	dt.config.driver = _tx.drv
	return dt
}

// String implements the fmt.Stringer.
func (dt *DamageType) String() string {
	var builder strings.Builder
	builder.WriteString("DamageType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dt.ID))
	builder.WriteString("indx=")
	builder.WriteString(dt.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(dt.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(dt.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (dt *DamageType) MarshalJSON() ([]byte, error) {
	type Alias DamageType
	return json.Marshal(&struct {
		*Alias
		DamageTypeEdges
	}{
		Alias:           (*Alias)(dt),
		DamageTypeEdges: dt.Edges,
	})
}

// DamageTypes is a parsable slice of DamageType.
type DamageTypes []*DamageType
