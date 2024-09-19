// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
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
	Desc []string `json:"desc,omitempty"`
	// EquipmentCategory holds the value of the "equipment_category" field.
	EquipmentCategory equipment.EquipmentCategory `json:"equipment_category,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight float64 `json:"weight,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentQuery when eager-loading is set.
	Edges        EquipmentEdges `json:"-"`
	selectValues sql.SelectValues
}

// EquipmentEdges holds the relations/edges for other nodes in the graph.
type EquipmentEdges struct {
	// EquipmentCosts holds the value of the equipment_costs edge.
	EquipmentCosts []*EquipmentCost `json:"equipment_costs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedEquipmentCosts map[string][]*EquipmentCost
}

// EquipmentCostsOrErr returns the EquipmentCosts value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentEdges) EquipmentCostsOrErr() ([]*EquipmentCost, error) {
	if e.loadedTypes[0] {
		return e.EquipmentCosts, nil
	}
	return nil, &NotLoadedError{edge: "equipment_costs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Equipment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipment.FieldDesc:
			values[i] = new([]byte)
		case equipment.FieldWeight:
			values[i] = new(sql.NullFloat64)
		case equipment.FieldID:
			values[i] = new(sql.NullInt64)
		case equipment.FieldIndx, equipment.FieldName, equipment.FieldEquipmentCategory:
			values[i] = new(sql.NullString)
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
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Desc); err != nil {
					return fmt.Errorf("unmarshal field desc: %w", err)
				}
			}
		case equipment.FieldEquipmentCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_category", values[i])
			} else if value.Valid {
				e.EquipmentCategory = equipment.EquipmentCategory(value.String)
			}
		case equipment.FieldWeight:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				e.Weight = value.Float64
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

// QueryEquipmentCosts queries the "equipment_costs" edge of the Equipment entity.
func (e *Equipment) QueryEquipmentCosts() *EquipmentCostQuery {
	return NewEquipmentClient(e.config).QueryEquipmentCosts(e)
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
	builder.WriteString(fmt.Sprintf("%v", e.Desc))
	builder.WriteString(", ")
	builder.WriteString("equipment_category=")
	builder.WriteString(fmt.Sprintf("%v", e.EquipmentCategory))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", e.Weight))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (e *Equipment) MarshalJSON() ([]byte, error) {
	type Alias Equipment
	return json.Marshal(&struct {
		*Alias
		EquipmentEdges
	}{
		Alias:          (*Alias)(e),
		EquipmentEdges: e.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *Equipment) UnmarshalJSON(data []byte) error {
	type Alias Equipment
	aux := &struct {
		*Alias
		EquipmentEdges
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	e.Edges = aux.EquipmentEdges
	return nil
}

func (ec *EquipmentCreate) SetEquipment(input *Equipment) *EquipmentCreate {
	ec.SetIndx(input.Indx)
	ec.SetName(input.Name)
	ec.SetDesc(input.Desc)
	ec.SetEquipmentCategory(input.EquipmentCategory)
	ec.SetWeight(input.Weight)
	return ec
}

// NamedEquipmentCosts returns the EquipmentCosts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Equipment) NamedEquipmentCosts(name string) ([]*EquipmentCost, error) {
	if e.Edges.namedEquipmentCosts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedEquipmentCosts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Equipment) appendNamedEquipmentCosts(name string, edges ...*EquipmentCost) {
	if e.Edges.namedEquipmentCosts == nil {
		e.Edges.namedEquipmentCosts = make(map[string][]*EquipmentCost)
	}
	if len(edges) == 0 {
		e.Edges.namedEquipmentCosts[name] = []*EquipmentCost{}
	} else {
		e.Edges.namedEquipmentCosts[name] = append(e.Edges.namedEquipmentCosts[name], edges...)
	}
}

// EquipmentSlice is a parsable slice of Equipment.
type EquipmentSlice []*Equipment
