// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/abilityscore"
)

// AbilityScore is the model entity for the AbilityScore schema.
type AbilityScore struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AbilityScoreQuery when eager-loading is set.
	Edges                       AbilityScoreEdges `json:"edges"`
	ability_bonus_ability_score *int
	class_saving_throws         *int
	prerequisite_ability_score  *int
	selectValues                sql.SelectValues
}

// AbilityScoreEdges holds the relations/edges for other nodes in the graph.
type AbilityScoreEdges struct {
	// Skills holds the value of the skills edge.
	Skills []*Skill `json:"skills,omitempty"`
	// Proficiencies holds the value of the proficiencies edge.
	Proficiencies []*Proficiency `json:"proficiencies,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedSkills        map[string][]*Skill
	namedProficiencies map[string][]*Proficiency
}

// SkillsOrErr returns the Skills value or an error if the edge
// was not loaded in eager-loading.
func (e AbilityScoreEdges) SkillsOrErr() ([]*Skill, error) {
	if e.loadedTypes[0] {
		return e.Skills, nil
	}
	return nil, &NotLoadedError{edge: "skills"}
}

// ProficienciesOrErr returns the Proficiencies value or an error if the edge
// was not loaded in eager-loading.
func (e AbilityScoreEdges) ProficienciesOrErr() ([]*Proficiency, error) {
	if e.loadedTypes[1] {
		return e.Proficiencies, nil
	}
	return nil, &NotLoadedError{edge: "proficiencies"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AbilityScore) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case abilityscore.FieldID:
			values[i] = new(sql.NullInt64)
		case abilityscore.FieldIndx, abilityscore.FieldName, abilityscore.FieldDesc, abilityscore.FieldFullName:
			values[i] = new(sql.NullString)
		case abilityscore.ForeignKeys[0]: // ability_bonus_ability_score
			values[i] = new(sql.NullInt64)
		case abilityscore.ForeignKeys[1]: // class_saving_throws
			values[i] = new(sql.NullInt64)
		case abilityscore.ForeignKeys[2]: // prerequisite_ability_score
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AbilityScore fields.
func (as *AbilityScore) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case abilityscore.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = int(value.Int64)
		case abilityscore.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				as.Indx = value.String
			}
		case abilityscore.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				as.Name = value.String
			}
		case abilityscore.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				as.Desc = value.String
			}
		case abilityscore.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				as.FullName = value.String
			}
		case abilityscore.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field ability_bonus_ability_score", value)
			} else if value.Valid {
				as.ability_bonus_ability_score = new(int)
				*as.ability_bonus_ability_score = int(value.Int64)
			}
		case abilityscore.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field class_saving_throws", value)
			} else if value.Valid {
				as.class_saving_throws = new(int)
				*as.class_saving_throws = int(value.Int64)
			}
		case abilityscore.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field prerequisite_ability_score", value)
			} else if value.Valid {
				as.prerequisite_ability_score = new(int)
				*as.prerequisite_ability_score = int(value.Int64)
			}
		default:
			as.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the AbilityScore.
// This includes values selected through modifiers, order, etc.
func (as *AbilityScore) Value(name string) (ent.Value, error) {
	return as.selectValues.Get(name)
}

// QuerySkills queries the "skills" edge of the AbilityScore entity.
func (as *AbilityScore) QuerySkills() *SkillQuery {
	return NewAbilityScoreClient(as.config).QuerySkills(as)
}

// QueryProficiencies queries the "proficiencies" edge of the AbilityScore entity.
func (as *AbilityScore) QueryProficiencies() *ProficiencyQuery {
	return NewAbilityScoreClient(as.config).QueryProficiencies(as)
}

// Update returns a builder for updating this AbilityScore.
// Note that you need to call AbilityScore.Unwrap() before calling this method if this AbilityScore
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AbilityScore) Update() *AbilityScoreUpdateOne {
	return NewAbilityScoreClient(as.config).UpdateOne(as)
}

// Unwrap unwraps the AbilityScore entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AbilityScore) Unwrap() *AbilityScore {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: AbilityScore is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AbilityScore) String() string {
	var builder strings.Builder
	builder.WriteString("AbilityScore(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("indx=")
	builder.WriteString(as.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(as.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(as.Desc)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(as.FullName)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (as *AbilityScore) MarshalJSON() ([]byte, error) {
	type Alias AbilityScore
	return json.Marshal(&struct {
		*Alias
		AbilityScoreEdges
	}{
		Alias:             (*Alias)(as),
		AbilityScoreEdges: as.Edges,
	})
}

// NamedSkills returns the Skills named value or an error if the edge was not
// loaded in eager-loading with this name.
func (as *AbilityScore) NamedSkills(name string) ([]*Skill, error) {
	if as.Edges.namedSkills == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := as.Edges.namedSkills[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (as *AbilityScore) appendNamedSkills(name string, edges ...*Skill) {
	if as.Edges.namedSkills == nil {
		as.Edges.namedSkills = make(map[string][]*Skill)
	}
	if len(edges) == 0 {
		as.Edges.namedSkills[name] = []*Skill{}
	} else {
		as.Edges.namedSkills[name] = append(as.Edges.namedSkills[name], edges...)
	}
}

// NamedProficiencies returns the Proficiencies named value or an error if the edge was not
// loaded in eager-loading with this name.
func (as *AbilityScore) NamedProficiencies(name string) ([]*Proficiency, error) {
	if as.Edges.namedProficiencies == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := as.Edges.namedProficiencies[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (as *AbilityScore) appendNamedProficiencies(name string, edges ...*Proficiency) {
	if as.Edges.namedProficiencies == nil {
		as.Edges.namedProficiencies = make(map[string][]*Proficiency)
	}
	if len(edges) == 0 {
		as.Edges.namedProficiencies[name] = []*Proficiency{}
	} else {
		as.Edges.namedProficiencies[name] = append(as.Edges.namedProficiencies[name], edges...)
	}
}

// AbilityScores is a parsable slice of AbilityScore.
type AbilityScores []*AbilityScore
