// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
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
	Edges        ClassEdges `json:"-"`
	selectValues sql.SelectValues
}

// ClassEdges holds the relations/edges for other nodes in the graph.
type ClassEdges struct {
	// Proficiencies holds the value of the proficiencies edge.
	Proficiencies []*Proficiency `json:"proficiencies,omitempty"`
	// ProficiencyChoices holds the value of the proficiency_choices edge.
	ProficiencyChoices []*Choice `json:"proficiency_choices,omitempty"`
	// StartingEquipment holds the value of the starting_equipment edge.
	StartingEquipment []*Equipment `json:"starting_equipment,omitempty"`
	// ClassStartingEquipment holds the value of the class_starting_equipment edge.
	ClassStartingEquipment []*StartingEquipment `json:"class_starting_equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedProficiencies          map[string][]*Proficiency
	namedProficiencyChoices     map[string][]*Choice
	namedStartingEquipment      map[string][]*Equipment
	namedClassStartingEquipment map[string][]*StartingEquipment
}

// ProficienciesOrErr returns the Proficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) ProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[0] {
		return e.Proficiencies, nil
	}
	return nil, &NotLoadedError{edge: "proficiencies"}
}

// ProficiencyChoicesOrErr returns the ProficiencyChoices value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) ProficiencyChoicesOrErr() ([]*Choice, error) {
	if e.loadedTypes[1] {
		return e.ProficiencyChoices, nil
	}
	return nil, &NotLoadedError{edge: "proficiency_choices"}
}

// StartingEquipmentOrErr returns the StartingEquipment value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) StartingEquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[2] {
		return e.StartingEquipment, nil
	}
	return nil, &NotLoadedError{edge: "starting_equipment"}
}

// ClassStartingEquipmentOrErr returns the ClassStartingEquipment value or an error if the edge
// was not loaded in eager-loading.
func (e ClassEdges) ClassStartingEquipmentOrErr() ([]*StartingEquipment, error) {
	if e.loadedTypes[3] {
		return e.ClassStartingEquipment, nil
	}
	return nil, &NotLoadedError{edge: "class_starting_equipment"}
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

// QueryProficiencies queries the "proficiencies" edge of the Class entity.
func (c *Class) QueryProficiencies() *ProficiencyQuery {
	return NewClassClient(c.config).QueryProficiencies(c)
}

// QueryProficiencyChoices queries the "proficiency_choices" edge of the Class entity.
func (c *Class) QueryProficiencyChoices() *ChoiceQuery {
	return NewClassClient(c.config).QueryProficiencyChoices(c)
}

// QueryStartingEquipment queries the "starting_equipment" edge of the Class entity.
func (c *Class) QueryStartingEquipment() *EquipmentQuery {
	return NewClassClient(c.config).QueryStartingEquipment(c)
}

// QueryClassStartingEquipment queries the "class_starting_equipment" edge of the Class entity.
func (c *Class) QueryClassStartingEquipment() *StartingEquipmentQuery {
	return NewClassClient(c.config).QueryClassStartingEquipment(c)
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

// MarshalJSON implements the json.Marshaler interface.
func (c *Class) MarshalJSON() ([]byte, error) {
	type Alias Class
	return json.Marshal(&struct {
		*Alias
		ClassEdges
	}{
		Alias:      (*Alias)(c),
		ClassEdges: c.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Class) UnmarshalJSON(data []byte) error {
	type Alias Class
	aux := &struct {
		*Alias
		ClassEdges
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	c.Edges = aux.ClassEdges
	return nil
}

func (cc *ClassCreate) SetClass(input *Class) *ClassCreate {
	cc.SetIndx(input.Indx)
	cc.SetName(input.Name)
	cc.SetHitDie(input.HitDie)
	return cc
}

// NamedProficiencies returns the Proficiencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Class) NamedProficiencies(name string) ([]*Proficiency, error) {
	if c.Edges.namedProficiencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProficiencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Class) appendNamedProficiencies(name string, edges ...*Proficiency) {
	if c.Edges.namedProficiencies == nil {
		c.Edges.namedProficiencies = make(map[string][]*Proficiency)
	}
	if len(edges) == 0 {
		c.Edges.namedProficiencies[name] = []*Proficiency{}
	} else {
		c.Edges.namedProficiencies[name] = append(c.Edges.namedProficiencies[name], edges...)
	}
}

// NamedProficiencyChoices returns the ProficiencyChoices named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Class) NamedProficiencyChoices(name string) ([]*Choice, error) {
	if c.Edges.namedProficiencyChoices == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProficiencyChoices[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Class) appendNamedProficiencyChoices(name string, edges ...*Choice) {
	if c.Edges.namedProficiencyChoices == nil {
		c.Edges.namedProficiencyChoices = make(map[string][]*Choice)
	}
	if len(edges) == 0 {
		c.Edges.namedProficiencyChoices[name] = []*Choice{}
	} else {
		c.Edges.namedProficiencyChoices[name] = append(c.Edges.namedProficiencyChoices[name], edges...)
	}
}

// NamedStartingEquipment returns the StartingEquipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Class) NamedStartingEquipment(name string) ([]*Equipment, error) {
	if c.Edges.namedStartingEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedStartingEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Class) appendNamedStartingEquipment(name string, edges ...*Equipment) {
	if c.Edges.namedStartingEquipment == nil {
		c.Edges.namedStartingEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		c.Edges.namedStartingEquipment[name] = []*Equipment{}
	} else {
		c.Edges.namedStartingEquipment[name] = append(c.Edges.namedStartingEquipment[name], edges...)
	}
}

// NamedClassStartingEquipment returns the ClassStartingEquipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Class) NamedClassStartingEquipment(name string) ([]*StartingEquipment, error) {
	if c.Edges.namedClassStartingEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedClassStartingEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Class) appendNamedClassStartingEquipment(name string, edges ...*StartingEquipment) {
	if c.Edges.namedClassStartingEquipment == nil {
		c.Edges.namedClassStartingEquipment = make(map[string][]*StartingEquipment)
	}
	if len(edges) == 0 {
		c.Edges.namedClassStartingEquipment[name] = []*StartingEquipment{}
	} else {
		c.Edges.namedClassStartingEquipment[name] = append(c.Edges.namedClassStartingEquipment[name], edges...)
	}
}

// Classes is a parsable slice of Class.
type Classes []*Class
