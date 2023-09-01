// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/skill"
)

// Skill is the model entity for the Skill schema.
type Skill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc []string `json:"desc,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SkillQuery when eager-loading is set.
	Edges               SkillEdges `json:"-"`
	skill_ability_score *int
	selectValues        sql.SelectValues
}

// SkillEdges holds the relations/edges for other nodes in the graph.
type SkillEdges struct {
	// AbilityScore holds the value of the ability_score edge.
	AbilityScore *AbilityScore `json:"ability_score,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// AbilityScoreOrErr returns the AbilityScore value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SkillEdges) AbilityScoreOrErr() (*AbilityScore, error) {
	if e.loadedTypes[0] {
		if e.AbilityScore == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: abilityscore.Label}
		}
		return e.AbilityScore, nil
	}
	return nil, &NotLoadedError{edge: "ability_score"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Skill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case skill.FieldDesc:
			values[i] = new([]byte)
		case skill.FieldID:
			values[i] = new(sql.NullInt64)
		case skill.FieldIndx, skill.FieldName:
			values[i] = new(sql.NullString)
		case skill.ForeignKeys[0]: // skill_ability_score
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Skill fields.
func (s *Skill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case skill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case skill.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				s.Indx = value.String
			}
		case skill.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case skill.FieldDesc:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.Desc); err != nil {
					return fmt.Errorf("unmarshal field desc: %w", err)
				}
			}
		case skill.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field skill_ability_score", value)
			} else if value.Valid {
				s.skill_ability_score = new(int)
				*s.skill_ability_score = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Skill.
// This includes values selected through modifiers, order, etc.
func (s *Skill) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryAbilityScore queries the "ability_score" edge of the Skill entity.
func (s *Skill) QueryAbilityScore() *AbilityScoreQuery {
	return NewSkillClient(s.config).QueryAbilityScore(s)
}

// Update returns a builder for updating this Skill.
// Note that you need to call Skill.Unwrap() before calling this method if this Skill
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Skill) Update() *SkillUpdateOne {
	return NewSkillClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Skill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Skill) Unwrap() *Skill {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Skill is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Skill) String() string {
	var builder strings.Builder
	builder.WriteString("Skill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("indx=")
	builder.WriteString(s.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(fmt.Sprintf("%v", s.Desc))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (s *Skill) MarshalJSON() ([]byte, error) {
	type Alias Skill
	return json.Marshal(&struct {
		*Alias
		SkillEdges
	}{
		Alias:      (*Alias)(s),
		SkillEdges: s.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Skill) UnmarshalJSON(data []byte) error {
	type Alias Skill
	aux := &struct {
		*Alias
		SkillEdges
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	s.Edges = aux.SkillEdges
	return nil
}

func (sc *SkillCreate) SetSkill(input *Skill) *SkillCreate {
	sc.SetIndx(input.Indx)
	sc.SetName(input.Name)
	sc.SetDesc(input.Desc)
	return sc
}

// Skills is a parsable slice of Skill.
type Skills []*Skill
