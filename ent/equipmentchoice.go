// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipmentchoice"
)

// EquipmentChoice is the model entity for the EquipmentChoice schema.
type EquipmentChoice struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ClassID holds the value of the "class_id" field.
	ClassID int `json:"class_id,omitempty"`
	// Choose holds the value of the "choose" field.
	Choose int `json:"choose,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentChoiceQuery when eager-loading is set.
	Edges        EquipmentChoiceEdges `json:"-"`
	selectValues sql.SelectValues
}

// EquipmentChoiceEdges holds the relations/edges for other nodes in the graph.
type EquipmentChoiceEdges struct {
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedEquipment map[string][]*Equipment
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentChoiceEdges) ClassOrErr() (*Class, error) {
	if e.loadedTypes[0] {
		if e.Class == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: class.Label}
		}
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentChoiceEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[1] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentChoice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipmentchoice.FieldID, equipmentchoice.FieldClassID, equipmentchoice.FieldChoose:
			values[i] = new(sql.NullInt64)
		case equipmentchoice.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentChoice fields.
func (ec *EquipmentChoice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipmentchoice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case equipmentchoice.FieldClassID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field class_id", values[i])
			} else if value.Valid {
				ec.ClassID = int(value.Int64)
			}
		case equipmentchoice.FieldChoose:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field choose", values[i])
			} else if value.Valid {
				ec.Choose = int(value.Int64)
			}
		case equipmentchoice.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				ec.Desc = value.String
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EquipmentChoice.
// This includes values selected through modifiers, order, etc.
func (ec *EquipmentChoice) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// QueryClass queries the "class" edge of the EquipmentChoice entity.
func (ec *EquipmentChoice) QueryClass() *ClassQuery {
	return NewEquipmentChoiceClient(ec.config).QueryClass(ec)
}

// QueryEquipment queries the "equipment" edge of the EquipmentChoice entity.
func (ec *EquipmentChoice) QueryEquipment() *EquipmentQuery {
	return NewEquipmentChoiceClient(ec.config).QueryEquipment(ec)
}

// Update returns a builder for updating this EquipmentChoice.
// Note that you need to call EquipmentChoice.Unwrap() before calling this method if this EquipmentChoice
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EquipmentChoice) Update() *EquipmentChoiceUpdateOne {
	return NewEquipmentChoiceClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the EquipmentChoice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EquipmentChoice) Unwrap() *EquipmentChoice {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentChoice is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EquipmentChoice) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentChoice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("class_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.ClassID))
	builder.WriteString(", ")
	builder.WriteString("choose=")
	builder.WriteString(fmt.Sprintf("%v", ec.Choose))
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(ec.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ec *EquipmentChoice) MarshalJSON() ([]byte, error) {
	type Alias EquipmentChoice
	return json.Marshal(&struct {
		*Alias
		EquipmentChoiceEdges
	}{
		Alias:                (*Alias)(ec),
		EquipmentChoiceEdges: ec.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ec *EquipmentChoice) UnmarshalJSON(data []byte) error {
	type Alias EquipmentChoice
	aux := &struct {
		*Alias
		EquipmentChoiceEdges
	}{
		Alias: (*Alias)(ec),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ec.Edges = aux.EquipmentChoiceEdges
	return nil
}

func (ecc *EquipmentChoiceCreate) SetEquipmentChoice(input *EquipmentChoice) *EquipmentChoiceCreate {
	ecc.SetClassID(input.ClassID)
	ecc.SetChoose(input.Choose)
	ecc.SetDesc(input.Desc)
	return ecc
}

// NamedEquipment returns the Equipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ec *EquipmentChoice) NamedEquipment(name string) ([]*Equipment, error) {
	if ec.Edges.namedEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ec.Edges.namedEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ec *EquipmentChoice) appendNamedEquipment(name string, edges ...*Equipment) {
	if ec.Edges.namedEquipment == nil {
		ec.Edges.namedEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		ec.Edges.namedEquipment[name] = []*Equipment{}
	} else {
		ec.Edges.namedEquipment[name] = append(ec.Edges.namedEquipment[name], edges...)
	}
}

// EquipmentChoices is a parsable slice of EquipmentChoice.
type EquipmentChoices []*EquipmentChoice