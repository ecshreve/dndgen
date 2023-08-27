// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/vehicle"
)

// VehicleUpdate is the builder for updating Vehicle entities.
type VehicleUpdate struct {
	config
	hooks    []Hook
	mutation *VehicleMutation
}

// Where appends a list predicates to the VehicleUpdate builder.
func (vu *VehicleUpdate) Where(ps ...predicate.Vehicle) *VehicleUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetIndx sets the "indx" field.
func (vu *VehicleUpdate) SetIndx(s string) *VehicleUpdate {
	vu.mutation.SetIndx(s)
	return vu
}

// SetName sets the "name" field.
func (vu *VehicleUpdate) SetName(s string) *VehicleUpdate {
	vu.mutation.SetName(s)
	return vu
}

// SetVehicleCategory sets the "vehicle_category" field.
func (vu *VehicleUpdate) SetVehicleCategory(s string) *VehicleUpdate {
	vu.mutation.SetVehicleCategory(s)
	return vu
}

// SetCapacity sets the "capacity" field.
func (vu *VehicleUpdate) SetCapacity(s string) *VehicleUpdate {
	vu.mutation.SetCapacity(s)
	return vu
}

// SetEquipmentID sets the "equipment_id" field.
func (vu *VehicleUpdate) SetEquipmentID(i int) *VehicleUpdate {
	vu.mutation.SetEquipmentID(i)
	return vu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (vu *VehicleUpdate) SetEquipment(e *Equipment) *VehicleUpdate {
	return vu.SetEquipmentID(e.ID)
}

// Mutation returns the VehicleMutation object of the builder.
func (vu *VehicleUpdate) Mutation() *VehicleMutation {
	return vu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (vu *VehicleUpdate) ClearEquipment() *VehicleUpdate {
	vu.mutation.ClearEquipment()
	return vu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VehicleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VehicleUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VehicleUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VehicleUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VehicleUpdate) check() error {
	if v, ok := vu.mutation.Indx(); ok {
		if err := vehicle.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Vehicle.indx": %w`, err)}
		}
	}
	if v, ok := vu.mutation.Name(); ok {
		if err := vehicle.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vehicle.name": %w`, err)}
		}
	}
	if _, ok := vu.mutation.EquipmentID(); vu.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Vehicle.equipment"`)
	}
	return nil
}

func (vu *VehicleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(vehicle.Table, vehicle.Columns, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Indx(); ok {
		_spec.SetField(vehicle.FieldIndx, field.TypeString, value)
	}
	if value, ok := vu.mutation.Name(); ok {
		_spec.SetField(vehicle.FieldName, field.TypeString, value)
	}
	if value, ok := vu.mutation.VehicleCategory(); ok {
		_spec.SetField(vehicle.FieldVehicleCategory, field.TypeString, value)
	}
	if value, ok := vu.mutation.Capacity(); ok {
		_spec.SetField(vehicle.FieldCapacity, field.TypeString, value)
	}
	if vu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vehicle.EquipmentTable,
			Columns: []string{vehicle.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vehicle.EquipmentTable,
			Columns: []string{vehicle.EquipmentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VehicleUpdateOne is the builder for updating a single Vehicle entity.
type VehicleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VehicleMutation
}

// SetIndx sets the "indx" field.
func (vuo *VehicleUpdateOne) SetIndx(s string) *VehicleUpdateOne {
	vuo.mutation.SetIndx(s)
	return vuo
}

// SetName sets the "name" field.
func (vuo *VehicleUpdateOne) SetName(s string) *VehicleUpdateOne {
	vuo.mutation.SetName(s)
	return vuo
}

// SetVehicleCategory sets the "vehicle_category" field.
func (vuo *VehicleUpdateOne) SetVehicleCategory(s string) *VehicleUpdateOne {
	vuo.mutation.SetVehicleCategory(s)
	return vuo
}

// SetCapacity sets the "capacity" field.
func (vuo *VehicleUpdateOne) SetCapacity(s string) *VehicleUpdateOne {
	vuo.mutation.SetCapacity(s)
	return vuo
}

// SetEquipmentID sets the "equipment_id" field.
func (vuo *VehicleUpdateOne) SetEquipmentID(i int) *VehicleUpdateOne {
	vuo.mutation.SetEquipmentID(i)
	return vuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (vuo *VehicleUpdateOne) SetEquipment(e *Equipment) *VehicleUpdateOne {
	return vuo.SetEquipmentID(e.ID)
}

// Mutation returns the VehicleMutation object of the builder.
func (vuo *VehicleUpdateOne) Mutation() *VehicleMutation {
	return vuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (vuo *VehicleUpdateOne) ClearEquipment() *VehicleUpdateOne {
	vuo.mutation.ClearEquipment()
	return vuo
}

// Where appends a list predicates to the VehicleUpdate builder.
func (vuo *VehicleUpdateOne) Where(ps ...predicate.Vehicle) *VehicleUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VehicleUpdateOne) Select(field string, fields ...string) *VehicleUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vehicle entity.
func (vuo *VehicleUpdateOne) Save(ctx context.Context) (*Vehicle, error) {
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VehicleUpdateOne) SaveX(ctx context.Context) *Vehicle {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VehicleUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VehicleUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VehicleUpdateOne) check() error {
	if v, ok := vuo.mutation.Indx(); ok {
		if err := vehicle.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Vehicle.indx": %w`, err)}
		}
	}
	if v, ok := vuo.mutation.Name(); ok {
		if err := vehicle.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Vehicle.name": %w`, err)}
		}
	}
	if _, ok := vuo.mutation.EquipmentID(); vuo.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Vehicle.equipment"`)
	}
	return nil
}

func (vuo *VehicleUpdateOne) sqlSave(ctx context.Context) (_node *Vehicle, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(vehicle.Table, vehicle.Columns, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Vehicle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicle.FieldID)
		for _, f := range fields {
			if !vehicle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vehicle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Indx(); ok {
		_spec.SetField(vehicle.FieldIndx, field.TypeString, value)
	}
	if value, ok := vuo.mutation.Name(); ok {
		_spec.SetField(vehicle.FieldName, field.TypeString, value)
	}
	if value, ok := vuo.mutation.VehicleCategory(); ok {
		_spec.SetField(vehicle.FieldVehicleCategory, field.TypeString, value)
	}
	if value, ok := vuo.mutation.Capacity(); ok {
		_spec.SetField(vehicle.FieldCapacity, field.TypeString, value)
	}
	if vuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vehicle.EquipmentTable,
			Columns: []string{vehicle.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   vehicle.EquipmentTable,
			Columns: []string{vehicle.EquipmentColumn},
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
	_node = &Vehicle{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
