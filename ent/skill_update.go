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
func (su *SkillUpdate) SetDesc(s string) *SkillUpdate {
	su.mutation.SetDesc(s)
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

// Mutation returns the SkillMutation object of the builder.
func (su *SkillUpdate) Mutation() *SkillMutation {
	return su.mutation
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (su *SkillUpdate) ClearAbilityScore() *SkillUpdate {
	su.mutation.ClearAbilityScore()
	return su
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

func (su *SkillUpdate) sqlSave(ctx context.Context) (n int, err error) {
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
		_spec.SetField(skill.FieldDesc, field.TypeString, value)
	}
	if su.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
			Inverse: true,
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
func (suo *SkillUpdateOne) SetDesc(s string) *SkillUpdateOne {
	suo.mutation.SetDesc(s)
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

// Mutation returns the SkillMutation object of the builder.
func (suo *SkillUpdateOne) Mutation() *SkillMutation {
	return suo.mutation
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (suo *SkillUpdateOne) ClearAbilityScore() *SkillUpdateOne {
	suo.mutation.ClearAbilityScore()
	return suo
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

func (suo *SkillUpdateOne) sqlSave(ctx context.Context) (_node *Skill, err error) {
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
		_spec.SetField(skill.FieldDesc, field.TypeString, value)
	}
	if suo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
			Inverse: true,
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
