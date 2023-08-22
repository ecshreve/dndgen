// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armorclass"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ArmorClassDelete is the builder for deleting a ArmorClass entity.
type ArmorClassDelete struct {
	config
	hooks    []Hook
	mutation *ArmorClassMutation
}

// Where appends a list predicates to the ArmorClassDelete builder.
func (acd *ArmorClassDelete) Where(ps ...predicate.ArmorClass) *ArmorClassDelete {
	acd.mutation.Where(ps...)
	return acd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acd *ArmorClassDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ArmorClassMutation](ctx, acd.sqlExec, acd.mutation, acd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acd *ArmorClassDelete) ExecX(ctx context.Context) int {
	n, err := acd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acd *ArmorClassDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(armorclass.Table, sqlgraph.NewFieldSpec(armorclass.FieldID, field.TypeInt))
	if ps := acd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acd.mutation.done = true
	return affected, err
}

// ArmorClassDeleteOne is the builder for deleting a single ArmorClass entity.
type ArmorClassDeleteOne struct {
	acd *ArmorClassDelete
}

// Where appends a list predicates to the ArmorClassDelete builder.
func (acdo *ArmorClassDeleteOne) Where(ps ...predicate.ArmorClass) *ArmorClassDeleteOne {
	acdo.acd.mutation.Where(ps...)
	return acdo
}

// Exec executes the deletion query.
func (acdo *ArmorClassDeleteOne) Exec(ctx context.Context) error {
	n, err := acdo.acd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{armorclass.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acdo *ArmorClassDeleteOne) ExecX(ctx context.Context) {
	if err := acdo.Exec(ctx); err != nil {
		panic(err)
	}
}
