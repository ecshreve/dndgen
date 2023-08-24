// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/tool"
)

// Tool is the model entity for the Tool schema.
type Tool struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ToolCategory holds the value of the "tool_category" field.
	ToolCategory string `json:"tool_category,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ToolQuery when eager-loading is set.
	Edges          ToolEdges `json:"edges"`
	equipment_tool *int
	selectValues   sql.SelectValues
}

// ToolEdges holds the relations/edges for other nodes in the graph.
type ToolEdges struct {
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
func (e ToolEdges) EquipmentOrErr() (*Equipment, error) {
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
func (*Tool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			values[i] = new(sql.NullInt64)
		case tool.FieldIndx, tool.FieldName, tool.FieldToolCategory:
			values[i] = new(sql.NullString)
		case tool.ForeignKeys[0]: // equipment_tool
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tool fields.
func (t *Tool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tool.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				t.Indx = value.String
			}
		case tool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tool.FieldToolCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tool_category", values[i])
			} else if value.Valid {
				t.ToolCategory = value.String
			}
		case tool.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_tool", value)
			} else if value.Valid {
				t.equipment_tool = new(int)
				*t.equipment_tool = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tool.
// This includes values selected through modifiers, order, etc.
func (t *Tool) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryEquipment queries the "equipment" edge of the Tool entity.
func (t *Tool) QueryEquipment() *EquipmentQuery {
	return NewToolClient(t.config).QueryEquipment(t)
}

// Update returns a builder for updating this Tool.
// Note that you need to call Tool.Unwrap() before calling this method if this Tool
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tool) Update() *ToolUpdateOne {
	return NewToolClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tool) Unwrap() *Tool {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tool is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tool) String() string {
	var builder strings.Builder
	builder.WriteString("Tool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("indx=")
	builder.WriteString(t.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("tool_category=")
	builder.WriteString(t.ToolCategory)
	builder.WriteByte(')')
	return builder.String()
}

func (tc *ToolCreate) SetTool(input *Tool) *ToolCreate {
	tc.SetIndx(input.Indx)
	tc.SetName(input.Name)
	tc.SetToolCategory(input.ToolCategory)
	return tc
}

// Tools is a parsable slice of Tool.
type Tools []*Tool
