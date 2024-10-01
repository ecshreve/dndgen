// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/race"
)

// Character is the model entity for the Character schema.
type Character struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// ProficiencyBonus holds the value of the "proficiency_bonus" field.
	ProficiencyBonus int `json:"proficiency_bonus,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CharacterQuery when eager-loading is set.
	Edges               CharacterEdges `json:"-"`
	character_race      *int
	character_class     *int
	character_alignment *int
	selectValues        sql.SelectValues
}

// CharacterEdges holds the relations/edges for other nodes in the graph.
type CharacterEdges struct {
	// Race holds the value of the race edge.
	Race *Race `json:"race,omitempty"`
	// Class holds the value of the class edge.
	Class *Class `json:"class,omitempty"`
	// Alignment holds the value of the alignment edge.
	Alignment *Alignment `json:"alignment,omitempty"`
	// CharacterAbilityScores holds the value of the character_ability_scores edge.
	CharacterAbilityScores []*CharacterAbilityScore `json:"character_ability_scores,omitempty"`
	// CharacterSkills holds the value of the character_skills edge.
	CharacterSkills []*CharacterSkill `json:"character_skills,omitempty"`
	// CharacterProficiencies holds the value of the character_proficiencies edge.
	CharacterProficiencies []*CharacterProficiency `json:"character_proficiencies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedCharacterAbilityScores map[string][]*CharacterAbilityScore
	namedCharacterSkills        map[string][]*CharacterSkill
	namedCharacterProficiencies map[string][]*CharacterProficiency
}

// RaceOrErr returns the Race value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterEdges) RaceOrErr() (*Race, error) {
	if e.Race != nil {
		return e.Race, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: race.Label}
	}
	return nil, &NotLoadedError{edge: "race"}
}

// ClassOrErr returns the Class value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterEdges) ClassOrErr() (*Class, error) {
	if e.Class != nil {
		return e.Class, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: class.Label}
	}
	return nil, &NotLoadedError{edge: "class"}
}

// AlignmentOrErr returns the Alignment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterEdges) AlignmentOrErr() (*Alignment, error) {
	if e.Alignment != nil {
		return e.Alignment, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: alignment.Label}
	}
	return nil, &NotLoadedError{edge: "alignment"}
}

// CharacterAbilityScoresOrErr returns the CharacterAbilityScores value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) CharacterAbilityScoresOrErr() ([]*CharacterAbilityScore, error) {
	if e.loadedTypes[3] {
		return e.CharacterAbilityScores, nil
	}
	return nil, &NotLoadedError{edge: "character_ability_scores"}
}

// CharacterSkillsOrErr returns the CharacterSkills value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) CharacterSkillsOrErr() ([]*CharacterSkill, error) {
	if e.loadedTypes[4] {
		return e.CharacterSkills, nil
	}
	return nil, &NotLoadedError{edge: "character_skills"}
}

// CharacterProficienciesOrErr returns the CharacterProficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e CharacterEdges) CharacterProficienciesOrErr() ([]*CharacterProficiency, error) {
	if e.loadedTypes[5] {
		return e.CharacterProficiencies, nil
	}
	return nil, &NotLoadedError{edge: "character_proficiencies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Character) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case character.FieldID, character.FieldAge, character.FieldLevel, character.FieldProficiencyBonus:
			values[i] = new(sql.NullInt64)
		case character.FieldName:
			values[i] = new(sql.NullString)
		case character.ForeignKeys[0]: // character_race
			values[i] = new(sql.NullInt64)
		case character.ForeignKeys[1]: // character_class
			values[i] = new(sql.NullInt64)
		case character.ForeignKeys[2]: // character_alignment
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Character fields.
func (c *Character) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case character.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case character.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case character.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				c.Age = int(value.Int64)
			}
		case character.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				c.Level = int(value.Int64)
			}
		case character.FieldProficiencyBonus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field proficiency_bonus", values[i])
			} else if value.Valid {
				c.ProficiencyBonus = int(value.Int64)
			}
		case character.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field character_race", value)
			} else if value.Valid {
				c.character_race = new(int)
				*c.character_race = int(value.Int64)
			}
		case character.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field character_class", value)
			} else if value.Valid {
				c.character_class = new(int)
				*c.character_class = int(value.Int64)
			}
		case character.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field character_alignment", value)
			} else if value.Valid {
				c.character_alignment = new(int)
				*c.character_alignment = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Character.
// This includes values selected through modifiers, order, etc.
func (c *Character) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryRace queries the "race" edge of the Character entity.
func (c *Character) QueryRace() *RaceQuery {
	return NewCharacterClient(c.config).QueryRace(c)
}

// QueryClass queries the "class" edge of the Character entity.
func (c *Character) QueryClass() *ClassQuery {
	return NewCharacterClient(c.config).QueryClass(c)
}

// QueryAlignment queries the "alignment" edge of the Character entity.
func (c *Character) QueryAlignment() *AlignmentQuery {
	return NewCharacterClient(c.config).QueryAlignment(c)
}

// QueryCharacterAbilityScores queries the "character_ability_scores" edge of the Character entity.
func (c *Character) QueryCharacterAbilityScores() *CharacterAbilityScoreQuery {
	return NewCharacterClient(c.config).QueryCharacterAbilityScores(c)
}

// QueryCharacterSkills queries the "character_skills" edge of the Character entity.
func (c *Character) QueryCharacterSkills() *CharacterSkillQuery {
	return NewCharacterClient(c.config).QueryCharacterSkills(c)
}

// QueryCharacterProficiencies queries the "character_proficiencies" edge of the Character entity.
func (c *Character) QueryCharacterProficiencies() *CharacterProficiencyQuery {
	return NewCharacterClient(c.config).QueryCharacterProficiencies(c)
}

// Update returns a builder for updating this Character.
// Note that you need to call Character.Unwrap() before calling this method if this Character
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Character) Update() *CharacterUpdateOne {
	return NewCharacterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Character entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Character) Unwrap() *Character {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Character is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Character) String() string {
	var builder strings.Builder
	builder.WriteString("Character(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", c.Age))
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", c.Level))
	builder.WriteString(", ")
	builder.WriteString("proficiency_bonus=")
	builder.WriteString(fmt.Sprintf("%v", c.ProficiencyBonus))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (c *Character) MarshalJSON() ([]byte, error) {
	type Alias Character
	return json.Marshal(&struct {
		*Alias
		CharacterEdges
	}{
		Alias:          (*Alias)(c),
		CharacterEdges: c.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Character) UnmarshalJSON(data []byte) error {
	type Alias Character
	aux := &struct {
		*Alias
		CharacterEdges
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	c.Edges = aux.CharacterEdges
	return nil
}

func (cc *CharacterCreate) SetCharacter(input *Character) *CharacterCreate {
	cc.SetName(input.Name)
	cc.SetAge(input.Age)
	cc.SetLevel(input.Level)
	cc.SetProficiencyBonus(input.ProficiencyBonus)
	return cc
}

// NamedCharacterAbilityScores returns the CharacterAbilityScores named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Character) NamedCharacterAbilityScores(name string) ([]*CharacterAbilityScore, error) {
	if c.Edges.namedCharacterAbilityScores == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCharacterAbilityScores[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Character) appendNamedCharacterAbilityScores(name string, edges ...*CharacterAbilityScore) {
	if c.Edges.namedCharacterAbilityScores == nil {
		c.Edges.namedCharacterAbilityScores = make(map[string][]*CharacterAbilityScore)
	}
	if len(edges) == 0 {
		c.Edges.namedCharacterAbilityScores[name] = []*CharacterAbilityScore{}
	} else {
		c.Edges.namedCharacterAbilityScores[name] = append(c.Edges.namedCharacterAbilityScores[name], edges...)
	}
}

// NamedCharacterSkills returns the CharacterSkills named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Character) NamedCharacterSkills(name string) ([]*CharacterSkill, error) {
	if c.Edges.namedCharacterSkills == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCharacterSkills[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Character) appendNamedCharacterSkills(name string, edges ...*CharacterSkill) {
	if c.Edges.namedCharacterSkills == nil {
		c.Edges.namedCharacterSkills = make(map[string][]*CharacterSkill)
	}
	if len(edges) == 0 {
		c.Edges.namedCharacterSkills[name] = []*CharacterSkill{}
	} else {
		c.Edges.namedCharacterSkills[name] = append(c.Edges.namedCharacterSkills[name], edges...)
	}
}

// NamedCharacterProficiencies returns the CharacterProficiencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Character) NamedCharacterProficiencies(name string) ([]*CharacterProficiency, error) {
	if c.Edges.namedCharacterProficiencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedCharacterProficiencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Character) appendNamedCharacterProficiencies(name string, edges ...*CharacterProficiency) {
	if c.Edges.namedCharacterProficiencies == nil {
		c.Edges.namedCharacterProficiencies = make(map[string][]*CharacterProficiency)
	}
	if len(edges) == 0 {
		c.Edges.namedCharacterProficiencies[name] = []*CharacterProficiency{}
	} else {
		c.Edges.namedCharacterProficiencies[name] = append(c.Edges.namedCharacterProficiencies[name], edges...)
	}
}

// Characters is a parsable slice of Character.
type Characters []*Character
