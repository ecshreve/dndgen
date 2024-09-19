// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damagetype"
)

// DamageTypeCreate is the builder for creating a DamageType entity.
type DamageTypeCreate struct {
	config
	mutation *DamageTypeMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (dtc *DamageTypeCreate) SetIndx(s string) *DamageTypeCreate {
	dtc.mutation.SetIndx(s)
	return dtc
}

// SetName sets the "name" field.
func (dtc *DamageTypeCreate) SetName(s string) *DamageTypeCreate {
	dtc.mutation.SetName(s)
	return dtc
}

// SetDesc sets the "desc" field.
func (dtc *DamageTypeCreate) SetDesc(s []string) *DamageTypeCreate {
	dtc.mutation.SetDesc(s)
	return dtc
}

// Mutation returns the DamageTypeMutation object of the builder.
func (dtc *DamageTypeCreate) Mutation() *DamageTypeMutation {
	return dtc.mutation
}

// Save creates the DamageType in the database.
func (dtc *DamageTypeCreate) Save(ctx context.Context) (*DamageType, error) {
	return withHooks(ctx, dtc.sqlSave, dtc.mutation, dtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dtc *DamageTypeCreate) SaveX(ctx context.Context) *DamageType {
	v, err := dtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtc *DamageTypeCreate) Exec(ctx context.Context) error {
	_, err := dtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtc *DamageTypeCreate) ExecX(ctx context.Context) {
	if err := dtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtc *DamageTypeCreate) check() error {
	if _, ok := dtc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "DamageType.indx"`)}
	}
	if v, ok := dtc.mutation.Indx(); ok {
		if err := damagetype.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "DamageType.indx": %w`, err)}
		}
	}
	if _, ok := dtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DamageType.name"`)}
	}
	if v, ok := dtc.mutation.Name(); ok {
		if err := damagetype.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "DamageType.name": %w`, err)}
		}
	}
	return nil
}

func (dtc *DamageTypeCreate) sqlSave(ctx context.Context) (*DamageType, error) {
	if err := dtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dtc.mutation.id = &_node.ID
	dtc.mutation.done = true
	return _node, nil
}

func (dtc *DamageTypeCreate) createSpec() (*DamageType, *sqlgraph.CreateSpec) {
	var (
		_node = &DamageType{config: dtc.config}
		_spec = sqlgraph.NewCreateSpec(damagetype.Table, sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt))
	)
	if value, ok := dtc.mutation.Indx(); ok {
		_spec.SetField(damagetype.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := dtc.mutation.Name(); ok {
		_spec.SetField(damagetype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dtc.mutation.Desc(); ok {
		_spec.SetField(damagetype.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	return _node, _spec
}

// DamageTypeCreateBulk is the builder for creating many DamageType entities in bulk.
type DamageTypeCreateBulk struct {
	config
	err      error
	builders []*DamageTypeCreate
}

// Save creates the DamageType entities in the database.
func (dtcb *DamageTypeCreateBulk) Save(ctx context.Context) ([]*DamageType, error) {
	if dtcb.err != nil {
		return nil, dtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dtcb.builders))
	nodes := make([]*DamageType, len(dtcb.builders))
	mutators := make([]Mutator, len(dtcb.builders))
	for i := range dtcb.builders {
		func(i int, root context.Context) {
			builder := dtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DamageTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, dtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtcb *DamageTypeCreateBulk) SaveX(ctx context.Context) []*DamageType {
	v, err := dtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtcb *DamageTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := dtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtcb *DamageTypeCreateBulk) ExecX(ctx context.Context) {
	if err := dtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
