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
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/skill"
)

// SkillUpdate is the builder for updating Skill entities.
type SkillUpdate struct {
	config
	hooks    []Hook
	mutation *SkillMutation
}

// Where appends a list predicates to the SkillUpdate builder.
func (su *SkillUpdate) Where(ps ...predicate.Skill) *SkillUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetIndx sets the "indx" field.
func (su *SkillUpdate) SetIndx(s string) *SkillUpdate {
	su.mutation.SetIndx(s)
	return su
}

// SetName sets the "name" field.
func (su *SkillUpdate) SetName(s string) *SkillUpdate {
	su.mutation.SetName(s)
	return su
}

// SetDesc sets the "desc" field.
func (su *SkillUpdate) SetDesc(s []string) *SkillUpdate {
	su.mutation.SetDesc(s)
	return su
}

// AppendDesc appends s to the "desc" field.
func (su *SkillUpdate) AppendDesc(s []string) *SkillUpdate {
	su.mutation.AppendDesc(s)
	return su
}

// SetAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID.
func (su *SkillUpdate) SetAbilityScoreID(id int) *SkillUpdate {
	su.mutation.SetAbilityScoreID(id)
	return su
}

// SetNillableAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID if the given value is not nil.
func (su *SkillUpdate) SetNillableAbilityScoreID(id *int) *SkillUpdate {
	if id != nil {
		su = su.SetAbilityScoreID(*id)
	}
	return su
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (su *SkillUpdate) SetAbilityScore(a *AbilityScore) *SkillUpdate {
	return su.SetAbilityScoreID(a.ID)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (su *SkillUpdate) AddProficiencyIDs(ids ...int) *SkillUpdate {
	su.mutation.AddProficiencyIDs(ids...)
	return su
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (su *SkillUpdate) AddProficiencies(p ...*Proficiency) *SkillUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddProficiencyIDs(ids...)
}

// Mutation returns the SkillMutation object of the builder.
func (su *SkillUpdate) Mutation() *SkillMutation {
	return su.mutation
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (su *SkillUpdate) ClearAbilityScore() *SkillUpdate {
	su.mutation.ClearAbilityScore()
	return su
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (su *SkillUpdate) ClearProficiencies() *SkillUpdate {
	su.mutation.ClearProficiencies()
	return su
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (su *SkillUpdate) RemoveProficiencyIDs(ids ...int) *SkillUpdate {
	su.mutation.RemoveProficiencyIDs(ids...)
	return su
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (su *SkillUpdate) RemoveProficiencies(p ...*Proficiency) *SkillUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveProficiencyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SkillUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SkillMutation](ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SkillUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SkillUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SkillUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SkillUpdate) check() error {
	if v, ok := su.mutation.Indx(); ok {
		if err := skill.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Skill.indx": %w`, err)}
		}
	}
	if v, ok := su.mutation.Name(); ok {
		if err := skill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Skill.name": %w`, err)}
		}
	}
	return nil
}

func (su *SkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(skill.Table, skill.Columns, sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Indx(); ok {
		_spec.SetField(skill.FieldIndx, field.TypeString, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(skill.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Desc(); ok {
		_spec.SetField(skill.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := su.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, skill.FieldDesc, value)
		})
	}
	if su.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   skill.AbilityScoreTable,
			Columns: []string{skill.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   skill.AbilityScoreTable,
			Columns: []string{skill.AbilityScoreColumn},
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
	if su.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   skill.ProficienciesTable,
			Columns: []string{skill.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !su.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   skill.ProficienciesTable,
			Columns: []string{skill.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   skill.ProficienciesTable,
			Columns: []string{skill.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SkillUpdateOne is the builder for updating a single Skill entity.
type SkillUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkillMutation
}

// SetIndx sets the "indx" field.
func (suo *SkillUpdateOne) SetIndx(s string) *SkillUpdateOne {
	suo.mutation.SetIndx(s)
	return suo
}

// SetName sets the "name" field.
func (suo *SkillUpdateOne) SetName(s string) *SkillUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetDesc sets the "desc" field.
func (suo *SkillUpdateOne) SetDesc(s []string) *SkillUpdateOne {
	suo.mutation.SetDesc(s)
	return suo
}

// AppendDesc appends s to the "desc" field.
func (suo *SkillUpdateOne) AppendDesc(s []string) *SkillUpdateOne {
	suo.mutation.AppendDesc(s)
	return suo
}

// SetAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID.
func (suo *SkillUpdateOne) SetAbilityScoreID(id int) *SkillUpdateOne {
	suo.mutation.SetAbilityScoreID(id)
	return suo
}

// SetNillableAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID if the given value is not nil.
func (suo *SkillUpdateOne) SetNillableAbilityScoreID(id *int) *SkillUpdateOne {
	if id != nil {
		suo = suo.SetAbilityScoreID(*id)
	}
	return suo
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (suo *SkillUpdateOne) SetAbilityScore(a *AbilityScore) *SkillUpdateOne {
	return suo.SetAbilityScoreID(a.ID)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (suo *SkillUpdateOne) AddProficiencyIDs(ids ...int) *SkillUpdateOne {
	suo.mutation.AddProficiencyIDs(ids...)
	return suo
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (suo *SkillUpdateOne) AddProficiencies(p ...*Proficiency) *SkillUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddProficiencyIDs(ids...)
}

// Mutation returns the SkillMutation object of the builder.
func (suo *SkillUpdateOne) Mutation() *SkillMutation {
	return suo.mutation
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (suo *SkillUpdateOne) ClearAbilityScore() *SkillUpdateOne {
	suo.mutation.ClearAbilityScore()
	return suo
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (suo *SkillUpdateOne) ClearProficiencies() *SkillUpdateOne {
	suo.mutation.ClearProficiencies()
	return suo
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (suo *SkillUpdateOne) RemoveProficiencyIDs(ids ...int) *SkillUpdateOne {
	suo.mutation.RemoveProficiencyIDs(ids...)
	return suo
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (suo *SkillUpdateOne) RemoveProficiencies(p ...*Proficiency) *SkillUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveProficiencyIDs(ids...)
}

// Where appends a list predicates to the SkillUpdate builder.
func (suo *SkillUpdateOne) Where(ps ...predicate.Skill) *SkillUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SkillUpdateOne) Select(field string, fields ...string) *SkillUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Skill entity.
func (suo *SkillUpdateOne) Save(ctx context.Context) (*Skill, error) {
	return withHooks[*Skill, SkillMutation](ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SkillUpdateOne) SaveX(ctx context.Context) *Skill {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SkillUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SkillUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SkillUpdateOne) check() error {
	if v, ok := suo.mutation.Indx(); ok {
		if err := skill.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Skill.indx": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Name(); ok {
		if err := skill.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Skill.name": %w`, err)}
		}
	}
	return nil
}

func (suo *SkillUpdateOne) sqlSave(ctx context.Context) (_node *Skill, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(skill.Table, skill.Columns, sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Skill.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skill.FieldID)
		for _, f := range fields {
			if !skill.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Indx(); ok {
		_spec.SetField(skill.FieldIndx, field.TypeString, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(skill.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Desc(); ok {
		_spec.SetField(skill.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := suo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, skill.FieldDesc, value)
		})
	}
	if suo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   skill.AbilityScoreTable,
			Columns: []string{skill.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   skill.AbilityScoreTable,
			Columns: []string{skill.AbilityScoreColumn},
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
	if suo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   skill.ProficienciesTable,
			Columns: []string{skill.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !suo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   skill.ProficienciesTable,
			Columns: []string{skill.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   skill.ProficienciesTable,
			Columns: []string{skill.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Skill{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skill.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
