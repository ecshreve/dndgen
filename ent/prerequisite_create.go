// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/prerequisite"
)

// PrerequisiteCreate is the builder for creating a Prerequisite entity.
type PrerequisiteCreate struct {
	config
	mutation *PrerequisiteMutation
	hooks    []Hook
}

// SetMinimum sets the "minimum" field.
func (pc *PrerequisiteCreate) SetMinimum(i int) *PrerequisiteCreate {
	pc.mutation.SetMinimum(i)
	return pc
}

// AddAbilityScoreIDs adds the "ability_score" edge to the AbilityScore entity by IDs.
func (pc *PrerequisiteCreate) AddAbilityScoreIDs(ids ...int) *PrerequisiteCreate {
	pc.mutation.AddAbilityScoreIDs(ids...)
	return pc
}

// AddAbilityScore adds the "ability_score" edges to the AbilityScore entity.
func (pc *PrerequisiteCreate) AddAbilityScore(a ...*AbilityScore) *PrerequisiteCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddAbilityScoreIDs(ids...)
}

// Mutation returns the PrerequisiteMutation object of the builder.
func (pc *PrerequisiteCreate) Mutation() *PrerequisiteMutation {
	return pc.mutation
}

// Save creates the Prerequisite in the database.
func (pc *PrerequisiteCreate) Save(ctx context.Context) (*Prerequisite, error) {
	return withHooks[*Prerequisite, PrerequisiteMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PrerequisiteCreate) SaveX(ctx context.Context) *Prerequisite {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PrerequisiteCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PrerequisiteCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PrerequisiteCreate) check() error {
	if _, ok := pc.mutation.Minimum(); !ok {
		return &ValidationError{Name: "minimum", err: errors.New(`ent: missing required field "Prerequisite.minimum"`)}
	}
	return nil
}

func (pc *PrerequisiteCreate) sqlSave(ctx context.Context) (*Prerequisite, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PrerequisiteCreate) createSpec() (*Prerequisite, *sqlgraph.CreateSpec) {
	var (
		_node = &Prerequisite{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(prerequisite.Table, sqlgraph.NewFieldSpec(prerequisite.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Minimum(); ok {
		_spec.SetField(prerequisite.FieldMinimum, field.TypeInt, value)
		_node.Minimum = value
	}
	if nodes := pc.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PrerequisiteCreateBulk is the builder for creating many Prerequisite entities in bulk.
type PrerequisiteCreateBulk struct {
	config
	builders []*PrerequisiteCreate
}

// Save creates the Prerequisite entities in the database.
func (pcb *PrerequisiteCreateBulk) Save(ctx context.Context) ([]*Prerequisite, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Prerequisite, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrerequisiteMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PrerequisiteCreateBulk) SaveX(ctx context.Context) []*Prerequisite {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PrerequisiteCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PrerequisiteCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
