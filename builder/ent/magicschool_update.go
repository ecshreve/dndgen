// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/builder/ent/magicschool"
	"github.com/ecshreve/dndgen/builder/ent/predicate"
)

// MagicSchoolUpdate is the builder for updating MagicSchool entities.
type MagicSchoolUpdate struct {
	config
	hooks    []Hook
	mutation *MagicSchoolMutation
}

// Where appends a list predicates to the MagicSchoolUpdate builder.
func (msu *MagicSchoolUpdate) Where(ps ...predicate.MagicSchool) *MagicSchoolUpdate {
	msu.mutation.Where(ps...)
	return msu
}

// SetIndx sets the "indx" field.
func (msu *MagicSchoolUpdate) SetIndx(s string) *MagicSchoolUpdate {
	msu.mutation.SetIndx(s)
	return msu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (msu *MagicSchoolUpdate) SetNillableIndx(s *string) *MagicSchoolUpdate {
	if s != nil {
		msu.SetIndx(*s)
	}
	return msu
}

// SetName sets the "name" field.
func (msu *MagicSchoolUpdate) SetName(s string) *MagicSchoolUpdate {
	msu.mutation.SetName(s)
	return msu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (msu *MagicSchoolUpdate) SetNillableName(s *string) *MagicSchoolUpdate {
	if s != nil {
		msu.SetName(*s)
	}
	return msu
}

// SetDesc sets the "desc" field.
func (msu *MagicSchoolUpdate) SetDesc(s string) *MagicSchoolUpdate {
	msu.mutation.SetDesc(s)
	return msu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (msu *MagicSchoolUpdate) SetNillableDesc(s *string) *MagicSchoolUpdate {
	if s != nil {
		msu.SetDesc(*s)
	}
	return msu
}

// Mutation returns the MagicSchoolMutation object of the builder.
func (msu *MagicSchoolUpdate) Mutation() *MagicSchoolMutation {
	return msu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (msu *MagicSchoolUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, msu.sqlSave, msu.mutation, msu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msu *MagicSchoolUpdate) SaveX(ctx context.Context) int {
	affected, err := msu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (msu *MagicSchoolUpdate) Exec(ctx context.Context) error {
	_, err := msu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msu *MagicSchoolUpdate) ExecX(ctx context.Context) {
	if err := msu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msu *MagicSchoolUpdate) check() error {
	if v, ok := msu.mutation.Indx(); ok {
		if err := magicschool.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "MagicSchool.indx": %w`, err)}
		}
	}
	if v, ok := msu.mutation.Name(); ok {
		if err := magicschool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MagicSchool.name": %w`, err)}
		}
	}
	return nil
}

func (msu *MagicSchoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := msu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(magicschool.Table, magicschool.Columns, sqlgraph.NewFieldSpec(magicschool.FieldID, field.TypeInt))
	if ps := msu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msu.mutation.Indx(); ok {
		_spec.SetField(magicschool.FieldIndx, field.TypeString, value)
	}
	if value, ok := msu.mutation.Name(); ok {
		_spec.SetField(magicschool.FieldName, field.TypeString, value)
	}
	if value, ok := msu.mutation.Desc(); ok {
		_spec.SetField(magicschool.FieldDesc, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, msu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{magicschool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	msu.mutation.done = true
	return n, nil
}

// MagicSchoolUpdateOne is the builder for updating a single MagicSchool entity.
type MagicSchoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MagicSchoolMutation
}

// SetIndx sets the "indx" field.
func (msuo *MagicSchoolUpdateOne) SetIndx(s string) *MagicSchoolUpdateOne {
	msuo.mutation.SetIndx(s)
	return msuo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (msuo *MagicSchoolUpdateOne) SetNillableIndx(s *string) *MagicSchoolUpdateOne {
	if s != nil {
		msuo.SetIndx(*s)
	}
	return msuo
}

// SetName sets the "name" field.
func (msuo *MagicSchoolUpdateOne) SetName(s string) *MagicSchoolUpdateOne {
	msuo.mutation.SetName(s)
	return msuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (msuo *MagicSchoolUpdateOne) SetNillableName(s *string) *MagicSchoolUpdateOne {
	if s != nil {
		msuo.SetName(*s)
	}
	return msuo
}

// SetDesc sets the "desc" field.
func (msuo *MagicSchoolUpdateOne) SetDesc(s string) *MagicSchoolUpdateOne {
	msuo.mutation.SetDesc(s)
	return msuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (msuo *MagicSchoolUpdateOne) SetNillableDesc(s *string) *MagicSchoolUpdateOne {
	if s != nil {
		msuo.SetDesc(*s)
	}
	return msuo
}

// Mutation returns the MagicSchoolMutation object of the builder.
func (msuo *MagicSchoolUpdateOne) Mutation() *MagicSchoolMutation {
	return msuo.mutation
}

// Where appends a list predicates to the MagicSchoolUpdate builder.
func (msuo *MagicSchoolUpdateOne) Where(ps ...predicate.MagicSchool) *MagicSchoolUpdateOne {
	msuo.mutation.Where(ps...)
	return msuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (msuo *MagicSchoolUpdateOne) Select(field string, fields ...string) *MagicSchoolUpdateOne {
	msuo.fields = append([]string{field}, fields...)
	return msuo
}

// Save executes the query and returns the updated MagicSchool entity.
func (msuo *MagicSchoolUpdateOne) Save(ctx context.Context) (*MagicSchool, error) {
	return withHooks(ctx, msuo.sqlSave, msuo.mutation, msuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (msuo *MagicSchoolUpdateOne) SaveX(ctx context.Context) *MagicSchool {
	node, err := msuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (msuo *MagicSchoolUpdateOne) Exec(ctx context.Context) error {
	_, err := msuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (msuo *MagicSchoolUpdateOne) ExecX(ctx context.Context) {
	if err := msuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (msuo *MagicSchoolUpdateOne) check() error {
	if v, ok := msuo.mutation.Indx(); ok {
		if err := magicschool.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "MagicSchool.indx": %w`, err)}
		}
	}
	if v, ok := msuo.mutation.Name(); ok {
		if err := magicschool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "MagicSchool.name": %w`, err)}
		}
	}
	return nil
}

func (msuo *MagicSchoolUpdateOne) sqlSave(ctx context.Context) (_node *MagicSchool, err error) {
	if err := msuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(magicschool.Table, magicschool.Columns, sqlgraph.NewFieldSpec(magicschool.FieldID, field.TypeInt))
	id, ok := msuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MagicSchool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := msuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, magicschool.FieldID)
		for _, f := range fields {
			if !magicschool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != magicschool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := msuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := msuo.mutation.Indx(); ok {
		_spec.SetField(magicschool.FieldIndx, field.TypeString, value)
	}
	if value, ok := msuo.mutation.Name(); ok {
		_spec.SetField(magicschool.FieldName, field.TypeString, value)
	}
	if value, ok := msuo.mutation.Desc(); ok {
		_spec.SetField(magicschool.FieldDesc, field.TypeString, value)
	}
	_node = &MagicSchool{config: msuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, msuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{magicschool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	msuo.mutation.done = true
	return _node, nil
}
