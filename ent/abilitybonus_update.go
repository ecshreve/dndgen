// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// AbilityBonusUpdate is the builder for updating AbilityBonus entities.
type AbilityBonusUpdate struct {
	config
	hooks    []Hook
	mutation *AbilityBonusMutation
}

// Where appends a list predicates to the AbilityBonusUpdate builder.
func (abu *AbilityBonusUpdate) Where(ps ...predicate.AbilityBonus) *AbilityBonusUpdate {
	abu.mutation.Where(ps...)
	return abu
}

// SetValue sets the "value" field.
func (abu *AbilityBonusUpdate) SetValue(i int) *AbilityBonusUpdate {
	abu.mutation.ResetValue()
	abu.mutation.SetValue(i)
	return abu
}

// AddValue adds i to the "value" field.
func (abu *AbilityBonusUpdate) AddValue(i int) *AbilityBonusUpdate {
	abu.mutation.AddValue(i)
	return abu
}

// AddAbilityScoreIDs adds the "ability_score" edge to the AbilityScore entity by IDs.
func (abu *AbilityBonusUpdate) AddAbilityScoreIDs(ids ...int) *AbilityBonusUpdate {
	abu.mutation.AddAbilityScoreIDs(ids...)
	return abu
}

// AddAbilityScore adds the "ability_score" edges to the AbilityScore entity.
func (abu *AbilityBonusUpdate) AddAbilityScore(a ...*AbilityScore) *AbilityBonusUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return abu.AddAbilityScoreIDs(ids...)
}

// Mutation returns the AbilityBonusMutation object of the builder.
func (abu *AbilityBonusUpdate) Mutation() *AbilityBonusMutation {
	return abu.mutation
}

// ClearAbilityScore clears all "ability_score" edges to the AbilityScore entity.
func (abu *AbilityBonusUpdate) ClearAbilityScore() *AbilityBonusUpdate {
	abu.mutation.ClearAbilityScore()
	return abu
}

// RemoveAbilityScoreIDs removes the "ability_score" edge to AbilityScore entities by IDs.
func (abu *AbilityBonusUpdate) RemoveAbilityScoreIDs(ids ...int) *AbilityBonusUpdate {
	abu.mutation.RemoveAbilityScoreIDs(ids...)
	return abu
}

// RemoveAbilityScore removes "ability_score" edges to AbilityScore entities.
func (abu *AbilityBonusUpdate) RemoveAbilityScore(a ...*AbilityScore) *AbilityBonusUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return abu.RemoveAbilityScoreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (abu *AbilityBonusUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, abu.sqlSave, abu.mutation, abu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (abu *AbilityBonusUpdate) SaveX(ctx context.Context) int {
	affected, err := abu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (abu *AbilityBonusUpdate) Exec(ctx context.Context) error {
	_, err := abu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abu *AbilityBonusUpdate) ExecX(ctx context.Context) {
	if err := abu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (abu *AbilityBonusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(abilitybonus.Table, abilitybonus.Columns, sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt))
	if ps := abu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := abu.mutation.Value(); ok {
		_spec.SetField(abilitybonus.FieldValue, field.TypeInt, value)
	}
	if value, ok := abu.mutation.AddedValue(); ok {
		_spec.AddField(abilitybonus.FieldValue, field.TypeInt, value)
	}
	if abu.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: abilitybonus.AbilityScorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abu.mutation.RemovedAbilityScoreIDs(); len(nodes) > 0 && !abu.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: abilitybonus.AbilityScorePrimaryKey,
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
	if nodes := abu.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: abilitybonus.AbilityScorePrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, abu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{abilitybonus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	abu.mutation.done = true
	return n, nil
}

// AbilityBonusUpdateOne is the builder for updating a single AbilityBonus entity.
type AbilityBonusUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AbilityBonusMutation
}

// SetValue sets the "value" field.
func (abuo *AbilityBonusUpdateOne) SetValue(i int) *AbilityBonusUpdateOne {
	abuo.mutation.ResetValue()
	abuo.mutation.SetValue(i)
	return abuo
}

// AddValue adds i to the "value" field.
func (abuo *AbilityBonusUpdateOne) AddValue(i int) *AbilityBonusUpdateOne {
	abuo.mutation.AddValue(i)
	return abuo
}

// AddAbilityScoreIDs adds the "ability_score" edge to the AbilityScore entity by IDs.
func (abuo *AbilityBonusUpdateOne) AddAbilityScoreIDs(ids ...int) *AbilityBonusUpdateOne {
	abuo.mutation.AddAbilityScoreIDs(ids...)
	return abuo
}

// AddAbilityScore adds the "ability_score" edges to the AbilityScore entity.
func (abuo *AbilityBonusUpdateOne) AddAbilityScore(a ...*AbilityScore) *AbilityBonusUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return abuo.AddAbilityScoreIDs(ids...)
}

// Mutation returns the AbilityBonusMutation object of the builder.
func (abuo *AbilityBonusUpdateOne) Mutation() *AbilityBonusMutation {
	return abuo.mutation
}

// ClearAbilityScore clears all "ability_score" edges to the AbilityScore entity.
func (abuo *AbilityBonusUpdateOne) ClearAbilityScore() *AbilityBonusUpdateOne {
	abuo.mutation.ClearAbilityScore()
	return abuo
}

// RemoveAbilityScoreIDs removes the "ability_score" edge to AbilityScore entities by IDs.
func (abuo *AbilityBonusUpdateOne) RemoveAbilityScoreIDs(ids ...int) *AbilityBonusUpdateOne {
	abuo.mutation.RemoveAbilityScoreIDs(ids...)
	return abuo
}

// RemoveAbilityScore removes "ability_score" edges to AbilityScore entities.
func (abuo *AbilityBonusUpdateOne) RemoveAbilityScore(a ...*AbilityScore) *AbilityBonusUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return abuo.RemoveAbilityScoreIDs(ids...)
}

// Where appends a list predicates to the AbilityBonusUpdate builder.
func (abuo *AbilityBonusUpdateOne) Where(ps ...predicate.AbilityBonus) *AbilityBonusUpdateOne {
	abuo.mutation.Where(ps...)
	return abuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (abuo *AbilityBonusUpdateOne) Select(field string, fields ...string) *AbilityBonusUpdateOne {
	abuo.fields = append([]string{field}, fields...)
	return abuo
}

// Save executes the query and returns the updated AbilityBonus entity.
func (abuo *AbilityBonusUpdateOne) Save(ctx context.Context) (*AbilityBonus, error) {
	return withHooks(ctx, abuo.sqlSave, abuo.mutation, abuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (abuo *AbilityBonusUpdateOne) SaveX(ctx context.Context) *AbilityBonus {
	node, err := abuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (abuo *AbilityBonusUpdateOne) Exec(ctx context.Context) error {
	_, err := abuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (abuo *AbilityBonusUpdateOne) ExecX(ctx context.Context) {
	if err := abuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (abuo *AbilityBonusUpdateOne) sqlSave(ctx context.Context) (_node *AbilityBonus, err error) {
	_spec := sqlgraph.NewUpdateSpec(abilitybonus.Table, abilitybonus.Columns, sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt))
	id, ok := abuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AbilityBonus.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := abuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, abilitybonus.FieldID)
		for _, f := range fields {
			if !abilitybonus.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != abilitybonus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := abuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := abuo.mutation.Value(); ok {
		_spec.SetField(abilitybonus.FieldValue, field.TypeInt, value)
	}
	if value, ok := abuo.mutation.AddedValue(); ok {
		_spec.AddField(abilitybonus.FieldValue, field.TypeInt, value)
	}
	if abuo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: abilitybonus.AbilityScorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abuo.mutation.RemovedAbilityScoreIDs(); len(nodes) > 0 && !abuo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: abilitybonus.AbilityScorePrimaryKey,
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
	if nodes := abuo.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: abilitybonus.AbilityScorePrimaryKey,
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
	_node = &AbilityBonus{config: abuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, abuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{abilitybonus.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	abuo.mutation.done = true
	return _node, nil
}
