// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/prerequisite"
)

// PrerequisiteUpdate is the builder for updating Prerequisite entities.
type PrerequisiteUpdate struct {
	config
	hooks    []Hook
	mutation *PrerequisiteMutation
}

// Where appends a list predicates to the PrerequisiteUpdate builder.
func (pu *PrerequisiteUpdate) Where(ps ...predicate.Prerequisite) *PrerequisiteUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetMinimum sets the "minimum" field.
func (pu *PrerequisiteUpdate) SetMinimum(i int) *PrerequisiteUpdate {
	pu.mutation.ResetMinimum()
	pu.mutation.SetMinimum(i)
	return pu
}

// AddMinimum adds i to the "minimum" field.
func (pu *PrerequisiteUpdate) AddMinimum(i int) *PrerequisiteUpdate {
	pu.mutation.AddMinimum(i)
	return pu
}

// AddAbilityScoreIDs adds the "ability_score" edge to the AbilityScore entity by IDs.
func (pu *PrerequisiteUpdate) AddAbilityScoreIDs(ids ...int) *PrerequisiteUpdate {
	pu.mutation.AddAbilityScoreIDs(ids...)
	return pu
}

// AddAbilityScore adds the "ability_score" edges to the AbilityScore entity.
func (pu *PrerequisiteUpdate) AddAbilityScore(a ...*AbilityScore) *PrerequisiteUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.AddAbilityScoreIDs(ids...)
}

// Mutation returns the PrerequisiteMutation object of the builder.
func (pu *PrerequisiteUpdate) Mutation() *PrerequisiteMutation {
	return pu.mutation
}

// ClearAbilityScore clears all "ability_score" edges to the AbilityScore entity.
func (pu *PrerequisiteUpdate) ClearAbilityScore() *PrerequisiteUpdate {
	pu.mutation.ClearAbilityScore()
	return pu
}

// RemoveAbilityScoreIDs removes the "ability_score" edge to AbilityScore entities by IDs.
func (pu *PrerequisiteUpdate) RemoveAbilityScoreIDs(ids ...int) *PrerequisiteUpdate {
	pu.mutation.RemoveAbilityScoreIDs(ids...)
	return pu
}

// RemoveAbilityScore removes "ability_score" edges to AbilityScore entities.
func (pu *PrerequisiteUpdate) RemoveAbilityScore(a ...*AbilityScore) *PrerequisiteUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pu.RemoveAbilityScoreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PrerequisiteUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PrerequisiteMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrerequisiteUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrerequisiteUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrerequisiteUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrerequisiteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(prerequisite.Table, prerequisite.Columns, sqlgraph.NewFieldSpec(prerequisite.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Minimum(); ok {
		_spec.SetField(prerequisite.FieldMinimum, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedMinimum(); ok {
		_spec.AddField(prerequisite.FieldMinimum, field.TypeInt, value)
	}
	if pu.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedAbilityScoreIDs(); len(nodes) > 0 && !pu.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prerequisite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PrerequisiteUpdateOne is the builder for updating a single Prerequisite entity.
type PrerequisiteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrerequisiteMutation
}

// SetMinimum sets the "minimum" field.
func (puo *PrerequisiteUpdateOne) SetMinimum(i int) *PrerequisiteUpdateOne {
	puo.mutation.ResetMinimum()
	puo.mutation.SetMinimum(i)
	return puo
}

// AddMinimum adds i to the "minimum" field.
func (puo *PrerequisiteUpdateOne) AddMinimum(i int) *PrerequisiteUpdateOne {
	puo.mutation.AddMinimum(i)
	return puo
}

// AddAbilityScoreIDs adds the "ability_score" edge to the AbilityScore entity by IDs.
func (puo *PrerequisiteUpdateOne) AddAbilityScoreIDs(ids ...int) *PrerequisiteUpdateOne {
	puo.mutation.AddAbilityScoreIDs(ids...)
	return puo
}

// AddAbilityScore adds the "ability_score" edges to the AbilityScore entity.
func (puo *PrerequisiteUpdateOne) AddAbilityScore(a ...*AbilityScore) *PrerequisiteUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.AddAbilityScoreIDs(ids...)
}

// Mutation returns the PrerequisiteMutation object of the builder.
func (puo *PrerequisiteUpdateOne) Mutation() *PrerequisiteMutation {
	return puo.mutation
}

// ClearAbilityScore clears all "ability_score" edges to the AbilityScore entity.
func (puo *PrerequisiteUpdateOne) ClearAbilityScore() *PrerequisiteUpdateOne {
	puo.mutation.ClearAbilityScore()
	return puo
}

// RemoveAbilityScoreIDs removes the "ability_score" edge to AbilityScore entities by IDs.
func (puo *PrerequisiteUpdateOne) RemoveAbilityScoreIDs(ids ...int) *PrerequisiteUpdateOne {
	puo.mutation.RemoveAbilityScoreIDs(ids...)
	return puo
}

// RemoveAbilityScore removes "ability_score" edges to AbilityScore entities.
func (puo *PrerequisiteUpdateOne) RemoveAbilityScore(a ...*AbilityScore) *PrerequisiteUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return puo.RemoveAbilityScoreIDs(ids...)
}

// Where appends a list predicates to the PrerequisiteUpdate builder.
func (puo *PrerequisiteUpdateOne) Where(ps ...predicate.Prerequisite) *PrerequisiteUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PrerequisiteUpdateOne) Select(field string, fields ...string) *PrerequisiteUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Prerequisite entity.
func (puo *PrerequisiteUpdateOne) Save(ctx context.Context) (*Prerequisite, error) {
	return withHooks[*Prerequisite, PrerequisiteMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrerequisiteUpdateOne) SaveX(ctx context.Context) *Prerequisite {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PrerequisiteUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrerequisiteUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrerequisiteUpdateOne) sqlSave(ctx context.Context) (_node *Prerequisite, err error) {
	_spec := sqlgraph.NewUpdateSpec(prerequisite.Table, prerequisite.Columns, sqlgraph.NewFieldSpec(prerequisite.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Prerequisite.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prerequisite.FieldID)
		for _, f := range fields {
			if !prerequisite.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != prerequisite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Minimum(); ok {
		_spec.SetField(prerequisite.FieldMinimum, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedMinimum(); ok {
		_spec.AddField(prerequisite.FieldMinimum, field.TypeInt, value)
	}
	if puo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedAbilityScoreIDs(); len(nodes) > 0 && !puo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prerequisite.AbilityScoreTable,
			Columns: []string{prerequisite.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Prerequisite{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prerequisite.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
