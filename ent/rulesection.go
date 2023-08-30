// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/rulesection"
)

// RuleSection is the model entity for the RuleSection schema.
type RuleSection struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc         string `json:"desc,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RuleSection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rulesection.FieldID:
			values[i] = new(sql.NullInt64)
		case rulesection.FieldIndx, rulesection.FieldName, rulesection.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RuleSection fields.
func (rs *RuleSection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rulesection.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rs.ID = int(value.Int64)
		case rulesection.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				rs.Indx = value.String
			}
		case rulesection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rs.Name = value.String
			}
		case rulesection.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				rs.Desc = value.String
			}
		default:
			rs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RuleSection.
// This includes values selected through modifiers, order, etc.
func (rs *RuleSection) Value(name string) (ent.Value, error) {
	return rs.selectValues.Get(name)
}

// Update returns a builder for updating this RuleSection.
// Note that you need to call RuleSection.Unwrap() before calling this method if this RuleSection
// was returned from a transaction, and the transaction was committed or rolled back.
func (rs *RuleSection) Update() *RuleSectionUpdateOne {
	return NewRuleSectionClient(rs.config).UpdateOne(rs)
}

// Unwrap unwraps the RuleSection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rs *RuleSection) Unwrap() *RuleSection {
	_tx, ok := rs.config.driver.(*txDriver)
	if !ok {
		panic("ent: RuleSection is not a transactional entity")
	}
	rs.config.driver = _tx.drv
	return rs
}

// String implements the fmt.Stringer.
func (rs *RuleSection) String() string {
	var builder strings.Builder
	builder.WriteString("RuleSection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rs.ID))
	builder.WriteString("indx=")
	builder.WriteString(rs.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(rs.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(rs.Desc)
	builder.WriteByte(')')
	return builder.String()
}

func (rsc *RuleSectionCreate) SetRuleSection(input *RuleSection) *RuleSectionCreate {
	rsc.SetIndx(input.Indx)
	rsc.SetName(input.Name)
	rsc.SetDesc(input.Desc)
	return rsc
}

// RuleSections is a parsable slice of RuleSection.
type RuleSections []*RuleSection
