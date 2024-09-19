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
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/skill"
)

// AbilityScoreUpdate is the builder for updating AbilityScore entities.
type AbilityScoreUpdate struct {
	config
	hooks    []Hook
	mutation *AbilityScoreMutation
}

// Where appends a list predicates to the AbilityScoreUpdate builder.
func (asu *AbilityScoreUpdate) Where(ps ...predicate.AbilityScore) *AbilityScoreUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetIndx sets the "indx" field.
func (asu *AbilityScoreUpdate) SetIndx(s string) *AbilityScoreUpdate {
	asu.mutation.SetIndx(s)
	return asu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (asu *AbilityScoreUpdate) SetNillableIndx(s *string) *AbilityScoreUpdate {
	if s != nil {
		asu.SetIndx(*s)
	}
	return asu
}

// SetName sets the "name" field.
func (asu *AbilityScoreUpdate) SetName(s string) *AbilityScoreUpdate {
	asu.mutation.SetName(s)
	return asu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (asu *AbilityScoreUpdate) SetNillableName(s *string) *AbilityScoreUpdate {
	if s != nil {
		asu.SetName(*s)
	}
	return asu
}

// SetDesc sets the "desc" field.
func (asu *AbilityScoreUpdate) SetDesc(s []string) *AbilityScoreUpdate {
	asu.mutation.SetDesc(s)
	return asu
}

// AppendDesc appends s to the "desc" field.
func (asu *AbilityScoreUpdate) AppendDesc(s []string) *AbilityScoreUpdate {
	asu.mutation.AppendDesc(s)
	return asu
}

// ClearDesc clears the value of the "desc" field.
func (asu *AbilityScoreUpdate) ClearDesc() *AbilityScoreUpdate {
	asu.mutation.ClearDesc()
	return asu
}

// SetFullName sets the "full_name" field.
func (asu *AbilityScoreUpdate) SetFullName(s string) *AbilityScoreUpdate {
	asu.mutation.SetFullName(s)
	return asu
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (asu *AbilityScoreUpdate) SetNillableFullName(s *string) *AbilityScoreUpdate {
	if s != nil {
		asu.SetFullName(*s)
	}
	return asu
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (asu *AbilityScoreUpdate) AddSkillIDs(ids ...int) *AbilityScoreUpdate {
	asu.mutation.AddSkillIDs(ids...)
	return asu
}

// AddSkills adds the "skills" edges to the Skill entity.
func (asu *AbilityScoreUpdate) AddSkills(s ...*Skill) *AbilityScoreUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return asu.AddSkillIDs(ids...)
}

// AddAbilityBonuseIDs adds the "ability_bonuses" edge to the AbilityBonus entity by IDs.
func (asu *AbilityScoreUpdate) AddAbilityBonuseIDs(ids ...int) *AbilityScoreUpdate {
	asu.mutation.AddAbilityBonuseIDs(ids...)
	return asu
}

// AddAbilityBonuses adds the "ability_bonuses" edges to the AbilityBonus entity.
func (asu *AbilityScoreUpdate) AddAbilityBonuses(a ...*AbilityBonus) *AbilityScoreUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asu.AddAbilityBonuseIDs(ids...)
}

// Mutation returns the AbilityScoreMutation object of the builder.
func (asu *AbilityScoreUpdate) Mutation() *AbilityScoreMutation {
	return asu.mutation
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (asu *AbilityScoreUpdate) ClearSkills() *AbilityScoreUpdate {
	asu.mutation.ClearSkills()
	return asu
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (asu *AbilityScoreUpdate) RemoveSkillIDs(ids ...int) *AbilityScoreUpdate {
	asu.mutation.RemoveSkillIDs(ids...)
	return asu
}

// RemoveSkills removes "skills" edges to Skill entities.
func (asu *AbilityScoreUpdate) RemoveSkills(s ...*Skill) *AbilityScoreUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return asu.RemoveSkillIDs(ids...)
}

// ClearAbilityBonuses clears all "ability_bonuses" edges to the AbilityBonus entity.
func (asu *AbilityScoreUpdate) ClearAbilityBonuses() *AbilityScoreUpdate {
	asu.mutation.ClearAbilityBonuses()
	return asu
}

// RemoveAbilityBonuseIDs removes the "ability_bonuses" edge to AbilityBonus entities by IDs.
func (asu *AbilityScoreUpdate) RemoveAbilityBonuseIDs(ids ...int) *AbilityScoreUpdate {
	asu.mutation.RemoveAbilityBonuseIDs(ids...)
	return asu
}

// RemoveAbilityBonuses removes "ability_bonuses" edges to AbilityBonus entities.
func (asu *AbilityScoreUpdate) RemoveAbilityBonuses(a ...*AbilityBonus) *AbilityScoreUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asu.RemoveAbilityBonuseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *AbilityScoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, asu.sqlSave, asu.mutation, asu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asu *AbilityScoreUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *AbilityScoreUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *AbilityScoreUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asu *AbilityScoreUpdate) check() error {
	if v, ok := asu.mutation.Indx(); ok {
		if err := abilityscore.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.indx": %w`, err)}
		}
	}
	if v, ok := asu.mutation.Name(); ok {
		if err := abilityscore.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.name": %w`, err)}
		}
	}
	return nil
}

func (asu *AbilityScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := asu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(abilityscore.Table, abilityscore.Columns, sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt))
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.Indx(); ok {
		_spec.SetField(abilityscore.FieldIndx, field.TypeString, value)
	}
	if value, ok := asu.mutation.Name(); ok {
		_spec.SetField(abilityscore.FieldName, field.TypeString, value)
	}
	if value, ok := asu.mutation.Desc(); ok {
		_spec.SetField(abilityscore.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := asu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, abilityscore.FieldDesc, value)
		})
	}
	if asu.mutation.DescCleared() {
		_spec.ClearField(abilityscore.FieldDesc, field.TypeJSON)
	}
	if value, ok := asu.mutation.FullName(); ok {
		_spec.SetField(abilityscore.FieldFullName, field.TypeString, value)
	}
	if asu.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.SkillsTable,
			Columns: []string{abilityscore.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !asu.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.SkillsTable,
			Columns: []string{abilityscore.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.SkillsTable,
			Columns: []string{abilityscore.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if asu.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.AbilityBonusesTable,
			Columns: []string{abilityscore.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.RemovedAbilityBonusesIDs(); len(nodes) > 0 && !asu.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.AbilityBonusesTable,
			Columns: []string{abilityscore.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.AbilityBonusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.AbilityBonusesTable,
			Columns: []string{abilityscore.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{abilityscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	asu.mutation.done = true
	return n, nil
}

// AbilityScoreUpdateOne is the builder for updating a single AbilityScore entity.
type AbilityScoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AbilityScoreMutation
}

// SetIndx sets the "indx" field.
func (asuo *AbilityScoreUpdateOne) SetIndx(s string) *AbilityScoreUpdateOne {
	asuo.mutation.SetIndx(s)
	return asuo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (asuo *AbilityScoreUpdateOne) SetNillableIndx(s *string) *AbilityScoreUpdateOne {
	if s != nil {
		asuo.SetIndx(*s)
	}
	return asuo
}

// SetName sets the "name" field.
func (asuo *AbilityScoreUpdateOne) SetName(s string) *AbilityScoreUpdateOne {
	asuo.mutation.SetName(s)
	return asuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (asuo *AbilityScoreUpdateOne) SetNillableName(s *string) *AbilityScoreUpdateOne {
	if s != nil {
		asuo.SetName(*s)
	}
	return asuo
}

// SetDesc sets the "desc" field.
func (asuo *AbilityScoreUpdateOne) SetDesc(s []string) *AbilityScoreUpdateOne {
	asuo.mutation.SetDesc(s)
	return asuo
}

// AppendDesc appends s to the "desc" field.
func (asuo *AbilityScoreUpdateOne) AppendDesc(s []string) *AbilityScoreUpdateOne {
	asuo.mutation.AppendDesc(s)
	return asuo
}

// ClearDesc clears the value of the "desc" field.
func (asuo *AbilityScoreUpdateOne) ClearDesc() *AbilityScoreUpdateOne {
	asuo.mutation.ClearDesc()
	return asuo
}

// SetFullName sets the "full_name" field.
func (asuo *AbilityScoreUpdateOne) SetFullName(s string) *AbilityScoreUpdateOne {
	asuo.mutation.SetFullName(s)
	return asuo
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (asuo *AbilityScoreUpdateOne) SetNillableFullName(s *string) *AbilityScoreUpdateOne {
	if s != nil {
		asuo.SetFullName(*s)
	}
	return asuo
}

// AddSkillIDs adds the "skills" edge to the Skill entity by IDs.
func (asuo *AbilityScoreUpdateOne) AddSkillIDs(ids ...int) *AbilityScoreUpdateOne {
	asuo.mutation.AddSkillIDs(ids...)
	return asuo
}

// AddSkills adds the "skills" edges to the Skill entity.
func (asuo *AbilityScoreUpdateOne) AddSkills(s ...*Skill) *AbilityScoreUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return asuo.AddSkillIDs(ids...)
}

// AddAbilityBonuseIDs adds the "ability_bonuses" edge to the AbilityBonus entity by IDs.
func (asuo *AbilityScoreUpdateOne) AddAbilityBonuseIDs(ids ...int) *AbilityScoreUpdateOne {
	asuo.mutation.AddAbilityBonuseIDs(ids...)
	return asuo
}

// AddAbilityBonuses adds the "ability_bonuses" edges to the AbilityBonus entity.
func (asuo *AbilityScoreUpdateOne) AddAbilityBonuses(a ...*AbilityBonus) *AbilityScoreUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asuo.AddAbilityBonuseIDs(ids...)
}

// Mutation returns the AbilityScoreMutation object of the builder.
func (asuo *AbilityScoreUpdateOne) Mutation() *AbilityScoreMutation {
	return asuo.mutation
}

// ClearSkills clears all "skills" edges to the Skill entity.
func (asuo *AbilityScoreUpdateOne) ClearSkills() *AbilityScoreUpdateOne {
	asuo.mutation.ClearSkills()
	return asuo
}

// RemoveSkillIDs removes the "skills" edge to Skill entities by IDs.
func (asuo *AbilityScoreUpdateOne) RemoveSkillIDs(ids ...int) *AbilityScoreUpdateOne {
	asuo.mutation.RemoveSkillIDs(ids...)
	return asuo
}

// RemoveSkills removes "skills" edges to Skill entities.
func (asuo *AbilityScoreUpdateOne) RemoveSkills(s ...*Skill) *AbilityScoreUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return asuo.RemoveSkillIDs(ids...)
}

// ClearAbilityBonuses clears all "ability_bonuses" edges to the AbilityBonus entity.
func (asuo *AbilityScoreUpdateOne) ClearAbilityBonuses() *AbilityScoreUpdateOne {
	asuo.mutation.ClearAbilityBonuses()
	return asuo
}

// RemoveAbilityBonuseIDs removes the "ability_bonuses" edge to AbilityBonus entities by IDs.
func (asuo *AbilityScoreUpdateOne) RemoveAbilityBonuseIDs(ids ...int) *AbilityScoreUpdateOne {
	asuo.mutation.RemoveAbilityBonuseIDs(ids...)
	return asuo
}

// RemoveAbilityBonuses removes "ability_bonuses" edges to AbilityBonus entities.
func (asuo *AbilityScoreUpdateOne) RemoveAbilityBonuses(a ...*AbilityBonus) *AbilityScoreUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asuo.RemoveAbilityBonuseIDs(ids...)
}

// Where appends a list predicates to the AbilityScoreUpdate builder.
func (asuo *AbilityScoreUpdateOne) Where(ps ...predicate.AbilityScore) *AbilityScoreUpdateOne {
	asuo.mutation.Where(ps...)
	return asuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *AbilityScoreUpdateOne) Select(field string, fields ...string) *AbilityScoreUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated AbilityScore entity.
func (asuo *AbilityScoreUpdateOne) Save(ctx context.Context) (*AbilityScore, error) {
	return withHooks(ctx, asuo.sqlSave, asuo.mutation, asuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *AbilityScoreUpdateOne) SaveX(ctx context.Context) *AbilityScore {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *AbilityScoreUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *AbilityScoreUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asuo *AbilityScoreUpdateOne) check() error {
	if v, ok := asuo.mutation.Indx(); ok {
		if err := abilityscore.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.indx": %w`, err)}
		}
	}
	if v, ok := asuo.mutation.Name(); ok {
		if err := abilityscore.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "AbilityScore.name": %w`, err)}
		}
	}
	return nil
}

func (asuo *AbilityScoreUpdateOne) sqlSave(ctx context.Context) (_node *AbilityScore, err error) {
	if err := asuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(abilityscore.Table, abilityscore.Columns, sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt))
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AbilityScore.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, abilityscore.FieldID)
		for _, f := range fields {
			if !abilityscore.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != abilityscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asuo.mutation.Indx(); ok {
		_spec.SetField(abilityscore.FieldIndx, field.TypeString, value)
	}
	if value, ok := asuo.mutation.Name(); ok {
		_spec.SetField(abilityscore.FieldName, field.TypeString, value)
	}
	if value, ok := asuo.mutation.Desc(); ok {
		_spec.SetField(abilityscore.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := asuo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, abilityscore.FieldDesc, value)
		})
	}
	if asuo.mutation.DescCleared() {
		_spec.ClearField(abilityscore.FieldDesc, field.TypeJSON)
	}
	if value, ok := asuo.mutation.FullName(); ok {
		_spec.SetField(abilityscore.FieldFullName, field.TypeString, value)
	}
	if asuo.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.SkillsTable,
			Columns: []string{abilityscore.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.RemovedSkillsIDs(); len(nodes) > 0 && !asuo.mutation.SkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.SkillsTable,
			Columns: []string{abilityscore.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.SkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.SkillsTable,
			Columns: []string{abilityscore.SkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(skill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if asuo.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.AbilityBonusesTable,
			Columns: []string{abilityscore.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.RemovedAbilityBonusesIDs(); len(nodes) > 0 && !asuo.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.AbilityBonusesTable,
			Columns: []string{abilityscore.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.AbilityBonusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   abilityscore.AbilityBonusesTable,
			Columns: []string{abilityscore.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AbilityScore{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{abilityscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	asuo.mutation.done = true
	return _node, nil
}
