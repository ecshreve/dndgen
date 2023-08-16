// Code generated by ent, DO NOT EDIT.

package ent

import (
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
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// Speed holds the value of the "speed" field.
	Speed int `json:"speed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RaceQuery when eager-loading is set.
	Edges        RaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RaceEdges holds the relations/edges for other nodes in the graph.
type RaceEdges struct {
	// Languages holds the value of the languages edge.
	Languages []*Language `json:"languages,omitempty"`
	// AbilityBonuses holds the value of the ability_bonuses edge.
	AbilityBonuses []*AbilityBonus `json:"ability_bonuses,omitempty"`
	// StartingProficiencies holds the value of the starting_proficiencies edge.
	StartingProficiencies []*Proficiency `json:"starting_proficiencies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// LanguagesOrErr returns the Languages value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) LanguagesOrErr() ([]*Language, error) {
	if e.loadedTypes[0] {
		return e.Languages, nil
	}
	return nil, &NotLoadedError{edge: "languages"}
}

// AbilityBonusesOrErr returns the AbilityBonuses value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) AbilityBonusesOrErr() ([]*AbilityBonus, error) {
	if e.loadedTypes[1] {
		return e.AbilityBonuses, nil
	}
	return nil, &NotLoadedError{edge: "ability_bonuses"}
}

// StartingProficienciesOrErr returns the StartingProficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) StartingProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[2] {
		return e.StartingProficiencies, nil
	}
	return nil, &NotLoadedError{edge: "starting_proficiencies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Race) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case race.FieldID, race.FieldSpeed:
			values[i] = new(sql.NullInt64)
		case race.FieldIndx, race.FieldName, race.FieldDesc:
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
		case race.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				r.Desc = value.String
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

// QueryLanguages queries the "languages" edge of the Race entity.
func (r *Race) QueryLanguages() *LanguageQuery {
	return NewRaceClient(r.config).QueryLanguages(r)
}

// QueryAbilityBonuses queries the "ability_bonuses" edge of the Race entity.
func (r *Race) QueryAbilityBonuses() *AbilityBonusQuery {
	return NewRaceClient(r.config).QueryAbilityBonuses(r)
}

// QueryStartingProficiencies queries the "starting_proficiencies" edge of the Race entity.
func (r *Race) QueryStartingProficiencies() *ProficiencyQuery {
	return NewRaceClient(r.config).QueryStartingProficiencies(r)
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
	builder.WriteString("desc=")
	builder.WriteString(r.Desc)
	builder.WriteString(", ")
	builder.WriteString("speed=")
	builder.WriteString(fmt.Sprintf("%v", r.Speed))
	builder.WriteByte(')')
	return builder.String()
}

// Races is a parsable slice of Race.
type Races []*Race
