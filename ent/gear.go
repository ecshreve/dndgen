// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/gear"
)

// Gear is the model entity for the Gear schema.
type Gear struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GearQuery when eager-loading is set.
	Edges        GearEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GearEdges holds the relations/edges for other nodes in the graph.
type GearEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedEquipment map[string][]*Equipment
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e GearEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[0] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Gear) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gear.FieldID:
			values[i] = new(sql.NullInt64)
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
	builder.WriteString(fmt.Sprintf("id=%v", ge.ID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedEquipment returns the Equipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ge *Gear) NamedEquipment(name string) ([]*Equipment, error) {
	if ge.Edges.namedEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ge.Edges.namedEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ge *Gear) appendNamedEquipment(name string, edges ...*Equipment) {
	if ge.Edges.namedEquipment == nil {
		ge.Edges.namedEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		ge.Edges.namedEquipment[name] = []*Equipment{}
	} else {
		ge.Edges.namedEquipment[name] = append(ge.Edges.namedEquipment[name], edges...)
	}
}

// Gears is a parsable slice of Gear.
type Gears []*Gear
