// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
	"github.com/ecshreve/dndgen/ent/race"
)

// ProficiencyCreate is the builder for creating a Proficiency entity.
type ProficiencyCreate struct {
	config
	mutation *ProficiencyMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (pc *ProficiencyCreate) SetIndx(s string) *ProficiencyCreate {
	pc.mutation.SetIndx(s)
	return pc
}

// SetName sets the "name" field.
func (pc *ProficiencyCreate) SetName(s string) *ProficiencyCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetReference sets the "reference" field.
func (pc *ProficiencyCreate) SetReference(s string) *ProficiencyCreate {
	pc.mutation.SetReference(s)
	return pc
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (pc *ProficiencyCreate) AddRaceIDs(ids ...int) *ProficiencyCreate {
	pc.mutation.AddRaceIDs(ids...)
	return pc
}

// AddRace adds the "race" edges to the Race entity.
func (pc *ProficiencyCreate) AddRace(r ...*Race) *ProficiencyCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pc.AddRaceIDs(ids...)
}

// AddOptionIDs adds the "options" edge to the ProficiencyChoice entity by IDs.
func (pc *ProficiencyCreate) AddOptionIDs(ids ...int) *ProficiencyCreate {
	pc.mutation.AddOptionIDs(ids...)
	return pc
}

// AddOptions adds the "options" edges to the ProficiencyChoice entity.
func (pc *ProficiencyCreate) AddOptions(p ...*ProficiencyChoice) *ProficiencyCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddOptionIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (pc *ProficiencyCreate) AddClasIDs(ids ...int) *ProficiencyCreate {
	pc.mutation.AddClasIDs(ids...)
	return pc
}

// AddClass adds the "class" edges to the Class entity.
func (pc *ProficiencyCreate) AddClass(c ...*Class) *ProficiencyCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddClasIDs(ids...)
}

// AddCharacterIDs adds the "character" edge to the Character entity by IDs.
func (pc *ProficiencyCreate) AddCharacterIDs(ids ...int) *ProficiencyCreate {
	pc.mutation.AddCharacterIDs(ids...)
	return pc
}

// AddCharacter adds the "character" edges to the Character entity.
func (pc *ProficiencyCreate) AddCharacter(c ...*Character) *ProficiencyCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCharacterIDs(ids...)
}

// Mutation returns the ProficiencyMutation object of the builder.
func (pc *ProficiencyCreate) Mutation() *ProficiencyMutation {
	return pc.mutation
}

// Save creates the Proficiency in the database.
func (pc *ProficiencyCreate) Save(ctx context.Context) (*Proficiency, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProficiencyCreate) SaveX(ctx context.Context) *Proficiency {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProficiencyCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProficiencyCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProficiencyCreate) check() error {
	if _, ok := pc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Proficiency.indx"`)}
	}
	if v, ok := pc.mutation.Indx(); ok {
		if err := proficiency.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Proficiency.indx": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Proficiency.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := proficiency.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Proficiency.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Reference(); !ok {
		return &ValidationError{Name: "reference", err: errors.New(`ent: missing required field "Proficiency.reference"`)}
	}
	if v, ok := pc.mutation.Reference(); ok {
		if err := proficiency.ReferenceValidator(v); err != nil {
			return &ValidationError{Name: "reference", err: fmt.Errorf(`ent: validator failed for field "Proficiency.reference": %w`, err)}
		}
	}
	return nil
}

func (pc *ProficiencyCreate) sqlSave(ctx context.Context) (*Proficiency, error) {
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

func (pc *ProficiencyCreate) createSpec() (*Proficiency, *sqlgraph.CreateSpec) {
	var (
		_node = &Proficiency{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(proficiency.Table, sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Indx(); ok {
		_spec.SetField(proficiency.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(proficiency.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Reference(); ok {
		_spec.SetField(proficiency.FieldReference, field.TypeString, value)
		_node.Reference = value
	}
	if nodes := pc.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiency.RaceTable,
			Columns: proficiency.RacePrimaryKey,
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
	if nodes := pc.mutation.OptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiency.OptionsTable,
			Columns: proficiency.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiency.ClassTable,
			Columns: proficiency.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiency.CharacterTable,
			Columns: proficiency.CharacterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProficiencyCreateBulk is the builder for creating many Proficiency entities in bulk.
type ProficiencyCreateBulk struct {
	config
	err      error
	builders []*ProficiencyCreate
}

// Save creates the Proficiency entities in the database.
func (pcb *ProficiencyCreateBulk) Save(ctx context.Context) ([]*Proficiency, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Proficiency, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProficiencyMutation)
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
func (pcb *ProficiencyCreateBulk) SaveX(ctx context.Context) []*Proficiency {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProficiencyCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProficiencyCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
