// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/class"
)

// Class is the model entity for the Class schema.
type Class struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"indx,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// HitDie holds the value of the "hit_die" field.
	HitDie int `json:"hit_die,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges        ClassEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// StartingProficiencies holds the value of the starting_proficiencies edge.
	StartingProficiencies []*Proficiency `json:"starting_proficiencies,omitempty"`
	// StartingEquipment holds the value of the starting_equipment edge.
	StartingEquipment []*Equipment `json:"starting_equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StartingProficienciesOrErr returns the StartingProficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StartingProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[0] {
		return e.StartingProficiencies, nil
	}
	return nil, &NotLoadedError{edge: "starting_proficiencies"}
}

// StartingEquipmentOrErr returns the StartingEquipment value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StartingEquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[1] {
		return e.StartingEquipment, nil
	}
	return nil, &NotLoadedError{edge: "starting_equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldID, class.FieldHitDie:
			values[i] = new(sql.NullInt64)
		case class.FieldIndx, class.FieldName, class.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Class fields.
func (c *Class) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case class.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case class.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				c.Indx = value.String
			}
		case class.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case class.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				c.Desc = value.String
			}
		case class.FieldHitDie:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field hit_die", values[i])
			} else if value.Valid {
				c.HitDie = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Class.
// This includes values selected through modifiers, order, etc.
func (c *Class) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryStartingProficiencies queries the "starting_proficiencies" edge of the Class entity.
func (c *Class) QueryStartingProficiencies() *ProficiencyQuery {
	return NewClassClient(c.config).QueryStartingProficiencies(c)
}

// QueryStartingEquipment queries the "starting_equipment" edge of the Class entity.
func (c *Class) QueryStartingEquipment() *EquipmentQuery {
	return NewClassClient(c.config).QueryStartingEquipment(c)
}

// Update returns a builder for updating this Class.
// Note that you need to call Class.Unwrap() before calling this method if this Class
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Class) Update() *ClassUpdateOne {
	return NewClassClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Class entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Class) Unwrap() *Class {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Class is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Class) String() string {
	var builder strings.Builder
	builder.WriteString("Class(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("indx=")
	builder.WriteString(c.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(c.Desc)
	builder.WriteString(", ")
	builder.WriteString("hit_die=")
	builder.WriteString(fmt.Sprintf("%v", c.HitDie))
	builder.WriteByte(')')
	return builder.String()
}

// Classes is a parsable slice of Class.
type Classes []*Class