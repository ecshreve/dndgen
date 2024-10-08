// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
)

// Gear is the model entity for the Gear schema.
type Gear struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GearCategory holds the value of the "gear_category" field.
	GearCategory string `json:"gear_category,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// EquipmentID holds the value of the "equipment_id" field.
	EquipmentID int `json:"equipment_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GearQuery when eager-loading is set.
	Edges        GearEdges `json:"-"`
	selectValues sql.SelectValues
}

// GearEdges holds the relations/edges for other nodes in the graph.
type GearEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GearEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Gear) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gear.FieldID, gear.FieldQuantity, gear.FieldEquipmentID:
			values[i] = new(sql.NullInt64)
		case gear.FieldIndx, gear.FieldName, gear.FieldGearCategory:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Gear fields.
func (ge *Gear) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gear.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ge.ID = int(value.Int64)
		case gear.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				ge.Indx = value.String
			}
		case gear.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ge.Name = value.String
			}
		case gear.FieldGearCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gear_category", values[i])
			} else if value.Valid {
				ge.GearCategory = value.String
			}
		case gear.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				ge.Quantity = int(value.Int64)
			}
		case gear.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				ge.EquipmentID = int(value.Int64)
			}
		default:
			ge.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Gear.
// This includes values selected through modifiers, order, etc.
func (ge *Gear) Value(name string) (ent.Value, error) {
	return ge.selectValues.Get(name)
}

// QueryEquipment queries the "equipment" edge of the Gear entity.
func (ge *Gear) QueryEquipment() *EquipmentQuery {
	return NewGearClient(ge.config).QueryEquipment(ge)
}

// Update returns a builder for updating this Gear.
// Note that you need to call Gear.Unwrap() before calling this method if this Gear
// was returned from a transaction, and the transaction was committed or rolled back.
func (ge *Gear) Update() *GearUpdateOne {
	return NewGearClient(ge.config).UpdateOne(ge)
}

// Unwrap unwraps the Gear entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ge *Gear) Unwrap() *Gear {
	_tx, ok := ge.config.driver.(*txDriver)
	if !ok {
		panic("ent: Gear is not a transactional entity")
	}
	ge.config.driver = _tx.drv
	return ge
}

// String implements the fmt.Stringer.
func (ge *Gear) String() string {
	var builder strings.Builder
	builder.WriteString("Gear(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ge.ID))
	builder.WriteString("indx=")
	builder.WriteString(ge.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ge.Name)
	builder.WriteString(", ")
	builder.WriteString("gear_category=")
	builder.WriteString(ge.GearCategory)
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", ge.Quantity))
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", ge.EquipmentID))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ge *Gear) MarshalJSON() ([]byte, error) {
	type Alias Gear
	return json.Marshal(&struct {
		*Alias
		GearEdges
	}{
		Alias:     (*Alias)(ge),
		GearEdges: ge.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ge *Gear) UnmarshalJSON(data []byte) error {
	type Alias Gear
	aux := &struct {
		*Alias
		GearEdges
	}{
		Alias: (*Alias)(ge),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ge.Edges = aux.GearEdges
	return nil
}

func (gc *GearCreate) SetGear(input *Gear) *GearCreate {
	gc.SetIndx(input.Indx)
	gc.SetName(input.Name)
	gc.SetGearCategory(input.GearCategory)
	gc.SetQuantity(input.Quantity)
	gc.SetEquipmentID(input.EquipmentID)
	return gc
}

// Gears is a parsable slice of Gear.
type Gears []*Gear
