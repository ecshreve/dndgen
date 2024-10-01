// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterskill"
)

// CharacterAbilityScoreCreate is the builder for creating a CharacterAbilityScore entity.
type CharacterAbilityScoreCreate struct {
	config
	mutation *CharacterAbilityScoreMutation
	hooks    []Hook
}

// SetScore sets the "score" field.
func (casc *CharacterAbilityScoreCreate) SetScore(i int) *CharacterAbilityScoreCreate {
	casc.mutation.SetScore(i)
	return casc
}

// SetModifier sets the "modifier" field.
func (casc *CharacterAbilityScoreCreate) SetModifier(i int) *CharacterAbilityScoreCreate {
	casc.mutation.SetModifier(i)
	return casc
}

// SetCharacterID sets the "character" edge to the Character entity by ID.
func (casc *CharacterAbilityScoreCreate) SetCharacterID(id int) *CharacterAbilityScoreCreate {
	casc.mutation.SetCharacterID(id)
	return casc
}

// SetCharacter sets the "character" edge to the Character entity.
func (casc *CharacterAbilityScoreCreate) SetCharacter(c *Character) *CharacterAbilityScoreCreate {
	return casc.SetCharacterID(c.ID)
}

// SetAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID.
func (casc *CharacterAbilityScoreCreate) SetAbilityScoreID(id int) *CharacterAbilityScoreCreate {
	casc.mutation.SetAbilityScoreID(id)
	return casc
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (casc *CharacterAbilityScoreCreate) SetAbilityScore(a *AbilityScore) *CharacterAbilityScoreCreate {
	return casc.SetAbilityScoreID(a.ID)
}

// AddCharacterSkillIDs adds the "character_skills" edge to the CharacterSkill entity by IDs.
func (casc *CharacterAbilityScoreCreate) AddCharacterSkillIDs(ids ...int) *CharacterAbilityScoreCreate {
	casc.mutation.AddCharacterSkillIDs(ids...)
	return casc
}

// AddCharacterSkills adds the "character_skills" edges to the CharacterSkill entity.
func (casc *CharacterAbilityScoreCreate) AddCharacterSkills(c ...*CharacterSkill) *CharacterAbilityScoreCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return casc.AddCharacterSkillIDs(ids...)
}

// Mutation returns the CharacterAbilityScoreMutation object of the builder.
func (casc *CharacterAbilityScoreCreate) Mutation() *CharacterAbilityScoreMutation {
	return casc.mutation
}

// Save creates the CharacterAbilityScore in the database.
func (casc *CharacterAbilityScoreCreate) Save(ctx context.Context) (*CharacterAbilityScore, error) {
	return withHooks(ctx, casc.sqlSave, casc.mutation, casc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (casc *CharacterAbilityScoreCreate) SaveX(ctx context.Context) *CharacterAbilityScore {
	v, err := casc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (casc *CharacterAbilityScoreCreate) Exec(ctx context.Context) error {
	_, err := casc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (casc *CharacterAbilityScoreCreate) ExecX(ctx context.Context) {
	if err := casc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (casc *CharacterAbilityScoreCreate) check() error {
	if _, ok := casc.mutation.Score(); !ok {
		return &ValidationError{Name: "score", err: errors.New(`ent: missing required field "CharacterAbilityScore.score"`)}
	}
	if v, ok := casc.mutation.Score(); ok {
		if err := characterabilityscore.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "CharacterAbilityScore.score": %w`, err)}
		}
	}
	if _, ok := casc.mutation.Modifier(); !ok {
		return &ValidationError{Name: "modifier", err: errors.New(`ent: missing required field "CharacterAbilityScore.modifier"`)}
	}
	if v, ok := casc.mutation.Modifier(); ok {
		if err := characterabilityscore.ModifierValidator(v); err != nil {
			return &ValidationError{Name: "modifier", err: fmt.Errorf(`ent: validator failed for field "CharacterAbilityScore.modifier": %w`, err)}
		}
	}
	if len(casc.mutation.CharacterIDs()) == 0 {
		return &ValidationError{Name: "character", err: errors.New(`ent: missing required edge "CharacterAbilityScore.character"`)}
	}
	if len(casc.mutation.AbilityScoreIDs()) == 0 {
		return &ValidationError{Name: "ability_score", err: errors.New(`ent: missing required edge "CharacterAbilityScore.ability_score"`)}
	}
	return nil
}

func (casc *CharacterAbilityScoreCreate) sqlSave(ctx context.Context) (*CharacterAbilityScore, error) {
	if err := casc.check(); err != nil {
		return nil, err
	}
	_node, _spec := casc.createSpec()
	if err := sqlgraph.CreateNode(ctx, casc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	casc.mutation.id = &_node.ID
	casc.mutation.done = true
	return _node, nil
}

func (casc *CharacterAbilityScoreCreate) createSpec() (*CharacterAbilityScore, *sqlgraph.CreateSpec) {
	var (
		_node = &CharacterAbilityScore{config: casc.config}
		_spec = sqlgraph.NewCreateSpec(characterabilityscore.Table, sqlgraph.NewFieldSpec(characterabilityscore.FieldID, field.TypeInt))
	)
	if value, ok := casc.mutation.Score(); ok {
		_spec.SetField(characterabilityscore.FieldScore, field.TypeInt, value)
		_node.Score = value
	}
	if value, ok := casc.mutation.Modifier(); ok {
		_spec.SetField(characterabilityscore.FieldModifier, field.TypeInt, value)
		_node.Modifier = value
	}
	if nodes := casc.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   characterabilityscore.CharacterTable,
			Columns: []string{characterabilityscore.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_character_ability_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := casc.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterabilityscore.AbilityScoreTable,
			Columns: []string{characterabilityscore.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_ability_score_ability_score = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := casc.mutation.CharacterSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CharacterAbilityScoreCreateBulk is the builder for creating many CharacterAbilityScore entities in bulk.
type CharacterAbilityScoreCreateBulk struct {
	config
	err      error
	builders []*CharacterAbilityScoreCreate
}

// Save creates the CharacterAbilityScore entities in the database.
func (cascb *CharacterAbilityScoreCreateBulk) Save(ctx context.Context) ([]*CharacterAbilityScore, error) {
	if cascb.err != nil {
		return nil, cascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cascb.builders))
	nodes := make([]*CharacterAbilityScore, len(cascb.builders))
	mutators := make([]Mutator, len(cascb.builders))
	for i := range cascb.builders {
		func(i int, root context.Context) {
			builder := cascb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterAbilityScoreMutation)
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
					_, err = mutators[i+1].Mutate(root, cascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cascb *CharacterAbilityScoreCreateBulk) SaveX(ctx context.Context) []*CharacterAbilityScore {
	v, err := cascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cascb *CharacterAbilityScoreCreateBulk) Exec(ctx context.Context) error {
	_, err := cascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cascb *CharacterAbilityScoreCreateBulk) ExecX(ctx context.Context) {
	if err := cascb.Exec(ctx); err != nil {
		panic(err)
	}
}
