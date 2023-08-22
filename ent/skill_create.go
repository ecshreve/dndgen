// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/skill"
)

// SkillCreate is the builder for creating a Skill entity.
type SkillCreate struct {
	config
	mutation *SkillMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (sc *SkillCreate) SetIndx(s string) *SkillCreate {
	sc.mutation.SetIndx(s)
	return sc
}

// SetName sets the "name" field.
func (sc *SkillCreate) SetName(s string) *SkillCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDesc sets the "desc" field.
func (sc *SkillCreate) SetDesc(s []string) *SkillCreate {
	sc.mutation.SetDesc(s)
	return sc
}

// SetAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID.
func (sc *SkillCreate) SetAbilityScoreID(id int) *SkillCreate {
	sc.mutation.SetAbilityScoreID(id)
	return sc
}

// SetNillableAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID if the given value is not nil.
func (sc *SkillCreate) SetNillableAbilityScoreID(id *int) *SkillCreate {
	if id != nil {
		sc = sc.SetAbilityScoreID(*id)
	}
	return sc
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (sc *SkillCreate) SetAbilityScore(a *AbilityScore) *SkillCreate {
	return sc.SetAbilityScoreID(a.ID)
}

// Mutation returns the SkillMutation object of the builder.
func (sc *SkillCreate) Mutation() *SkillMutation {
	return sc.mutation
}

// Save creates the Skill in the database.
func (sc *SkillCreate) Save(ctx context.Context) (*Skill, error) {
	return withHooks[*Skill, SkillMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SkillCreate) SaveX(ctx context.Context) *Skill {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SkillCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SkillCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SkillCreate) check() error {
	if _, ok := sc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Skill.indx"`)}
	}
	if v, ok := sc.mutation.Indx(); ok {
		if err := skill.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Skill.indx": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Skill.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := skill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Skill.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "Skill.desc"`)}
	}
	return nil
}

func (sc *SkillCreate) sqlSave(ctx context.Context) (*Skill, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SkillCreate) createSpec() (*Skill, *sqlgraph.CreateSpec) {
	var (
		_node = &Skill{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(skill.Table, sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Indx(); ok {
		_spec.SetField(skill.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(skill.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Desc(); ok {
		_spec.SetField(skill.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if nodes := sc.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   skill.AbilityScoreTable,
			Columns: []string{skill.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.skill_ability_score = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SkillCreateBulk is the builder for creating many Skill entities in bulk.
type SkillCreateBulk struct {
	config
	builders []*SkillCreate
}

// Save creates the Skill entities in the database.
func (scb *SkillCreateBulk) Save(ctx context.Context) ([]*Skill, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Skill, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SkillMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SkillCreateBulk) SaveX(ctx context.Context) []*Skill {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SkillCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SkillCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
