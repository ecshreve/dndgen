// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipmententry"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// EquipmentEntryDelete is the builder for deleting a EquipmentEntry entity.
type EquipmentEntryDelete struct {
	config
	hooks    []Hook
	mutation *EquipmentEntryMutation
}

// Where appends a list predicates to the EquipmentEntryDelete builder.
func (eed *EquipmentEntryDelete) Where(ps ...predicate.EquipmentEntry) *EquipmentEntryDelete {
	eed.mutation.Where(ps...)
	return eed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eed *EquipmentEntryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eed.sqlExec, eed.mutation, eed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eed *EquipmentEntryDelete) ExecX(ctx context.Context) int {
	n, err := eed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eed *EquipmentEntryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(equipmententry.Table, sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt))
	if ps := eed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eed.mutation.done = true
	return affected, err
}

// EquipmentEntryDeleteOne is the builder for deleting a single EquipmentEntry entity.
type EquipmentEntryDeleteOne struct {
	eed *EquipmentEntryDelete
}

// Where appends a list predicates to the EquipmentEntryDelete builder.
func (eedo *EquipmentEntryDeleteOne) Where(ps ...predicate.EquipmentEntry) *EquipmentEntryDeleteOne {
	eedo.eed.mutation.Where(ps...)
	return eedo
}

// Exec executes the deletion query.
func (eedo *EquipmentEntryDeleteOne) Exec(ctx context.Context) error {
	n, err := eedo.eed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{equipmententry.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eedo *EquipmentEntryDeleteOne) ExecX(ctx context.Context) {
	if err := eedo.Exec(ctx); err != nil {
		panic(err)
	}
}
