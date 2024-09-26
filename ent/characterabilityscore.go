// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
)

// CharacterAbilityScore is the model entity for the CharacterAbilityScore schema.
type CharacterAbilityScore struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Score holds the value of the "score" field.
	Score int `json:"score,omitempty"`
	// Modifier holds the value of the "modifier" field.
	Modifier int `json:"modifier,omitempty"`
	// CharacterID holds the value of the "character_id" field.
	CharacterID int `json:"character_id,omitempty"`
	// AbilityScoreID holds the value of the "ability_score_id" field.
	AbilityScoreID int `json:"ability_score_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CharacterAbilityScoreQuery when eager-loading is set.
	Edges        CharacterAbilityScoreEdges `json:"-"`
	selectValues sql.SelectValues
}

// CharacterAbilityScoreEdges holds the relations/edges for other nodes in the graph.
type CharacterAbilityScoreEdges struct {
	// Character holds the value of the character edge.
	Character *Character `json:"character,omitempty"`
	// AbilityScore holds the value of the ability_score edge.
	AbilityScore *AbilityScore `json:"ability_score,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// CharacterOrErr returns the Character value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterAbilityScoreEdges) CharacterOrErr() (*Character, error) {
	if e.Character != nil {
		return e.Character, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: character.Label}
	}
	return nil, &NotLoadedError{edge: "character"}
}

// AbilityScoreOrErr returns the AbilityScore value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CharacterAbilityScoreEdges) AbilityScoreOrErr() (*AbilityScore, error) {
	if e.AbilityScore != nil {
		return e.AbilityScore, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: abilityscore.Label}
	}
	return nil, &NotLoadedError{edge: "ability_score"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CharacterAbilityScore) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case characterabilityscore.FieldID, characterabilityscore.FieldScore, characterabilityscore.FieldModifier, characterabilityscore.FieldCharacterID, characterabilityscore.FieldAbilityScoreID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CharacterAbilityScore fields.
func (cas *CharacterAbilityScore) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case characterabilityscore.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cas.ID = int(value.Int64)
		case characterabilityscore.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				cas.Score = int(value.Int64)
			}
		case characterabilityscore.FieldModifier:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field modifier", values[i])
			} else if value.Valid {
				cas.Modifier = int(value.Int64)
			}
		case characterabilityscore.FieldCharacterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field character_id", values[i])
			} else if value.Valid {
				cas.CharacterID = int(value.Int64)
			}
		case characterabilityscore.FieldAbilityScoreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ability_score_id", values[i])
			} else if value.Valid {
				cas.AbilityScoreID = int(value.Int64)
			}
		default:
			cas.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CharacterAbilityScore.
// This includes values selected through modifiers, order, etc.
func (cas *CharacterAbilityScore) Value(name string) (ent.Value, error) {
	return cas.selectValues.Get(name)
}

// QueryCharacter queries the "character" edge of the CharacterAbilityScore entity.
func (cas *CharacterAbilityScore) QueryCharacter() *CharacterQuery {
	return NewCharacterAbilityScoreClient(cas.config).QueryCharacter(cas)
}

// QueryAbilityScore queries the "ability_score" edge of the CharacterAbilityScore entity.
func (cas *CharacterAbilityScore) QueryAbilityScore() *AbilityScoreQuery {
	return NewCharacterAbilityScoreClient(cas.config).QueryAbilityScore(cas)
}

// Update returns a builder for updating this CharacterAbilityScore.
// Note that you need to call CharacterAbilityScore.Unwrap() before calling this method if this CharacterAbilityScore
// was returned from a transaction, and the transaction was committed or rolled back.
func (cas *CharacterAbilityScore) Update() *CharacterAbilityScoreUpdateOne {
	return NewCharacterAbilityScoreClient(cas.config).UpdateOne(cas)
}

// Unwrap unwraps the CharacterAbilityScore entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cas *CharacterAbilityScore) Unwrap() *CharacterAbilityScore {
	_tx, ok := cas.config.driver.(*txDriver)
	if !ok {
		panic("ent: CharacterAbilityScore is not a transactional entity")
	}
	cas.config.driver = _tx.drv
	return cas
}

// String implements the fmt.Stringer.
func (cas *CharacterAbilityScore) String() string {
	var builder strings.Builder
	builder.WriteString("CharacterAbilityScore(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cas.ID))
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", cas.Score))
	builder.WriteString(", ")
	builder.WriteString("modifier=")
	builder.WriteString(fmt.Sprintf("%v", cas.Modifier))
	builder.WriteString(", ")
	builder.WriteString("character_id=")
	builder.WriteString(fmt.Sprintf("%v", cas.CharacterID))
	builder.WriteString(", ")
	builder.WriteString("ability_score_id=")
	builder.WriteString(fmt.Sprintf("%v", cas.AbilityScoreID))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (cas *CharacterAbilityScore) MarshalJSON() ([]byte, error) {
	type Alias CharacterAbilityScore
	return json.Marshal(&struct {
		*Alias
		CharacterAbilityScoreEdges
	}{
		Alias:                      (*Alias)(cas),
		CharacterAbilityScoreEdges: cas.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (cas *CharacterAbilityScore) UnmarshalJSON(data []byte) error {
	type Alias CharacterAbilityScore
	aux := &struct {
		*Alias
		CharacterAbilityScoreEdges
	}{
		Alias: (*Alias)(cas),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	cas.Edges = aux.CharacterAbilityScoreEdges
	return nil
}

func (casc *CharacterAbilityScoreCreate) SetCharacterAbilityScore(input *CharacterAbilityScore) *CharacterAbilityScoreCreate {
	casc.SetScore(input.Score)
	casc.SetModifier(input.Modifier)
	casc.SetCharacterID(input.CharacterID)
	casc.SetAbilityScoreID(input.AbilityScoreID)
	return casc
}

// CharacterAbilityScores is a parsable slice of CharacterAbilityScore.
type CharacterAbilityScores []*CharacterAbilityScore
