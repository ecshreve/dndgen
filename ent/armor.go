// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/equipment"
)

// Armor is the model entity for the Armor schema.
type Armor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// StealthDisadvantage holds the value of the "stealth_disadvantage" field.
	StealthDisadvantage bool `json:"stealth_disadvantage,omitempty"`
	// MinStrength holds the value of the "min_strength" field.
	MinStrength int `json:"min_strength,omitempty"`
	// EquipmentID holds the value of the "equipment_id" field.
	EquipmentID int `json:"equipment_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ArmorQuery when eager-loading is set.
	Edges        ArmorEdges `json:"-"`
	selectValues sql.SelectValues
}

// ArmorEdges holds the relations/edges for other nodes in the graph.
type ArmorEdges struct {
	// Equipment holds the value of the equipment edge.
	Equipment *Equipment `json:"equipment,omitempty"`
	// ArmorClass holds the value of the armor_class edge.
	ArmorClass []*ArmorClass `json:"armor_class,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedArmorClass map[string][]*ArmorClass
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ArmorEdges) EquipmentOrErr() (*Equipment, error) {
	if e.loadedTypes[0] {
		if e.Equipment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipment.Label}
		}
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// ArmorClassOrErr returns the ArmorClass value or an error if the edge
// was not loaded in eager-loading.
func (e ArmorEdges) ArmorClassOrErr() ([]*ArmorClass, error) {
	if e.loadedTypes[1] {
		return e.ArmorClass, nil
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
		case armor.FieldID, armor.FieldMinStrength, armor.FieldEquipmentID:
			values[i] = new(sql.NullInt64)
		case armor.FieldIndx, armor.FieldName:
			values[i] = new(sql.NullString)
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
		case armor.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				a.Indx = value.String
			}
		case armor.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case armor.FieldStealthDisadvantage:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field stealth_disadvantage", values[i])
			} else if value.Valid {
				a.StealthDisadvantage = value.Bool
			}
		case armor.FieldMinStrength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_strength", values[i])
			} else if value.Valid {
				a.MinStrength = int(value.Int64)
			}
		case armor.FieldEquipmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field equipment_id", values[i])
			} else if value.Valid {
				a.EquipmentID = int(value.Int64)
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
	builder.WriteString("indx=")
	builder.WriteString(a.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("stealth_disadvantage=")
	builder.WriteString(fmt.Sprintf("%v", a.StealthDisadvantage))
	builder.WriteString(", ")
	builder.WriteString("min_strength=")
	builder.WriteString(fmt.Sprintf("%v", a.MinStrength))
	builder.WriteString(", ")
	builder.WriteString("equipment_id=")
	builder.WriteString(fmt.Sprintf("%v", a.EquipmentID))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
// func (a *Armor) MarshalJSON() ([]byte, error) {
// 		type Alias Armor
// 		return json.Marshal(&struct {
// 				*Alias
// 				ArmorEdges
// 		}{
// 				Alias: (*Alias)(a),
// 				ArmorEdges: a.Edges,
// 		})
// }

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
	ac.SetIndx(input.Indx)
	ac.SetName(input.Name)
	ac.SetStealthDisadvantage(input.StealthDisadvantage)
	ac.SetMinStrength(input.MinStrength)
	ac.SetEquipmentID(input.EquipmentID)
	return ac
}

// NamedArmorClass returns the ArmorClass named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Armor) NamedArmorClass(name string) ([]*ArmorClass, error) {
	if a.Edges.namedArmorClass == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedArmorClass[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Armor) appendNamedArmorClass(name string, edges ...*ArmorClass) {
	if a.Edges.namedArmorClass == nil {
		a.Edges.namedArmorClass = make(map[string][]*ArmorClass)
	}
	if len(edges) == 0 {
		a.Edges.namedArmorClass[name] = []*ArmorClass{}
	} else {
		a.Edges.namedArmorClass[name] = append(a.Edges.namedArmorClass[name], edges...)
	}
}

// Armors is a parsable slice of Armor.
type Armors []*Armor
