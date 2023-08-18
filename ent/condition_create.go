// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/condition"
)

// ConditionCreate is the builder for creating a Condition entity.
type ConditionCreate struct {
	config
	mutation *ConditionMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (cc *ConditionCreate) SetIndx(s string) *ConditionCreate {
	cc.mutation.SetIndx(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ConditionCreate) SetName(s string) *ConditionCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDesc sets the "desc" field.
func (cc *ConditionCreate) SetDesc(s []string) *ConditionCreate {
	cc.mutation.SetDesc(s)
	return cc
}

// Mutation returns the ConditionMutation object of the builder.
func (cc *ConditionCreate) Mutation() *ConditionMutation {
	return cc.mutation
}

// Save creates the Condition in the database.
func (cc *ConditionCreate) Save(ctx context.Context) (*Condition, error) {
	return withHooks[*Condition, ConditionMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConditionCreate) SaveX(ctx context.Context) *Condition {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConditionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConditionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConditionCreate) check() error {
	if _, ok := cc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Condition.indx"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Condition.name"`)}
	}
	return nil
}

func (cc *ConditionCreate) sqlSave(ctx context.Context) (*Condition, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ConditionCreate) createSpec() (*Condition, *sqlgraph.CreateSpec) {
	var (
		_node = &Condition{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(condition.Table, sqlgraph.NewFieldSpec(condition.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Indx(); ok {
		_spec.SetField(condition.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(condition.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Desc(); ok {
		_spec.SetField(condition.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	return _node, _spec
}

// ConditionCreateBulk is the builder for creating many Condition entities in bulk.
type ConditionCreateBulk struct {
	config
	builders []*ConditionCreate
}

// Save creates the Condition entities in the database.
func (ccb *ConditionCreateBulk) Save(ctx context.Context) ([]*Condition, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Condition, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConditionMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConditionCreateBulk) SaveX(ctx context.Context) []*Condition {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConditionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConditionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
