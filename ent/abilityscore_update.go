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
	"github.com/ecshreve/dndgen/ent/proficiency"
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

// SetName sets the "name" field.
func (asu *AbilityScoreUpdate) SetName(s string) *AbilityScoreUpdate {
	asu.mutation.SetName(s)
	return asu
}

// SetDesc sets the "desc" field.
func (asu *AbilityScoreUpdate) SetDesc(s string) *AbilityScoreUpdate {
	asu.mutation.SetDesc(s)
	return asu
}

// SetFullName sets the "full_name" field.
func (asu *AbilityScoreUpdate) SetFullName(s string) *AbilityScoreUpdate {
	asu.mutation.SetFullName(s)
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

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (asu *AbilityScoreUpdate) AddProficiencyIDs(ids ...int) *AbilityScoreUpdate {
	asu.mutation.AddProficiencyIDs(ids...)
	return asu
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (asu *AbilityScoreUpdate) AddProficiencies(p ...*Proficiency) *AbilityScoreUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asu.AddProficiencyIDs(ids...)
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

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (asu *AbilityScoreUpdate) ClearProficiencies() *AbilityScoreUpdate {
	asu.mutation.ClearProficiencies()
	return asu
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (asu *AbilityScoreUpdate) RemoveProficiencyIDs(ids ...int) *AbilityScoreUpdate {
	asu.mutation.RemoveProficiencyIDs(ids...)
	return asu
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (asu *AbilityScoreUpdate) RemoveProficiencies(p ...*Proficiency) *AbilityScoreUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asu.RemoveProficiencyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *AbilityScoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AbilityScoreMutation](ctx, asu.sqlSave, asu.mutation, asu.hooks)
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

func (asu *AbilityScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
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
		_spec.SetField(abilityscore.FieldDesc, field.TypeString, value)
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
	if asu.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   abilityscore.ProficienciesTable,
			Columns: abilityscore.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asu.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !asu.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   abilityscore.ProficienciesTable,
			Columns: abilityscore.ProficienciesPrimaryKey,
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
	if nodes := asu.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   abilityscore.ProficienciesTable,
			Columns: abilityscore.ProficienciesPrimaryKey,
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

// SetName sets the "name" field.
func (asuo *AbilityScoreUpdateOne) SetName(s string) *AbilityScoreUpdateOne {
	asuo.mutation.SetName(s)
	return asuo
}

// SetDesc sets the "desc" field.
func (asuo *AbilityScoreUpdateOne) SetDesc(s string) *AbilityScoreUpdateOne {
	asuo.mutation.SetDesc(s)
	return asuo
}

// SetFullName sets the "full_name" field.
func (asuo *AbilityScoreUpdateOne) SetFullName(s string) *AbilityScoreUpdateOne {
	asuo.mutation.SetFullName(s)
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

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (asuo *AbilityScoreUpdateOne) AddProficiencyIDs(ids ...int) *AbilityScoreUpdateOne {
	asuo.mutation.AddProficiencyIDs(ids...)
	return asuo
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (asuo *AbilityScoreUpdateOne) AddProficiencies(p ...*Proficiency) *AbilityScoreUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asuo.AddProficiencyIDs(ids...)
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

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (asuo *AbilityScoreUpdateOne) ClearProficiencies() *AbilityScoreUpdateOne {
	asuo.mutation.ClearProficiencies()
	return asuo
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (asuo *AbilityScoreUpdateOne) RemoveProficiencyIDs(ids ...int) *AbilityScoreUpdateOne {
	asuo.mutation.RemoveProficiencyIDs(ids...)
	return asuo
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (asuo *AbilityScoreUpdateOne) RemoveProficiencies(p ...*Proficiency) *AbilityScoreUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return asuo.RemoveProficiencyIDs(ids...)
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
	return withHooks[*AbilityScore, AbilityScoreMutation](ctx, asuo.sqlSave, asuo.mutation, asuo.hooks)
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

func (asuo *AbilityScoreUpdateOne) sqlSave(ctx context.Context) (_node *AbilityScore, err error) {
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
		_spec.SetField(abilityscore.FieldDesc, field.TypeString, value)
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
	if asuo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   abilityscore.ProficienciesTable,
			Columns: abilityscore.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := asuo.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !asuo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   abilityscore.ProficienciesTable,
			Columns: abilityscore.ProficienciesPrimaryKey,
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
	if nodes := asuo.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   abilityscore.ProficienciesTable,
			Columns: abilityscore.ProficienciesPrimaryKey,
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
