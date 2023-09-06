// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/equipmentcategory"
)

// EquipmentCategory is the model entity for the EquipmentCategory schema.
type EquipmentCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ParentCategoryID holds the value of the "parent_category_id" field.
	ParentCategoryID int `json:"parent_category_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EquipmentCategoryQuery when eager-loading is set.
	Edges                          EquipmentCategoryEdges `json:"-"`
	proficiency_equipment_category *int
	selectValues                   sql.SelectValues
}

// EquipmentCategoryEdges holds the relations/edges for other nodes in the graph.
type EquipmentCategoryEdges struct {
	// Parent holds the value of the parent edge.
	Parent *EquipmentCategory `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*EquipmentCategory `json:"children,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `json:"equipment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedChildren  map[string][]*EquipmentCategory
	namedEquipment map[string][]*Equipment
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EquipmentCategoryEdges) ParentOrErr() (*EquipmentCategory, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: equipmentcategory.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentCategoryEdges) ChildrenOrErr() ([]*EquipmentCategory, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e EquipmentCategoryEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[2] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EquipmentCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case equipmentcategory.FieldID, equipmentcategory.FieldParentCategoryID:
			values[i] = new(sql.NullInt64)
		case equipmentcategory.FieldName:
			values[i] = new(sql.NullString)
		case equipmentcategory.ForeignKeys[0]: // proficiency_equipment_category
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EquipmentCategory fields.
func (ec *EquipmentCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case equipmentcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case equipmentcategory.FieldParentCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_category_id", values[i])
			} else if value.Valid {
				ec.ParentCategoryID = int(value.Int64)
			}
		case equipmentcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ec.Name = value.String
			}
		case equipmentcategory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field proficiency_equipment_category", value)
			} else if value.Valid {
				ec.proficiency_equipment_category = new(int)
				*ec.proficiency_equipment_category = int(value.Int64)
			}
		default:
			ec.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EquipmentCategory.
// This includes values selected through modifiers, order, etc.
func (ec *EquipmentCategory) Value(name string) (ent.Value, error) {
	return ec.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the EquipmentCategory entity.
func (ec *EquipmentCategory) QueryParent() *EquipmentCategoryQuery {
	return NewEquipmentCategoryClient(ec.config).QueryParent(ec)
}

// QueryChildren queries the "children" edge of the EquipmentCategory entity.
func (ec *EquipmentCategory) QueryChildren() *EquipmentCategoryQuery {
	return NewEquipmentCategoryClient(ec.config).QueryChildren(ec)
}

// QueryEquipment queries the "equipment" edge of the EquipmentCategory entity.
func (ec *EquipmentCategory) QueryEquipment() *EquipmentQuery {
	return NewEquipmentCategoryClient(ec.config).QueryEquipment(ec)
}

// Update returns a builder for updating this EquipmentCategory.
// Note that you need to call EquipmentCategory.Unwrap() before calling this method if this EquipmentCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *EquipmentCategory) Update() *EquipmentCategoryUpdateOne {
	return NewEquipmentCategoryClient(ec.config).UpdateOne(ec)
}

// Unwrap unwraps the EquipmentCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *EquipmentCategory) Unwrap() *EquipmentCategory {
	_tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: EquipmentCategory is not a transactional entity")
	}
	ec.config.driver = _tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *EquipmentCategory) String() string {
	var builder strings.Builder
	builder.WriteString("EquipmentCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ec.ID))
	builder.WriteString("parent_category_id=")
	builder.WriteString(fmt.Sprintf("%v", ec.ParentCategoryID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ec.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ec *EquipmentCategory) MarshalJSON() ([]byte, error) {
	type Alias EquipmentCategory
	return json.Marshal(&struct {
		*Alias
		EquipmentCategoryEdges
	}{
		Alias:                  (*Alias)(ec),
		EquipmentCategoryEdges: ec.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ec *EquipmentCategory) UnmarshalJSON(data []byte) error {
	type Alias EquipmentCategory
	aux := &struct {
		*Alias
		EquipmentCategoryEdges
	}{
		Alias: (*Alias)(ec),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	ec.Edges = aux.EquipmentCategoryEdges
	return nil
}

func (ecc *EquipmentCategoryCreate) SetEquipmentCategory(input *EquipmentCategory) *EquipmentCategoryCreate {
	ecc.SetParentCategoryID(input.ParentCategoryID)
	ecc.SetName(input.Name)
	return ecc
}

// NamedChildren returns the Children named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ec *EquipmentCategory) NamedChildren(name string) ([]*EquipmentCategory, error) {
	if ec.Edges.namedChildren == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ec.Edges.namedChildren[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ec *EquipmentCategory) appendNamedChildren(name string, edges ...*EquipmentCategory) {
	if ec.Edges.namedChildren == nil {
		ec.Edges.namedChildren = make(map[string][]*EquipmentCategory)
	}
	if len(edges) == 0 {
		ec.Edges.namedChildren[name] = []*EquipmentCategory{}
	} else {
		ec.Edges.namedChildren[name] = append(ec.Edges.namedChildren[name], edges...)
	}
}

// NamedEquipment returns the Equipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ec *EquipmentCategory) NamedEquipment(name string) ([]*Equipment, error) {
	if ec.Edges.namedEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ec.Edges.namedEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ec *EquipmentCategory) appendNamedEquipment(name string, edges ...*Equipment) {
	if ec.Edges.namedEquipment == nil {
		ec.Edges.namedEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		ec.Edges.namedEquipment[name] = []*Equipment{}
	} else {
		ec.Edges.namedEquipment[name] = append(ec.Edges.namedEquipment[name], edges...)
	}
}

// EquipmentCategories is a parsable slice of EquipmentCategory.
type EquipmentCategories []*EquipmentCategory
