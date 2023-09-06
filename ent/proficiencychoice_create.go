// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
)

// ProficiencyChoiceCreate is the builder for creating a ProficiencyChoice entity.
type ProficiencyChoiceCreate struct {
	config
	mutation *ProficiencyChoiceMutation
	hooks    []Hook
}

// SetChoose sets the "choose" field.
func (pcc *ProficiencyChoiceCreate) SetChoose(i int) *ProficiencyChoiceCreate {
	pcc.mutation.SetChoose(i)
	return pcc
}

// SetDesc sets the "desc" field.
func (pcc *ProficiencyChoiceCreate) SetDesc(s string) *ProficiencyChoiceCreate {
	pcc.mutation.SetDesc(s)
	return pcc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (pcc *ProficiencyChoiceCreate) SetNillableDesc(s *string) *ProficiencyChoiceCreate {
	if s != nil {
		pcc.SetDesc(*s)
	}
	return pcc
}

// AddProficiencyIDs adds the "proficiency" edge to the Proficiency entity by IDs.
func (pcc *ProficiencyChoiceCreate) AddProficiencyIDs(ids ...int) *ProficiencyChoiceCreate {
	pcc.mutation.AddProficiencyIDs(ids...)
	return pcc
}

// AddProficiency adds the "proficiency" edges to the Proficiency entity.
func (pcc *ProficiencyChoiceCreate) AddProficiency(p ...*Proficiency) *ProficiencyChoiceCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcc.AddProficiencyIDs(ids...)
}

// SetParentChoiceID sets the "parent_choice" edge to the ProficiencyChoice entity by ID.
func (pcc *ProficiencyChoiceCreate) SetParentChoiceID(id int) *ProficiencyChoiceCreate {
	pcc.mutation.SetParentChoiceID(id)
	return pcc
}

// SetNillableParentChoiceID sets the "parent_choice" edge to the ProficiencyChoice entity by ID if the given value is not nil.
func (pcc *ProficiencyChoiceCreate) SetNillableParentChoiceID(id *int) *ProficiencyChoiceCreate {
	if id != nil {
		pcc = pcc.SetParentChoiceID(*id)
	}
	return pcc
}

// SetParentChoice sets the "parent_choice" edge to the ProficiencyChoice entity.
func (pcc *ProficiencyChoiceCreate) SetParentChoice(p *ProficiencyChoice) *ProficiencyChoiceCreate {
	return pcc.SetParentChoiceID(p.ID)
}

// AddSubChoiceIDs adds the "sub_choice" edge to the ProficiencyChoice entity by IDs.
func (pcc *ProficiencyChoiceCreate) AddSubChoiceIDs(ids ...int) *ProficiencyChoiceCreate {
	pcc.mutation.AddSubChoiceIDs(ids...)
	return pcc
}

// AddSubChoice adds the "sub_choice" edges to the ProficiencyChoice entity.
func (pcc *ProficiencyChoiceCreate) AddSubChoice(p ...*ProficiencyChoice) *ProficiencyChoiceCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcc.AddSubChoiceIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (pcc *ProficiencyChoiceCreate) AddClasIDs(ids ...int) *ProficiencyChoiceCreate {
	pcc.mutation.AddClasIDs(ids...)
	return pcc
}

// AddClass adds the "class" edges to the Class entity.
func (pcc *ProficiencyChoiceCreate) AddClass(c ...*Class) *ProficiencyChoiceCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pcc.AddClasIDs(ids...)
}

// Mutation returns the ProficiencyChoiceMutation object of the builder.
func (pcc *ProficiencyChoiceCreate) Mutation() *ProficiencyChoiceMutation {
	return pcc.mutation
}

// Save creates the ProficiencyChoice in the database.
func (pcc *ProficiencyChoiceCreate) Save(ctx context.Context) (*ProficiencyChoice, error) {
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *ProficiencyChoiceCreate) SaveX(ctx context.Context) *ProficiencyChoice {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *ProficiencyChoiceCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *ProficiencyChoiceCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *ProficiencyChoiceCreate) check() error {
	if _, ok := pcc.mutation.Choose(); !ok {
		return &ValidationError{Name: "choose", err: errors.New(`ent: missing required field "ProficiencyChoice.choose"`)}
	}
	return nil
}

func (pcc *ProficiencyChoiceCreate) sqlSave(ctx context.Context) (*ProficiencyChoice, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *ProficiencyChoiceCreate) createSpec() (*ProficiencyChoice, *sqlgraph.CreateSpec) {
	var (
		_node = &ProficiencyChoice{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(proficiencychoice.Table, sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt))
	)
	if value, ok := pcc.mutation.Choose(); ok {
		_spec.SetField(proficiencychoice.FieldChoose, field.TypeInt, value)
		_node.Choose = value
	}
	if value, ok := pcc.mutation.Desc(); ok {
		_spec.SetField(proficiencychoice.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if nodes := pcc.mutation.ProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
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
	if nodes := pcc.mutation.ParentChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proficiencychoice.ParentChoiceTable,
			Columns: []string{proficiencychoice.ParentChoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.proficiency_choice_sub_choice = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.SubChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
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
	if nodes := pcc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
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
	return _node, _spec
}

// ProficiencyChoiceCreateBulk is the builder for creating many ProficiencyChoice entities in bulk.
type ProficiencyChoiceCreateBulk struct {
	config
	builders []*ProficiencyChoiceCreate
}

// Save creates the ProficiencyChoice entities in the database.
func (pccb *ProficiencyChoiceCreateBulk) Save(ctx context.Context) ([]*ProficiencyChoice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*ProficiencyChoice, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProficiencyChoiceMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *ProficiencyChoiceCreateBulk) SaveX(ctx context.Context) []*ProficiencyChoice {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *ProficiencyChoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *ProficiencyChoiceCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
