// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/coin"
)

// Coin is the model entity for the Coin schema.
type Coin struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc []string `json:"desc,omitempty"`
	// GoldConversionRate holds the value of the "gold_conversion_rate" field.
	GoldConversionRate float64 `json:"gold_conversion_rate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CoinQuery when eager-loading is set.
	Edges        CoinEdges `json:"-"`
	selectValues sql.SelectValues
}

// CoinEdges holds the relations/edges for other nodes in the graph.
type CoinEdges struct {
	// EquipmentCosts holds the value of the equipment_costs edge.
	EquipmentCosts []*EquipmentCost `json:"equipment_costs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool

	namedEquipmentCosts map[string][]*EquipmentCost
}

// EquipmentCostsOrErr returns the EquipmentCosts value or an error if the edge
// was not loaded in eager-loading.
func (e CoinEdges) EquipmentCostsOrErr() ([]*EquipmentCost, error) {
	if e.loadedTypes[0] {
		return e.EquipmentCosts, nil
	}
	return nil, &NotLoadedError{edge: "equipment_costs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Coin) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case coin.FieldDesc:
			values[i] = new([]byte)
		case coin.FieldGoldConversionRate:
			values[i] = new(sql.NullFloat64)
		case coin.FieldID:
			values[i] = new(sql.NullInt64)
		case coin.FieldIndx, coin.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Coin fields.
func (c *Coin) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coin.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case coin.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				c.Indx = value.String
			}
		case coin.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case coin.FieldDesc:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Desc); err != nil {
					return fmt.Errorf("unmarshal field desc: %w", err)
				}
			}
		case coin.FieldGoldConversionRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field gold_conversion_rate", values[i])
			} else if value.Valid {
				c.GoldConversionRate = value.Float64
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Coin.
// This includes values selected through modifiers, order, etc.
func (c *Coin) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryEquipmentCosts queries the "equipment_costs" edge of the Coin entity.
func (c *Coin) QueryEquipmentCosts() *EquipmentCostQuery {
	return NewCoinClient(c.config).QueryEquipmentCosts(c)
}

// Update returns a builder for updating this Coin.
// Note that you need to call Coin.Unwrap() before calling this method if this Coin
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Coin) Update() *CoinUpdateOne {
	return NewCoinClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Coin entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Coin) Unwrap() *Coin {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Coin is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Coin) String() string {
	var builder strings.Builder
	builder.WriteString("Coin(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("indx=")
	builder.WriteString(c.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("desc=")
	builder.WriteString(fmt.Sprintf("%v", c.Desc))
	builder.WriteString(", ")
	builder.WriteString("gold_conversion_rate=")
	builder.WriteString(fmt.Sprintf("%v", c.GoldConversionRate))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (c *Coin) MarshalJSON() ([]byte, error) {
	type Alias Coin
	return json.Marshal(&struct {
		*Alias
		CoinEdges
	}{
		Alias:     (*Alias)(c),
		CoinEdges: c.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (c *Coin) UnmarshalJSON(data []byte) error {
	type Alias Coin
	aux := &struct {
		*Alias
		CoinEdges
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	c.Edges = aux.CoinEdges
	return nil
}

func (cc *CoinCreate) SetCoin(input *Coin) *CoinCreate {
	cc.SetIndx(input.Indx)
	cc.SetName(input.Name)
	cc.SetDesc(input.Desc)
	cc.SetGoldConversionRate(input.GoldConversionRate)
	return cc
}

// NamedEquipmentCosts returns the EquipmentCosts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Coin) NamedEquipmentCosts(name string) ([]*EquipmentCost, error) {
	if c.Edges.namedEquipmentCosts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedEquipmentCosts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Coin) appendNamedEquipmentCosts(name string, edges ...*EquipmentCost) {
	if c.Edges.namedEquipmentCosts == nil {
		c.Edges.namedEquipmentCosts = make(map[string][]*EquipmentCost)
	}
	if len(edges) == 0 {
		c.Edges.namedEquipmentCosts[name] = []*EquipmentCost{}
	} else {
		c.Edges.namedEquipmentCosts[name] = append(c.Edges.namedEquipmentCosts[name], edges...)
	}
}

// Coins is a parsable slice of Coin.
type Coins []*Coin
