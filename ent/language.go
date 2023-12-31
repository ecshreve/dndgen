// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/language"
)

// Language is the model entity for the Language schema.
type Language struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// LanguageType holds the value of the "language_type" field.
	LanguageType language.LanguageType `json:"type"`
	// Script holds the value of the "script" field.
	Script language.Script `json:"script,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LanguageQuery when eager-loading is set.
	Edges        LanguageEdges `json:"-"`
	selectValues sql.SelectValues
}

// LanguageEdges holds the relations/edges for other nodes in the graph.
type LanguageEdges struct {
	// RaceSpeakers holds the value of the race_speakers edge.
	RaceSpeakers []*Race `json:"race_speakers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedRaceSpeakers map[string][]*Race
}

// RaceSpeakersOrErr returns the RaceSpeakers value or an error if the edge
// was not loaded in eager-loading.
func (e LanguageEdges) RaceSpeakersOrErr() ([]*Race, error) {
	if e.loadedTypes[0] {
		return e.RaceSpeakers, nil
	}
	return nil, &NotLoadedError{edge: "race_speakers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Language) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case language.FieldID:
			values[i] = new(sql.NullInt64)
		case language.FieldIndx, language.FieldName, language.FieldDesc, language.FieldLanguageType, language.FieldScript:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Language fields.
func (l *Language) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case language.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case language.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				l.Indx = value.String
			}
		case language.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		case language.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				l.Desc = value.String
			}
		case language.FieldLanguageType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language_type", values[i])
			} else if value.Valid {
				l.LanguageType = language.LanguageType(value.String)
			}
		case language.FieldScript:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field script", values[i])
			} else if value.Valid {
				l.Script = language.Script(value.String)
			}
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Language.
// This includes values selected through modifiers, order, etc.
func (l *Language) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// QueryRaceSpeakers queries the "race_speakers" edge of the Language entity.
func (l *Language) QueryRaceSpeakers() *RaceQuery {
	return NewLanguageClient(l.config).QueryRaceSpeakers(l)
}

// Update returns a builder for updating this Language.
// Note that you need to call Language.Unwrap() before calling this method if this Language
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Language) Update() *LanguageUpdateOne {
	return NewLanguageClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Language entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Language) Unwrap() *Language {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Language is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Language) String() string {
	var builder strings.Builder
	builder.WriteString("Language(")
	builder.WriteString(fmt.Sprintf("id=%v, ", l.ID))
	builder.WriteString("indx=")
	builder.WriteString(l.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(l.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(l.Desc)
	builder.WriteString(", ")
	builder.WriteString("language_type=")
	builder.WriteString(fmt.Sprintf("%v", l.LanguageType))
	builder.WriteString(", ")
	builder.WriteString("script=")
	builder.WriteString(fmt.Sprintf("%v", l.Script))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (l *Language) MarshalJSON() ([]byte, error) {
	type Alias Language
	return json.Marshal(&struct {
		*Alias
		LanguageEdges
	}{
		Alias:         (*Alias)(l),
		LanguageEdges: l.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (l *Language) UnmarshalJSON(data []byte) error {
	type Alias Language
	aux := &struct {
		*Alias
		LanguageEdges
	}{
		Alias: (*Alias)(l),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	l.Edges = aux.LanguageEdges
	return nil
}

func (lc *LanguageCreate) SetLanguage(input *Language) *LanguageCreate {
	lc.SetIndx(input.Indx)
	lc.SetName(input.Name)
	lc.SetDesc(input.Desc)
	lc.SetLanguageType(input.LanguageType)
	lc.SetScript(input.Script)
	return lc
}

// NamedRaceSpeakers returns the RaceSpeakers named value or an error if the edge was not
// loaded in eager-loading with this name.
func (l *Language) NamedRaceSpeakers(name string) ([]*Race, error) {
	if l.Edges.namedRaceSpeakers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := l.Edges.namedRaceSpeakers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (l *Language) appendNamedRaceSpeakers(name string, edges ...*Race) {
	if l.Edges.namedRaceSpeakers == nil {
		l.Edges.namedRaceSpeakers = make(map[string][]*Race)
	}
	if len(edges) == 0 {
		l.Edges.namedRaceSpeakers[name] = []*Race{}
	} else {
		l.Edges.namedRaceSpeakers[name] = append(l.Edges.namedRaceSpeakers[name], edges...)
	}
}

// Languages is a parsable slice of Language.
type Languages []*Language
