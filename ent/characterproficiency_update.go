// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterproficiency"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
)

// CharacterProficiencyUpdate is the builder for updating CharacterProficiency entities.
type CharacterProficiencyUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterProficiencyMutation
}

// Where appends a list predicates to the CharacterProficiencyUpdate builder.
func (cpu *CharacterProficiencyUpdate) Where(ps ...predicate.CharacterProficiency) *CharacterProficiencyUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetCharacterID sets the "character_id" field.
func (cpu *CharacterProficiencyUpdate) SetCharacterID(i int) *CharacterProficiencyUpdate {
	cpu.mutation.SetCharacterID(i)
	return cpu
}

// SetNillableCharacterID sets the "character_id" field if the given value is not nil.
func (cpu *CharacterProficiencyUpdate) SetNillableCharacterID(i *int) *CharacterProficiencyUpdate {
	if i != nil {
		cpu.SetCharacterID(*i)
	}
	return cpu
}

// SetProficiencyID sets the "proficiency_id" field.
func (cpu *CharacterProficiencyUpdate) SetProficiencyID(i int) *CharacterProficiencyUpdate {
	cpu.mutation.SetProficiencyID(i)
	return cpu
}

// SetNillableProficiencyID sets the "proficiency_id" field if the given value is not nil.
func (cpu *CharacterProficiencyUpdate) SetNillableProficiencyID(i *int) *CharacterProficiencyUpdate {
	if i != nil {
		cpu.SetProficiencyID(*i)
	}
	return cpu
}

// SetCharacter sets the "character" edge to the Character entity.
func (cpu *CharacterProficiencyUpdate) SetCharacter(c *Character) *CharacterProficiencyUpdate {
	return cpu.SetCharacterID(c.ID)
}

// SetProficiency sets the "proficiency" edge to the Proficiency entity.
func (cpu *CharacterProficiencyUpdate) SetProficiency(p *Proficiency) *CharacterProficiencyUpdate {
	return cpu.SetProficiencyID(p.ID)
}

// Mutation returns the CharacterProficiencyMutation object of the builder.
func (cpu *CharacterProficiencyUpdate) Mutation() *CharacterProficiencyMutation {
	return cpu.mutation
}

// ClearCharacter clears the "character" edge to the Character entity.
func (cpu *CharacterProficiencyUpdate) ClearCharacter() *CharacterProficiencyUpdate {
	cpu.mutation.ClearCharacter()
	return cpu
}

// ClearProficiency clears the "proficiency" edge to the Proficiency entity.
func (cpu *CharacterProficiencyUpdate) ClearProficiency() *CharacterProficiencyUpdate {
	cpu.mutation.ClearProficiency()
	return cpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *CharacterProficiencyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cpu.sqlSave, cpu.mutation, cpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *CharacterProficiencyUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *CharacterProficiencyUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *CharacterProficiencyUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpu *CharacterProficiencyUpdate) check() error {
	if cpu.mutation.CharacterCleared() && len(cpu.mutation.CharacterIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterProficiency.character"`)
	}
	if cpu.mutation.ProficiencyCleared() && len(cpu.mutation.ProficiencyIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterProficiency.proficiency"`)
	}
	return nil
}

func (cpu *CharacterProficiencyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(characterproficiency.Table, characterproficiency.Columns, sqlgraph.NewFieldSpec(characterproficiency.FieldID, field.TypeInt))
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cpu.mutation.CharacterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.CharacterTable,
			Columns: []string{characterproficiency.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.CharacterTable,
			Columns: []string{characterproficiency.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpu.mutation.ProficiencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.ProficiencyTable,
			Columns: []string{characterproficiency.ProficiencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.ProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.ProficiencyTable,
			Columns: []string{characterproficiency.ProficiencyColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{characterproficiency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cpu.mutation.done = true
	return n, nil
}

// CharacterProficiencyUpdateOne is the builder for updating a single CharacterProficiency entity.
type CharacterProficiencyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterProficiencyMutation
}

// SetCharacterID sets the "character_id" field.
func (cpuo *CharacterProficiencyUpdateOne) SetCharacterID(i int) *CharacterProficiencyUpdateOne {
	cpuo.mutation.SetCharacterID(i)
	return cpuo
}

// SetNillableCharacterID sets the "character_id" field if the given value is not nil.
func (cpuo *CharacterProficiencyUpdateOne) SetNillableCharacterID(i *int) *CharacterProficiencyUpdateOne {
	if i != nil {
		cpuo.SetCharacterID(*i)
	}
	return cpuo
}

// SetProficiencyID sets the "proficiency_id" field.
func (cpuo *CharacterProficiencyUpdateOne) SetProficiencyID(i int) *CharacterProficiencyUpdateOne {
	cpuo.mutation.SetProficiencyID(i)
	return cpuo
}

// SetNillableProficiencyID sets the "proficiency_id" field if the given value is not nil.
func (cpuo *CharacterProficiencyUpdateOne) SetNillableProficiencyID(i *int) *CharacterProficiencyUpdateOne {
	if i != nil {
		cpuo.SetProficiencyID(*i)
	}
	return cpuo
}

// SetCharacter sets the "character" edge to the Character entity.
func (cpuo *CharacterProficiencyUpdateOne) SetCharacter(c *Character) *CharacterProficiencyUpdateOne {
	return cpuo.SetCharacterID(c.ID)
}

// SetProficiency sets the "proficiency" edge to the Proficiency entity.
func (cpuo *CharacterProficiencyUpdateOne) SetProficiency(p *Proficiency) *CharacterProficiencyUpdateOne {
	return cpuo.SetProficiencyID(p.ID)
}

// Mutation returns the CharacterProficiencyMutation object of the builder.
func (cpuo *CharacterProficiencyUpdateOne) Mutation() *CharacterProficiencyMutation {
	return cpuo.mutation
}

// ClearCharacter clears the "character" edge to the Character entity.
func (cpuo *CharacterProficiencyUpdateOne) ClearCharacter() *CharacterProficiencyUpdateOne {
	cpuo.mutation.ClearCharacter()
	return cpuo
}

// ClearProficiency clears the "proficiency" edge to the Proficiency entity.
func (cpuo *CharacterProficiencyUpdateOne) ClearProficiency() *CharacterProficiencyUpdateOne {
	cpuo.mutation.ClearProficiency()
	return cpuo
}

// Where appends a list predicates to the CharacterProficiencyUpdate builder.
func (cpuo *CharacterProficiencyUpdateOne) Where(ps ...predicate.CharacterProficiency) *CharacterProficiencyUpdateOne {
	cpuo.mutation.Where(ps...)
	return cpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *CharacterProficiencyUpdateOne) Select(field string, fields ...string) *CharacterProficiencyUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated CharacterProficiency entity.
func (cpuo *CharacterProficiencyUpdateOne) Save(ctx context.Context) (*CharacterProficiency, error) {
	return withHooks(ctx, cpuo.sqlSave, cpuo.mutation, cpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *CharacterProficiencyUpdateOne) SaveX(ctx context.Context) *CharacterProficiency {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *CharacterProficiencyUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *CharacterProficiencyUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *CharacterProficiencyUpdateOne) check() error {
	if cpuo.mutation.CharacterCleared() && len(cpuo.mutation.CharacterIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterProficiency.character"`)
	}
	if cpuo.mutation.ProficiencyCleared() && len(cpuo.mutation.ProficiencyIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterProficiency.proficiency"`)
	}
	return nil
}

func (cpuo *CharacterProficiencyUpdateOne) sqlSave(ctx context.Context) (_node *CharacterProficiency, err error) {
	if err := cpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(characterproficiency.Table, characterproficiency.Columns, sqlgraph.NewFieldSpec(characterproficiency.FieldID, field.TypeInt))
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CharacterProficiency.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, characterproficiency.FieldID)
		for _, f := range fields {
			if !characterproficiency.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != characterproficiency.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cpuo.mutation.CharacterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.CharacterTable,
			Columns: []string{characterproficiency.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.CharacterTable,
			Columns: []string{characterproficiency.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cpuo.mutation.ProficiencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.ProficiencyTable,
			Columns: []string{characterproficiency.ProficiencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.ProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterproficiency.ProficiencyTable,
			Columns: []string{characterproficiency.ProficiencyColumn},
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
	_node = &CharacterProficiency{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{characterproficiency.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cpuo.mutation.done = true
	return _node, nil
}
