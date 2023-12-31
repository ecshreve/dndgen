// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/magicschool"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// MagicSchoolDelete is the builder for deleting a MagicSchool entity.
type MagicSchoolDelete struct {
	config
	hooks    []Hook
	mutation *MagicSchoolMutation
}

// Where appends a list predicates to the MagicSchoolDelete builder.
func (msd *MagicSchoolDelete) Where(ps ...predicate.MagicSchool) *MagicSchoolDelete {
	msd.mutation.Where(ps...)
	return msd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (msd *MagicSchoolDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, msd.sqlExec, msd.mutation, msd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (msd *MagicSchoolDelete) ExecX(ctx context.Context) int {
	n, err := msd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (msd *MagicSchoolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(magicschool.Table, sqlgraph.NewFieldSpec(magicschool.FieldID, field.TypeInt))
	if ps := msd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, msd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	msd.mutation.done = true
	return affected, err
}

// MagicSchoolDeleteOne is the builder for deleting a single MagicSchool entity.
type MagicSchoolDeleteOne struct {
	msd *MagicSchoolDelete
}

// Where appends a list predicates to the MagicSchoolDelete builder.
func (msdo *MagicSchoolDeleteOne) Where(ps ...predicate.MagicSchool) *MagicSchoolDeleteOne {
	msdo.msd.mutation.Where(ps...)
	return msdo
}

// Exec executes the deletion query.
func (msdo *MagicSchoolDeleteOne) Exec(ctx context.Context) error {
	n, err := msdo.msd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{magicschool.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (msdo *MagicSchoolDeleteOne) ExecX(ctx context.Context) {
	if err := msdo.Exec(ctx); err != nil {
		panic(err)
	}
}
