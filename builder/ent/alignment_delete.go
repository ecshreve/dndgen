// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/builder/ent/alignment"
	"github.com/ecshreve/dndgen/builder/ent/predicate"
)

// AlignmentDelete is the builder for deleting a Alignment entity.
type AlignmentDelete struct {
	config
	hooks    []Hook
	mutation *AlignmentMutation
}

// Where appends a list predicates to the AlignmentDelete builder.
func (ad *AlignmentDelete) Where(ps ...predicate.Alignment) *AlignmentDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AlignmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AlignmentDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AlignmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(alignment.Table, sqlgraph.NewFieldSpec(alignment.FieldID, field.TypeInt))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AlignmentDeleteOne is the builder for deleting a single Alignment entity.
type AlignmentDeleteOne struct {
	ad *AlignmentDelete
}

// Where appends a list predicates to the AlignmentDelete builder.
func (ado *AlignmentDeleteOne) Where(ps ...predicate.Alignment) *AlignmentDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AlignmentDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{alignment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AlignmentDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
