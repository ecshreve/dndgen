// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// AbilityScoreDelete is the builder for deleting a AbilityScore entity.
type AbilityScoreDelete struct {
	config
	hooks    []Hook
	mutation *AbilityScoreMutation
}

// Where appends a list predicates to the AbilityScoreDelete builder.
func (asd *AbilityScoreDelete) Where(ps ...predicate.AbilityScore) *AbilityScoreDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AbilityScoreDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AbilityScoreDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AbilityScoreDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(abilityscore.Table, sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// AbilityScoreDeleteOne is the builder for deleting a single AbilityScore entity.
type AbilityScoreDeleteOne struct {
	asd *AbilityScoreDelete
}

// Where appends a list predicates to the AbilityScoreDelete builder.
func (asdo *AbilityScoreDeleteOne) Where(ps ...predicate.AbilityScore) *AbilityScoreDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *AbilityScoreDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{abilityscore.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AbilityScoreDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}
