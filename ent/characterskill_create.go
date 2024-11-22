// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterproficiency"
	"github.com/ecshreve/dndgen/ent/characterskill"
	"github.com/ecshreve/dndgen/ent/skill"
)

// CharacterSkillCreate is the builder for creating a CharacterSkill entity.
type CharacterSkillCreate struct {
	config
	mutation *CharacterSkillMutation
	hooks    []Hook
}

// SetProficient sets the "proficient" field.
func (csc *CharacterSkillCreate) SetProficient(b bool) *CharacterSkillCreate {
	csc.mutation.SetProficient(b)
	return csc
}

// SetNillableProficient sets the "proficient" field if the given value is not nil.
func (csc *CharacterSkillCreate) SetNillableProficient(b *bool) *CharacterSkillCreate {
	if b != nil {
		csc.SetProficient(*b)
	}
	return csc
}

// SetCharacterID sets the "character" edge to the Character entity by ID.
func (csc *CharacterSkillCreate) SetCharacterID(id int) *CharacterSkillCreate {
	csc.mutation.SetCharacterID(id)
	return csc
}

// SetCharacter sets the "character" edge to the Character entity.
func (csc *CharacterSkillCreate) SetCharacter(c *Character) *CharacterSkillCreate {
	return csc.SetCharacterID(c.ID)
}

// SetSkillID sets the "skill" edge to the Skill entity by ID.
func (csc *CharacterSkillCreate) SetSkillID(id int) *CharacterSkillCreate {
	csc.mutation.SetSkillID(id)
	return csc
}

// SetSkill sets the "skill" edge to the Skill entity.
func (csc *CharacterSkillCreate) SetSkill(s *Skill) *CharacterSkillCreate {
	return csc.SetSkillID(s.ID)
}

// SetCharacterAbilityScoreID sets the "character_ability_score" edge to the CharacterAbilityScore entity by ID.
func (csc *CharacterSkillCreate) SetCharacterAbilityScoreID(id int) *CharacterSkillCreate {
	csc.mutation.SetCharacterAbilityScoreID(id)
	return csc
}

// SetCharacterAbilityScore sets the "character_ability_score" edge to the CharacterAbilityScore entity.
func (csc *CharacterSkillCreate) SetCharacterAbilityScore(c *CharacterAbilityScore) *CharacterSkillCreate {
	return csc.SetCharacterAbilityScoreID(c.ID)
}

// SetCharacterProficiencyID sets the "character_proficiency" edge to the CharacterProficiency entity by ID.
func (csc *CharacterSkillCreate) SetCharacterProficiencyID(id int) *CharacterSkillCreate {
	csc.mutation.SetCharacterProficiencyID(id)
	return csc
}

// SetNillableCharacterProficiencyID sets the "character_proficiency" edge to the CharacterProficiency entity by ID if the given value is not nil.
func (csc *CharacterSkillCreate) SetNillableCharacterProficiencyID(id *int) *CharacterSkillCreate {
	if id != nil {
		csc = csc.SetCharacterProficiencyID(*id)
	}
	return csc
}

// SetCharacterProficiency sets the "character_proficiency" edge to the CharacterProficiency entity.
func (csc *CharacterSkillCreate) SetCharacterProficiency(c *CharacterProficiency) *CharacterSkillCreate {
	return csc.SetCharacterProficiencyID(c.ID)
}

// Mutation returns the CharacterSkillMutation object of the builder.
func (csc *CharacterSkillCreate) Mutation() *CharacterSkillMutation {
	return csc.mutation
}

// Save creates the CharacterSkill in the database.
func (csc *CharacterSkillCreate) Save(ctx context.Context) (*CharacterSkill, error) {
	csc.defaults()
	return withHooks(ctx, csc.sqlSave, csc.mutation, csc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (csc *CharacterSkillCreate) SaveX(ctx context.Context) *CharacterSkill {
	v, err := csc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (csc *CharacterSkillCreate) Exec(ctx context.Context) error {
	_, err := csc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (csc *CharacterSkillCreate) ExecX(ctx context.Context) {
	if err := csc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (csc *CharacterSkillCreate) defaults() {
	if _, ok := csc.mutation.Proficient(); !ok {
		v := characterskill.DefaultProficient
		csc.mutation.SetProficient(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (csc *CharacterSkillCreate) check() error {
	if _, ok := csc.mutation.Proficient(); !ok {
		return &ValidationError{Name: "proficient", err: errors.New(`ent: missing required field "CharacterSkill.proficient"`)}
	}
	if len(csc.mutation.CharacterIDs()) == 0 {
		return &ValidationError{Name: "character", err: errors.New(`ent: missing required edge "CharacterSkill.character"`)}
	}
	if len(csc.mutation.SkillIDs()) == 0 {
		return &ValidationError{Name: "skill", err: errors.New(`ent: missing required edge "CharacterSkill.skill"`)}
	}
	if len(csc.mutation.CharacterAbilityScoreIDs()) == 0 {
		return &ValidationError{Name: "character_ability_score", err: errors.New(`ent: missing required edge "CharacterSkill.character_ability_score"`)}
	}
	return nil
}

func (csc *CharacterSkillCreate) sqlSave(ctx context.Context) (*CharacterSkill, error) {
	if err := csc.check(); err != nil {
		return nil, err
	}
	_node, _spec := csc.createSpec()
	if err := sqlgraph.CreateNode(ctx, csc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	csc.mutation.id = &_node.ID
	csc.mutation.done = true
	return _node, nil
}

func (csc *CharacterSkillCreate) createSpec() (*CharacterSkill, *sqlgraph.CreateSpec) {
	var (
		_node = &CharacterSkill{config: csc.config}
		_spec = sqlgraph.NewCreateSpec(characterskill.Table, sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt))
	)
	if value, ok := csc.mutation.Proficient(); ok {
		_spec.SetField(characterskill.FieldProficient, field.TypeBool, value)
		_node.Proficient = value
	}
	if nodes := csc.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   characterskill.CharacterTable,
			Columns: []string{characterskill.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_character_skills = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csc.mutation.SkillIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterskill.SkillTable,
			Columns: []string{characterskill.SkillColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_skill_skill = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csc.mutation.CharacterAbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterskill.CharacterAbilityScoreTable,
			Columns: []string{characterskill.CharacterAbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterabilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_skill_character_ability_score = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := csc.mutation.CharacterProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   characterskill.CharacterProficiencyTable,
			Columns: []string{characterskill.CharacterProficiencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterproficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CharacterSkillCreateBulk is the builder for creating many CharacterSkill entities in bulk.
type CharacterSkillCreateBulk struct {
	config
	err      error
	builders []*CharacterSkillCreate
}

// Save creates the CharacterSkill entities in the database.
func (cscb *CharacterSkillCreateBulk) Save(ctx context.Context) ([]*CharacterSkill, error) {
	if cscb.err != nil {
		return nil, cscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cscb.builders))
	nodes := make([]*CharacterSkill, len(cscb.builders))
	mutators := make([]Mutator, len(cscb.builders))
	for i := range cscb.builders {
		func(i int, root context.Context) {
			builder := cscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterSkillMutation)
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
					_, err = mutators[i+1].Mutate(root, cscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cscb *CharacterSkillCreateBulk) SaveX(ctx context.Context) []*CharacterSkill {
	v, err := cscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cscb *CharacterSkillCreateBulk) Exec(ctx context.Context) error {
	_, err := cscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cscb *CharacterSkillCreateBulk) ExecX(ctx context.Context) {
	if err := cscb.Exec(ctx); err != nil {
		panic(err)
	}
}
