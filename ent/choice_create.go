// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/choice"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/race"
)

// ChoiceCreate is the builder for creating a Choice entity.
type ChoiceCreate struct {
	config
	mutation *ChoiceMutation
	hooks    []Hook
}

// SetChoose sets the "choose" field.
func (cc *ChoiceCreate) SetChoose(i int) *ChoiceCreate {
	cc.mutation.SetChoose(i)
	return cc
}

// SetDesc sets the "desc" field.
func (cc *ChoiceCreate) SetDesc(s string) *ChoiceCreate {
	cc.mutation.SetDesc(s)
	return cc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (cc *ChoiceCreate) SetNillableDesc(s *string) *ChoiceCreate {
	if s != nil {
		cc.SetDesc(*s)
	}
	return cc
}

// SetParentChoiceID sets the "parent_choice" edge to the Choice entity by ID.
func (cc *ChoiceCreate) SetParentChoiceID(id int) *ChoiceCreate {
	cc.mutation.SetParentChoiceID(id)
	return cc
}

// SetNillableParentChoiceID sets the "parent_choice" edge to the Choice entity by ID if the given value is not nil.
func (cc *ChoiceCreate) SetNillableParentChoiceID(id *int) *ChoiceCreate {
	if id != nil {
		cc = cc.SetParentChoiceID(*id)
	}
	return cc
}

// SetParentChoice sets the "parent_choice" edge to the Choice entity.
func (cc *ChoiceCreate) SetParentChoice(c *Choice) *ChoiceCreate {
	return cc.SetParentChoiceID(c.ID)
}

// AddChoiceIDs adds the "choices" edge to the Choice entity by IDs.
func (cc *ChoiceCreate) AddChoiceIDs(ids ...int) *ChoiceCreate {
	cc.mutation.AddChoiceIDs(ids...)
	return cc
}

// AddChoices adds the "choices" edges to the Choice entity.
func (cc *ChoiceCreate) AddChoices(c ...*Choice) *ChoiceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddChoiceIDs(ids...)
}

// AddProficiencyOptionIDs adds the "proficiency_options" edge to the Proficiency entity by IDs.
func (cc *ChoiceCreate) AddProficiencyOptionIDs(ids ...int) *ChoiceCreate {
	cc.mutation.AddProficiencyOptionIDs(ids...)
	return cc
}

// AddProficiencyOptions adds the "proficiency_options" edges to the Proficiency entity.
func (cc *ChoiceCreate) AddProficiencyOptions(p ...*Proficiency) *ChoiceCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProficiencyOptionIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (cc *ChoiceCreate) AddClasIDs(ids ...int) *ChoiceCreate {
	cc.mutation.AddClasIDs(ids...)
	return cc
}

// AddClass adds the "class" edges to the Class entity.
func (cc *ChoiceCreate) AddClass(c ...*Class) *ChoiceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddClasIDs(ids...)
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (cc *ChoiceCreate) AddRaceIDs(ids ...int) *ChoiceCreate {
	cc.mutation.AddRaceIDs(ids...)
	return cc
}

// AddRace adds the "race" edges to the Race entity.
func (cc *ChoiceCreate) AddRace(r ...*Race) *ChoiceCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRaceIDs(ids...)
}

// Mutation returns the ChoiceMutation object of the builder.
func (cc *ChoiceCreate) Mutation() *ChoiceMutation {
	return cc.mutation
}

// Save creates the Choice in the database.
func (cc *ChoiceCreate) Save(ctx context.Context) (*Choice, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChoiceCreate) SaveX(ctx context.Context) *Choice {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChoiceCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChoiceCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChoiceCreate) check() error {
	if _, ok := cc.mutation.Choose(); !ok {
		return &ValidationError{Name: "choose", err: errors.New(`ent: missing required field "Choice.choose"`)}
	}
	return nil
}

func (cc *ChoiceCreate) sqlSave(ctx context.Context) (*Choice, error) {
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

func (cc *ChoiceCreate) createSpec() (*Choice, *sqlgraph.CreateSpec) {
	var (
		_node = &Choice{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(choice.Table, sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Choose(); ok {
		_spec.SetField(choice.FieldChoose, field.TypeInt, value)
		_node.Choose = value
	}
	if value, ok := cc.mutation.Desc(); ok {
		_spec.SetField(choice.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if nodes := cc.mutation.ParentChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   choice.ParentChoiceTable,
			Columns: []string{choice.ParentChoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.choice_choices = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ChoicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   choice.ChoicesTable,
			Columns: []string{choice.ChoicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ProficiencyOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   choice.ProficiencyOptionsTable,
			Columns: choice.ProficiencyOptionsPrimaryKey,
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
	if nodes := cc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   choice.ClassTable,
			Columns: choice.ClassPrimaryKey,
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
	if nodes := cc.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   choice.RaceTable,
			Columns: []string{choice.RaceColumn},
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

// ChoiceCreateBulk is the builder for creating many Choice entities in bulk.
type ChoiceCreateBulk struct {
	config
	builders []*ChoiceCreate
}

// Save creates the Choice entities in the database.
func (ccb *ChoiceCreateBulk) Save(ctx context.Context) ([]*Choice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Choice, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChoiceMutation)
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
func (ccb *ChoiceCreateBulk) SaveX(ctx context.Context) []*Choice {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChoiceCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
