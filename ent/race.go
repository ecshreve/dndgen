// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/race"
)

// Race is the model entity for the Race schema.
type Race struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Speed holds the value of the "speed" field.
	Speed int `json:"speed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RaceQuery when eager-loading is set.
	Edges        RaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RaceEdges holds the relations/edges for other nodes in the graph.
type RaceEdges struct {
	// Proficiencies holds the value of the proficiencies edge.
	Proficiencies []*Proficiency `json:"proficiencies,omitempty"`
	// Languages holds the value of the languages edge.
	Languages []*Language `json:"languages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedProficiencies map[string][]*Proficiency
	namedLanguages     map[string][]*Language
}

// ProficienciesOrErr returns the Proficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) ProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[0] {
		return e.Proficiencies, nil
	}
	return nil, &NotLoadedError{edge: "proficiencies"}
}

// LanguagesOrErr returns the Languages value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) LanguagesOrErr() ([]*Language, error) {
	if e.loadedTypes[1] {
		return e.Languages, nil
	}
	return nil, &NotLoadedError{edge: "languages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Race) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case race.FieldID, race.FieldSpeed:
			values[i] = new(sql.NullInt64)
		case race.FieldIndx, race.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Race fields.
func (r *Race) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case race.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case race.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				r.Indx = value.String
			}
		case race.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case race.FieldSpeed:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field speed", values[i])
			} else if value.Valid {
				r.Speed = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Race.
// This includes values selected through modifiers, order, etc.
func (r *Race) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryProficiencies queries the "proficiencies" edge of the Race entity.
func (r *Race) QueryProficiencies() *ProficiencyQuery {
	return NewRaceClient(r.config).QueryProficiencies(r)
}

// QueryLanguages queries the "languages" edge of the Race entity.
func (r *Race) QueryLanguages() *LanguageQuery {
	return NewRaceClient(r.config).QueryLanguages(r)
}

// Update returns a builder for updating this Race.
// Note that you need to call Race.Unwrap() before calling this method if this Race
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Race) Update() *RaceUpdateOne {
	return NewRaceClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Race entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Race) Unwrap() *Race {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Race is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Race) String() string {
	var builder strings.Builder
	builder.WriteString("Race(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("indx=")
	builder.WriteString(r.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("speed=")
	builder.WriteString(fmt.Sprintf("%v", r.Speed))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
// func (r *Race) MarshalJSON() ([]byte, error) {
// 		type Alias Race
// 		return json.Marshal(&struct {
// 				*Alias
// 				RaceEdges
// 		}{
// 				Alias: (*Alias)(r),
// 				RaceEdges: r.Edges,
// 		})
// }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (r *Race) UnmarshalJSON(data []byte) error {
	type Alias Race
	aux := &struct {
		*Alias
		RaceEdges
	}{
		Alias: (*Alias)(r),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	r.Edges = aux.RaceEdges
	return nil
}

func (rc *RaceCreate) SetRace(input *Race) *RaceCreate {
	rc.SetIndx(input.Indx)
	rc.SetName(input.Name)
	rc.SetSpeed(input.Speed)
	return rc
}

// NamedProficiencies returns the Proficiencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Race) NamedProficiencies(name string) ([]*Proficiency, error) {
	if r.Edges.namedProficiencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedProficiencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Race) appendNamedProficiencies(name string, edges ...*Proficiency) {
	if r.Edges.namedProficiencies == nil {
		r.Edges.namedProficiencies = make(map[string][]*Proficiency)
	}
	if len(edges) == 0 {
		r.Edges.namedProficiencies[name] = []*Proficiency{}
	} else {
		r.Edges.namedProficiencies[name] = append(r.Edges.namedProficiencies[name], edges...)
	}
}

// NamedLanguages returns the Languages named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Race) NamedLanguages(name string) ([]*Language, error) {
	if r.Edges.namedLanguages == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedLanguages[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Race) appendNamedLanguages(name string, edges ...*Language) {
	if r.Edges.namedLanguages == nil {
		r.Edges.namedLanguages = make(map[string][]*Language)
	}
	if len(edges) == 0 {
		r.Edges.namedLanguages[name] = []*Language{}
	} else {
		r.Edges.namedLanguages[name] = append(r.Edges.namedLanguages[name], edges...)
	}
}

// Races is a parsable slice of Race.
type Races []*Race
