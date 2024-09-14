// Code generated by ent, DO NOT EDIT.

package ent

import (
	"builder/ent/predicate"
	"builder/ent/skill"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
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

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (su *SkillUpdate) SetNillableIndx(s *string) *SkillUpdate {
	if s != nil {
		su.SetIndx(*s)
	}
	return su
}

// SetName sets the "name" field.
func (su *SkillUpdate) SetName(s string) *SkillUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SkillUpdate) SetNillableName(s *string) *SkillUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// Mutation returns the SkillMutation object of the builder.
func (su *SkillUpdate) Mutation() *SkillMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SkillUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
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

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (suo *SkillUpdateOne) SetNillableIndx(s *string) *SkillUpdateOne {
	if s != nil {
		suo.SetIndx(*s)
	}
	return suo
}

// SetName sets the "name" field.
func (suo *SkillUpdateOne) SetName(s string) *SkillUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SkillUpdateOne) SetNillableName(s *string) *SkillUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// Mutation returns the SkillMutation object of the builder.
func (suo *SkillUpdateOne) Mutation() *SkillMutation {
	return suo.mutation
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
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
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