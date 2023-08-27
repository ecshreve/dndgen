// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// GearUpdate is the builder for updating Gear entities.
type GearUpdate struct {
	config
	hooks    []Hook
	mutation *GearMutation
}

// Where appends a list predicates to the GearUpdate builder.
func (gu *GearUpdate) Where(ps ...predicate.Gear) *GearUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetIndx sets the "indx" field.
func (gu *GearUpdate) SetIndx(s string) *GearUpdate {
	gu.mutation.SetIndx(s)
	return gu
}

// SetName sets the "name" field.
func (gu *GearUpdate) SetName(s string) *GearUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetGearCategory sets the "gear_category" field.
func (gu *GearUpdate) SetGearCategory(gc gear.GearCategory) *GearUpdate {
	gu.mutation.SetGearCategory(gc)
	return gu
}

// SetNillableGearCategory sets the "gear_category" field if the given value is not nil.
func (gu *GearUpdate) SetNillableGearCategory(gc *gear.GearCategory) *GearUpdate {
	if gc != nil {
		gu.SetGearCategory(*gc)
	}
	return gu
}

// SetDesc sets the "desc" field.
func (gu *GearUpdate) SetDesc(s []string) *GearUpdate {
	gu.mutation.SetDesc(s)
	return gu
}

// AppendDesc appends s to the "desc" field.
func (gu *GearUpdate) AppendDesc(s []string) *GearUpdate {
	gu.mutation.AppendDesc(s)
	return gu
}

// SetQuantity sets the "quantity" field.
func (gu *GearUpdate) SetQuantity(i int) *GearUpdate {
	gu.mutation.ResetQuantity()
	gu.mutation.SetQuantity(i)
	return gu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (gu *GearUpdate) SetNillableQuantity(i *int) *GearUpdate {
	if i != nil {
		gu.SetQuantity(*i)
	}
	return gu
}

// AddQuantity adds i to the "quantity" field.
func (gu *GearUpdate) AddQuantity(i int) *GearUpdate {
	gu.mutation.AddQuantity(i)
	return gu
}

// ClearQuantity clears the value of the "quantity" field.
func (gu *GearUpdate) ClearQuantity() *GearUpdate {
	gu.mutation.ClearQuantity()
	return gu
}

// SetEquipmentID sets the "equipment_id" field.
func (gu *GearUpdate) SetEquipmentID(i int) *GearUpdate {
	gu.mutation.SetEquipmentID(i)
	return gu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (gu *GearUpdate) SetEquipment(e *Equipment) *GearUpdate {
	return gu.SetEquipmentID(e.ID)
}

// Mutation returns the GearMutation object of the builder.
func (gu *GearUpdate) Mutation() *GearMutation {
	return gu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (gu *GearUpdate) ClearEquipment() *GearUpdate {
	gu.mutation.ClearEquipment()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GearUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GearUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GearUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GearUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GearUpdate) check() error {
	if v, ok := gu.mutation.Indx(); ok {
		if err := gear.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Gear.indx": %w`, err)}
		}
	}
	if v, ok := gu.mutation.Name(); ok {
		if err := gear.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Gear.name": %w`, err)}
		}
	}
	if v, ok := gu.mutation.GearCategory(); ok {
		if err := gear.GearCategoryValidator(v); err != nil {
			return &ValidationError{Name: "gear_category", err: fmt.Errorf(`ent: validator failed for field "Gear.gear_category": %w`, err)}
		}
	}
	if _, ok := gu.mutation.EquipmentID(); gu.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Gear.equipment"`)
	}
	return nil
}

func (gu *GearUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(gear.Table, gear.Columns, sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Indx(); ok {
		_spec.SetField(gear.FieldIndx, field.TypeString, value)
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(gear.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.GearCategory(); ok {
		_spec.SetField(gear.FieldGearCategory, field.TypeEnum, value)
	}
	if value, ok := gu.mutation.Desc(); ok {
		_spec.SetField(gear.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := gu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, gear.FieldDesc, value)
		})
	}
	if value, ok := gu.mutation.Quantity(); ok {
		_spec.SetField(gear.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := gu.mutation.AddedQuantity(); ok {
		_spec.AddField(gear.FieldQuantity, field.TypeInt, value)
	}
	if gu.mutation.QuantityCleared() {
		_spec.ClearField(gear.FieldQuantity, field.TypeInt)
	}
	if gu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   gear.EquipmentTable,
			Columns: []string{gear.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   gear.EquipmentTable,
			Columns: []string{gear.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gear.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GearUpdateOne is the builder for updating a single Gear entity.
type GearUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GearMutation
}

// SetIndx sets the "indx" field.
func (guo *GearUpdateOne) SetIndx(s string) *GearUpdateOne {
	guo.mutation.SetIndx(s)
	return guo
}

// SetName sets the "name" field.
func (guo *GearUpdateOne) SetName(s string) *GearUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetGearCategory sets the "gear_category" field.
func (guo *GearUpdateOne) SetGearCategory(gc gear.GearCategory) *GearUpdateOne {
	guo.mutation.SetGearCategory(gc)
	return guo
}

// SetNillableGearCategory sets the "gear_category" field if the given value is not nil.
func (guo *GearUpdateOne) SetNillableGearCategory(gc *gear.GearCategory) *GearUpdateOne {
	if gc != nil {
		guo.SetGearCategory(*gc)
	}
	return guo
}

// SetDesc sets the "desc" field.
func (guo *GearUpdateOne) SetDesc(s []string) *GearUpdateOne {
	guo.mutation.SetDesc(s)
	return guo
}

// AppendDesc appends s to the "desc" field.
func (guo *GearUpdateOne) AppendDesc(s []string) *GearUpdateOne {
	guo.mutation.AppendDesc(s)
	return guo
}

// SetQuantity sets the "quantity" field.
func (guo *GearUpdateOne) SetQuantity(i int) *GearUpdateOne {
	guo.mutation.ResetQuantity()
	guo.mutation.SetQuantity(i)
	return guo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (guo *GearUpdateOne) SetNillableQuantity(i *int) *GearUpdateOne {
	if i != nil {
		guo.SetQuantity(*i)
	}
	return guo
}

// AddQuantity adds i to the "quantity" field.
func (guo *GearUpdateOne) AddQuantity(i int) *GearUpdateOne {
	guo.mutation.AddQuantity(i)
	return guo
}

// ClearQuantity clears the value of the "quantity" field.
func (guo *GearUpdateOne) ClearQuantity() *GearUpdateOne {
	guo.mutation.ClearQuantity()
	return guo
}

// SetEquipmentID sets the "equipment_id" field.
func (guo *GearUpdateOne) SetEquipmentID(i int) *GearUpdateOne {
	guo.mutation.SetEquipmentID(i)
	return guo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (guo *GearUpdateOne) SetEquipment(e *Equipment) *GearUpdateOne {
	return guo.SetEquipmentID(e.ID)
}

// Mutation returns the GearMutation object of the builder.
func (guo *GearUpdateOne) Mutation() *GearMutation {
	return guo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (guo *GearUpdateOne) ClearEquipment() *GearUpdateOne {
	guo.mutation.ClearEquipment()
	return guo
}

// Where appends a list predicates to the GearUpdate builder.
func (guo *GearUpdateOne) Where(ps ...predicate.Gear) *GearUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GearUpdateOne) Select(field string, fields ...string) *GearUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Gear entity.
func (guo *GearUpdateOne) Save(ctx context.Context) (*Gear, error) {
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GearUpdateOne) SaveX(ctx context.Context) *Gear {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GearUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GearUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GearUpdateOne) check() error {
	if v, ok := guo.mutation.Indx(); ok {
		if err := gear.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Gear.indx": %w`, err)}
		}
	}
	if v, ok := guo.mutation.Name(); ok {
		if err := gear.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Gear.name": %w`, err)}
		}
	}
	if v, ok := guo.mutation.GearCategory(); ok {
		if err := gear.GearCategoryValidator(v); err != nil {
			return &ValidationError{Name: "gear_category", err: fmt.Errorf(`ent: validator failed for field "Gear.gear_category": %w`, err)}
		}
	}
	if _, ok := guo.mutation.EquipmentID(); guo.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Gear.equipment"`)
	}
	return nil
}

func (guo *GearUpdateOne) sqlSave(ctx context.Context) (_node *Gear, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(gear.Table, gear.Columns, sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Gear.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gear.FieldID)
		for _, f := range fields {
			if !gear.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gear.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Indx(); ok {
		_spec.SetField(gear.FieldIndx, field.TypeString, value)
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(gear.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.GearCategory(); ok {
		_spec.SetField(gear.FieldGearCategory, field.TypeEnum, value)
	}
	if value, ok := guo.mutation.Desc(); ok {
		_spec.SetField(gear.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := guo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, gear.FieldDesc, value)
		})
	}
	if value, ok := guo.mutation.Quantity(); ok {
		_spec.SetField(gear.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := guo.mutation.AddedQuantity(); ok {
		_spec.AddField(gear.FieldQuantity, field.TypeInt, value)
	}
	if guo.mutation.QuantityCleared() {
		_spec.ClearField(gear.FieldQuantity, field.TypeInt)
	}
	if guo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   gear.EquipmentTable,
			Columns: []string{gear.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   gear.EquipmentTable,
			Columns: []string{gear.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Gear{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gear.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
