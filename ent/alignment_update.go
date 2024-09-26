// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// AlignmentUpdate is the builder for updating Alignment entities.
type AlignmentUpdate struct {
	config
	hooks    []Hook
	mutation *AlignmentMutation
}

// Where appends a list predicates to the AlignmentUpdate builder.
func (au *AlignmentUpdate) Where(ps ...predicate.Alignment) *AlignmentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetIndx sets the "indx" field.
func (au *AlignmentUpdate) SetIndx(s string) *AlignmentUpdate {
	au.mutation.SetIndx(s)
	return au
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (au *AlignmentUpdate) SetNillableIndx(s *string) *AlignmentUpdate {
	if s != nil {
		au.SetIndx(*s)
	}
	return au
}

// SetName sets the "name" field.
func (au *AlignmentUpdate) SetName(s string) *AlignmentUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AlignmentUpdate) SetNillableName(s *string) *AlignmentUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetDesc sets the "desc" field.
func (au *AlignmentUpdate) SetDesc(s []string) *AlignmentUpdate {
	au.mutation.SetDesc(s)
	return au
}

// AppendDesc appends s to the "desc" field.
func (au *AlignmentUpdate) AppendDesc(s []string) *AlignmentUpdate {
	au.mutation.AppendDesc(s)
	return au
}

// ClearDesc clears the value of the "desc" field.
func (au *AlignmentUpdate) ClearDesc() *AlignmentUpdate {
	au.mutation.ClearDesc()
	return au
}

// SetAbbr sets the "abbr" field.
func (au *AlignmentUpdate) SetAbbr(a alignment.Abbr) *AlignmentUpdate {
	au.mutation.SetAbbr(a)
	return au
}

// SetNillableAbbr sets the "abbr" field if the given value is not nil.
func (au *AlignmentUpdate) SetNillableAbbr(a *alignment.Abbr) *AlignmentUpdate {
	if a != nil {
		au.SetAbbr(*a)
	}
	return au
}

// Mutation returns the AlignmentMutation object of the builder.
func (au *AlignmentUpdate) Mutation() *AlignmentMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AlignmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AlignmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AlignmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AlignmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AlignmentUpdate) check() error {
	if v, ok := au.mutation.Indx(); ok {
		if err := alignment.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Alignment.indx": %w`, err)}
		}
	}
	if v, ok := au.mutation.Name(); ok {
		if err := alignment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Alignment.name": %w`, err)}
		}
	}
	if v, ok := au.mutation.Abbr(); ok {
		if err := alignment.AbbrValidator(v); err != nil {
			return &ValidationError{Name: "abbr", err: fmt.Errorf(`ent: validator failed for field "Alignment.abbr": %w`, err)}
		}
	}
	return nil
}

func (au *AlignmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(alignment.Table, alignment.Columns, sqlgraph.NewFieldSpec(alignment.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Indx(); ok {
		_spec.SetField(alignment.FieldIndx, field.TypeString, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(alignment.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Desc(); ok {
		_spec.SetField(alignment.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := au.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, alignment.FieldDesc, value)
		})
	}
	if au.mutation.DescCleared() {
		_spec.ClearField(alignment.FieldDesc, field.TypeJSON)
	}
	if value, ok := au.mutation.Abbr(); ok {
		_spec.SetField(alignment.FieldAbbr, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AlignmentUpdateOne is the builder for updating a single Alignment entity.
type AlignmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AlignmentMutation
}

// SetIndx sets the "indx" field.
func (auo *AlignmentUpdateOne) SetIndx(s string) *AlignmentUpdateOne {
	auo.mutation.SetIndx(s)
	return auo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (auo *AlignmentUpdateOne) SetNillableIndx(s *string) *AlignmentUpdateOne {
	if s != nil {
		auo.SetIndx(*s)
	}
	return auo
}

// SetName sets the "name" field.
func (auo *AlignmentUpdateOne) SetName(s string) *AlignmentUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AlignmentUpdateOne) SetNillableName(s *string) *AlignmentUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetDesc sets the "desc" field.
func (auo *AlignmentUpdateOne) SetDesc(s []string) *AlignmentUpdateOne {
	auo.mutation.SetDesc(s)
	return auo
}

// AppendDesc appends s to the "desc" field.
func (auo *AlignmentUpdateOne) AppendDesc(s []string) *AlignmentUpdateOne {
	auo.mutation.AppendDesc(s)
	return auo
}

// ClearDesc clears the value of the "desc" field.
func (auo *AlignmentUpdateOne) ClearDesc() *AlignmentUpdateOne {
	auo.mutation.ClearDesc()
	return auo
}

// SetAbbr sets the "abbr" field.
func (auo *AlignmentUpdateOne) SetAbbr(a alignment.Abbr) *AlignmentUpdateOne {
	auo.mutation.SetAbbr(a)
	return auo
}

// SetNillableAbbr sets the "abbr" field if the given value is not nil.
func (auo *AlignmentUpdateOne) SetNillableAbbr(a *alignment.Abbr) *AlignmentUpdateOne {
	if a != nil {
		auo.SetAbbr(*a)
	}
	return auo
}

// Mutation returns the AlignmentMutation object of the builder.
func (auo *AlignmentUpdateOne) Mutation() *AlignmentMutation {
	return auo.mutation
}

// Where appends a list predicates to the AlignmentUpdate builder.
func (auo *AlignmentUpdateOne) Where(ps ...predicate.Alignment) *AlignmentUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AlignmentUpdateOne) Select(field string, fields ...string) *AlignmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Alignment entity.
func (auo *AlignmentUpdateOne) Save(ctx context.Context) (*Alignment, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AlignmentUpdateOne) SaveX(ctx context.Context) *Alignment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AlignmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AlignmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AlignmentUpdateOne) check() error {
	if v, ok := auo.mutation.Indx(); ok {
		if err := alignment.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Alignment.indx": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Name(); ok {
		if err := alignment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Alignment.name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Abbr(); ok {
		if err := alignment.AbbrValidator(v); err != nil {
			return &ValidationError{Name: "abbr", err: fmt.Errorf(`ent: validator failed for field "Alignment.abbr": %w`, err)}
		}
	}
	return nil
}

func (auo *AlignmentUpdateOne) sqlSave(ctx context.Context) (_node *Alignment, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(alignment.Table, alignment.Columns, sqlgraph.NewFieldSpec(alignment.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Alignment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, alignment.FieldID)
		for _, f := range fields {
			if !alignment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != alignment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Indx(); ok {
		_spec.SetField(alignment.FieldIndx, field.TypeString, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(alignment.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Desc(); ok {
		_spec.SetField(alignment.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := auo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, alignment.FieldDesc, value)
		})
	}
	if auo.mutation.DescCleared() {
		_spec.ClearField(alignment.FieldDesc, field.TypeJSON)
	}
	if value, ok := auo.mutation.Abbr(); ok {
		_spec.SetField(alignment.FieldAbbr, field.TypeEnum, value)
	}
	_node = &Alignment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{alignment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
