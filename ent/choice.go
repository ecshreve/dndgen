// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/choice"
	"github.com/ecshreve/dndgen/ent/race"
)

// Choice is the model entity for the Choice schema.
type Choice struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RaceID holds the value of the "race_id" field.
	RaceID int `json:"race_id,omitempty"`
	// Choose holds the value of the "choose" field.
	Choose int `json:"choose,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChoiceQuery when eager-loading is set.
	Edges        ChoiceEdges `json:"-"`
	selectValues sql.SelectValues
}

// ChoiceEdges holds the relations/edges for other nodes in the graph.
type ChoiceEdges struct {
	// Proficiencies holds the value of the proficiencies edge.
	Proficiencies []*Proficiency `json:"proficiencies,omitempty"`
	// Race holds the value of the race edge.
	Race *Race `json:"race,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedProficiencies map[string][]*Proficiency
}

// ProficienciesOrErr returns the Proficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e ChoiceEdges) ProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[0] {
		return e.Proficiencies, nil
	}
	return nil, &NotLoadedError{edge: "proficiencies"}
}

// RaceOrErr returns the Race value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChoiceEdges) RaceOrErr() (*Race, error) {
	if e.loadedTypes[1] {
		if e.Race == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: race.Label}
		}
		return e.Race, nil
	}
	return nil, &NotLoadedError{edge: "race"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Choice) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case choice.FieldID, choice.FieldRaceID, choice.FieldChoose:
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
		case choice.FieldRaceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field race_id", values[i])
			} else if value.Valid {
				c.RaceID = int(value.Int64)
			}
		case choice.FieldChoose:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field choose", values[i])
			} else if value.Valid {
				c.Choose = int(value.Int64)
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

// QueryProficiencies queries the "proficiencies" edge of the Choice entity.
func (c *Choice) QueryProficiencies() *ProficiencyQuery {
	return NewChoiceClient(c.config).QueryProficiencies(c)
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
	builder.WriteString("race_id=")
	builder.WriteString(fmt.Sprintf("%v", c.RaceID))
	builder.WriteString(", ")
	builder.WriteString("choose=")
	builder.WriteString(fmt.Sprintf("%v", c.Choose))
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
	cc.SetRaceID(input.RaceID)
	cc.SetChoose(input.Choose)
	return cc
}

// NamedProficiencies returns the Proficiencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Choice) NamedProficiencies(name string) ([]*Proficiency, error) {
	if c.Edges.namedProficiencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedProficiencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Choice) appendNamedProficiencies(name string, edges ...*Proficiency) {
	if c.Edges.namedProficiencies == nil {
		c.Edges.namedProficiencies = make(map[string][]*Proficiency)
	}
	if len(edges) == 0 {
		c.Edges.namedProficiencies[name] = []*Proficiency{}
	} else {
		c.Edges.namedProficiencies[name] = append(c.Edges.namedProficiencies[name], edges...)
	}
}

// Choices is a parsable slice of Choice.
type Choices []*Choice
