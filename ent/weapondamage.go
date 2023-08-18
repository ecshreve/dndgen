// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// WeaponDamage is the model entity for the WeaponDamage schema.
type WeaponDamage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Dice holds the value of the "dice" field.
	Dice string `json:"dice,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WeaponDamageQuery when eager-loading is set.
	Edges                    WeaponDamageEdges `json:"edges"`
	weapon_two_handed_damage *int
	selectValues             sql.SelectValues
}

// WeaponDamageEdges holds the relations/edges for other nodes in the graph.
type WeaponDamageEdges struct {
	// DamageType holds the value of the damage_type edge.
	DamageType []*DamageType `json:"damage_type,omitempty"`
	// Weapon holds the value of the weapon edge.
	Weapon []*Weapon `json:"weapon,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedDamageType map[string][]*DamageType
	namedWeapon     map[string][]*Weapon
}

// DamageTypeOrErr returns the DamageType value or an error if the edge
// was not loaded in eager-loading.
func (e WeaponDamageEdges) DamageTypeOrErr() ([]*DamageType, error) {
	if e.loadedTypes[0] {
		return e.DamageType, nil
	}
	return nil, &NotLoadedError{edge: "damage_type"}
}

// WeaponOrErr returns the Weapon value or an error if the edge
// was not loaded in eager-loading.
func (e WeaponDamageEdges) WeaponOrErr() ([]*Weapon, error) {
	if e.loadedTypes[1] {
		return e.Weapon, nil
	}
	return nil, &NotLoadedError{edge: "weapon"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WeaponDamage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case weapondamage.FieldID:
			values[i] = new(sql.NullInt64)
		case weapondamage.FieldDice:
			values[i] = new(sql.NullString)
		case weapondamage.ForeignKeys[0]: // weapon_two_handed_damage
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WeaponDamage fields.
func (wd *WeaponDamage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case weapondamage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wd.ID = int(value.Int64)
		case weapondamage.FieldDice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dice", values[i])
			} else if value.Valid {
				wd.Dice = value.String
			}
		case weapondamage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field weapon_two_handed_damage", value)
			} else if value.Valid {
				wd.weapon_two_handed_damage = new(int)
				*wd.weapon_two_handed_damage = int(value.Int64)
			}
		default:
			wd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WeaponDamage.
// This includes values selected through modifiers, order, etc.
func (wd *WeaponDamage) Value(name string) (ent.Value, error) {
	return wd.selectValues.Get(name)
}

// QueryDamageType queries the "damage_type" edge of the WeaponDamage entity.
func (wd *WeaponDamage) QueryDamageType() *DamageTypeQuery {
	return NewWeaponDamageClient(wd.config).QueryDamageType(wd)
}

// QueryWeapon queries the "weapon" edge of the WeaponDamage entity.
func (wd *WeaponDamage) QueryWeapon() *WeaponQuery {
	return NewWeaponDamageClient(wd.config).QueryWeapon(wd)
}

// Update returns a builder for updating this WeaponDamage.
// Note that you need to call WeaponDamage.Unwrap() before calling this method if this WeaponDamage
// was returned from a transaction, and the transaction was committed or rolled back.
func (wd *WeaponDamage) Update() *WeaponDamageUpdateOne {
	return NewWeaponDamageClient(wd.config).UpdateOne(wd)
}

// Unwrap unwraps the WeaponDamage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wd *WeaponDamage) Unwrap() *WeaponDamage {
	_tx, ok := wd.config.driver.(*txDriver)
	if !ok {
		panic("ent: WeaponDamage is not a transactional entity")
	}
	wd.config.driver = _tx.drv
	return wd
}

// String implements the fmt.Stringer.
func (wd *WeaponDamage) String() string {
	var builder strings.Builder
	builder.WriteString("WeaponDamage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wd.ID))
	builder.WriteString("dice=")
	builder.WriteString(wd.Dice)
	builder.WriteByte(')')
	return builder.String()
}

func (wdc *WeaponDamageCreate) SetWeaponDamage(input *WeaponDamage) *WeaponDamageCreate {
	wdc.SetDice(input.Dice)
	return wdc
}

// NamedDamageType returns the DamageType named value or an error if the edge was not
// loaded in eager-loading with this name.
func (wd *WeaponDamage) NamedDamageType(name string) ([]*DamageType, error) {
	if wd.Edges.namedDamageType == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := wd.Edges.namedDamageType[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (wd *WeaponDamage) appendNamedDamageType(name string, edges ...*DamageType) {
	if wd.Edges.namedDamageType == nil {
		wd.Edges.namedDamageType = make(map[string][]*DamageType)
	}
	if len(edges) == 0 {
		wd.Edges.namedDamageType[name] = []*DamageType{}
	} else {
		wd.Edges.namedDamageType[name] = append(wd.Edges.namedDamageType[name], edges...)
	}
}

// NamedWeapon returns the Weapon named value or an error if the edge was not
// loaded in eager-loading with this name.
func (wd *WeaponDamage) NamedWeapon(name string) ([]*Weapon, error) {
	if wd.Edges.namedWeapon == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := wd.Edges.namedWeapon[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (wd *WeaponDamage) appendNamedWeapon(name string, edges ...*Weapon) {
	if wd.Edges.namedWeapon == nil {
		wd.Edges.namedWeapon = make(map[string][]*Weapon)
	}
	if len(edges) == 0 {
		wd.Edges.namedWeapon[name] = []*Weapon{}
	} else {
		wd.Edges.namedWeapon[name] = append(wd.Edges.namedWeapon[name], edges...)
	}
}

// WeaponDamages is a parsable slice of WeaponDamage.
type WeaponDamages []*WeaponDamage
