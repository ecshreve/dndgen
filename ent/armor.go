// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/armorclass"
	"github.com/ecshreve/dndgen/ent/equipment"
)

// Armor is the model entity for the Armor schema.
type Armor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ArmorCategory holds the value of the "armor_category" field.
	ArmorCategory armor.ArmorCategory `json:"armor_category,omitempty"`
	// StrMinimum holds the value of the "str_minimum" field.
	StrMinimum int `json:"str_minimum,omitempty"`
	// StealthDisadvantage holds the value of the "stealth_disadvantage" field.
	StealthDisadvantage bool `json:"stealth_disadvantage,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArmorQuery when eager-loading is set.
	Edges           ArmorEdges `json:"-"`
	equipment_armor *int
	selectValues    sql.SelectValues
}

// ArmorEdges holds the relations/edges for other nodes in the graph.
type ArmorEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// ArmorClass holds the value of the armor_class edge.
	ArmorClass *ArmorClass `json:"armor_class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArmorEdges) EquipmentOrErr() (*Equipment, error) {
	if e.Equipment != nil {
		return e.Equipment, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: equipment.Label}
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// ArmorClassOrErr returns the ArmorClass value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArmorEdges) ArmorClassOrErr() (*ArmorClass, error) {
	if e.ArmorClass != nil {
		return e.ArmorClass, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: armorclass.Label}
	}
	return nil, &NotLoadedError{edge: "armor_class"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Armor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case armor.FieldStealthDisadvantage:
			values[i] = new(sql.NullBool)
		case armor.FieldID, armor.FieldStrMinimum:
			values[i] = new(sql.NullInt64)
		case armor.FieldArmorCategory:
			values[i] = new(sql.NullString)
		case armor.ForeignKeys[0]: // equipment_armor
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Armor fields.
func (a *Armor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case armor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case armor.FieldArmorCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field armor_category", values[i])
			} else if value.Valid {
				a.ArmorCategory = armor.ArmorCategory(value.String)
			}
		case armor.FieldStrMinimum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field str_minimum", values[i])
			} else if value.Valid {
				a.StrMinimum = int(value.Int64)
			}
		case armor.FieldStealthDisadvantage:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field stealth_disadvantage", values[i])
			} else if value.Valid {
				a.StealthDisadvantage = value.Bool
			}
		case armor.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field equipment_armor", value)
			} else if value.Valid {
				a.equipment_armor = new(int)
				*a.equipment_armor = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Armor.
// This includes values selected through modifiers, order, etc.
func (a *Armor) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryEquipment queries the "equipment" edge of the Armor entity.
func (a *Armor) QueryEquipment() *EquipmentQuery {
	return NewArmorClient(a.config).QueryEquipment(a)
}

// QueryArmorClass queries the "armor_class" edge of the Armor entity.
func (a *Armor) QueryArmorClass() *ArmorClassQuery {
	return NewArmorClient(a.config).QueryArmorClass(a)
}

// Update returns a builder for updating this Armor.
// Note that you need to call Armor.Unwrap() before calling this method if this Armor
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Armor) Update() *ArmorUpdateOne {
	return NewArmorClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Armor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Armor) Unwrap() *Armor {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Armor is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Armor) String() string {
	var builder strings.Builder
	builder.WriteString("Armor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("armor_category=")
	builder.WriteString(fmt.Sprintf("%v", a.ArmorCategory))
	builder.WriteString(", ")
	builder.WriteString("str_minimum=")
	builder.WriteString(fmt.Sprintf("%v", a.StrMinimum))
	builder.WriteString(", ")
	builder.WriteString("stealth_disadvantage=")
	builder.WriteString(fmt.Sprintf("%v", a.StealthDisadvantage))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (a *Armor) MarshalJSON() ([]byte, error) {
	type Alias Armor
	return json.Marshal(&struct {
		*Alias
		ArmorEdges
	}{
		Alias:      (*Alias)(a),
		ArmorEdges: a.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Armor) UnmarshalJSON(data []byte) error {
	type Alias Armor
	aux := &struct {
		*Alias
		ArmorEdges
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	a.Edges = aux.ArmorEdges
	return nil
}

func (ac *ArmorCreate) SetArmor(input *Armor) *ArmorCreate {
	ac.SetArmorCategory(input.ArmorCategory)
	ac.SetStrMinimum(input.StrMinimum)
	ac.SetStealthDisadvantage(input.StealthDisadvantage)
	return ac
}

// Armors is a parsable slice of Armor.
type Armors []*Armor
