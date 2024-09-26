// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterproficiency"
	"github.com/ecshreve/dndgen/ent/proficiency"
)

// CharacterProficiencyCreate is the builder for creating a CharacterProficiency entity.
type CharacterProficiencyCreate struct {
	config
	mutation *CharacterProficiencyMutation
	hooks    []Hook
}

// SetCharacterID sets the "character_id" field.
func (cpc *CharacterProficiencyCreate) SetCharacterID(i int) *CharacterProficiencyCreate {
	cpc.mutation.SetCharacterID(i)
	return cpc
}

// SetProficiencyID sets the "proficiency_id" field.
func (cpc *CharacterProficiencyCreate) SetProficiencyID(i int) *CharacterProficiencyCreate {
	cpc.mutation.SetProficiencyID(i)
	return cpc
}

// SetCharacter sets the "character" edge to the Character entity.
func (cpc *CharacterProficiencyCreate) SetCharacter(c *Character) *CharacterProficiencyCreate {
	return cpc.SetCharacterID(c.ID)
}

// SetProficiency sets the "proficiency" edge to the Proficiency entity.
func (cpc *CharacterProficiencyCreate) SetProficiency(p *Proficiency) *CharacterProficiencyCreate {
	return cpc.SetProficiencyID(p.ID)
}

// Mutation returns the CharacterProficiencyMutation object of the builder.
func (cpc *CharacterProficiencyCreate) Mutation() *CharacterProficiencyMutation {
	return cpc.mutation
}

// Save creates the CharacterProficiency in the database.
func (cpc *CharacterProficiencyCreate) Save(ctx context.Context) (*CharacterProficiency, error) {
	return withHooks(ctx, cpc.sqlSave, cpc.mutation, cpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *CharacterProficiencyCreate) SaveX(ctx context.Context) *CharacterProficiency {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *CharacterProficiencyCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *CharacterProficiencyCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *CharacterProficiencyCreate) check() error {
	if _, ok := cpc.mutation.CharacterID(); !ok {
		return &ValidationError{Name: "character_id", err: errors.New(`ent: missing required field "CharacterProficiency.character_id"`)}
	}
	if _, ok := cpc.mutation.ProficiencyID(); !ok {
		return &ValidationError{Name: "proficiency_id", err: errors.New(`ent: missing required field "CharacterProficiency.proficiency_id"`)}
	}
	if len(cpc.mutation.CharacterIDs()) == 0 {
		return &ValidationError{Name: "character", err: errors.New(`ent: missing required edge "CharacterProficiency.character"`)}
	}
	if len(cpc.mutation.ProficiencyIDs()) == 0 {
		return &ValidationError{Name: "proficiency", err: errors.New(`ent: missing required edge "CharacterProficiency.proficiency"`)}
	}
	return nil
}

func (cpc *CharacterProficiencyCreate) sqlSave(ctx context.Context) (*CharacterProficiency, error) {
	if err := cpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (cpc *CharacterProficiencyCreate) createSpec() (*CharacterProficiency, *sqlgraph.CreateSpec) {
	var (
		_node = &CharacterProficiency{config: cpc.config}
		_spec = sqlgraph.NewCreateSpec(characterproficiency.Table, nil)
	)
	if nodes := cpc.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.CharacterTable,
			Columns: []string{characterproficiency.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CharacterID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cpc.mutation.ProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.ProficiencyTable,
			Columns: []string{characterproficiency.ProficiencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProficiencyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CharacterProficiencyCreateBulk is the builder for creating many CharacterProficiency entities in bulk.
type CharacterProficiencyCreateBulk struct {
	config
	err      error
	builders []*CharacterProficiencyCreate
}

// Save creates the CharacterProficiency entities in the database.
func (cpcb *CharacterProficiencyCreateBulk) Save(ctx context.Context) ([]*CharacterProficiency, error) {
	if cpcb.err != nil {
		return nil, cpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*CharacterProficiency, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterProficiencyMutation)
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
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *CharacterProficiencyCreateBulk) SaveX(ctx context.Context) []*CharacterProficiency {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *CharacterProficiencyCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *CharacterProficiencyCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}