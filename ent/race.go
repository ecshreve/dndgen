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

// Race is the model entity for the Race schema.
type Race struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Alignment holds the value of the "alignment" field.
	Alignment string `json:"alignment,omitempty"`
	// Age holds the value of the "age" field.
	Age string `json:"age,omitempty"`
	// Size holds the value of the "size" field.
	Size string `json:"size,omitempty"`
	// SizeDescription holds the value of the "size_description" field.
	SizeDescription string `json:"size_description,omitempty"`
	// LanguageDesc holds the value of the "language_desc" field.
	LanguageDesc string `json:"language_desc,omitempty"`
	// Speed holds the value of the "speed" field.
	Speed int `json:"speed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RaceQuery when eager-loading is set.
	Edges        RaceEdges `json:"-"`
	selectValues sql.SelectValues
}

// RaceEdges holds the relations/edges for other nodes in the graph.
type RaceEdges struct {
	// Languages holds the value of the languages edge.
	Languages []*Language `json:"languages,omitempty"`
	// Proficiencies holds the value of the proficiencies edge.
	Proficiencies []*Proficiency `json:"proficiencies,omitempty"`
	// Subraces holds the value of the subraces edge.
	Subraces []*Subrace `json:"subraces,omitempty"`
	// Traits holds the value of the traits edge.
	Traits []*Trait `json:"traits,omitempty"`
	// AbilityBonuses holds the value of the ability_bonuses edge.
	AbilityBonuses []*AbilityBonus `json:"ability_bonuses,omitempty"`
	// StartingProficiencyOption holds the value of the starting_proficiency_option edge.
	StartingProficiencyOption *Choice `json:"starting_proficiency_option,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedLanguages      map[string][]*Language
	namedProficiencies  map[string][]*Proficiency
	namedSubraces       map[string][]*Subrace
	namedTraits         map[string][]*Trait
	namedAbilityBonuses map[string][]*AbilityBonus
}

// LanguagesOrErr returns the Languages value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) LanguagesOrErr() ([]*Language, error) {
	if e.loadedTypes[0] {
		return e.Languages, nil
	}
	return nil, &NotLoadedError{edge: "languages"}
}

// ProficienciesOrErr returns the Proficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) ProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[1] {
		return e.Proficiencies, nil
	}
	return nil, &NotLoadedError{edge: "proficiencies"}
}

// SubracesOrErr returns the Subraces value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) SubracesOrErr() ([]*Subrace, error) {
	if e.loadedTypes[2] {
		return e.Subraces, nil
	}
	return nil, &NotLoadedError{edge: "subraces"}
}

// TraitsOrErr returns the Traits value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) TraitsOrErr() ([]*Trait, error) {
	if e.loadedTypes[3] {
		return e.Traits, nil
	}
	return nil, &NotLoadedError{edge: "traits"}
}

// AbilityBonusesOrErr returns the AbilityBonuses value or an error if the edge
// was not loaded in eager-loading.
func (e RaceEdges) AbilityBonusesOrErr() ([]*AbilityBonus, error) {
	if e.loadedTypes[4] {
		return e.AbilityBonuses, nil
	}
	return nil, &NotLoadedError{edge: "ability_bonuses"}
}

// StartingProficiencyOptionOrErr returns the StartingProficiencyOption value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RaceEdges) StartingProficiencyOptionOrErr() (*Choice, error) {
	if e.loadedTypes[5] {
		if e.StartingProficiencyOption == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: choice.Label}
		}
		return e.StartingProficiencyOption, nil
	}
	return nil, &NotLoadedError{edge: "starting_proficiency_option"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Race) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case race.FieldID, race.FieldSpeed:
			values[i] = new(sql.NullInt64)
		case race.FieldIndx, race.FieldName, race.FieldAlignment, race.FieldAge, race.FieldSize, race.FieldSizeDescription, race.FieldLanguageDesc:
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
		case race.FieldAlignment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alignment", values[i])
			} else if value.Valid {
				r.Alignment = value.String
			}
		case race.FieldAge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				r.Age = value.String
			}
		case race.FieldSize:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				r.Size = value.String
			}
		case race.FieldSizeDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field size_description", values[i])
			} else if value.Valid {
				r.SizeDescription = value.String
			}
		case race.FieldLanguageDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field language_desc", values[i])
			} else if value.Valid {
				r.LanguageDesc = value.String
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

// QueryProficiencies queries the "proficiencies" edge of the Race entity.
func (r *Race) QueryProficiencies() *ProficiencyQuery {
	return NewRaceClient(r.config).QueryProficiencies(r)
}

// QuerySubraces queries the "subraces" edge of the Race entity.
func (r *Race) QuerySubraces() *SubraceQuery {
	return NewRaceClient(r.config).QuerySubraces(r)
}

// QueryTraits queries the "traits" edge of the Race entity.
func (r *Race) QueryTraits() *TraitQuery {
	return NewRaceClient(r.config).QueryTraits(r)
}

// QueryAbilityBonuses queries the "ability_bonuses" edge of the Race entity.
func (r *Race) QueryAbilityBonuses() *AbilityBonusQuery {
	return NewRaceClient(r.config).QueryAbilityBonuses(r)
}

// QueryStartingProficiencyOption queries the "starting_proficiency_option" edge of the Race entity.
func (r *Race) QueryStartingProficiencyOption() *ChoiceQuery {
	return NewRaceClient(r.config).QueryStartingProficiencyOption(r)
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
	builder.WriteString("alignment=")
	builder.WriteString(r.Alignment)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(r.Age)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(r.Size)
	builder.WriteString(", ")
	builder.WriteString("size_description=")
	builder.WriteString(r.SizeDescription)
	builder.WriteString(", ")
	builder.WriteString("language_desc=")
	builder.WriteString(r.LanguageDesc)
	builder.WriteString(", ")
	builder.WriteString("speed=")
	builder.WriteString(fmt.Sprintf("%v", r.Speed))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (r *Race) MarshalJSON() ([]byte, error) {
	type Alias Race
	return json.Marshal(&struct {
		*Alias
		RaceEdges
	}{
		Alias:     (*Alias)(r),
		RaceEdges: r.Edges,
	})
}

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
	rc.SetAlignment(input.Alignment)
	rc.SetAge(input.Age)
	rc.SetSize(input.Size)
	rc.SetSizeDescription(input.SizeDescription)
	rc.SetLanguageDesc(input.LanguageDesc)
	rc.SetSpeed(input.Speed)
	return rc
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

// NamedSubraces returns the Subraces named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Race) NamedSubraces(name string) ([]*Subrace, error) {
	if r.Edges.namedSubraces == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedSubraces[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Race) appendNamedSubraces(name string, edges ...*Subrace) {
	if r.Edges.namedSubraces == nil {
		r.Edges.namedSubraces = make(map[string][]*Subrace)
	}
	if len(edges) == 0 {
		r.Edges.namedSubraces[name] = []*Subrace{}
	} else {
		r.Edges.namedSubraces[name] = append(r.Edges.namedSubraces[name], edges...)
	}
}

// NamedTraits returns the Traits named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Race) NamedTraits(name string) ([]*Trait, error) {
	if r.Edges.namedTraits == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedTraits[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Race) appendNamedTraits(name string, edges ...*Trait) {
	if r.Edges.namedTraits == nil {
		r.Edges.namedTraits = make(map[string][]*Trait)
	}
	if len(edges) == 0 {
		r.Edges.namedTraits[name] = []*Trait{}
	} else {
		r.Edges.namedTraits[name] = append(r.Edges.namedTraits[name], edges...)
	}
}

// NamedAbilityBonuses returns the AbilityBonuses named value or an error if the edge was not
// loaded in eager-loading with this name.
func (r *Race) NamedAbilityBonuses(name string) ([]*AbilityBonus, error) {
	if r.Edges.namedAbilityBonuses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := r.Edges.namedAbilityBonuses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (r *Race) appendNamedAbilityBonuses(name string, edges ...*AbilityBonus) {
	if r.Edges.namedAbilityBonuses == nil {
		r.Edges.namedAbilityBonuses = make(map[string][]*AbilityBonus)
	}
	if len(edges) == 0 {
		r.Edges.namedAbilityBonuses[name] = []*AbilityBonus{}
	} else {
		r.Edges.namedAbilityBonuses[name] = append(r.Edges.namedAbilityBonuses[name], edges...)
	}
}

// Races is a parsable slice of Race.
type Races []*Race
