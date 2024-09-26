// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterskill"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/skill"
)

// CharacterCreate is the builder for creating a Character entity.
type CharacterCreate struct {
	config
	mutation *CharacterMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CharacterCreate) SetName(s string) *CharacterCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetAge sets the "age" field.
func (cc *CharacterCreate) SetAge(i int) *CharacterCreate {
	cc.mutation.SetAge(i)
	return cc
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (cc *CharacterCreate) SetNillableAge(i *int) *CharacterCreate {
	if i != nil {
		cc.SetAge(*i)
	}
	return cc
}

// SetLevel sets the "level" field.
func (cc *CharacterCreate) SetLevel(i int) *CharacterCreate {
	cc.mutation.SetLevel(i)
	return cc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (cc *CharacterCreate) SetNillableLevel(i *int) *CharacterCreate {
	if i != nil {
		cc.SetLevel(*i)
	}
	return cc
}

// SetProficiencyBonus sets the "proficiency_bonus" field.
func (cc *CharacterCreate) SetProficiencyBonus(i int) *CharacterCreate {
	cc.mutation.SetProficiencyBonus(i)
	return cc
}

// SetNillableProficiencyBonus sets the "proficiency_bonus" field if the given value is not nil.
func (cc *CharacterCreate) SetNillableProficiencyBonus(i *int) *CharacterCreate {
	if i != nil {
		cc.SetProficiencyBonus(*i)
	}
	return cc
}

// SetRaceID sets the "race" edge to the Race entity by ID.
func (cc *CharacterCreate) SetRaceID(id int) *CharacterCreate {
	cc.mutation.SetRaceID(id)
	return cc
}

// SetNillableRaceID sets the "race" edge to the Race entity by ID if the given value is not nil.
func (cc *CharacterCreate) SetNillableRaceID(id *int) *CharacterCreate {
	if id != nil {
		cc = cc.SetRaceID(*id)
	}
	return cc
}

// SetRace sets the "race" edge to the Race entity.
func (cc *CharacterCreate) SetRace(r *Race) *CharacterCreate {
	return cc.SetRaceID(r.ID)
}

// SetClassID sets the "class" edge to the Class entity by ID.
func (cc *CharacterCreate) SetClassID(id int) *CharacterCreate {
	cc.mutation.SetClassID(id)
	return cc
}

// SetNillableClassID sets the "class" edge to the Class entity by ID if the given value is not nil.
func (cc *CharacterCreate) SetNillableClassID(id *int) *CharacterCreate {
	if id != nil {
		cc = cc.SetClassID(*id)
	}
	return cc
}

// SetClass sets the "class" edge to the Class entity.
func (cc *CharacterCreate) SetClass(c *Class) *CharacterCreate {
	return cc.SetClassID(c.ID)
}

// SetAlignmentID sets the "alignment" edge to the Alignment entity by ID.
func (cc *CharacterCreate) SetAlignmentID(id int) *CharacterCreate {
	cc.mutation.SetAlignmentID(id)
	return cc
}

// SetNillableAlignmentID sets the "alignment" edge to the Alignment entity by ID if the given value is not nil.
func (cc *CharacterCreate) SetNillableAlignmentID(id *int) *CharacterCreate {
	if id != nil {
		cc = cc.SetAlignmentID(*id)
	}
	return cc
}

// SetAlignment sets the "alignment" edge to the Alignment entity.
func (cc *CharacterCreate) SetAlignment(a *Alignment) *CharacterCreate {
	return cc.SetAlignmentID(a.ID)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (cc *CharacterCreate) AddProficiencyIDs(ids ...int) *CharacterCreate {
	cc.mutation.AddProficiencyIDs(ids...)
	return cc
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (cc *CharacterCreate) AddProficiencies(p ...*Proficiency) *CharacterCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProficiencyIDs(ids...)
}

// AddAbilityScoreIDs adds the "ability_scores" edge to the AbilityScore entity by IDs.
func (cc *CharacterCreate) AddAbilityScoreIDs(ids ...int) *CharacterCreate {
	cc.mutation.AddAbilityScoreIDs(ids...)
	return cc
}

// AddAbilityScores adds the "ability_scores" edges to the AbilityScore entity.
func (cc *CharacterCreate) AddAbilityScores(a ...*AbilityScore) *CharacterCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cc.AddAbilityScoreIDs(ids...)
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (cc *CharacterCreate) AddSkillIDs(ids ...int) *CharacterCreate {
	cc.mutation.AddSkillIDs(ids...)
	return cc
}

// AddSkills adds the "skills" edges to the Skill entity.
func (cc *CharacterCreate) AddSkills(s ...*Skill) *CharacterCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSkillIDs(ids...)
}

// AddCharacterAbilityScoreIDs adds the "character_ability_scores" edge to the CharacterAbilityScore entity by IDs.
func (cc *CharacterCreate) AddCharacterAbilityScoreIDs(ids ...int) *CharacterCreate {
	cc.mutation.AddCharacterAbilityScoreIDs(ids...)
	return cc
}

// AddCharacterAbilityScores adds the "character_ability_scores" edges to the CharacterAbilityScore entity.
func (cc *CharacterCreate) AddCharacterAbilityScores(c ...*CharacterAbilityScore) *CharacterCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCharacterAbilityScoreIDs(ids...)
}

// AddCharacterSkillIDs adds the "character_skills" edge to the CharacterSkill entity by IDs.
func (cc *CharacterCreate) AddCharacterSkillIDs(ids ...int) *CharacterCreate {
	cc.mutation.AddCharacterSkillIDs(ids...)
	return cc
}

// AddCharacterSkills adds the "character_skills" edges to the CharacterSkill entity.
func (cc *CharacterCreate) AddCharacterSkills(c ...*CharacterSkill) *CharacterCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCharacterSkillIDs(ids...)
}

// Mutation returns the CharacterMutation object of the builder.
func (cc *CharacterCreate) Mutation() *CharacterMutation {
	return cc.mutation
}

// Save creates the Character in the database.
func (cc *CharacterCreate) Save(ctx context.Context) (*Character, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CharacterCreate) SaveX(ctx context.Context) *Character {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CharacterCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CharacterCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CharacterCreate) defaults() {
	if _, ok := cc.mutation.Age(); !ok {
		v := character.DefaultAge
		cc.mutation.SetAge(v)
	}
	if _, ok := cc.mutation.Level(); !ok {
		v := character.DefaultLevel
		cc.mutation.SetLevel(v)
	}
	if _, ok := cc.mutation.ProficiencyBonus(); !ok {
		v := character.DefaultProficiencyBonus
		cc.mutation.SetProficiencyBonus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CharacterCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Character.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := character.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Character.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Character.age"`)}
	}
	if v, ok := cc.mutation.Age(); ok {
		if err := character.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "Character.age": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Character.level"`)}
	}
	if v, ok := cc.mutation.Level(); ok {
		if err := character.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Character.level": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ProficiencyBonus(); !ok {
		return &ValidationError{Name: "proficiency_bonus", err: errors.New(`ent: missing required field "Character.proficiency_bonus"`)}
	}
	if v, ok := cc.mutation.ProficiencyBonus(); ok {
		if err := character.ProficiencyBonusValidator(v); err != nil {
			return &ValidationError{Name: "proficiency_bonus", err: fmt.Errorf(`ent: validator failed for field "Character.proficiency_bonus": %w`, err)}
		}
	}
	return nil
}

func (cc *CharacterCreate) sqlSave(ctx context.Context) (*Character, error) {
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

func (cc *CharacterCreate) createSpec() (*Character, *sqlgraph.CreateSpec) {
	var (
		_node = &Character{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(character.Table, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(character.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Age(); ok {
		_spec.SetField(character.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := cc.mutation.Level(); ok {
		_spec.SetField(character.FieldLevel, field.TypeInt, value)
		_node.Level = value
	}
	if value, ok := cc.mutation.ProficiencyBonus(); ok {
		_spec.SetField(character.FieldProficiencyBonus, field.TypeInt, value)
		_node.ProficiencyBonus = value
	}
	if nodes := cc.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   character.RaceTable,
			Columns: []string{character.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_race = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   character.ClassTable,
			Columns: []string{character.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_class = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AlignmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   character.AlignmentTable,
			Columns: []string{character.AlignmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(alignment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.character_alignment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.ProficienciesTable,
			Columns: character.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AbilityScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.AbilityScoresTable,
			Columns: character.AbilityScoresPrimaryKey,
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
	if nodes := cc.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   character.SkillsTable,
			Columns: character.SkillsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &CharacterSkillCreate{config: cc.config, mutation: newCharacterSkillMutation(cc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CharacterAbilityScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.CharacterAbilityScoresTable,
			Columns: []string{character.CharacterAbilityScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterabilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CharacterSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   character.CharacterSkillsTable,
			Columns: []string{character.CharacterSkillsColumn},
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

// CharacterCreateBulk is the builder for creating many Character entities in bulk.
type CharacterCreateBulk struct {
	config
	err      error
	builders []*CharacterCreate
}

// Save creates the Character entities in the database.
func (ccb *CharacterCreateBulk) Save(ctx context.Context) ([]*Character, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Character, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CharacterMutation)
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
func (ccb *CharacterCreateBulk) SaveX(ctx context.Context) []*Character {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CharacterCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CharacterCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
