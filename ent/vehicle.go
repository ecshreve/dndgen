// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/vehicle"
)

// Vehicle is the model entity for the Vehicle schema.
type Vehicle struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// VehicleCategory holds the value of the "vehicle_category" field.
	VehicleCategory string `json:"vehicle_category,omitempty"`
	// Capacity holds the value of the "capacity" field.
	Capacity string `json:"capacity,omitempty"`
	// EquipmentID holds the value of the "equipment_id" field.
	EquipmentID int `json:"equipment_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VehicleQuery when eager-loading is set.
	Edges        VehicleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VehicleEdges holds the relations/edges for other nodes in the graph.
type VehicleEdges struct {
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
func (e VehicleEdges) EquipmentOrErr() (*Equipment, error) {
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
func (*Vehicle) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldID, vehicle.FieldEquipmentID:
			values[i] = new(sql.NullInt64)
		case vehicle.FieldIndx, vehicle.FieldName, vehicle.FieldVehicleCategory, vehicle.FieldCapacity:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vehicle fields.
func (v *Vehicle) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vehicle.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				v.Indx = value.String
			}
		case vehicle.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case vehicle.FieldVehicleCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vehicle_category", values[i])
			} else if value.Valid {
				v.VehicleCategory = value.String
			}
		case vehicle.FieldCapacity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field capacity", values[i])
			} else if value.Valid {
				v.Capacity = value.String
			}
		case vehicle.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				v.EquipmentID = int(value.Int64)
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vehicle.
// This includes values selected through modifiers, order, etc.
func (v *Vehicle) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryEquipment queries the "equipment" edge of the Vehicle entity.
func (v *Vehicle) QueryEquipment() *EquipmentQuery {
	return NewVehicleClient(v.config).QueryEquipment(v)
}

// Update returns a builder for updating this Vehicle.
// Note that you need to call Vehicle.Unwrap() before calling this method if this Vehicle
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vehicle) Update() *VehicleUpdateOne {
	return NewVehicleClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vehicle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vehicle) Unwrap() *Vehicle {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vehicle is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vehicle) String() string {
	var builder strings.Builder
	builder.WriteString("Vehicle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("indx=")
	builder.WriteString(v.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(v.Name)
	builder.WriteString(", ")
	builder.WriteString("vehicle_category=")
	builder.WriteString(v.VehicleCategory)
	builder.WriteString(", ")
	builder.WriteString("capacity=")
	builder.WriteString(v.Capacity)
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", v.EquipmentID))
	builder.WriteByte(')')
	return builder.String()
}

func (vc *VehicleCreate) SetVehicle(input *Vehicle) *VehicleCreate {
	vc.SetIndx(input.Indx)
	vc.SetName(input.Name)
	vc.SetVehicleCategory(input.VehicleCategory)
	vc.SetCapacity(input.Capacity)
	vc.SetEquipmentID(input.EquipmentID)
	return vc
}

// Vehicles is a parsable slice of Vehicle.
type Vehicles []*Vehicle
