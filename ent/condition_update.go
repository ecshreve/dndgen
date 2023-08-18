// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/condition"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ConditionUpdate is the builder for updating Condition entities.
type ConditionUpdate struct {
	config
	hooks    []Hook
	mutation *ConditionMutation
}

// Where appends a list predicates to the ConditionUpdate builder.
func (cu *ConditionUpdate) Where(ps ...predicate.Condition) *ConditionUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetIndx sets the "indx" field.
func (cu *ConditionUpdate) SetIndx(s string) *ConditionUpdate {
	cu.mutation.SetIndx(s)
	return cu
}

// SetName sets the "name" field.
func (cu *ConditionUpdate) SetName(s string) *ConditionUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDesc sets the "desc" field.
func (cu *ConditionUpdate) SetDesc(s string) *ConditionUpdate {
	cu.mutation.SetDesc(s)
	return cu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (cu *ConditionUpdate) SetNillableDesc(s *string) *ConditionUpdate {
	if s != nil {
		cu.SetDesc(*s)
	}
	return cu
}

// ClearDesc clears the value of the "desc" field.
func (cu *ConditionUpdate) ClearDesc() *ConditionUpdate {
	cu.mutation.ClearDesc()
	return cu
}

// Mutation returns the ConditionMutation object of the builder.
func (cu *ConditionUpdate) Mutation() *ConditionMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConditionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ConditionMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConditionUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConditionUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConditionUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *ConditionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(condition.Table, condition.Columns, sqlgraph.NewFieldSpec(condition.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Indx(); ok {
		_spec.SetField(condition.FieldIndx, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(condition.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Desc(); ok {
		_spec.SetField(condition.FieldDesc, field.TypeString, value)
	}
	if cu.mutation.DescCleared() {
		_spec.ClearField(condition.FieldDesc, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{condition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConditionUpdateOne is the builder for updating a single Condition entity.
type ConditionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConditionMutation
}

// SetIndx sets the "indx" field.
func (cuo *ConditionUpdateOne) SetIndx(s string) *ConditionUpdateOne {
	cuo.mutation.SetIndx(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ConditionUpdateOne) SetName(s string) *ConditionUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDesc sets the "desc" field.
func (cuo *ConditionUpdateOne) SetDesc(s string) *ConditionUpdateOne {
	cuo.mutation.SetDesc(s)
	return cuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (cuo *ConditionUpdateOne) SetNillableDesc(s *string) *ConditionUpdateOne {
	if s != nil {
		cuo.SetDesc(*s)
	}
	return cuo
}

// ClearDesc clears the value of the "desc" field.
func (cuo *ConditionUpdateOne) ClearDesc() *ConditionUpdateOne {
	cuo.mutation.ClearDesc()
	return cuo
}

// Mutation returns the ConditionMutation object of the builder.
func (cuo *ConditionUpdateOne) Mutation() *ConditionMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConditionUpdate builder.
func (cuo *ConditionUpdateOne) Where(ps ...predicate.Condition) *ConditionUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConditionUpdateOne) Select(field string, fields ...string) *ConditionUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Condition entity.
func (cuo *ConditionUpdateOne) Save(ctx context.Context) (*Condition, error) {
	return withHooks[*Condition, ConditionMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConditionUpdateOne) SaveX(ctx context.Context) *Condition {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConditionUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConditionUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *ConditionUpdateOne) sqlSave(ctx context.Context) (_node *Condition, err error) {
	_spec := sqlgraph.NewUpdateSpec(condition.Table, condition.Columns, sqlgraph.NewFieldSpec(condition.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Condition.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, condition.FieldID)
		for _, f := range fields {
			if !condition.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != condition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Indx(); ok {
		_spec.SetField(condition.FieldIndx, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(condition.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Desc(); ok {
		_spec.SetField(condition.FieldDesc, field.TypeString, value)
	}
	if cuo.mutation.DescCleared() {
		_spec.ClearField(condition.FieldDesc, field.TypeString)
	}
	_node = &Condition{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{condition.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
