// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/race"
)

// AbilityBonusCreate is the builder for creating a AbilityBonus entity.
type AbilityBonusCreate struct {
	config
	mutation *AbilityBonusMutation
	hooks    []Hook
}

// SetBonus sets the "bonus" field.
func (abc *AbilityBonusCreate) SetBonus(i int) *AbilityBonusCreate {
	abc.mutation.SetBonus(i)
	return abc
}

// SetRaceID sets the "race_id" field.
func (abc *AbilityBonusCreate) SetRaceID(i int) *AbilityBonusCreate {
	abc.mutation.SetRaceID(i)
	return abc
}

// SetAbilityScoreID sets the "ability_score_id" field.
func (abc *AbilityBonusCreate) SetAbilityScoreID(i int) *AbilityBonusCreate {
	abc.mutation.SetAbilityScoreID(i)
	return abc
}

// SetRace sets the "race" edge to the Race entity.
func (abc *AbilityBonusCreate) SetRace(r *Race) *AbilityBonusCreate {
	return abc.SetRaceID(r.ID)
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (abc *AbilityBonusCreate) SetAbilityScore(a *AbilityScore) *AbilityBonusCreate {
	return abc.SetAbilityScoreID(a.ID)
}

// Mutation returns the AbilityBonusMutation object of the builder.
func (abc *AbilityBonusCreate) Mutation() *AbilityBonusMutation {
	return abc.mutation
}

// Save creates the AbilityBonus in the database.
func (abc *AbilityBonusCreate) Save(ctx context.Context) (*AbilityBonus, error) {
	return withHooks(ctx, abc.sqlSave, abc.mutation, abc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (abc *AbilityBonusCreate) SaveX(ctx context.Context) *AbilityBonus {
	v, err := abc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (abc *AbilityBonusCreate) Exec(ctx context.Context) error {
	_, err := abc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abc *AbilityBonusCreate) ExecX(ctx context.Context) {
	if err := abc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (abc *AbilityBonusCreate) check() error {
	if _, ok := abc.mutation.Bonus(); !ok {
		return &ValidationError{Name: "bonus", err: errors.New(`ent: missing required field "AbilityBonus.bonus"`)}
	}
	if v, ok := abc.mutation.Bonus(); ok {
		if err := abilitybonus.BonusValidator(v); err != nil {
			return &ValidationError{Name: "bonus", err: fmt.Errorf(`ent: validator failed for field "AbilityBonus.bonus": %w`, err)}
		}
	}
	if _, ok := abc.mutation.RaceID(); !ok {
		return &ValidationError{Name: "race_id", err: errors.New(`ent: missing required field "AbilityBonus.race_id"`)}
	}
	if _, ok := abc.mutation.AbilityScoreID(); !ok {
		return &ValidationError{Name: "ability_score_id", err: errors.New(`ent: missing required field "AbilityBonus.ability_score_id"`)}
	}
	if len(abc.mutation.RaceIDs()) == 0 {
		return &ValidationError{Name: "race", err: errors.New(`ent: missing required edge "AbilityBonus.race"`)}
	}
	if len(abc.mutation.AbilityScoreIDs()) == 0 {
		return &ValidationError{Name: "ability_score", err: errors.New(`ent: missing required edge "AbilityBonus.ability_score"`)}
	}
	return nil
}

func (abc *AbilityBonusCreate) sqlSave(ctx context.Context) (*AbilityBonus, error) {
	if err := abc.check(); err != nil {
		return nil, err
	}
	_node, _spec := abc.createSpec()
	if err := sqlgraph.CreateNode(ctx, abc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (abc *AbilityBonusCreate) createSpec() (*AbilityBonus, *sqlgraph.CreateSpec) {
	var (
		_node = &AbilityBonus{config: abc.config}
		_spec = sqlgraph.NewCreateSpec(abilitybonus.Table, nil)
	)
	if value, ok := abc.mutation.Bonus(); ok {
		_spec.SetField(abilitybonus.FieldBonus, field.TypeInt, value)
		_node.Bonus = value
	}
	if nodes := abc.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   abilitybonus.RaceTable,
			Columns: []string{abilitybonus.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RaceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := abc.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: []string{abilitybonus.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AbilityScoreID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AbilityBonusCreateBulk is the builder for creating many AbilityBonus entities in bulk.
type AbilityBonusCreateBulk struct {
	config
	err      error
	builders []*AbilityBonusCreate
}

// Save creates the AbilityBonus entities in the database.
func (abcb *AbilityBonusCreateBulk) Save(ctx context.Context) ([]*AbilityBonus, error) {
	if abcb.err != nil {
		return nil, abcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(abcb.builders))
	nodes := make([]*AbilityBonus, len(abcb.builders))
	mutators := make([]Mutator, len(abcb.builders))
	for i := range abcb.builders {
		func(i int, root context.Context) {
			builder := abcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AbilityBonusMutation)
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
					_, err = mutators[i+1].Mutate(root, abcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, abcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, abcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (abcb *AbilityBonusCreateBulk) SaveX(ctx context.Context) []*AbilityBonus {
	v, err := abcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (abcb *AbilityBonusCreateBulk) Exec(ctx context.Context) error {
	_, err := abcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abcb *AbilityBonusCreateBulk) ExecX(ctx context.Context) {
	if err := abcb.Exec(ctx); err != nil {
		panic(err)
	}
}
