// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/languagechoice"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
)

// LanguageChoiceUpdate is the builder for updating LanguageChoice entities.
type LanguageChoiceUpdate struct {
	config
	hooks    []Hook
	mutation *LanguageChoiceMutation
}

// Where appends a list predicates to the LanguageChoiceUpdate builder.
func (lcu *LanguageChoiceUpdate) Where(ps ...predicate.LanguageChoice) *LanguageChoiceUpdate {
	lcu.mutation.Where(ps...)
	return lcu
}

// SetChoose sets the "choose" field.
func (lcu *LanguageChoiceUpdate) SetChoose(i int) *LanguageChoiceUpdate {
	lcu.mutation.ResetChoose()
	lcu.mutation.SetChoose(i)
	return lcu
}

// SetNillableChoose sets the "choose" field if the given value is not nil.
func (lcu *LanguageChoiceUpdate) SetNillableChoose(i *int) *LanguageChoiceUpdate {
	if i != nil {
		lcu.SetChoose(*i)
	}
	return lcu
}

// AddChoose adds i to the "choose" field.
func (lcu *LanguageChoiceUpdate) AddChoose(i int) *LanguageChoiceUpdate {
	lcu.mutation.AddChoose(i)
	return lcu
}

// AddLanguageIDs adds the "languages" edge to the Language entity by IDs.
func (lcu *LanguageChoiceUpdate) AddLanguageIDs(ids ...int) *LanguageChoiceUpdate {
	lcu.mutation.AddLanguageIDs(ids...)
	return lcu
}

// AddLanguages adds the "languages" edges to the Language entity.
func (lcu *LanguageChoiceUpdate) AddLanguages(l ...*Language) *LanguageChoiceUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lcu.AddLanguageIDs(ids...)
}

// SetRaceID sets the "race" edge to the Race entity by ID.
func (lcu *LanguageChoiceUpdate) SetRaceID(id int) *LanguageChoiceUpdate {
	lcu.mutation.SetRaceID(id)
	return lcu
}

// SetNillableRaceID sets the "race" edge to the Race entity by ID if the given value is not nil.
func (lcu *LanguageChoiceUpdate) SetNillableRaceID(id *int) *LanguageChoiceUpdate {
	if id != nil {
		lcu = lcu.SetRaceID(*id)
	}
	return lcu
}

// SetRace sets the "race" edge to the Race entity.
func (lcu *LanguageChoiceUpdate) SetRace(r *Race) *LanguageChoiceUpdate {
	return lcu.SetRaceID(r.ID)
}

// SetSubraceID sets the "subrace" edge to the Subrace entity by ID.
func (lcu *LanguageChoiceUpdate) SetSubraceID(id int) *LanguageChoiceUpdate {
	lcu.mutation.SetSubraceID(id)
	return lcu
}

// SetNillableSubraceID sets the "subrace" edge to the Subrace entity by ID if the given value is not nil.
func (lcu *LanguageChoiceUpdate) SetNillableSubraceID(id *int) *LanguageChoiceUpdate {
	if id != nil {
		lcu = lcu.SetSubraceID(*id)
	}
	return lcu
}

// SetSubrace sets the "subrace" edge to the Subrace entity.
func (lcu *LanguageChoiceUpdate) SetSubrace(s *Subrace) *LanguageChoiceUpdate {
	return lcu.SetSubraceID(s.ID)
}

// Mutation returns the LanguageChoiceMutation object of the builder.
func (lcu *LanguageChoiceUpdate) Mutation() *LanguageChoiceMutation {
	return lcu.mutation
}

// ClearLanguages clears all "languages" edges to the Language entity.
func (lcu *LanguageChoiceUpdate) ClearLanguages() *LanguageChoiceUpdate {
	lcu.mutation.ClearLanguages()
	return lcu
}

// RemoveLanguageIDs removes the "languages" edge to Language entities by IDs.
func (lcu *LanguageChoiceUpdate) RemoveLanguageIDs(ids ...int) *LanguageChoiceUpdate {
	lcu.mutation.RemoveLanguageIDs(ids...)
	return lcu
}

// RemoveLanguages removes "languages" edges to Language entities.
func (lcu *LanguageChoiceUpdate) RemoveLanguages(l ...*Language) *LanguageChoiceUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lcu.RemoveLanguageIDs(ids...)
}

// ClearRace clears the "race" edge to the Race entity.
func (lcu *LanguageChoiceUpdate) ClearRace() *LanguageChoiceUpdate {
	lcu.mutation.ClearRace()
	return lcu
}

// ClearSubrace clears the "subrace" edge to the Subrace entity.
func (lcu *LanguageChoiceUpdate) ClearSubrace() *LanguageChoiceUpdate {
	lcu.mutation.ClearSubrace()
	return lcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lcu *LanguageChoiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lcu.sqlSave, lcu.mutation, lcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lcu *LanguageChoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := lcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lcu *LanguageChoiceUpdate) Exec(ctx context.Context) error {
	_, err := lcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcu *LanguageChoiceUpdate) ExecX(ctx context.Context) {
	if err := lcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcu *LanguageChoiceUpdate) check() error {
	if v, ok := lcu.mutation.Choose(); ok {
		if err := languagechoice.ChooseValidator(v); err != nil {
			return &ValidationError{Name: "choose", err: fmt.Errorf(`ent: validator failed for field "LanguageChoice.choose": %w`, err)}
		}
	}
	return nil
}

func (lcu *LanguageChoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(languagechoice.Table, languagechoice.Columns, sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt))
	if ps := lcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lcu.mutation.Choose(); ok {
		_spec.SetField(languagechoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := lcu.mutation.AddedChoose(); ok {
		_spec.AddField(languagechoice.FieldChoose, field.TypeInt, value)
	}
	if lcu.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   languagechoice.LanguagesTable,
			Columns: languagechoice.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.RemovedLanguagesIDs(); len(nodes) > 0 && !lcu.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   languagechoice.LanguagesTable,
			Columns: languagechoice.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.LanguagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   languagechoice.LanguagesTable,
			Columns: languagechoice.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lcu.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   languagechoice.RaceTable,
			Columns: []string{languagechoice.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   languagechoice.RaceTable,
			Columns: []string{languagechoice.RaceColumn},
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
	if lcu.mutation.SubraceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   languagechoice.SubraceTable,
			Columns: []string{languagechoice.SubraceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcu.mutation.SubraceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   languagechoice.SubraceTable,
			Columns: []string{languagechoice.SubraceColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, lcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{languagechoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lcu.mutation.done = true
	return n, nil
}

// LanguageChoiceUpdateOne is the builder for updating a single LanguageChoice entity.
type LanguageChoiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LanguageChoiceMutation
}

// SetChoose sets the "choose" field.
func (lcuo *LanguageChoiceUpdateOne) SetChoose(i int) *LanguageChoiceUpdateOne {
	lcuo.mutation.ResetChoose()
	lcuo.mutation.SetChoose(i)
	return lcuo
}

// SetNillableChoose sets the "choose" field if the given value is not nil.
func (lcuo *LanguageChoiceUpdateOne) SetNillableChoose(i *int) *LanguageChoiceUpdateOne {
	if i != nil {
		lcuo.SetChoose(*i)
	}
	return lcuo
}

// AddChoose adds i to the "choose" field.
func (lcuo *LanguageChoiceUpdateOne) AddChoose(i int) *LanguageChoiceUpdateOne {
	lcuo.mutation.AddChoose(i)
	return lcuo
}

// AddLanguageIDs adds the "languages" edge to the Language entity by IDs.
func (lcuo *LanguageChoiceUpdateOne) AddLanguageIDs(ids ...int) *LanguageChoiceUpdateOne {
	lcuo.mutation.AddLanguageIDs(ids...)
	return lcuo
}

// AddLanguages adds the "languages" edges to the Language entity.
func (lcuo *LanguageChoiceUpdateOne) AddLanguages(l ...*Language) *LanguageChoiceUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lcuo.AddLanguageIDs(ids...)
}

// SetRaceID sets the "race" edge to the Race entity by ID.
func (lcuo *LanguageChoiceUpdateOne) SetRaceID(id int) *LanguageChoiceUpdateOne {
	lcuo.mutation.SetRaceID(id)
	return lcuo
}

// SetNillableRaceID sets the "race" edge to the Race entity by ID if the given value is not nil.
func (lcuo *LanguageChoiceUpdateOne) SetNillableRaceID(id *int) *LanguageChoiceUpdateOne {
	if id != nil {
		lcuo = lcuo.SetRaceID(*id)
	}
	return lcuo
}

// SetRace sets the "race" edge to the Race entity.
func (lcuo *LanguageChoiceUpdateOne) SetRace(r *Race) *LanguageChoiceUpdateOne {
	return lcuo.SetRaceID(r.ID)
}

// SetSubraceID sets the "subrace" edge to the Subrace entity by ID.
func (lcuo *LanguageChoiceUpdateOne) SetSubraceID(id int) *LanguageChoiceUpdateOne {
	lcuo.mutation.SetSubraceID(id)
	return lcuo
}

// SetNillableSubraceID sets the "subrace" edge to the Subrace entity by ID if the given value is not nil.
func (lcuo *LanguageChoiceUpdateOne) SetNillableSubraceID(id *int) *LanguageChoiceUpdateOne {
	if id != nil {
		lcuo = lcuo.SetSubraceID(*id)
	}
	return lcuo
}

// SetSubrace sets the "subrace" edge to the Subrace entity.
func (lcuo *LanguageChoiceUpdateOne) SetSubrace(s *Subrace) *LanguageChoiceUpdateOne {
	return lcuo.SetSubraceID(s.ID)
}

// Mutation returns the LanguageChoiceMutation object of the builder.
func (lcuo *LanguageChoiceUpdateOne) Mutation() *LanguageChoiceMutation {
	return lcuo.mutation
}

// ClearLanguages clears all "languages" edges to the Language entity.
func (lcuo *LanguageChoiceUpdateOne) ClearLanguages() *LanguageChoiceUpdateOne {
	lcuo.mutation.ClearLanguages()
	return lcuo
}

// RemoveLanguageIDs removes the "languages" edge to Language entities by IDs.
func (lcuo *LanguageChoiceUpdateOne) RemoveLanguageIDs(ids ...int) *LanguageChoiceUpdateOne {
	lcuo.mutation.RemoveLanguageIDs(ids...)
	return lcuo
}

// RemoveLanguages removes "languages" edges to Language entities.
func (lcuo *LanguageChoiceUpdateOne) RemoveLanguages(l ...*Language) *LanguageChoiceUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lcuo.RemoveLanguageIDs(ids...)
}

// ClearRace clears the "race" edge to the Race entity.
func (lcuo *LanguageChoiceUpdateOne) ClearRace() *LanguageChoiceUpdateOne {
	lcuo.mutation.ClearRace()
	return lcuo
}

// ClearSubrace clears the "subrace" edge to the Subrace entity.
func (lcuo *LanguageChoiceUpdateOne) ClearSubrace() *LanguageChoiceUpdateOne {
	lcuo.mutation.ClearSubrace()
	return lcuo
}

// Where appends a list predicates to the LanguageChoiceUpdate builder.
func (lcuo *LanguageChoiceUpdateOne) Where(ps ...predicate.LanguageChoice) *LanguageChoiceUpdateOne {
	lcuo.mutation.Where(ps...)
	return lcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lcuo *LanguageChoiceUpdateOne) Select(field string, fields ...string) *LanguageChoiceUpdateOne {
	lcuo.fields = append([]string{field}, fields...)
	return lcuo
}

// Save executes the query and returns the updated LanguageChoice entity.
func (lcuo *LanguageChoiceUpdateOne) Save(ctx context.Context) (*LanguageChoice, error) {
	return withHooks(ctx, lcuo.sqlSave, lcuo.mutation, lcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lcuo *LanguageChoiceUpdateOne) SaveX(ctx context.Context) *LanguageChoice {
	node, err := lcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lcuo *LanguageChoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := lcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcuo *LanguageChoiceUpdateOne) ExecX(ctx context.Context) {
	if err := lcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lcuo *LanguageChoiceUpdateOne) check() error {
	if v, ok := lcuo.mutation.Choose(); ok {
		if err := languagechoice.ChooseValidator(v); err != nil {
			return &ValidationError{Name: "choose", err: fmt.Errorf(`ent: validator failed for field "LanguageChoice.choose": %w`, err)}
		}
	}
	return nil
}

func (lcuo *LanguageChoiceUpdateOne) sqlSave(ctx context.Context) (_node *LanguageChoice, err error) {
	if err := lcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(languagechoice.Table, languagechoice.Columns, sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt))
	id, ok := lcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LanguageChoice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, languagechoice.FieldID)
		for _, f := range fields {
			if !languagechoice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != languagechoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lcuo.mutation.Choose(); ok {
		_spec.SetField(languagechoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := lcuo.mutation.AddedChoose(); ok {
		_spec.AddField(languagechoice.FieldChoose, field.TypeInt, value)
	}
	if lcuo.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   languagechoice.LanguagesTable,
			Columns: languagechoice.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.RemovedLanguagesIDs(); len(nodes) > 0 && !lcuo.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   languagechoice.LanguagesTable,
			Columns: languagechoice.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.LanguagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   languagechoice.LanguagesTable,
			Columns: languagechoice.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lcuo.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   languagechoice.RaceTable,
			Columns: []string{languagechoice.RaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   languagechoice.RaceTable,
			Columns: []string{languagechoice.RaceColumn},
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
	if lcuo.mutation.SubraceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   languagechoice.SubraceTable,
			Columns: []string{languagechoice.SubraceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lcuo.mutation.SubraceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   languagechoice.SubraceTable,
			Columns: []string{languagechoice.SubraceColumn},
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
	_node = &LanguageChoice{config: lcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{languagechoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lcuo.mutation.done = true
	return _node, nil
}