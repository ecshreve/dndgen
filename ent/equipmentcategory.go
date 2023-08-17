// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipmentcategory"
)

// EquipmentCategory is the model entity for the EquipmentCategory schema.
type EquipmentCategory struct {
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
	// The values are being populated by the EquipmentCategoryQuery when eager-loading is set.
	Edges                 EquipmentCategoryEdges `json:"edges"`
	equipment_subcategory *int
	selectValues          sql.SelectValues
}

// EquipmentCategoryEdges holds the relations/edges for other nodes in the graph.
type EquipmentCategoryEdges struct {
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
func (e EquipmentCategoryEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[0] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipmentcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case equipmentcategory.FieldIndx, equipmentcategory.FieldName, equipmentcategory.FieldDesc:
			values[i] = new(sql.NullString)
		case equipmentcategory.ForeignKeys[0]: // equipment_subcategory
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentCategory fields.
func (ec *EquipmentCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipmentcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case equipmentcategory.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				ec.Indx = value.String
			}
		case equipmentcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ec.Name = value.String
			}
		case equipmentcategory.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				ec.Desc = value.String
			}
		case equipmentcategory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_subcategory", value)
			} else if value.Valid {
				ec.equipment_subcategory = new(int)
				*ec.equipment_subcategory = int(value.Int64)
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EquipmentCategory.
// This includes values selected through modifiers, order, etc.
func (ec *EquipmentCategory) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// QueryEquipment queries the "equipment" edge of the EquipmentCategory entity.
func (ec *EquipmentCategory) QueryEquipment() *EquipmentQuery {
	return NewEquipmentCategoryClient(ec.config).QueryEquipment(ec)
}

// Update returns a builder for updating this EquipmentCategory.
// Note that you need to call EquipmentCategory.Unwrap() before calling this method if this EquipmentCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EquipmentCategory) Update() *EquipmentCategoryUpdateOne {
	return NewEquipmentCategoryClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the EquipmentCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EquipmentCategory) Unwrap() *EquipmentCategory {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentCategory is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EquipmentCategory) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("indx=")
	builder.WriteString(ec.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ec.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(ec.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// NamedEquipment returns the Equipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ec *EquipmentCategory) NamedEquipment(name string) ([]*Equipment, error) {
	if ec.Edges.namedEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ec.Edges.namedEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ec *EquipmentCategory) appendNamedEquipment(name string, edges ...*Equipment) {
	if ec.Edges.namedEquipment == nil {
		ec.Edges.namedEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		ec.Edges.namedEquipment[name] = []*Equipment{}
	} else {
		ec.Edges.namedEquipment[name] = append(ec.Edges.namedEquipment[name], edges...)
	}
}

// EquipmentCategories is a parsable slice of EquipmentCategory.
type EquipmentCategories []*EquipmentCategory
