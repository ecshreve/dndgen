// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/magicschool"
)

// MagicSchool is the model entity for the MagicSchool schema.
type MagicSchool struct {
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
func (*MagicSchool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case magicschool.FieldID:
			values[i] = new(sql.NullInt64)
		case magicschool.FieldIndx, magicschool.FieldName, magicschool.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MagicSchool fields.
func (ms *MagicSchool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case magicschool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ms.ID = int(value.Int64)
		case magicschool.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				ms.Indx = value.String
			}
		case magicschool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ms.Name = value.String
			}
		case magicschool.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				ms.Desc = value.String
			}
		default:
			ms.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MagicSchool.
// This includes values selected through modifiers, order, etc.
func (ms *MagicSchool) Value(name string) (ent.Value, error) {
	return ms.selectValues.Get(name)
}

// Update returns a builder for updating this MagicSchool.
// Note that you need to call MagicSchool.Unwrap() before calling this method if this MagicSchool
// was returned from a transaction, and the transaction was committed or rolled back.
func (ms *MagicSchool) Update() *MagicSchoolUpdateOne {
	return NewMagicSchoolClient(ms.config).UpdateOne(ms)
}

// Unwrap unwraps the MagicSchool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ms *MagicSchool) Unwrap() *MagicSchool {
	_tx, ok := ms.config.driver.(*txDriver)
	if !ok {
		panic("ent: MagicSchool is not a transactional entity")
	}
	ms.config.driver = _tx.drv
	return ms
}

// String implements the fmt.Stringer.
func (ms *MagicSchool) String() string {
	var builder strings.Builder
	builder.WriteString("MagicSchool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ms.ID))
	builder.WriteString("indx=")
	builder.WriteString(ms.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ms.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(ms.Desc)
	builder.WriteByte(')')
	return builder.String()
}

// MagicSchools is a parsable slice of MagicSchool.
type MagicSchools []*MagicSchool
