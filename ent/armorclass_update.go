// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armorclass"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ArmorClassUpdate is the builder for updating ArmorClass entities.
type ArmorClassUpdate struct {
	config
	hooks    []Hook
	mutation *ArmorClassMutation
}

// Where appends a list predicates to the ArmorClassUpdate builder.
func (acu *ArmorClassUpdate) Where(ps ...predicate.ArmorClass) *ArmorClassUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetBase sets the "base" field.
func (acu *ArmorClassUpdate) SetBase(i int) *ArmorClassUpdate {
	acu.mutation.ResetBase()
	acu.mutation.SetBase(i)
	return acu
}

// AddBase adds i to the "base" field.
func (acu *ArmorClassUpdate) AddBase(i int) *ArmorClassUpdate {
	acu.mutation.AddBase(i)
	return acu
}

// SetDexBonus sets the "dex_bonus" field.
func (acu *ArmorClassUpdate) SetDexBonus(b bool) *ArmorClassUpdate {
	acu.mutation.SetDexBonus(b)
	return acu
}

// SetMaxBonus sets the "max_bonus" field.
func (acu *ArmorClassUpdate) SetMaxBonus(i int) *ArmorClassUpdate {
	acu.mutation.ResetMaxBonus()
	acu.mutation.SetMaxBonus(i)
	return acu
}

// SetNillableMaxBonus sets the "max_bonus" field if the given value is not nil.
func (acu *ArmorClassUpdate) SetNillableMaxBonus(i *int) *ArmorClassUpdate {
	if i != nil {
		acu.SetMaxBonus(*i)
	}
	return acu
}

// AddMaxBonus adds i to the "max_bonus" field.
func (acu *ArmorClassUpdate) AddMaxBonus(i int) *ArmorClassUpdate {
	acu.mutation.AddMaxBonus(i)
	return acu
}

// ClearMaxBonus clears the value of the "max_bonus" field.
func (acu *ArmorClassUpdate) ClearMaxBonus() *ArmorClassUpdate {
	acu.mutation.ClearMaxBonus()
	return acu
}

// Mutation returns the ArmorClassMutation object of the builder.
func (acu *ArmorClassUpdate) Mutation() *ArmorClassMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *ArmorClassUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ArmorClassMutation](ctx, acu.sqlSave, acu.mutation, acu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acu *ArmorClassUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *ArmorClassUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *ArmorClassUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acu *ArmorClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(armorclass.Table, armorclass.Columns, sqlgraph.NewFieldSpec(armorclass.FieldID, field.TypeInt))
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.Base(); ok {
		_spec.SetField(armorclass.FieldBase, field.TypeInt, value)
	}
	if value, ok := acu.mutation.AddedBase(); ok {
		_spec.AddField(armorclass.FieldBase, field.TypeInt, value)
	}
	if value, ok := acu.mutation.DexBonus(); ok {
		_spec.SetField(armorclass.FieldDexBonus, field.TypeBool, value)
	}
	if value, ok := acu.mutation.MaxBonus(); ok {
		_spec.SetField(armorclass.FieldMaxBonus, field.TypeInt, value)
	}
	if value, ok := acu.mutation.AddedMaxBonus(); ok {
		_spec.AddField(armorclass.FieldMaxBonus, field.TypeInt, value)
	}
	if acu.mutation.MaxBonusCleared() {
		_spec.ClearField(armorclass.FieldMaxBonus, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{armorclass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	acu.mutation.done = true
	return n, nil
}

// ArmorClassUpdateOne is the builder for updating a single ArmorClass entity.
type ArmorClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArmorClassMutation
}

// SetBase sets the "base" field.
func (acuo *ArmorClassUpdateOne) SetBase(i int) *ArmorClassUpdateOne {
	acuo.mutation.ResetBase()
	acuo.mutation.SetBase(i)
	return acuo
}

// AddBase adds i to the "base" field.
func (acuo *ArmorClassUpdateOne) AddBase(i int) *ArmorClassUpdateOne {
	acuo.mutation.AddBase(i)
	return acuo
}

// SetDexBonus sets the "dex_bonus" field.
func (acuo *ArmorClassUpdateOne) SetDexBonus(b bool) *ArmorClassUpdateOne {
	acuo.mutation.SetDexBonus(b)
	return acuo
}

// SetMaxBonus sets the "max_bonus" field.
func (acuo *ArmorClassUpdateOne) SetMaxBonus(i int) *ArmorClassUpdateOne {
	acuo.mutation.ResetMaxBonus()
	acuo.mutation.SetMaxBonus(i)
	return acuo
}

// SetNillableMaxBonus sets the "max_bonus" field if the given value is not nil.
func (acuo *ArmorClassUpdateOne) SetNillableMaxBonus(i *int) *ArmorClassUpdateOne {
	if i != nil {
		acuo.SetMaxBonus(*i)
	}
	return acuo
}

// AddMaxBonus adds i to the "max_bonus" field.
func (acuo *ArmorClassUpdateOne) AddMaxBonus(i int) *ArmorClassUpdateOne {
	acuo.mutation.AddMaxBonus(i)
	return acuo
}

// ClearMaxBonus clears the value of the "max_bonus" field.
func (acuo *ArmorClassUpdateOne) ClearMaxBonus() *ArmorClassUpdateOne {
	acuo.mutation.ClearMaxBonus()
	return acuo
}

// Mutation returns the ArmorClassMutation object of the builder.
func (acuo *ArmorClassUpdateOne) Mutation() *ArmorClassMutation {
	return acuo.mutation
}

// Where appends a list predicates to the ArmorClassUpdate builder.
func (acuo *ArmorClassUpdateOne) Where(ps ...predicate.ArmorClass) *ArmorClassUpdateOne {
	acuo.mutation.Where(ps...)
	return acuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *ArmorClassUpdateOne) Select(field string, fields ...string) *ArmorClassUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated ArmorClass entity.
func (acuo *ArmorClassUpdateOne) Save(ctx context.Context) (*ArmorClass, error) {
	return withHooks[*ArmorClass, ArmorClassMutation](ctx, acuo.sqlSave, acuo.mutation, acuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *ArmorClassUpdateOne) SaveX(ctx context.Context) *ArmorClass {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *ArmorClassUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *ArmorClassUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (acuo *ArmorClassUpdateOne) sqlSave(ctx context.Context) (_node *ArmorClass, err error) {
	_spec := sqlgraph.NewUpdateSpec(armorclass.Table, armorclass.Columns, sqlgraph.NewFieldSpec(armorclass.FieldID, field.TypeInt))
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArmorClass.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, armorclass.FieldID)
		for _, f := range fields {
			if !armorclass.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != armorclass.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.Base(); ok {
		_spec.SetField(armorclass.FieldBase, field.TypeInt, value)
	}
	if value, ok := acuo.mutation.AddedBase(); ok {
		_spec.AddField(armorclass.FieldBase, field.TypeInt, value)
	}
	if value, ok := acuo.mutation.DexBonus(); ok {
		_spec.SetField(armorclass.FieldDexBonus, field.TypeBool, value)
	}
	if value, ok := acuo.mutation.MaxBonus(); ok {
		_spec.SetField(armorclass.FieldMaxBonus, field.TypeInt, value)
	}
	if value, ok := acuo.mutation.AddedMaxBonus(); ok {
		_spec.AddField(armorclass.FieldMaxBonus, field.TypeInt, value)
	}
	if acuo.mutation.MaxBonusCleared() {
		_spec.ClearField(armorclass.FieldMaxBonus, field.TypeInt)
	}
	_node = &ArmorClass{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{armorclass.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	acuo.mutation.done = true
	return _node, nil
}