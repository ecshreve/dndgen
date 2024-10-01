// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/proficiency"
)

// Proficiency is the model entity for the Proficiency schema.
type Proficiency struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"indx,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Reference holds the value of the "reference" field.
	Reference string `json:"reference,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProficiencyQuery when eager-loading is set.
	Edges        ProficiencyEdges `json:"-"`
	selectValues sql.SelectValues
}

// ProficiencyEdges holds the relations/edges for other nodes in the graph.
type ProficiencyEdges struct {
	// Race holds the value of the race edge.
	Race []*Race `json:"race,omitempty"`
	// Options holds the value of the options edge.
	Options []*ProficiencyChoice `json:"options,omitempty"`
	// Class holds the value of the class edge.
	Class []*Class `json:"class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedRace    map[string][]*Race
	namedOptions map[string][]*ProficiencyChoice
	namedClass   map[string][]*Class
}

// RaceOrErr returns the Race value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) RaceOrErr() ([]*Race, error) {
	if e.loadedTypes[0] {
		return e.Race, nil
	}
	return nil, &NotLoadedError{edge: "race"}
}

// OptionsOrErr returns the Options value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) OptionsOrErr() ([]*ProficiencyChoice, error) {
	if e.loadedTypes[1] {
		return e.Options, nil
	}
	return nil, &NotLoadedError{edge: "options"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) ClassOrErr() ([]*Class, error) {
	if e.loadedTypes[2] {
		return e.Class, nil
	}
	return nil, &NotLoadedError{edge: "class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Proficiency) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case proficiency.FieldID:
			values[i] = new(sql.NullInt64)
		case proficiency.FieldIndx, proficiency.FieldName, proficiency.FieldReference:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Proficiency fields.
func (pr *Proficiency) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case proficiency.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case proficiency.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				pr.Indx = value.String
			}
		case proficiency.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case proficiency.FieldReference:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference", values[i])
			} else if value.Valid {
				pr.Reference = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Proficiency.
// This includes values selected through modifiers, order, etc.
func (pr *Proficiency) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryRace queries the "race" edge of the Proficiency entity.
func (pr *Proficiency) QueryRace() *RaceQuery {
	return NewProficiencyClient(pr.config).QueryRace(pr)
}

// QueryOptions queries the "options" edge of the Proficiency entity.
func (pr *Proficiency) QueryOptions() *ProficiencyChoiceQuery {
	return NewProficiencyClient(pr.config).QueryOptions(pr)
}

// QueryClass queries the "class" edge of the Proficiency entity.
func (pr *Proficiency) QueryClass() *ClassQuery {
	return NewProficiencyClient(pr.config).QueryClass(pr)
}

// Update returns a builder for updating this Proficiency.
// Note that you need to call Proficiency.Unwrap() before calling this method if this Proficiency
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Proficiency) Update() *ProficiencyUpdateOne {
	return NewProficiencyClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Proficiency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Proficiency) Unwrap() *Proficiency {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Proficiency is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Proficiency) String() string {
	var builder strings.Builder
	builder.WriteString("Proficiency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("indx=")
	builder.WriteString(pr.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("reference=")
	builder.WriteString(pr.Reference)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (pr *Proficiency) MarshalJSON() ([]byte, error) {
	type Alias Proficiency
	return json.Marshal(&struct {
		*Alias
		ProficiencyEdges
	}{
		Alias:            (*Alias)(pr),
		ProficiencyEdges: pr.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (pr *Proficiency) UnmarshalJSON(data []byte) error {
	type Alias Proficiency
	aux := &struct {
		*Alias
		ProficiencyEdges
	}{
		Alias: (*Alias)(pr),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	pr.Edges = aux.ProficiencyEdges
	return nil
}

func (pc *ProficiencyCreate) SetProficiency(input *Proficiency) *ProficiencyCreate {
	pc.SetIndx(input.Indx)
	pc.SetName(input.Name)
	pc.SetReference(input.Reference)
	return pc
}

// NamedRace returns the Race named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedRace(name string) ([]*Race, error) {
	if pr.Edges.namedRace == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedRace[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedRace(name string, edges ...*Race) {
	if pr.Edges.namedRace == nil {
		pr.Edges.namedRace = make(map[string][]*Race)
	}
	if len(edges) == 0 {
		pr.Edges.namedRace[name] = []*Race{}
	} else {
		pr.Edges.namedRace[name] = append(pr.Edges.namedRace[name], edges...)
	}
}

// NamedOptions returns the Options named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedOptions(name string) ([]*ProficiencyChoice, error) {
	if pr.Edges.namedOptions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedOptions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedOptions(name string, edges ...*ProficiencyChoice) {
	if pr.Edges.namedOptions == nil {
		pr.Edges.namedOptions = make(map[string][]*ProficiencyChoice)
	}
	if len(edges) == 0 {
		pr.Edges.namedOptions[name] = []*ProficiencyChoice{}
	} else {
		pr.Edges.namedOptions[name] = append(pr.Edges.namedOptions[name], edges...)
	}
}

// NamedClass returns the Class named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedClass(name string) ([]*Class, error) {
	if pr.Edges.namedClass == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedClass[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedClass(name string, edges ...*Class) {
	if pr.Edges.namedClass == nil {
		pr.Edges.namedClass = make(map[string][]*Class)
	}
	if len(edges) == 0 {
		pr.Edges.namedClass[name] = []*Class{}
	} else {
		pr.Edges.namedClass[name] = append(pr.Edges.namedClass[name], edges...)
	}
}

// Proficiencies is a parsable slice of Proficiency.
type Proficiencies []*Proficiency
