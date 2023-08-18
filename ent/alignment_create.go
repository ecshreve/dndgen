// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/alignment"
)

// AlignmentCreate is the builder for creating a Alignment entity.
type AlignmentCreate struct {
	config
	mutation *AlignmentMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (ac *AlignmentCreate) SetIndx(s string) *AlignmentCreate {
	ac.mutation.SetIndx(s)
	return ac
}

// SetName sets the "name" field.
func (ac *AlignmentCreate) SetName(s string) *AlignmentCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDesc sets the "desc" field.
func (ac *AlignmentCreate) SetDesc(s []string) *AlignmentCreate {
	ac.mutation.SetDesc(s)
	return ac
}

// SetAbbr sets the "abbr" field.
func (ac *AlignmentCreate) SetAbbr(s string) *AlignmentCreate {
	ac.mutation.SetAbbr(s)
	return ac
}

// Mutation returns the AlignmentMutation object of the builder.
func (ac *AlignmentCreate) Mutation() *AlignmentMutation {
	return ac.mutation
}

// Save creates the Alignment in the database.
func (ac *AlignmentCreate) Save(ctx context.Context) (*Alignment, error) {
	return withHooks[*Alignment, AlignmentMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AlignmentCreate) SaveX(ctx context.Context) *Alignment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AlignmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AlignmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AlignmentCreate) check() error {
	if _, ok := ac.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Alignment.indx"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Alignment.name"`)}
	}
	if _, ok := ac.mutation.Abbr(); !ok {
		return &ValidationError{Name: "abbr", err: errors.New(`ent: missing required field "Alignment.abbr"`)}
	}
	return nil
}

func (ac *AlignmentCreate) sqlSave(ctx context.Context) (*Alignment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AlignmentCreate) createSpec() (*Alignment, *sqlgraph.CreateSpec) {
	var (
		_node = &Alignment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(alignment.Table, sqlgraph.NewFieldSpec(alignment.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Indx(); ok {
		_spec.SetField(alignment.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(alignment.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Desc(); ok {
		_spec.SetField(alignment.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if value, ok := ac.mutation.Abbr(); ok {
		_spec.SetField(alignment.FieldAbbr, field.TypeString, value)
		_node.Abbr = value
	}
	return _node, _spec
}

// AlignmentCreateBulk is the builder for creating many Alignment entities in bulk.
type AlignmentCreateBulk struct {
	config
	builders []*AlignmentCreate
}

// Save creates the Alignment entities in the database.
func (acb *AlignmentCreateBulk) Save(ctx context.Context) ([]*Alignment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Alignment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlignmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AlignmentCreateBulk) SaveX(ctx context.Context) []*Alignment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AlignmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AlignmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
