// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmententry"
)

// EquipmentEntryCreate is the builder for creating a EquipmentEntry entity.
type EquipmentEntryCreate struct {
	config
	mutation *EquipmentEntryMutation
	hooks    []Hook
}

// SetQuantity sets the "quantity" field.
func (eec *EquipmentEntryCreate) SetQuantity(i int) *EquipmentEntryCreate {
	eec.mutation.SetQuantity(i)
	return eec
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (eec *EquipmentEntryCreate) AddClasIDs(ids ...int) *EquipmentEntryCreate {
	eec.mutation.AddClasIDs(ids...)
	return eec
}

// AddClass adds the "class" edges to the Class entity.
func (eec *EquipmentEntryCreate) AddClass(c ...*Class) *EquipmentEntryCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eec.AddClasIDs(ids...)
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eec *EquipmentEntryCreate) SetEquipmentID(id int) *EquipmentEntryCreate {
	eec.mutation.SetEquipmentID(id)
	return eec
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eec *EquipmentEntryCreate) SetEquipment(e *Equipment) *EquipmentEntryCreate {
	return eec.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentEntryMutation object of the builder.
func (eec *EquipmentEntryCreate) Mutation() *EquipmentEntryMutation {
	return eec.mutation
}

// Save creates the EquipmentEntry in the database.
func (eec *EquipmentEntryCreate) Save(ctx context.Context) (*EquipmentEntry, error) {
	return withHooks(ctx, eec.sqlSave, eec.mutation, eec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eec *EquipmentEntryCreate) SaveX(ctx context.Context) *EquipmentEntry {
	v, err := eec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eec *EquipmentEntryCreate) Exec(ctx context.Context) error {
	_, err := eec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eec *EquipmentEntryCreate) ExecX(ctx context.Context) {
	if err := eec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eec *EquipmentEntryCreate) check() error {
	if _, ok := eec.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "EquipmentEntry.quantity"`)}
	}
	if v, ok := eec.mutation.Quantity(); ok {
		if err := equipmententry.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "EquipmentEntry.quantity": %w`, err)}
		}
	}
	if len(eec.mutation.EquipmentIDs()) == 0 {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "EquipmentEntry.equipment"`)}
	}
	return nil
}

func (eec *EquipmentEntryCreate) sqlSave(ctx context.Context) (*EquipmentEntry, error) {
	if err := eec.check(); err != nil {
		return nil, err
	}
	_node, _spec := eec.createSpec()
	if err := sqlgraph.CreateNode(ctx, eec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	eec.mutation.id = &_node.ID
	eec.mutation.done = true
	return _node, nil
}

func (eec *EquipmentEntryCreate) createSpec() (*EquipmentEntry, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentEntry{config: eec.config}
		_spec = sqlgraph.NewCreateSpec(equipmententry.Table, sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt))
	)
	if value, ok := eec.mutation.Quantity(); ok {
		_spec.SetField(equipmententry.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if nodes := eec.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
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
	if nodes := eec.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmententry.EquipmentTable,
			Columns: []string{equipmententry.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.equipment_entry_equipment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentEntryCreateBulk is the builder for creating many EquipmentEntry entities in bulk.
type EquipmentEntryCreateBulk struct {
	config
	err      error
	builders []*EquipmentEntryCreate
}

// Save creates the EquipmentEntry entities in the database.
func (eecb *EquipmentEntryCreateBulk) Save(ctx context.Context) ([]*EquipmentEntry, error) {
	if eecb.err != nil {
		return nil, eecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eecb.builders))
	nodes := make([]*EquipmentEntry, len(eecb.builders))
	mutators := make([]Mutator, len(eecb.builders))
	for i := range eecb.builders {
		func(i int, root context.Context) {
			builder := eecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentEntryMutation)
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
					_, err = mutators[i+1].Mutate(root, eecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eecb *EquipmentEntryCreateBulk) SaveX(ctx context.Context) []*EquipmentEntry {
	v, err := eecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eecb *EquipmentEntryCreateBulk) Exec(ctx context.Context) error {
	_, err := eecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eecb *EquipmentEntryCreateBulk) ExecX(ctx context.Context) {
	if err := eecb.Exec(ctx); err != nil {
		panic(err)
	}
}