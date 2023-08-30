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
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
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

// SetAbilityScoreID sets the "ability_score_id" field.
func (abu *AbilityBonusUpdate) SetAbilityScoreID(i int) *AbilityBonusUpdate {
	abu.mutation.SetAbilityScoreID(i)
	return abu
}

// SetBonus sets the "bonus" field.
func (abu *AbilityBonusUpdate) SetBonus(i int) *AbilityBonusUpdate {
	abu.mutation.ResetBonus()
	abu.mutation.SetBonus(i)
	return abu
}

// AddBonus adds i to the "bonus" field.
func (abu *AbilityBonusUpdate) AddBonus(i int) *AbilityBonusUpdate {
	abu.mutation.AddBonus(i)
	return abu
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (abu *AbilityBonusUpdate) SetAbilityScore(a *AbilityScore) *AbilityBonusUpdate {
	return abu.SetAbilityScoreID(a.ID)
}

// SetRaceID sets the "race" edge to the Race entity by ID.
func (abu *AbilityBonusUpdate) SetRaceID(id int) *AbilityBonusUpdate {
	abu.mutation.SetRaceID(id)
	return abu
}

// SetNillableRaceID sets the "race" edge to the Race entity by ID if the given value is not nil.
func (abu *AbilityBonusUpdate) SetNillableRaceID(id *int) *AbilityBonusUpdate {
	if id != nil {
		abu = abu.SetRaceID(*id)
	}
	return abu
}

// SetRace sets the "race" edge to the Race entity.
func (abu *AbilityBonusUpdate) SetRace(r *Race) *AbilityBonusUpdate {
	return abu.SetRaceID(r.ID)
}

// SetSubraceID sets the "subrace" edge to the Subrace entity by ID.
func (abu *AbilityBonusUpdate) SetSubraceID(id int) *AbilityBonusUpdate {
	abu.mutation.SetSubraceID(id)
	return abu
}

// SetNillableSubraceID sets the "subrace" edge to the Subrace entity by ID if the given value is not nil.
func (abu *AbilityBonusUpdate) SetNillableSubraceID(id *int) *AbilityBonusUpdate {
	if id != nil {
		abu = abu.SetSubraceID(*id)
	}
	return abu
}

// SetSubrace sets the "subrace" edge to the Subrace entity.
func (abu *AbilityBonusUpdate) SetSubrace(s *Subrace) *AbilityBonusUpdate {
	return abu.SetSubraceID(s.ID)
}

// Mutation returns the AbilityBonusMutation object of the builder.
func (abu *AbilityBonusUpdate) Mutation() *AbilityBonusMutation {
	return abu.mutation
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (abu *AbilityBonusUpdate) ClearAbilityScore() *AbilityBonusUpdate {
	abu.mutation.ClearAbilityScore()
	return abu
}

// ClearRace clears the "race" edge to the Race entity.
func (abu *AbilityBonusUpdate) ClearRace() *AbilityBonusUpdate {
	abu.mutation.ClearRace()
	return abu
}

// ClearSubrace clears the "subrace" edge to the Subrace entity.
func (abu *AbilityBonusUpdate) ClearSubrace() *AbilityBonusUpdate {
	abu.mutation.ClearSubrace()
	return abu
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

// check runs all checks and user-defined validators on the builder.
func (abu *AbilityBonusUpdate) check() error {
	if _, ok := abu.mutation.AbilityScoreID(); abu.mutation.AbilityScoreCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AbilityBonus.ability_score"`)
	}
	return nil
}

func (abu *AbilityBonusUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := abu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(abilitybonus.Table, abilitybonus.Columns, sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt))
	if ps := abu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := abu.mutation.Bonus(); ok {
		_spec.SetField(abilitybonus.FieldBonus, field.TypeInt, value)
	}
	if value, ok := abu.mutation.AddedBonus(); ok {
		_spec.AddField(abilitybonus.FieldBonus, field.TypeInt, value)
	}
	if abu.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: []string{abilitybonus.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abu.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: []string{abilitybonus.AbilityScoreColumn},
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
	if abu.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.RaceTable,
			Columns: []string{abilitybonus.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abu.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.RaceTable,
			Columns: []string{abilitybonus.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if abu.mutation.SubraceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.SubraceTable,
			Columns: []string{abilitybonus.SubraceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abu.mutation.SubraceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.SubraceTable,
			Columns: []string{abilitybonus.SubraceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
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

// SetAbilityScoreID sets the "ability_score_id" field.
func (abuo *AbilityBonusUpdateOne) SetAbilityScoreID(i int) *AbilityBonusUpdateOne {
	abuo.mutation.SetAbilityScoreID(i)
	return abuo
}

// SetBonus sets the "bonus" field.
func (abuo *AbilityBonusUpdateOne) SetBonus(i int) *AbilityBonusUpdateOne {
	abuo.mutation.ResetBonus()
	abuo.mutation.SetBonus(i)
	return abuo
}

// AddBonus adds i to the "bonus" field.
func (abuo *AbilityBonusUpdateOne) AddBonus(i int) *AbilityBonusUpdateOne {
	abuo.mutation.AddBonus(i)
	return abuo
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (abuo *AbilityBonusUpdateOne) SetAbilityScore(a *AbilityScore) *AbilityBonusUpdateOne {
	return abuo.SetAbilityScoreID(a.ID)
}

// SetRaceID sets the "race" edge to the Race entity by ID.
func (abuo *AbilityBonusUpdateOne) SetRaceID(id int) *AbilityBonusUpdateOne {
	abuo.mutation.SetRaceID(id)
	return abuo
}

// SetNillableRaceID sets the "race" edge to the Race entity by ID if the given value is not nil.
func (abuo *AbilityBonusUpdateOne) SetNillableRaceID(id *int) *AbilityBonusUpdateOne {
	if id != nil {
		abuo = abuo.SetRaceID(*id)
	}
	return abuo
}

// SetRace sets the "race" edge to the Race entity.
func (abuo *AbilityBonusUpdateOne) SetRace(r *Race) *AbilityBonusUpdateOne {
	return abuo.SetRaceID(r.ID)
}

// SetSubraceID sets the "subrace" edge to the Subrace entity by ID.
func (abuo *AbilityBonusUpdateOne) SetSubraceID(id int) *AbilityBonusUpdateOne {
	abuo.mutation.SetSubraceID(id)
	return abuo
}

// SetNillableSubraceID sets the "subrace" edge to the Subrace entity by ID if the given value is not nil.
func (abuo *AbilityBonusUpdateOne) SetNillableSubraceID(id *int) *AbilityBonusUpdateOne {
	if id != nil {
		abuo = abuo.SetSubraceID(*id)
	}
	return abuo
}

// SetSubrace sets the "subrace" edge to the Subrace entity.
func (abuo *AbilityBonusUpdateOne) SetSubrace(s *Subrace) *AbilityBonusUpdateOne {
	return abuo.SetSubraceID(s.ID)
}

// Mutation returns the AbilityBonusMutation object of the builder.
func (abuo *AbilityBonusUpdateOne) Mutation() *AbilityBonusMutation {
	return abuo.mutation
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (abuo *AbilityBonusUpdateOne) ClearAbilityScore() *AbilityBonusUpdateOne {
	abuo.mutation.ClearAbilityScore()
	return abuo
}

// ClearRace clears the "race" edge to the Race entity.
func (abuo *AbilityBonusUpdateOne) ClearRace() *AbilityBonusUpdateOne {
	abuo.mutation.ClearRace()
	return abuo
}

// ClearSubrace clears the "subrace" edge to the Subrace entity.
func (abuo *AbilityBonusUpdateOne) ClearSubrace() *AbilityBonusUpdateOne {
	abuo.mutation.ClearSubrace()
	return abuo
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

// check runs all checks and user-defined validators on the builder.
func (abuo *AbilityBonusUpdateOne) check() error {
	if _, ok := abuo.mutation.AbilityScoreID(); abuo.mutation.AbilityScoreCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AbilityBonus.ability_score"`)
	}
	return nil
}

func (abuo *AbilityBonusUpdateOne) sqlSave(ctx context.Context) (_node *AbilityBonus, err error) {
	if err := abuo.check(); err != nil {
		return _node, err
	}
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
	if value, ok := abuo.mutation.Bonus(); ok {
		_spec.SetField(abilitybonus.FieldBonus, field.TypeInt, value)
	}
	if value, ok := abuo.mutation.AddedBonus(); ok {
		_spec.AddField(abilitybonus.FieldBonus, field.TypeInt, value)
	}
	if abuo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: []string{abilitybonus.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abuo.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.AbilityScoreTable,
			Columns: []string{abilitybonus.AbilityScoreColumn},
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
	if abuo.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.RaceTable,
			Columns: []string{abilitybonus.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abuo.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.RaceTable,
			Columns: []string{abilitybonus.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if abuo.mutation.SubraceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.SubraceTable,
			Columns: []string{abilitybonus.SubraceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := abuo.mutation.SubraceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   abilitybonus.SubraceTable,
			Columns: []string{abilitybonus.SubraceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
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
