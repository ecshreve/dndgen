// Code generated by ent, DO NOT EDIT.

package ent

import (
	"builder/ent/abilityscore"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
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
	// Abbr holds the value of the "abbr" field.
	Abbr abilityscore.Abbr `json:"abbr,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc         []string `json:"desc,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AbilityScore) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case abilityscore.FieldDesc:
			values[i] = new([]byte)
		case abilityscore.FieldID:
			values[i] = new(sql.NullInt64)
		case abilityscore.FieldIndx, abilityscore.FieldName, abilityscore.FieldAbbr:
			values[i] = new(sql.NullString)
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
		case abilityscore.FieldAbbr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abbr", values[i])
			} else if value.Valid {
				as.Abbr = abilityscore.Abbr(value.String)
			}
		case abilityscore.FieldDesc:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &as.Desc); err != nil {
					return fmt.Errorf("unmarshal field desc: %w", err)
				}
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
	builder.WriteString("abbr=")
	builder.WriteString(fmt.Sprintf("%v", as.Abbr))
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(fmt.Sprintf("%v", as.Desc))
	builder.WriteByte(')')
	return builder.String()
}

func (asc *AbilityScoreCreate) SetAbilityScore(input *AbilityScore) *AbilityScoreCreate {
	asc.SetIndx(input.Indx)
	asc.SetName(input.Name)
	asc.SetAbbr(input.Abbr)
	asc.SetDesc(input.Desc)
	return asc
}

// AbilityScores is a parsable slice of AbilityScore.
type AbilityScores []*AbilityScore