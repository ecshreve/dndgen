// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
)

// AbilityBonus is the model entity for the AbilityBonus schema.
type AbilityBonus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AbilityBonusQuery when eager-loading is set.
	Edges        AbilityBonusEdges `json:"-"`
	selectValues sql.SelectValues
}

// AbilityBonusEdges holds the relations/edges for other nodes in the graph.
type AbilityBonusEdges struct {
	// AbilityScore holds the value of the ability_score edge.
	AbilityScore []*AbilityScore `json:"ability_score,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedAbilityScore map[string][]*AbilityScore
}

// AbilityScoreOrErr returns the AbilityScore value or an error if the edge
// was not loaded in eager-loading.
func (e AbilityBonusEdges) AbilityScoreOrErr() ([]*AbilityScore, error) {
	if e.loadedTypes[0] {
		return e.AbilityScore, nil
	}
	return nil, &NotLoadedError{edge: "ability_score"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AbilityBonus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case abilitybonus.FieldID, abilitybonus.FieldValue:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AbilityBonus fields.
func (ab *AbilityBonus) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case abilitybonus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ab.ID = int(value.Int64)
		case abilitybonus.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				ab.Value = int(value.Int64)
			}
		default:
			ab.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the AbilityBonus.
// This includes values selected through modifiers, order, etc.
func (ab *AbilityBonus) GetValue(name string) (ent.Value, error) {
	return ab.selectValues.Get(name)
}

// QueryAbilityScore queries the "ability_score" edge of the AbilityBonus entity.
func (ab *AbilityBonus) QueryAbilityScore() *AbilityScoreQuery {
	return NewAbilityBonusClient(ab.config).QueryAbilityScore(ab)
}

// Update returns a builder for updating this AbilityBonus.
// Note that you need to call AbilityBonus.Unwrap() before calling this method if this AbilityBonus
// was returned from a transaction, and the transaction was committed or rolled back.
func (ab *AbilityBonus) Update() *AbilityBonusUpdateOne {
	return NewAbilityBonusClient(ab.config).UpdateOne(ab)
}

// Unwrap unwraps the AbilityBonus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ab *AbilityBonus) Unwrap() *AbilityBonus {
	_tx, ok := ab.config.driver.(*txDriver)
	if !ok {
		panic("ent: AbilityBonus is not a transactional entity")
	}
	ab.config.driver = _tx.drv
	return ab
}

// String implements the fmt.Stringer.
func (ab *AbilityBonus) String() string {
	var builder strings.Builder
	builder.WriteString("AbilityBonus(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ab.ID))
	builder.WriteString("value=")
	builder.WriteString(fmt.Sprintf("%v", ab.Value))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
// func (ab *AbilityBonus) MarshalJSON() ([]byte, error) {
// 		type Alias AbilityBonus
// 		return json.Marshal(&struct {
// 				*Alias
// 				AbilityBonusEdges
// 		}{
// 				Alias: (*Alias)(ab),
// 				AbilityBonusEdges: ab.Edges,
// 		})
// }

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ab *AbilityBonus) UnmarshalJSON(data []byte) error {
	type Alias AbilityBonus
	aux := &struct {
		*Alias
		AbilityBonusEdges
	}{
		Alias: (*Alias)(ab),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ab.Edges = aux.AbilityBonusEdges
	return nil
}

func (abc *AbilityBonusCreate) SetAbilityBonus(input *AbilityBonus) *AbilityBonusCreate {
	abc.SetValue(input.Value)
	return abc
}

// NamedAbilityScore returns the AbilityScore named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ab *AbilityBonus) NamedAbilityScore(name string) ([]*AbilityScore, error) {
	if ab.Edges.namedAbilityScore == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ab.Edges.namedAbilityScore[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ab *AbilityBonus) appendNamedAbilityScore(name string, edges ...*AbilityScore) {
	if ab.Edges.namedAbilityScore == nil {
		ab.Edges.namedAbilityScore = make(map[string][]*AbilityScore)
	}
	if len(edges) == 0 {
		ab.Edges.namedAbilityScore[name] = []*AbilityScore{}
	} else {
		ab.Edges.namedAbilityScore[name] = append(ab.Edges.namedAbilityScore[name], edges...)
	}
}

// AbilityBonusSlice is a parsable slice of AbilityBonus.
type AbilityBonusSlice []*AbilityBonus
