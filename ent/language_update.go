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
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/languagechoice"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
)

// LanguageUpdate is the builder for updating Language entities.
type LanguageUpdate struct {
	config
	hooks    []Hook
	mutation *LanguageMutation
}

// Where appends a list predicates to the LanguageUpdate builder.
func (lu *LanguageUpdate) Where(ps ...predicate.Language) *LanguageUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetIndx sets the "indx" field.
func (lu *LanguageUpdate) SetIndx(s string) *LanguageUpdate {
	lu.mutation.SetIndx(s)
	return lu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableIndx(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetIndx(*s)
	}
	return lu
}

// SetName sets the "name" field.
func (lu *LanguageUpdate) SetName(s string) *LanguageUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableName(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// SetDesc sets the "desc" field.
func (lu *LanguageUpdate) SetDesc(s []string) *LanguageUpdate {
	lu.mutation.SetDesc(s)
	return lu
}

// AppendDesc appends s to the "desc" field.
func (lu *LanguageUpdate) AppendDesc(s []string) *LanguageUpdate {
	lu.mutation.AppendDesc(s)
	return lu
}

// ClearDesc clears the value of the "desc" field.
func (lu *LanguageUpdate) ClearDesc() *LanguageUpdate {
	lu.mutation.ClearDesc()
	return lu
}

// SetLanguageType sets the "language_type" field.
func (lu *LanguageUpdate) SetLanguageType(lt language.LanguageType) *LanguageUpdate {
	lu.mutation.SetLanguageType(lt)
	return lu
}

// SetNillableLanguageType sets the "language_type" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableLanguageType(lt *language.LanguageType) *LanguageUpdate {
	if lt != nil {
		lu.SetLanguageType(*lt)
	}
	return lu
}

// SetScript sets the "script" field.
func (lu *LanguageUpdate) SetScript(l language.Script) *LanguageUpdate {
	lu.mutation.SetScript(l)
	return lu
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableScript(l *language.Script) *LanguageUpdate {
	if l != nil {
		lu.SetScript(*l)
	}
	return lu
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (lu *LanguageUpdate) AddRaceIDs(ids ...int) *LanguageUpdate {
	lu.mutation.AddRaceIDs(ids...)
	return lu
}

// AddRace adds the "race" edges to the Race entity.
func (lu *LanguageUpdate) AddRace(r ...*Race) *LanguageUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return lu.AddRaceIDs(ids...)
}

// AddOptionIDs adds the "options" edge to the LanguageChoice entity by IDs.
func (lu *LanguageUpdate) AddOptionIDs(ids ...int) *LanguageUpdate {
	lu.mutation.AddOptionIDs(ids...)
	return lu
}

// AddOptions adds the "options" edges to the LanguageChoice entity.
func (lu *LanguageUpdate) AddOptions(l ...*LanguageChoice) *LanguageUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.AddOptionIDs(ids...)
}

// Mutation returns the LanguageMutation object of the builder.
func (lu *LanguageUpdate) Mutation() *LanguageMutation {
	return lu.mutation
}

// ClearRace clears all "race" edges to the Race entity.
func (lu *LanguageUpdate) ClearRace() *LanguageUpdate {
	lu.mutation.ClearRace()
	return lu
}

// RemoveRaceIDs removes the "race" edge to Race entities by IDs.
func (lu *LanguageUpdate) RemoveRaceIDs(ids ...int) *LanguageUpdate {
	lu.mutation.RemoveRaceIDs(ids...)
	return lu
}

// RemoveRace removes "race" edges to Race entities.
func (lu *LanguageUpdate) RemoveRace(r ...*Race) *LanguageUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return lu.RemoveRaceIDs(ids...)
}

// ClearOptions clears all "options" edges to the LanguageChoice entity.
func (lu *LanguageUpdate) ClearOptions() *LanguageUpdate {
	lu.mutation.ClearOptions()
	return lu
}

// RemoveOptionIDs removes the "options" edge to LanguageChoice entities by IDs.
func (lu *LanguageUpdate) RemoveOptionIDs(ids ...int) *LanguageUpdate {
	lu.mutation.RemoveOptionIDs(ids...)
	return lu
}

// RemoveOptions removes "options" edges to LanguageChoice entities.
func (lu *LanguageUpdate) RemoveOptions(l ...*LanguageChoice) *LanguageUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lu.RemoveOptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LanguageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LanguageUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LanguageUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LanguageUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lu *LanguageUpdate) check() error {
	if v, ok := lu.mutation.Indx(); ok {
		if err := language.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Language.indx": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Name(); ok {
		if err := language.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Language.name": %w`, err)}
		}
	}
	if v, ok := lu.mutation.LanguageType(); ok {
		if err := language.LanguageTypeValidator(v); err != nil {
			return &ValidationError{Name: "language_type", err: fmt.Errorf(`ent: validator failed for field "Language.language_type": %w`, err)}
		}
	}
	if v, ok := lu.mutation.Script(); ok {
		if err := language.ScriptValidator(v); err != nil {
			return &ValidationError{Name: "script", err: fmt.Errorf(`ent: validator failed for field "Language.script": %w`, err)}
		}
	}
	return nil
}

func (lu *LanguageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := lu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.Indx(); ok {
		_spec.SetField(language.FieldIndx, field.TypeString, value)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(language.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Desc(); ok {
		_spec.SetField(language.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := lu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, language.FieldDesc, value)
		})
	}
	if lu.mutation.DescCleared() {
		_spec.ClearField(language.FieldDesc, field.TypeJSON)
	}
	if value, ok := lu.mutation.LanguageType(); ok {
		_spec.SetField(language.FieldLanguageType, field.TypeEnum, value)
	}
	if value, ok := lu.mutation.Script(); ok {
		_spec.SetField(language.FieldScript, field.TypeEnum, value)
	}
	if lu.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.RaceTable,
			Columns: language.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedRaceIDs(); len(nodes) > 0 && !lu.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.RaceTable,
			Columns: language.RacePrimaryKey,
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
	if nodes := lu.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.RaceTable,
			Columns: language.RacePrimaryKey,
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
	if lu.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.OptionsTable,
			Columns: language.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedOptionsIDs(); len(nodes) > 0 && !lu.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.OptionsTable,
			Columns: language.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.OptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.OptionsTable,
			Columns: language.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LanguageUpdateOne is the builder for updating a single Language entity.
type LanguageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LanguageMutation
}

// SetIndx sets the "indx" field.
func (luo *LanguageUpdateOne) SetIndx(s string) *LanguageUpdateOne {
	luo.mutation.SetIndx(s)
	return luo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableIndx(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetIndx(*s)
	}
	return luo
}

// SetName sets the "name" field.
func (luo *LanguageUpdateOne) SetName(s string) *LanguageUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableName(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// SetDesc sets the "desc" field.
func (luo *LanguageUpdateOne) SetDesc(s []string) *LanguageUpdateOne {
	luo.mutation.SetDesc(s)
	return luo
}

// AppendDesc appends s to the "desc" field.
func (luo *LanguageUpdateOne) AppendDesc(s []string) *LanguageUpdateOne {
	luo.mutation.AppendDesc(s)
	return luo
}

// ClearDesc clears the value of the "desc" field.
func (luo *LanguageUpdateOne) ClearDesc() *LanguageUpdateOne {
	luo.mutation.ClearDesc()
	return luo
}

// SetLanguageType sets the "language_type" field.
func (luo *LanguageUpdateOne) SetLanguageType(lt language.LanguageType) *LanguageUpdateOne {
	luo.mutation.SetLanguageType(lt)
	return luo
}

// SetNillableLanguageType sets the "language_type" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableLanguageType(lt *language.LanguageType) *LanguageUpdateOne {
	if lt != nil {
		luo.SetLanguageType(*lt)
	}
	return luo
}

// SetScript sets the "script" field.
func (luo *LanguageUpdateOne) SetScript(l language.Script) *LanguageUpdateOne {
	luo.mutation.SetScript(l)
	return luo
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableScript(l *language.Script) *LanguageUpdateOne {
	if l != nil {
		luo.SetScript(*l)
	}
	return luo
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (luo *LanguageUpdateOne) AddRaceIDs(ids ...int) *LanguageUpdateOne {
	luo.mutation.AddRaceIDs(ids...)
	return luo
}

// AddRace adds the "race" edges to the Race entity.
func (luo *LanguageUpdateOne) AddRace(r ...*Race) *LanguageUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return luo.AddRaceIDs(ids...)
}

// AddOptionIDs adds the "options" edge to the LanguageChoice entity by IDs.
func (luo *LanguageUpdateOne) AddOptionIDs(ids ...int) *LanguageUpdateOne {
	luo.mutation.AddOptionIDs(ids...)
	return luo
}

// AddOptions adds the "options" edges to the LanguageChoice entity.
func (luo *LanguageUpdateOne) AddOptions(l ...*LanguageChoice) *LanguageUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.AddOptionIDs(ids...)
}

// Mutation returns the LanguageMutation object of the builder.
func (luo *LanguageUpdateOne) Mutation() *LanguageMutation {
	return luo.mutation
}

// ClearRace clears all "race" edges to the Race entity.
func (luo *LanguageUpdateOne) ClearRace() *LanguageUpdateOne {
	luo.mutation.ClearRace()
	return luo
}

// RemoveRaceIDs removes the "race" edge to Race entities by IDs.
func (luo *LanguageUpdateOne) RemoveRaceIDs(ids ...int) *LanguageUpdateOne {
	luo.mutation.RemoveRaceIDs(ids...)
	return luo
}

// RemoveRace removes "race" edges to Race entities.
func (luo *LanguageUpdateOne) RemoveRace(r ...*Race) *LanguageUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return luo.RemoveRaceIDs(ids...)
}

// ClearOptions clears all "options" edges to the LanguageChoice entity.
func (luo *LanguageUpdateOne) ClearOptions() *LanguageUpdateOne {
	luo.mutation.ClearOptions()
	return luo
}

// RemoveOptionIDs removes the "options" edge to LanguageChoice entities by IDs.
func (luo *LanguageUpdateOne) RemoveOptionIDs(ids ...int) *LanguageUpdateOne {
	luo.mutation.RemoveOptionIDs(ids...)
	return luo
}

// RemoveOptions removes "options" edges to LanguageChoice entities.
func (luo *LanguageUpdateOne) RemoveOptions(l ...*LanguageChoice) *LanguageUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return luo.RemoveOptionIDs(ids...)
}

// Where appends a list predicates to the LanguageUpdate builder.
func (luo *LanguageUpdateOne) Where(ps ...predicate.Language) *LanguageUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LanguageUpdateOne) Select(field string, fields ...string) *LanguageUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Language entity.
func (luo *LanguageUpdateOne) Save(ctx context.Context) (*Language, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LanguageUpdateOne) SaveX(ctx context.Context) *Language {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LanguageUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LanguageUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luo *LanguageUpdateOne) check() error {
	if v, ok := luo.mutation.Indx(); ok {
		if err := language.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Language.indx": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Name(); ok {
		if err := language.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Language.name": %w`, err)}
		}
	}
	if v, ok := luo.mutation.LanguageType(); ok {
		if err := language.LanguageTypeValidator(v); err != nil {
			return &ValidationError{Name: "language_type", err: fmt.Errorf(`ent: validator failed for field "Language.language_type": %w`, err)}
		}
	}
	if v, ok := luo.mutation.Script(); ok {
		if err := language.ScriptValidator(v); err != nil {
			return &ValidationError{Name: "script", err: fmt.Errorf(`ent: validator failed for field "Language.script": %w`, err)}
		}
	}
	return nil
}

func (luo *LanguageUpdateOne) sqlSave(ctx context.Context) (_node *Language, err error) {
	if err := luo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Language.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, language.FieldID)
		for _, f := range fields {
			if !language.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != language.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.Indx(); ok {
		_spec.SetField(language.FieldIndx, field.TypeString, value)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(language.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Desc(); ok {
		_spec.SetField(language.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := luo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, language.FieldDesc, value)
		})
	}
	if luo.mutation.DescCleared() {
		_spec.ClearField(language.FieldDesc, field.TypeJSON)
	}
	if value, ok := luo.mutation.LanguageType(); ok {
		_spec.SetField(language.FieldLanguageType, field.TypeEnum, value)
	}
	if value, ok := luo.mutation.Script(); ok {
		_spec.SetField(language.FieldScript, field.TypeEnum, value)
	}
	if luo.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.RaceTable,
			Columns: language.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedRaceIDs(); len(nodes) > 0 && !luo.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.RaceTable,
			Columns: language.RacePrimaryKey,
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
	if nodes := luo.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.RaceTable,
			Columns: language.RacePrimaryKey,
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
	if luo.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.OptionsTable,
			Columns: language.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedOptionsIDs(); len(nodes) > 0 && !luo.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.OptionsTable,
			Columns: language.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.OptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.OptionsTable,
			Columns: language.OptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(languagechoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Language{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
