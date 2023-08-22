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
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// HitDie holds the value of the "hit_die" field.
	HitDie int `json:"hit_die,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassQuery when eager-loading is set.
	Edges        ClassEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// SavingThrows holds the value of the saving_throws edge.
	SavingThrows []*AbilityScore `json:"saving_throws,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedSavingThrows map[string][]*AbilityScore
}

// SavingThrowsOrErr returns the SavingThrows value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) SavingThrowsOrErr() ([]*AbilityScore, error) {
	if e.loadedTypes[0] {
		return e.SavingThrows, nil
	}
	return nil, &NotLoadedError{edge: "saving_throws"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Class) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case class.FieldID, class.FieldHitDie:
			values[i] = new(sql.NullInt64)
		case class.FieldIndx, class.FieldName:
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

// QuerySavingThrows queries the "saving_throws" edge of the Class entity.
func (c *Class) QuerySavingThrows() *AbilityScoreQuery {
	return NewClassClient(c.config).QuerySavingThrows(c)
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
	builder.WriteString("hit_die=")
	builder.WriteString(fmt.Sprintf("%v", c.HitDie))
	builder.WriteByte(')')
	return builder.String()
}

func (cc *ClassCreate) SetClass(input *Class) *ClassCreate {
	cc.SetIndx(input.Indx)
	cc.SetName(input.Name)
	cc.SetHitDie(input.HitDie)
	return cc
}

// NamedSavingThrows returns the SavingThrows named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Class) NamedSavingThrows(name string) ([]*AbilityScore, error) {
	if c.Edges.namedSavingThrows == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedSavingThrows[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Class) appendNamedSavingThrows(name string, edges ...*AbilityScore) {
	if c.Edges.namedSavingThrows == nil {
		c.Edges.namedSavingThrows = make(map[string][]*AbilityScore)
	}
	if len(edges) == 0 {
		c.Edges.namedSavingThrows[name] = []*AbilityScore{}
	} else {
		c.Edges.namedSavingThrows[name] = append(c.Edges.namedSavingThrows[name], edges...)
	}
}

// Classes is a parsable slice of Class.
type Classes []*Class
