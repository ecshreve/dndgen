// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/choice"
)

// Choice is the model entity for the Choice schema.
type Choice struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Choose holds the value of the "choose" field.
	Choose int `json:"choose,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChoiceQuery when eager-loading is set.
	Edges          ChoiceEdges `json:"-"`
	choice_choices *int
	selectValues   sql.SelectValues
}

// ChoiceEdges holds the relations/edges for other nodes in the graph.
type ChoiceEdges struct {
	// ParentChoice holds the value of the parent_choice edge.
	ParentChoice *Choice `json:"parent_choice,omitempty"`
	// Choices holds the value of the choices edge.
	Choices []*Choice `json:"choices,omitempty"`
	// ProficiencyOptions holds the value of the proficiency_options edge.
	ProficiencyOptions []*Proficiency `json:"proficiency_options,omitempty"`
	// Class holds the value of the class edge.
	Class []*Class `json:"class,omitempty"`
	// Race holds the value of the race edge.
	Race []*Race `json:"race,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedChoices            map[string][]*Choice
	namedProficiencyOptions map[string][]*Proficiency
	namedClass              map[string][]*Class
	namedRace               map[string][]*Race
}

// ParentChoiceOrErr returns the ParentChoice value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChoiceEdges) ParentChoiceOrErr() (*Choice, error) {
	if e.loadedTypes[0] {
		if e.ParentChoice == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: choice.Label}
		}
		return e.ParentChoice, nil
	}
	return nil, &NotLoadedError{edge: "parent_choice"}
}

// ChoicesOrErr returns the Choices value or an error if the edge
// was not loaded in eager-loading.
func (e ChoiceEdges) ChoicesOrErr() ([]*Choice, error) {
	if e.loadedTypes[1] {
		return e.Choices, nil
	}
	return nil, &NotLoadedError{edge: "choices"}
}

// ProficiencyOptionsOrErr returns the ProficiencyOptions value or an error if the edge
// was not loaded in eager-loading.
func (e ChoiceEdges) ProficiencyOptionsOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[2] {
		return e.ProficiencyOptions, nil
	}
	return nil, &NotLoadedError{edge: "proficiency_options"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading.
func (e ChoiceEdges) ClassOrErr() ([]*Class, error) {
	if e.loadedTypes[3] {
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// RaceOrErr returns the Race value or an error if the edge
// was not loaded in eager-loading.
func (e ChoiceEdges) RaceOrErr() ([]*Race, error) {
	if e.loadedTypes[4] {
		return e.Race, nil
	}
	return nil, &NotLoadedError{edge: "race"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Choice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case choice.FieldID, choice.FieldChoose:
			values[i] = new(sql.NullInt64)
		case choice.FieldDesc:
			values[i] = new(sql.NullString)
		case choice.ForeignKeys[0]: // choice_choices
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Choice fields.
func (c *Choice) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case choice.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case choice.FieldChoose:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field choose", values[i])
			} else if value.Valid {
				c.Choose = int(value.Int64)
			}
		case choice.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				c.Desc = value.String
			}
		case choice.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field choice_choices", value)
			} else if value.Valid {
				c.choice_choices = new(int)
				*c.choice_choices = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Choice.
// This includes values selected through modifiers, order, etc.
func (c *Choice) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryParentChoice queries the "parent_choice" edge of the Choice entity.
func (c *Choice) QueryParentChoice() *ChoiceQuery {
	return NewChoiceClient(c.config).QueryParentChoice(c)
}

// QueryChoices queries the "choices" edge of the Choice entity.
func (c *Choice) QueryChoices() *ChoiceQuery {
	return NewChoiceClient(c.config).QueryChoices(c)
}

// QueryProficiencyOptions queries the "proficiency_options" edge of the Choice entity.
func (c *Choice) QueryProficiencyOptions() *ProficiencyQuery {
	return NewChoiceClient(c.config).QueryProficiencyOptions(c)
}

// QueryClass queries the "class" edge of the Choice entity.
func (c *Choice) QueryClass() *ClassQuery {
	return NewChoiceClient(c.config).QueryClass(c)
}

// QueryRace queries the "race" edge of the Choice entity.
func (c *Choice) QueryRace() *RaceQuery {
	return NewChoiceClient(c.config).QueryRace(c)
}

// Update returns a builder for updating this Choice.
// Note that you need to call Choice.Unwrap() before calling this method if this Choice
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Choice) Update() *ChoiceUpdateOne {
	return NewChoiceClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Choice entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Choice) Unwrap() *Choice {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Choice is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Choice) String() string {
	var builder strings.Builder
	builder.WriteString("Choice(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("choose=")
	builder.WriteString(fmt.Sprintf("%v", c.Choose))
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(c.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (c *Choice) MarshalJSON() ([]byte, error) {
	type Alias Choice
	return json.Marshal(&struct {
		*Alias
		ChoiceEdges
	}{
		Alias:       (*Alias)(c),
		ChoiceEdges: c.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Choice) UnmarshalJSON(data []byte) error {
	type Alias Choice
	aux := &struct {
		*Alias
		ChoiceEdges
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	c.Edges = aux.ChoiceEdges
	return nil
}

func (cc *ChoiceCreate) SetChoice(input *Choice) *ChoiceCreate {
	cc.SetChoose(input.Choose)
	cc.SetDesc(input.Desc)
	return cc
}

// NamedChoices returns the Choices named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Choice) NamedChoices(name string) ([]*Choice, error) {
	if c.Edges.namedChoices == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedChoices[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Choice) appendNamedChoices(name string, edges ...*Choice) {
	if c.Edges.namedChoices == nil {
		c.Edges.namedChoices = make(map[string][]*Choice)
	}
	if len(edges) == 0 {
		c.Edges.namedChoices[name] = []*Choice{}
	} else {
		c.Edges.namedChoices[name] = append(c.Edges.namedChoices[name], edges...)
	}
}

// NamedProficiencyOptions returns the ProficiencyOptions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Choice) NamedProficiencyOptions(name string) ([]*Proficiency, error) {
	if c.Edges.namedProficiencyOptions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProficiencyOptions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Choice) appendNamedProficiencyOptions(name string, edges ...*Proficiency) {
	if c.Edges.namedProficiencyOptions == nil {
		c.Edges.namedProficiencyOptions = make(map[string][]*Proficiency)
	}
	if len(edges) == 0 {
		c.Edges.namedProficiencyOptions[name] = []*Proficiency{}
	} else {
		c.Edges.namedProficiencyOptions[name] = append(c.Edges.namedProficiencyOptions[name], edges...)
	}
}

// NamedClass returns the Class named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Choice) NamedClass(name string) ([]*Class, error) {
	if c.Edges.namedClass == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedClass[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Choice) appendNamedClass(name string, edges ...*Class) {
	if c.Edges.namedClass == nil {
		c.Edges.namedClass = make(map[string][]*Class)
	}
	if len(edges) == 0 {
		c.Edges.namedClass[name] = []*Class{}
	} else {
		c.Edges.namedClass[name] = append(c.Edges.namedClass[name], edges...)
	}
}

// NamedRace returns the Race named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Choice) NamedRace(name string) ([]*Race, error) {
	if c.Edges.namedRace == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedRace[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Choice) appendNamedRace(name string, edges ...*Race) {
	if c.Edges.namedRace == nil {
		c.Edges.namedRace = make(map[string][]*Race)
	}
	if len(edges) == 0 {
		c.Edges.namedRace[name] = []*Race{}
	} else {
		c.Edges.namedRace[name] = append(c.Edges.namedRace[name], edges...)
	}
}

// Choices is a parsable slice of Choice.
type Choices []*Choice
