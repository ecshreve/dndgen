// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/feat"
)

// Feat is the model entity for the Feat schema.
type Feat struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc         []string `json:"desc,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Feat) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case feat.FieldDesc:
			values[i] = new([]byte)
		case feat.FieldID:
			values[i] = new(sql.NullInt64)
		case feat.FieldIndx, feat.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Feat fields.
func (f *Feat) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case feat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case feat.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				f.Indx = value.String
			}
		case feat.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case feat.FieldDesc:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &f.Desc); err != nil {
					return fmt.Errorf("unmarshal field desc: %w", err)
				}
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Feat.
// This includes values selected through modifiers, order, etc.
func (f *Feat) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Feat.
// Note that you need to call Feat.Unwrap() before calling this method if this Feat
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Feat) Update() *FeatUpdateOne {
	return NewFeatClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Feat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Feat) Unwrap() *Feat {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Feat is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Feat) String() string {
	var builder strings.Builder
	builder.WriteString("Feat(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("indx=")
	builder.WriteString(f.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(fmt.Sprintf("%v", f.Desc))
	builder.WriteByte(')')
	return builder.String()
}

func (fc *FeatCreate) SetFeat(input *Feat) *FeatCreate {
	fc.SetIndx(input.Indx)
	fc.SetName(input.Name)
	fc.SetDesc(input.Desc)
	return fc
}

// Feats is a parsable slice of Feat.
type Feats []*Feat
