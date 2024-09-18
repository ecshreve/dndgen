// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/builder/ent/abilityscore"
)

// AbilityScoreCreate is the builder for creating a AbilityScore entity.
type AbilityScoreCreate struct {
	config
	mutation *AbilityScoreMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (asc *AbilityScoreCreate) SetIndx(s string) *AbilityScoreCreate {
	asc.mutation.SetIndx(s)
	return asc
}

// SetName sets the "name" field.
func (asc *AbilityScoreCreate) SetName(s string) *AbilityScoreCreate {
	asc.mutation.SetName(s)
	return asc
}

// SetAbbr sets the "abbr" field.
func (asc *AbilityScoreCreate) SetAbbr(a abilityscore.Abbr) *AbilityScoreCreate {
	asc.mutation.SetAbbr(a)
	return asc
}

// SetDesc sets the "desc" field.
func (asc *AbilityScoreCreate) SetDesc(s []string) *AbilityScoreCreate {
	asc.mutation.SetDesc(s)
	return asc
}

// Mutation returns the AbilityScoreMutation object of the builder.
func (asc *AbilityScoreCreate) Mutation() *AbilityScoreMutation {
	return asc.mutation
}

// Save creates the AbilityScore in the database.
func (asc *AbilityScoreCreate) Save(ctx context.Context) (*AbilityScore, error) {
	return withHooks(ctx, asc.sqlSave, asc.mutation, asc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AbilityScoreCreate) SaveX(ctx context.Context) *AbilityScore {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AbilityScoreCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AbilityScoreCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AbilityScoreCreate) check() error {
	if _, ok := asc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "AbilityScore.indx"`)}
	}
	if v, ok := asc.mutation.Indx(); ok {
		if err := abilityscore.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.indx": %w`, err)}
		}
	}
	if _, ok := asc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "AbilityScore.name"`)}
	}
	if v, ok := asc.mutation.Name(); ok {
		if err := abilityscore.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.name": %w`, err)}
		}
	}
	if _, ok := asc.mutation.Abbr(); !ok {
		return &ValidationError{Name: "abbr", err: errors.New(`ent: missing required field "AbilityScore.abbr"`)}
	}
	if v, ok := asc.mutation.Abbr(); ok {
		if err := abilityscore.AbbrValidator(v); err != nil {
			return &ValidationError{Name: "abbr", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.abbr": %w`, err)}
		}
	}
	if _, ok := asc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "AbilityScore.desc"`)}
	}
	return nil
}

func (asc *AbilityScoreCreate) sqlSave(ctx context.Context) (*AbilityScore, error) {
	if err := asc.check(); err != nil {
		return nil, err
	}
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	asc.mutation.id = &_node.ID
	asc.mutation.done = true
	return _node, nil
}

func (asc *AbilityScoreCreate) createSpec() (*AbilityScore, *sqlgraph.CreateSpec) {
	var (
		_node = &AbilityScore{config: asc.config}
		_spec = sqlgraph.NewCreateSpec(abilityscore.Table, sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt))
	)
	if value, ok := asc.mutation.Indx(); ok {
		_spec.SetField(abilityscore.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := asc.mutation.Name(); ok {
		_spec.SetField(abilityscore.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := asc.mutation.Abbr(); ok {
		_spec.SetField(abilityscore.FieldAbbr, field.TypeEnum, value)
		_node.Abbr = value
	}
	if value, ok := asc.mutation.Desc(); ok {
		_spec.SetField(abilityscore.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	return _node, _spec
}

// AbilityScoreCreateBulk is the builder for creating many AbilityScore entities in bulk.
type AbilityScoreCreateBulk struct {
	config
	err      error
	builders []*AbilityScoreCreate
}

// Save creates the AbilityScore entities in the database.
func (ascb *AbilityScoreCreateBulk) Save(ctx context.Context) ([]*AbilityScore, error) {
	if ascb.err != nil {
		return nil, ascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AbilityScore, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AbilityScoreMutation)
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
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AbilityScoreCreateBulk) SaveX(ctx context.Context) []*AbilityScore {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AbilityScoreCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AbilityScoreCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
