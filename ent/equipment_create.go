// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentchoice"
	"github.com/ecshreve/dndgen/ent/equipmentcost"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/tool"
	"github.com/ecshreve/dndgen/ent/vehicle"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// EquipmentCreate is the builder for creating a Equipment entity.
type EquipmentCreate struct {
	config
	mutation *EquipmentMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (ec *EquipmentCreate) SetIndx(s string) *EquipmentCreate {
	ec.mutation.SetIndx(s)
	return ec
}

// SetName sets the "name" field.
func (ec *EquipmentCreate) SetName(s string) *EquipmentCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetEquipmentCategory sets the "equipment_category" field.
func (ec *EquipmentCreate) SetEquipmentCategory(value equipment.EquipmentCategory) *EquipmentCreate {
	ec.mutation.SetEquipmentCategory(value)
	return ec
}

// SetNillableEquipmentCategory sets the "equipment_category" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableEquipmentCategory(value *equipment.EquipmentCategory) *EquipmentCreate {
	if value != nil {
		ec.SetEquipmentCategory(*value)
	}
	return ec
}

// SetEquipmentSubcategory sets the "equipment_subcategory" field.
func (ec *EquipmentCreate) SetEquipmentSubcategory(s string) *EquipmentCreate {
	ec.mutation.SetEquipmentSubcategory(s)
	return ec
}

// SetNillableEquipmentSubcategory sets the "equipment_subcategory" field if the given value is not nil.
func (ec *EquipmentCreate) SetNillableEquipmentSubcategory(s *string) *EquipmentCreate {
	if s != nil {
		ec.SetEquipmentSubcategory(*s)
	}
	return ec
}

// SetCostID sets the "cost" edge to the EquipmentCost entity by ID.
func (ec *EquipmentCreate) SetCostID(id int) *EquipmentCreate {
	ec.mutation.SetCostID(id)
	return ec
}

// SetNillableCostID sets the "cost" edge to the EquipmentCost entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableCostID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetCostID(*id)
	}
	return ec
}

// SetCost sets the "cost" edge to the EquipmentCost entity.
func (ec *EquipmentCreate) SetCost(e *EquipmentCost) *EquipmentCreate {
	return ec.SetCostID(e.ID)
}

// SetWeaponID sets the "weapon" edge to the Weapon entity by ID.
func (ec *EquipmentCreate) SetWeaponID(id int) *EquipmentCreate {
	ec.mutation.SetWeaponID(id)
	return ec
}

// SetNillableWeaponID sets the "weapon" edge to the Weapon entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableWeaponID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetWeaponID(*id)
	}
	return ec
}

// SetWeapon sets the "weapon" edge to the Weapon entity.
func (ec *EquipmentCreate) SetWeapon(w *Weapon) *EquipmentCreate {
	return ec.SetWeaponID(w.ID)
}

// SetArmorID sets the "armor" edge to the Armor entity by ID.
func (ec *EquipmentCreate) SetArmorID(id int) *EquipmentCreate {
	ec.mutation.SetArmorID(id)
	return ec
}

// SetNillableArmorID sets the "armor" edge to the Armor entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableArmorID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetArmorID(*id)
	}
	return ec
}

// SetArmor sets the "armor" edge to the Armor entity.
func (ec *EquipmentCreate) SetArmor(a *Armor) *EquipmentCreate {
	return ec.SetArmorID(a.ID)
}

// SetGearID sets the "gear" edge to the Gear entity by ID.
func (ec *EquipmentCreate) SetGearID(id int) *EquipmentCreate {
	ec.mutation.SetGearID(id)
	return ec
}

// SetNillableGearID sets the "gear" edge to the Gear entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableGearID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetGearID(*id)
	}
	return ec
}

// SetGear sets the "gear" edge to the Gear entity.
func (ec *EquipmentCreate) SetGear(g *Gear) *EquipmentCreate {
	return ec.SetGearID(g.ID)
}

// SetToolID sets the "tool" edge to the Tool entity by ID.
func (ec *EquipmentCreate) SetToolID(id int) *EquipmentCreate {
	ec.mutation.SetToolID(id)
	return ec
}

// SetNillableToolID sets the "tool" edge to the Tool entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableToolID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetToolID(*id)
	}
	return ec
}

// SetTool sets the "tool" edge to the Tool entity.
func (ec *EquipmentCreate) SetTool(t *Tool) *EquipmentCreate {
	return ec.SetToolID(t.ID)
}

// SetVehicleID sets the "vehicle" edge to the Vehicle entity by ID.
func (ec *EquipmentCreate) SetVehicleID(id int) *EquipmentCreate {
	ec.mutation.SetVehicleID(id)
	return ec
}

// SetNillableVehicleID sets the "vehicle" edge to the Vehicle entity by ID if the given value is not nil.
func (ec *EquipmentCreate) SetNillableVehicleID(id *int) *EquipmentCreate {
	if id != nil {
		ec = ec.SetVehicleID(*id)
	}
	return ec
}

// SetVehicle sets the "vehicle" edge to the Vehicle entity.
func (ec *EquipmentCreate) SetVehicle(v *Vehicle) *EquipmentCreate {
	return ec.SetVehicleID(v.ID)
}

// AddClassEquipmentIDs adds the "class_equipment" edge to the Class entity by IDs.
func (ec *EquipmentCreate) AddClassEquipmentIDs(ids ...int) *EquipmentCreate {
	ec.mutation.AddClassEquipmentIDs(ids...)
	return ec
}

// AddClassEquipment adds the "class_equipment" edges to the Class entity.
func (ec *EquipmentCreate) AddClassEquipment(c ...*Class) *EquipmentCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddClassEquipmentIDs(ids...)
}

// AddChoiceIDs adds the "choice" edge to the EquipmentChoice entity by IDs.
func (ec *EquipmentCreate) AddChoiceIDs(ids ...int) *EquipmentCreate {
	ec.mutation.AddChoiceIDs(ids...)
	return ec
}

// AddChoice adds the "choice" edges to the EquipmentChoice entity.
func (ec *EquipmentCreate) AddChoice(e ...*EquipmentChoice) *EquipmentCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddChoiceIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (ec *EquipmentCreate) Mutation() *EquipmentMutation {
	return ec.mutation
}

// Save creates the Equipment in the database.
func (ec *EquipmentCreate) Save(ctx context.Context) (*Equipment, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EquipmentCreate) SaveX(ctx context.Context) *Equipment {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EquipmentCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EquipmentCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EquipmentCreate) defaults() {
	if _, ok := ec.mutation.EquipmentCategory(); !ok {
		v := equipment.DefaultEquipmentCategory
		ec.mutation.SetEquipmentCategory(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EquipmentCreate) check() error {
	if _, ok := ec.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Equipment.indx"`)}
	}
	if v, ok := ec.mutation.Indx(); ok {
		if err := equipment.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Equipment.indx": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Equipment.name"`)}
	}
	if v, ok := ec.mutation.Name(); ok {
		if err := equipment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Equipment.name": %w`, err)}
		}
	}
	if _, ok := ec.mutation.EquipmentCategory(); !ok {
		return &ValidationError{Name: "equipment_category", err: errors.New(`ent: missing required field "Equipment.equipment_category"`)}
	}
	if v, ok := ec.mutation.EquipmentCategory(); ok {
		if err := equipment.EquipmentCategoryValidator(v); err != nil {
			return &ValidationError{Name: "equipment_category", err: fmt.Errorf(`ent: validator failed for field "Equipment.equipment_category": %w`, err)}
		}
	}
	return nil
}

func (ec *EquipmentCreate) sqlSave(ctx context.Context) (*Equipment, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EquipmentCreate) createSpec() (*Equipment, *sqlgraph.CreateSpec) {
	var (
		_node = &Equipment{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(equipment.Table, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.Indx(); ok {
		_spec.SetField(equipment.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(equipment.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.EquipmentCategory(); ok {
		_spec.SetField(equipment.FieldEquipmentCategory, field.TypeEnum, value)
		_node.EquipmentCategory = value
	}
	if value, ok := ec.mutation.EquipmentSubcategory(); ok {
		_spec.SetField(equipment.FieldEquipmentSubcategory, field.TypeString, value)
		_node.EquipmentSubcategory = value
	}
	if nodes := ec.mutation.CostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.CostTable,
			Columns: []string{equipment.CostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmentcost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.WeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.WeaponTable,
			Columns: []string{equipment.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ArmorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ArmorTable,
			Columns: []string{equipment.ArmorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.GearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.GearTable,
			Columns: []string{equipment.GearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ToolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ToolTable,
			Columns: []string{equipment.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.VehicleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.VehicleTable,
			Columns: []string{equipment.VehicleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ClassEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipment.ClassEquipmentTable,
			Columns: equipment.ClassEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ChoiceTable,
			Columns: equipment.ChoicePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmentchoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentCreateBulk is the builder for creating many Equipment entities in bulk.
type EquipmentCreateBulk struct {
	config
	builders []*EquipmentCreate
}

// Save creates the Equipment entities in the database.
func (ecb *EquipmentCreateBulk) Save(ctx context.Context) ([]*Equipment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Equipment, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) SaveX(ctx context.Context) []*Equipment {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EquipmentCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EquipmentCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
