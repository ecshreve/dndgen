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
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/weaponproperty"
)

// WeaponPropertyUpdate is the builder for updating WeaponProperty entities.
type WeaponPropertyUpdate struct {
	config
	hooks    []Hook
	mutation *WeaponPropertyMutation
}

// Where appends a list predicates to the WeaponPropertyUpdate builder.
func (wpu *WeaponPropertyUpdate) Where(ps ...predicate.WeaponProperty) *WeaponPropertyUpdate {
	wpu.mutation.Where(ps...)
	return wpu
}

// SetIndx sets the "indx" field.
func (wpu *WeaponPropertyUpdate) SetIndx(s string) *WeaponPropertyUpdate {
	wpu.mutation.SetIndx(s)
	return wpu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (wpu *WeaponPropertyUpdate) SetNillableIndx(s *string) *WeaponPropertyUpdate {
	if s != nil {
		wpu.SetIndx(*s)
	}
	return wpu
}

// SetName sets the "name" field.
func (wpu *WeaponPropertyUpdate) SetName(s string) *WeaponPropertyUpdate {
	wpu.mutation.SetName(s)
	return wpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wpu *WeaponPropertyUpdate) SetNillableName(s *string) *WeaponPropertyUpdate {
	if s != nil {
		wpu.SetName(*s)
	}
	return wpu
}

// SetDesc sets the "desc" field.
func (wpu *WeaponPropertyUpdate) SetDesc(s []string) *WeaponPropertyUpdate {
	wpu.mutation.SetDesc(s)
	return wpu
}

// AppendDesc appends s to the "desc" field.
func (wpu *WeaponPropertyUpdate) AppendDesc(s []string) *WeaponPropertyUpdate {
	wpu.mutation.AppendDesc(s)
	return wpu
}

// ClearDesc clears the value of the "desc" field.
func (wpu *WeaponPropertyUpdate) ClearDesc() *WeaponPropertyUpdate {
	wpu.mutation.ClearDesc()
	return wpu
}

// Mutation returns the WeaponPropertyMutation object of the builder.
func (wpu *WeaponPropertyUpdate) Mutation() *WeaponPropertyMutation {
	return wpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wpu *WeaponPropertyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wpu.sqlSave, wpu.mutation, wpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpu *WeaponPropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := wpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wpu *WeaponPropertyUpdate) Exec(ctx context.Context) error {
	_, err := wpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpu *WeaponPropertyUpdate) ExecX(ctx context.Context) {
	if err := wpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wpu *WeaponPropertyUpdate) check() error {
	if v, ok := wpu.mutation.Indx(); ok {
		if err := weaponproperty.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "WeaponProperty.indx": %w`, err)}
		}
	}
	if v, ok := wpu.mutation.Name(); ok {
		if err := weaponproperty.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WeaponProperty.name": %w`, err)}
		}
	}
	return nil
}

func (wpu *WeaponPropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(weaponproperty.Table, weaponproperty.Columns, sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt))
	if ps := wpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wpu.mutation.Indx(); ok {
		_spec.SetField(weaponproperty.FieldIndx, field.TypeString, value)
	}
	if value, ok := wpu.mutation.Name(); ok {
		_spec.SetField(weaponproperty.FieldName, field.TypeString, value)
	}
	if value, ok := wpu.mutation.Desc(); ok {
		_spec.SetField(weaponproperty.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := wpu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, weaponproperty.FieldDesc, value)
		})
	}
	if wpu.mutation.DescCleared() {
		_spec.ClearField(weaponproperty.FieldDesc, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{weaponproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wpu.mutation.done = true
	return n, nil
}

// WeaponPropertyUpdateOne is the builder for updating a single WeaponProperty entity.
type WeaponPropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WeaponPropertyMutation
}

// SetIndx sets the "indx" field.
func (wpuo *WeaponPropertyUpdateOne) SetIndx(s string) *WeaponPropertyUpdateOne {
	wpuo.mutation.SetIndx(s)
	return wpuo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (wpuo *WeaponPropertyUpdateOne) SetNillableIndx(s *string) *WeaponPropertyUpdateOne {
	if s != nil {
		wpuo.SetIndx(*s)
	}
	return wpuo
}

// SetName sets the "name" field.
func (wpuo *WeaponPropertyUpdateOne) SetName(s string) *WeaponPropertyUpdateOne {
	wpuo.mutation.SetName(s)
	return wpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wpuo *WeaponPropertyUpdateOne) SetNillableName(s *string) *WeaponPropertyUpdateOne {
	if s != nil {
		wpuo.SetName(*s)
	}
	return wpuo
}

// SetDesc sets the "desc" field.
func (wpuo *WeaponPropertyUpdateOne) SetDesc(s []string) *WeaponPropertyUpdateOne {
	wpuo.mutation.SetDesc(s)
	return wpuo
}

// AppendDesc appends s to the "desc" field.
func (wpuo *WeaponPropertyUpdateOne) AppendDesc(s []string) *WeaponPropertyUpdateOne {
	wpuo.mutation.AppendDesc(s)
	return wpuo
}

// ClearDesc clears the value of the "desc" field.
func (wpuo *WeaponPropertyUpdateOne) ClearDesc() *WeaponPropertyUpdateOne {
	wpuo.mutation.ClearDesc()
	return wpuo
}

// Mutation returns the WeaponPropertyMutation object of the builder.
func (wpuo *WeaponPropertyUpdateOne) Mutation() *WeaponPropertyMutation {
	return wpuo.mutation
}

// Where appends a list predicates to the WeaponPropertyUpdate builder.
func (wpuo *WeaponPropertyUpdateOne) Where(ps ...predicate.WeaponProperty) *WeaponPropertyUpdateOne {
	wpuo.mutation.Where(ps...)
	return wpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wpuo *WeaponPropertyUpdateOne) Select(field string, fields ...string) *WeaponPropertyUpdateOne {
	wpuo.fields = append([]string{field}, fields...)
	return wpuo
}

// Save executes the query and returns the updated WeaponProperty entity.
func (wpuo *WeaponPropertyUpdateOne) Save(ctx context.Context) (*WeaponProperty, error) {
	return withHooks(ctx, wpuo.sqlSave, wpuo.mutation, wpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wpuo *WeaponPropertyUpdateOne) SaveX(ctx context.Context) *WeaponProperty {
	node, err := wpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wpuo *WeaponPropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := wpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpuo *WeaponPropertyUpdateOne) ExecX(ctx context.Context) {
	if err := wpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wpuo *WeaponPropertyUpdateOne) check() error {
	if v, ok := wpuo.mutation.Indx(); ok {
		if err := weaponproperty.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "WeaponProperty.indx": %w`, err)}
		}
	}
	if v, ok := wpuo.mutation.Name(); ok {
		if err := weaponproperty.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WeaponProperty.name": %w`, err)}
		}
	}
	return nil
}

func (wpuo *WeaponPropertyUpdateOne) sqlSave(ctx context.Context) (_node *WeaponProperty, err error) {
	if err := wpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(weaponproperty.Table, weaponproperty.Columns, sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt))
	id, ok := wpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WeaponProperty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, weaponproperty.FieldID)
		for _, f := range fields {
			if !weaponproperty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != weaponproperty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wpuo.mutation.Indx(); ok {
		_spec.SetField(weaponproperty.FieldIndx, field.TypeString, value)
	}
	if value, ok := wpuo.mutation.Name(); ok {
		_spec.SetField(weaponproperty.FieldName, field.TypeString, value)
	}
	if value, ok := wpuo.mutation.Desc(); ok {
		_spec.SetField(weaponproperty.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := wpuo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, weaponproperty.FieldDesc, value)
		})
	}
	if wpuo.mutation.DescCleared() {
		_spec.ClearField(weaponproperty.FieldDesc, field.TypeJSON)
	}
	_node = &WeaponProperty{config: wpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{weaponproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wpuo.mutation.done = true
	return _node, nil
}