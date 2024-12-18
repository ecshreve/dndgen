// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/trait"
)

// TraitCreate is the builder for creating a Trait entity.
type TraitCreate struct {
	config
	mutation *TraitMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (tc *TraitCreate) SetIndx(s string) *TraitCreate {
	tc.mutation.SetIndx(s)
	return tc
}

// SetName sets the "name" field.
func (tc *TraitCreate) SetName(s string) *TraitCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetDesc sets the "desc" field.
func (tc *TraitCreate) SetDesc(s []string) *TraitCreate {
	tc.mutation.SetDesc(s)
	return tc
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (tc *TraitCreate) AddRaceIDs(ids ...int) *TraitCreate {
	tc.mutation.AddRaceIDs(ids...)
	return tc
}

// AddRace adds the "race" edges to the Race entity.
func (tc *TraitCreate) AddRace(r ...*Race) *TraitCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tc.AddRaceIDs(ids...)
}

// Mutation returns the TraitMutation object of the builder.
func (tc *TraitCreate) Mutation() *TraitMutation {
	return tc.mutation
}

// Save creates the Trait in the database.
func (tc *TraitCreate) Save(ctx context.Context) (*Trait, error) {
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TraitCreate) SaveX(ctx context.Context) *Trait {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TraitCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TraitCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TraitCreate) check() error {
	if _, ok := tc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Trait.indx"`)}
	}
	if v, ok := tc.mutation.Indx(); ok {
		if err := trait.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Trait.indx": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Trait.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := trait.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Trait.name": %w`, err)}
		}
	}
	return nil
}

func (tc *TraitCreate) sqlSave(ctx context.Context) (*Trait, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TraitCreate) createSpec() (*Trait, *sqlgraph.CreateSpec) {
	var (
		_node = &Trait{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(trait.Table, sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Indx(); ok {
		_spec.SetField(trait.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(trait.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Desc(); ok {
		_spec.SetField(trait.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if nodes := tc.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RaceTable,
			Columns: trait.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TraitCreateBulk is the builder for creating many Trait entities in bulk.
type TraitCreateBulk struct {
	config
	err      error
	builders []*TraitCreate
}

// Save creates the Trait entities in the database.
func (tcb *TraitCreateBulk) Save(ctx context.Context) ([]*Trait, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Trait, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TraitMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TraitCreateBulk) SaveX(ctx context.Context) []*Trait {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TraitCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TraitCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
