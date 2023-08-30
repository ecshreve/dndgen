// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/trait"
)

// TraitUpdate is the builder for updating Trait entities.
type TraitUpdate struct {
	config
	hooks    []Hook
	mutation *TraitMutation
}

// Where appends a list predicates to the TraitUpdate builder.
func (tu *TraitUpdate) Where(ps ...predicate.Trait) *TraitUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetIndx sets the "indx" field.
func (tu *TraitUpdate) SetIndx(s string) *TraitUpdate {
	tu.mutation.SetIndx(s)
	return tu
}

// SetName sets the "name" field.
func (tu *TraitUpdate) SetName(s string) *TraitUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDesc sets the "desc" field.
func (tu *TraitUpdate) SetDesc(s []string) *TraitUpdate {
	tu.mutation.SetDesc(s)
	return tu
}

// AppendDesc appends s to the "desc" field.
func (tu *TraitUpdate) AppendDesc(s []string) *TraitUpdate {
	tu.mutation.AppendDesc(s)
	return tu
}

// AddRaceIDs adds the "races" edge to the Race entity by IDs.
func (tu *TraitUpdate) AddRaceIDs(ids ...int) *TraitUpdate {
	tu.mutation.AddRaceIDs(ids...)
	return tu
}

// AddRaces adds the "races" edges to the Race entity.
func (tu *TraitUpdate) AddRaces(r ...*Race) *TraitUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddRaceIDs(ids...)
}

// AddSubraceIDs adds the "subraces" edge to the Subrace entity by IDs.
func (tu *TraitUpdate) AddSubraceIDs(ids ...int) *TraitUpdate {
	tu.mutation.AddSubraceIDs(ids...)
	return tu
}

// AddSubraces adds the "subraces" edges to the Subrace entity.
func (tu *TraitUpdate) AddSubraces(s ...*Subrace) *TraitUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddSubraceIDs(ids...)
}

// Mutation returns the TraitMutation object of the builder.
func (tu *TraitUpdate) Mutation() *TraitMutation {
	return tu.mutation
}

// ClearRaces clears all "races" edges to the Race entity.
func (tu *TraitUpdate) ClearRaces() *TraitUpdate {
	tu.mutation.ClearRaces()
	return tu
}

// RemoveRaceIDs removes the "races" edge to Race entities by IDs.
func (tu *TraitUpdate) RemoveRaceIDs(ids ...int) *TraitUpdate {
	tu.mutation.RemoveRaceIDs(ids...)
	return tu
}

// RemoveRaces removes "races" edges to Race entities.
func (tu *TraitUpdate) RemoveRaces(r ...*Race) *TraitUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveRaceIDs(ids...)
}

// ClearSubraces clears all "subraces" edges to the Subrace entity.
func (tu *TraitUpdate) ClearSubraces() *TraitUpdate {
	tu.mutation.ClearSubraces()
	return tu
}

// RemoveSubraceIDs removes the "subraces" edge to Subrace entities by IDs.
func (tu *TraitUpdate) RemoveSubraceIDs(ids ...int) *TraitUpdate {
	tu.mutation.RemoveSubraceIDs(ids...)
	return tu
}

// RemoveSubraces removes "subraces" edges to Subrace entities.
func (tu *TraitUpdate) RemoveSubraces(s ...*Subrace) *TraitUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveSubraceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TraitUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TraitUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TraitUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TraitUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TraitUpdate) check() error {
	if v, ok := tu.mutation.Indx(); ok {
		if err := trait.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Trait.indx": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Name(); ok {
		if err := trait.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Trait.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TraitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(trait.Table, trait.Columns, sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Indx(); ok {
		_spec.SetField(trait.FieldIndx, field.TypeString, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(trait.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Desc(); ok {
		_spec.SetField(trait.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, trait.FieldDesc, value)
		})
	}
	if tu.mutation.RacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RacesTable,
			Columns: trait.RacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedRacesIDs(); len(nodes) > 0 && !tu.mutation.RacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RacesTable,
			Columns: trait.RacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RacesTable,
			Columns: trait.RacesPrimaryKey,
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
	if tu.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.SubracesTable,
			Columns: trait.SubracesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSubracesIDs(); len(nodes) > 0 && !tu.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.SubracesTable,
			Columns: trait.SubracesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SubracesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.SubracesTable,
			Columns: trait.SubracesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TraitUpdateOne is the builder for updating a single Trait entity.
type TraitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TraitMutation
}

// SetIndx sets the "indx" field.
func (tuo *TraitUpdateOne) SetIndx(s string) *TraitUpdateOne {
	tuo.mutation.SetIndx(s)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TraitUpdateOne) SetName(s string) *TraitUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDesc sets the "desc" field.
func (tuo *TraitUpdateOne) SetDesc(s []string) *TraitUpdateOne {
	tuo.mutation.SetDesc(s)
	return tuo
}

// AppendDesc appends s to the "desc" field.
func (tuo *TraitUpdateOne) AppendDesc(s []string) *TraitUpdateOne {
	tuo.mutation.AppendDesc(s)
	return tuo
}

// AddRaceIDs adds the "races" edge to the Race entity by IDs.
func (tuo *TraitUpdateOne) AddRaceIDs(ids ...int) *TraitUpdateOne {
	tuo.mutation.AddRaceIDs(ids...)
	return tuo
}

// AddRaces adds the "races" edges to the Race entity.
func (tuo *TraitUpdateOne) AddRaces(r ...*Race) *TraitUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddRaceIDs(ids...)
}

// AddSubraceIDs adds the "subraces" edge to the Subrace entity by IDs.
func (tuo *TraitUpdateOne) AddSubraceIDs(ids ...int) *TraitUpdateOne {
	tuo.mutation.AddSubraceIDs(ids...)
	return tuo
}

// AddSubraces adds the "subraces" edges to the Subrace entity.
func (tuo *TraitUpdateOne) AddSubraces(s ...*Subrace) *TraitUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddSubraceIDs(ids...)
}

// Mutation returns the TraitMutation object of the builder.
func (tuo *TraitUpdateOne) Mutation() *TraitMutation {
	return tuo.mutation
}

// ClearRaces clears all "races" edges to the Race entity.
func (tuo *TraitUpdateOne) ClearRaces() *TraitUpdateOne {
	tuo.mutation.ClearRaces()
	return tuo
}

// RemoveRaceIDs removes the "races" edge to Race entities by IDs.
func (tuo *TraitUpdateOne) RemoveRaceIDs(ids ...int) *TraitUpdateOne {
	tuo.mutation.RemoveRaceIDs(ids...)
	return tuo
}

// RemoveRaces removes "races" edges to Race entities.
func (tuo *TraitUpdateOne) RemoveRaces(r ...*Race) *TraitUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveRaceIDs(ids...)
}

// ClearSubraces clears all "subraces" edges to the Subrace entity.
func (tuo *TraitUpdateOne) ClearSubraces() *TraitUpdateOne {
	tuo.mutation.ClearSubraces()
	return tuo
}

// RemoveSubraceIDs removes the "subraces" edge to Subrace entities by IDs.
func (tuo *TraitUpdateOne) RemoveSubraceIDs(ids ...int) *TraitUpdateOne {
	tuo.mutation.RemoveSubraceIDs(ids...)
	return tuo
}

// RemoveSubraces removes "subraces" edges to Subrace entities.
func (tuo *TraitUpdateOne) RemoveSubraces(s ...*Subrace) *TraitUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveSubraceIDs(ids...)
}

// Where appends a list predicates to the TraitUpdate builder.
func (tuo *TraitUpdateOne) Where(ps ...predicate.Trait) *TraitUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TraitUpdateOne) Select(field string, fields ...string) *TraitUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Trait entity.
func (tuo *TraitUpdateOne) Save(ctx context.Context) (*Trait, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TraitUpdateOne) SaveX(ctx context.Context) *Trait {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TraitUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TraitUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TraitUpdateOne) check() error {
	if v, ok := tuo.mutation.Indx(); ok {
		if err := trait.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Trait.indx": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Name(); ok {
		if err := trait.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Trait.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TraitUpdateOne) sqlSave(ctx context.Context) (_node *Trait, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(trait.Table, trait.Columns, sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Trait.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trait.FieldID)
		for _, f := range fields {
			if !trait.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != trait.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Indx(); ok {
		_spec.SetField(trait.FieldIndx, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(trait.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Desc(); ok {
		_spec.SetField(trait.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, trait.FieldDesc, value)
		})
	}
	if tuo.mutation.RacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RacesTable,
			Columns: trait.RacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedRacesIDs(); len(nodes) > 0 && !tuo.mutation.RacesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RacesTable,
			Columns: trait.RacesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RacesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.RacesTable,
			Columns: trait.RacesPrimaryKey,
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
	if tuo.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.SubracesTable,
			Columns: trait.SubracesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSubracesIDs(); len(nodes) > 0 && !tuo.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.SubracesTable,
			Columns: trait.SubracesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SubracesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   trait.SubracesTable,
			Columns: trait.SubracesPrimaryKey,
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
	_node = &Trait{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}