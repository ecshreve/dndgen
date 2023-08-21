// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
)

// RaceUpdate is the builder for updating Race entities.
type RaceUpdate struct {
	config
	hooks    []Hook
	mutation *RaceMutation
}

// Where appends a list predicates to the RaceUpdate builder.
func (ru *RaceUpdate) Where(ps ...predicate.Race) *RaceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RaceUpdate) SetName(s string) *RaceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetSpeed sets the "speed" field.
func (ru *RaceUpdate) SetSpeed(i int) *RaceUpdate {
	ru.mutation.ResetSpeed()
	ru.mutation.SetSpeed(i)
	return ru
}

// AddSpeed adds i to the "speed" field.
func (ru *RaceUpdate) AddSpeed(i int) *RaceUpdate {
	ru.mutation.AddSpeed(i)
	return ru
}

// Mutation returns the RaceMutation object of the builder.
func (ru *RaceUpdate) Mutation() *RaceMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RaceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RaceMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RaceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RaceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RaceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(race.Table, race.Columns, sqlgraph.NewFieldSpec(race.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(race.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Speed(); ok {
		_spec.SetField(race.FieldSpeed, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedSpeed(); ok {
		_spec.AddField(race.FieldSpeed, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{race.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RaceUpdateOne is the builder for updating a single Race entity.
type RaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RaceMutation
}

// SetName sets the "name" field.
func (ruo *RaceUpdateOne) SetName(s string) *RaceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetSpeed sets the "speed" field.
func (ruo *RaceUpdateOne) SetSpeed(i int) *RaceUpdateOne {
	ruo.mutation.ResetSpeed()
	ruo.mutation.SetSpeed(i)
	return ruo
}

// AddSpeed adds i to the "speed" field.
func (ruo *RaceUpdateOne) AddSpeed(i int) *RaceUpdateOne {
	ruo.mutation.AddSpeed(i)
	return ruo
}

// Mutation returns the RaceMutation object of the builder.
func (ruo *RaceUpdateOne) Mutation() *RaceMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RaceUpdate builder.
func (ruo *RaceUpdateOne) Where(ps ...predicate.Race) *RaceUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RaceUpdateOne) Select(field string, fields ...string) *RaceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Race entity.
func (ruo *RaceUpdateOne) Save(ctx context.Context) (*Race, error) {
	return withHooks[*Race, RaceMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RaceUpdateOne) SaveX(ctx context.Context) *Race {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RaceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RaceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RaceUpdateOne) sqlSave(ctx context.Context) (_node *Race, err error) {
	_spec := sqlgraph.NewUpdateSpec(race.Table, race.Columns, sqlgraph.NewFieldSpec(race.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Race.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, race.FieldID)
		for _, f := range fields {
			if !race.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != race.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(race.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Speed(); ok {
		_spec.SetField(race.FieldSpeed, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedSpeed(); ok {
		_spec.AddField(race.FieldSpeed, field.TypeInt, value)
	}
	_node = &Race{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{race.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
