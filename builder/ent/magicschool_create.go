// Code generated by ent, DO NOT EDIT.

package ent

import (
	"builder/ent/magicschool"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MagicSchoolCreate is the builder for creating a MagicSchool entity.
type MagicSchoolCreate struct {
	config
	mutation *MagicSchoolMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (msc *MagicSchoolCreate) SetIndx(s string) *MagicSchoolCreate {
	msc.mutation.SetIndx(s)
	return msc
}

// SetName sets the "name" field.
func (msc *MagicSchoolCreate) SetName(s string) *MagicSchoolCreate {
	msc.mutation.SetName(s)
	return msc
}

// SetDesc sets the "desc" field.
func (msc *MagicSchoolCreate) SetDesc(s string) *MagicSchoolCreate {
	msc.mutation.SetDesc(s)
	return msc
}

// Mutation returns the MagicSchoolMutation object of the builder.
func (msc *MagicSchoolCreate) Mutation() *MagicSchoolMutation {
	return msc.mutation
}

// Save creates the MagicSchool in the database.
func (msc *MagicSchoolCreate) Save(ctx context.Context) (*MagicSchool, error) {
	return withHooks(ctx, msc.sqlSave, msc.mutation, msc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (msc *MagicSchoolCreate) SaveX(ctx context.Context) *MagicSchool {
	v, err := msc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (msc *MagicSchoolCreate) Exec(ctx context.Context) error {
	_, err := msc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msc *MagicSchoolCreate) ExecX(ctx context.Context) {
	if err := msc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msc *MagicSchoolCreate) check() error {
	if _, ok := msc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "MagicSchool.indx"`)}
	}
	if v, ok := msc.mutation.Indx(); ok {
		if err := magicschool.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "MagicSchool.indx": %w`, err)}
		}
	}
	if _, ok := msc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "MagicSchool.name"`)}
	}
	if v, ok := msc.mutation.Name(); ok {
		if err := magicschool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MagicSchool.name": %w`, err)}
		}
	}
	if _, ok := msc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "MagicSchool.desc"`)}
	}
	return nil
}

func (msc *MagicSchoolCreate) sqlSave(ctx context.Context) (*MagicSchool, error) {
	if err := msc.check(); err != nil {
		return nil, err
	}
	_node, _spec := msc.createSpec()
	if err := sqlgraph.CreateNode(ctx, msc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	msc.mutation.id = &_node.ID
	msc.mutation.done = true
	return _node, nil
}

func (msc *MagicSchoolCreate) createSpec() (*MagicSchool, *sqlgraph.CreateSpec) {
	var (
		_node = &MagicSchool{config: msc.config}
		_spec = sqlgraph.NewCreateSpec(magicschool.Table, sqlgraph.NewFieldSpec(magicschool.FieldID, field.TypeInt))
	)
	if value, ok := msc.mutation.Indx(); ok {
		_spec.SetField(magicschool.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := msc.mutation.Name(); ok {
		_spec.SetField(magicschool.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := msc.mutation.Desc(); ok {
		_spec.SetField(magicschool.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	return _node, _spec
}

// MagicSchoolCreateBulk is the builder for creating many MagicSchool entities in bulk.
type MagicSchoolCreateBulk struct {
	config
	err      error
	builders []*MagicSchoolCreate
}

// Save creates the MagicSchool entities in the database.
func (mscb *MagicSchoolCreateBulk) Save(ctx context.Context) ([]*MagicSchool, error) {
	if mscb.err != nil {
		return nil, mscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mscb.builders))
	nodes := make([]*MagicSchool, len(mscb.builders))
	mutators := make([]Mutator, len(mscb.builders))
	for i := range mscb.builders {
		func(i int, root context.Context) {
			builder := mscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MagicSchoolMutation)
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
					_, err = mutators[i+1].Mutate(root, mscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mscb *MagicSchoolCreateBulk) SaveX(ctx context.Context) []*MagicSchool {
	v, err := mscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mscb *MagicSchoolCreateBulk) Exec(ctx context.Context) error {
	_, err := mscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mscb *MagicSchoolCreateBulk) ExecX(ctx context.Context) {
	if err := mscb.Exec(ctx); err != nil {
		panic(err)
	}
}
