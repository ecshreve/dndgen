// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/vehicle"
)

// VehicleCreate is the builder for creating a Vehicle entity.
type VehicleCreate struct {
	config
	mutation *VehicleMutation
	hooks    []Hook
}

// SetVehicleCategory sets the "vehicle_category" field.
func (vc *VehicleCreate) SetVehicleCategory(value vehicle.VehicleCategory) *VehicleCreate {
	vc.mutation.SetVehicleCategory(value)
	return vc
}

// SetCapacity sets the "capacity" field.
func (vc *VehicleCreate) SetCapacity(s string) *VehicleCreate {
	vc.mutation.SetCapacity(s)
	return vc
}

// SetNillableCapacity sets the "capacity" field if the given value is not nil.
func (vc *VehicleCreate) SetNillableCapacity(s *string) *VehicleCreate {
	if s != nil {
		vc.SetCapacity(*s)
	}
	return vc
}

// SetDesc sets the "desc" field.
func (vc *VehicleCreate) SetDesc(s []string) *VehicleCreate {
	vc.mutation.SetDesc(s)
	return vc
}

// SetSpeedQuantity sets the "speed_quantity" field.
func (vc *VehicleCreate) SetSpeedQuantity(f float64) *VehicleCreate {
	vc.mutation.SetSpeedQuantity(f)
	return vc
}

// SetNillableSpeedQuantity sets the "speed_quantity" field if the given value is not nil.
func (vc *VehicleCreate) SetNillableSpeedQuantity(f *float64) *VehicleCreate {
	if f != nil {
		vc.SetSpeedQuantity(*f)
	}
	return vc
}

// SetSpeedUnits sets the "speed_units" field.
func (vc *VehicleCreate) SetSpeedUnits(vu vehicle.SpeedUnits) *VehicleCreate {
	vc.mutation.SetSpeedUnits(vu)
	return vc
}

// SetNillableSpeedUnits sets the "speed_units" field if the given value is not nil.
func (vc *VehicleCreate) SetNillableSpeedUnits(vu *vehicle.SpeedUnits) *VehicleCreate {
	if vu != nil {
		vc.SetSpeedUnits(*vu)
	}
	return vc
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (vc *VehicleCreate) SetEquipmentID(id int) *VehicleCreate {
	vc.mutation.SetEquipmentID(id)
	return vc
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (vc *VehicleCreate) SetEquipment(e *Equipment) *VehicleCreate {
	return vc.SetEquipmentID(e.ID)
}

// Mutation returns the VehicleMutation object of the builder.
func (vc *VehicleCreate) Mutation() *VehicleMutation {
	return vc.mutation
}

// Save creates the Vehicle in the database.
func (vc *VehicleCreate) Save(ctx context.Context) (*Vehicle, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VehicleCreate) SaveX(ctx context.Context) *Vehicle {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VehicleCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VehicleCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VehicleCreate) check() error {
	if _, ok := vc.mutation.VehicleCategory(); !ok {
		return &ValidationError{Name: "vehicle_category", err: errors.New(`ent: missing required field "Vehicle.vehicle_category"`)}
	}
	if v, ok := vc.mutation.VehicleCategory(); ok {
		if err := vehicle.VehicleCategoryValidator(v); err != nil {
			return &ValidationError{Name: "vehicle_category", err: fmt.Errorf(`ent: validator failed for field "Vehicle.vehicle_category": %w`, err)}
		}
	}
	if v, ok := vc.mutation.SpeedUnits(); ok {
		if err := vehicle.SpeedUnitsValidator(v); err != nil {
			return &ValidationError{Name: "speed_units", err: fmt.Errorf(`ent: validator failed for field "Vehicle.speed_units": %w`, err)}
		}
	}
	if len(vc.mutation.EquipmentIDs()) == 0 {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "Vehicle.equipment"`)}
	}
	return nil
}

func (vc *VehicleCreate) sqlSave(ctx context.Context) (*Vehicle, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VehicleCreate) createSpec() (*Vehicle, *sqlgraph.CreateSpec) {
	var (
		_node = &Vehicle{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(vehicle.Table, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.VehicleCategory(); ok {
		_spec.SetField(vehicle.FieldVehicleCategory, field.TypeEnum, value)
		_node.VehicleCategory = value
	}
	if value, ok := vc.mutation.Capacity(); ok {
		_spec.SetField(vehicle.FieldCapacity, field.TypeString, value)
		_node.Capacity = value
	}
	if value, ok := vc.mutation.Desc(); ok {
		_spec.SetField(vehicle.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if value, ok := vc.mutation.SpeedQuantity(); ok {
		_spec.SetField(vehicle.FieldSpeedQuantity, field.TypeFloat64, value)
		_node.SpeedQuantity = value
	}
	if value, ok := vc.mutation.SpeedUnits(); ok {
		_spec.SetField(vehicle.FieldSpeedUnits, field.TypeEnum, value)
		_node.SpeedUnits = value
	}
	if nodes := vc.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_node.equipment_vehicle = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VehicleCreateBulk is the builder for creating many Vehicle entities in bulk.
type VehicleCreateBulk struct {
	config
	err      error
	builders []*VehicleCreate
}

// Save creates the Vehicle entities in the database.
func (vcb *VehicleCreateBulk) Save(ctx context.Context) ([]*Vehicle, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vehicle, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VehicleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VehicleCreateBulk) SaveX(ctx context.Context) []*Vehicle {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VehicleCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VehicleCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
