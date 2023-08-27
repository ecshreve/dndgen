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
)

// ClassCreate is the builder for creating a Class entity.
type ClassCreate struct {
	config
	mutation *ClassMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (cc *ClassCreate) SetIndx(s string) *ClassCreate {
	cc.mutation.SetIndx(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ClassCreate) SetName(s string) *ClassCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetHitDie sets the "hit_die" field.
func (cc *ClassCreate) SetHitDie(i int) *ClassCreate {
	cc.mutation.SetHitDie(i)
	return cc
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (cc *ClassCreate) AddProficiencyIDs(ids ...int) *ClassCreate {
	cc.mutation.AddProficiencyIDs(ids...)
	return cc
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (cc *ClassCreate) AddProficiencies(p ...*Proficiency) *ClassCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProficiencyIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cc *ClassCreate) Mutation() *ClassMutation {
	return cc.mutation
}

// Save creates the Class in the database.
func (cc *ClassCreate) Save(ctx context.Context) (*Class, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassCreate) SaveX(ctx context.Context) *Class {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ClassCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ClassCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ClassCreate) check() error {
	if _, ok := cc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Class.indx"`)}
	}
	if v, ok := cc.mutation.Indx(); ok {
		if err := class.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Class.indx": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Class.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Class.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.HitDie(); !ok {
		return &ValidationError{Name: "hit_die", err: errors.New(`ent: missing required field "Class.hit_die"`)}
	}
	return nil
}

func (cc *ClassCreate) sqlSave(ctx context.Context) (*Class, error) {
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

func (cc *ClassCreate) createSpec() (*Class, *sqlgraph.CreateSpec) {
	var (
		_node = &Class{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(class.Table, sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Indx(); ok {
		_spec.SetField(class.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(class.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.HitDie(); ok {
		_spec.SetField(class.FieldHitDie, field.TypeInt, value)
		_node.HitDie = value
	}
	if nodes := cc.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
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
	return _node, _spec
}

// ClassCreateBulk is the builder for creating many Class entities in bulk.
type ClassCreateBulk struct {
	config
	builders []*ClassCreate
}

// Save creates the Class entities in the database.
func (ccb *ClassCreateBulk) Save(ctx context.Context) ([]*Class, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Class, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassMutation)
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
func (ccb *ClassCreateBulk) SaveX(ctx context.Context) []*Class {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ClassCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ClassCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
