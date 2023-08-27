// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// WeaponDamageUpdate is the builder for updating WeaponDamage entities.
type WeaponDamageUpdate struct {
	config
	hooks    []Hook
	mutation *WeaponDamageMutation
}

// Where appends a list predicates to the WeaponDamageUpdate builder.
func (wdu *WeaponDamageUpdate) Where(ps ...predicate.WeaponDamage) *WeaponDamageUpdate {
	wdu.mutation.Where(ps...)
	return wdu
}

// SetWeaponID sets the "weapon_id" field.
func (wdu *WeaponDamageUpdate) SetWeaponID(i int) *WeaponDamageUpdate {
	wdu.mutation.SetWeaponID(i)
	return wdu
}

// SetDamageTypeID sets the "damage_type_id" field.
func (wdu *WeaponDamageUpdate) SetDamageTypeID(i int) *WeaponDamageUpdate {
	wdu.mutation.SetDamageTypeID(i)
	return wdu
}

// SetDice sets the "dice" field.
func (wdu *WeaponDamageUpdate) SetDice(s string) *WeaponDamageUpdate {
	wdu.mutation.SetDice(s)
	return wdu
}

// SetWeapon sets the "weapon" edge to the Weapon entity.
func (wdu *WeaponDamageUpdate) SetWeapon(w *Weapon) *WeaponDamageUpdate {
	return wdu.SetWeaponID(w.ID)
}

// SetDamageType sets the "damage_type" edge to the DamageType entity.
func (wdu *WeaponDamageUpdate) SetDamageType(d *DamageType) *WeaponDamageUpdate {
	return wdu.SetDamageTypeID(d.ID)
}

// Mutation returns the WeaponDamageMutation object of the builder.
func (wdu *WeaponDamageUpdate) Mutation() *WeaponDamageMutation {
	return wdu.mutation
}

// ClearWeapon clears the "weapon" edge to the Weapon entity.
func (wdu *WeaponDamageUpdate) ClearWeapon() *WeaponDamageUpdate {
	wdu.mutation.ClearWeapon()
	return wdu
}

// ClearDamageType clears the "damage_type" edge to the DamageType entity.
func (wdu *WeaponDamageUpdate) ClearDamageType() *WeaponDamageUpdate {
	wdu.mutation.ClearDamageType()
	return wdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wdu *WeaponDamageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, WeaponDamageMutation](ctx, wdu.sqlSave, wdu.mutation, wdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wdu *WeaponDamageUpdate) SaveX(ctx context.Context) int {
	affected, err := wdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wdu *WeaponDamageUpdate) Exec(ctx context.Context) error {
	_, err := wdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdu *WeaponDamageUpdate) ExecX(ctx context.Context) {
	if err := wdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdu *WeaponDamageUpdate) check() error {
	if _, ok := wdu.mutation.WeaponID(); wdu.mutation.WeaponCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WeaponDamage.weapon"`)
	}
	if _, ok := wdu.mutation.DamageTypeID(); wdu.mutation.DamageTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WeaponDamage.damage_type"`)
	}
	return nil
}

func (wdu *WeaponDamageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wdu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(weapondamage.Table, weapondamage.Columns, sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt))
	if ps := wdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wdu.mutation.Dice(); ok {
		_spec.SetField(weapondamage.FieldDice, field.TypeString, value)
	}
	if wdu.mutation.WeaponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   weapondamage.WeaponTable,
			Columns: []string{weapondamage.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wdu.mutation.WeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   weapondamage.WeaponTable,
			Columns: []string{weapondamage.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wdu.mutation.DamageTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   weapondamage.DamageTypeTable,
			Columns: []string{weapondamage.DamageTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wdu.mutation.DamageTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   weapondamage.DamageTypeTable,
			Columns: []string{weapondamage.DamageTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{weapondamage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wdu.mutation.done = true
	return n, nil
}

// WeaponDamageUpdateOne is the builder for updating a single WeaponDamage entity.
type WeaponDamageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WeaponDamageMutation
}

// SetWeaponID sets the "weapon_id" field.
func (wduo *WeaponDamageUpdateOne) SetWeaponID(i int) *WeaponDamageUpdateOne {
	wduo.mutation.SetWeaponID(i)
	return wduo
}

// SetDamageTypeID sets the "damage_type_id" field.
func (wduo *WeaponDamageUpdateOne) SetDamageTypeID(i int) *WeaponDamageUpdateOne {
	wduo.mutation.SetDamageTypeID(i)
	return wduo
}

// SetDice sets the "dice" field.
func (wduo *WeaponDamageUpdateOne) SetDice(s string) *WeaponDamageUpdateOne {
	wduo.mutation.SetDice(s)
	return wduo
}

// SetWeapon sets the "weapon" edge to the Weapon entity.
func (wduo *WeaponDamageUpdateOne) SetWeapon(w *Weapon) *WeaponDamageUpdateOne {
	return wduo.SetWeaponID(w.ID)
}

// SetDamageType sets the "damage_type" edge to the DamageType entity.
func (wduo *WeaponDamageUpdateOne) SetDamageType(d *DamageType) *WeaponDamageUpdateOne {
	return wduo.SetDamageTypeID(d.ID)
}

// Mutation returns the WeaponDamageMutation object of the builder.
func (wduo *WeaponDamageUpdateOne) Mutation() *WeaponDamageMutation {
	return wduo.mutation
}

// ClearWeapon clears the "weapon" edge to the Weapon entity.
func (wduo *WeaponDamageUpdateOne) ClearWeapon() *WeaponDamageUpdateOne {
	wduo.mutation.ClearWeapon()
	return wduo
}

// ClearDamageType clears the "damage_type" edge to the DamageType entity.
func (wduo *WeaponDamageUpdateOne) ClearDamageType() *WeaponDamageUpdateOne {
	wduo.mutation.ClearDamageType()
	return wduo
}

// Where appends a list predicates to the WeaponDamageUpdate builder.
func (wduo *WeaponDamageUpdateOne) Where(ps ...predicate.WeaponDamage) *WeaponDamageUpdateOne {
	wduo.mutation.Where(ps...)
	return wduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wduo *WeaponDamageUpdateOne) Select(field string, fields ...string) *WeaponDamageUpdateOne {
	wduo.fields = append([]string{field}, fields...)
	return wduo
}

// Save executes the query and returns the updated WeaponDamage entity.
func (wduo *WeaponDamageUpdateOne) Save(ctx context.Context) (*WeaponDamage, error) {
	return withHooks[*WeaponDamage, WeaponDamageMutation](ctx, wduo.sqlSave, wduo.mutation, wduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wduo *WeaponDamageUpdateOne) SaveX(ctx context.Context) *WeaponDamage {
	node, err := wduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wduo *WeaponDamageUpdateOne) Exec(ctx context.Context) error {
	_, err := wduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wduo *WeaponDamageUpdateOne) ExecX(ctx context.Context) {
	if err := wduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wduo *WeaponDamageUpdateOne) check() error {
	if _, ok := wduo.mutation.WeaponID(); wduo.mutation.WeaponCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WeaponDamage.weapon"`)
	}
	if _, ok := wduo.mutation.DamageTypeID(); wduo.mutation.DamageTypeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WeaponDamage.damage_type"`)
	}
	return nil
}

func (wduo *WeaponDamageUpdateOne) sqlSave(ctx context.Context) (_node *WeaponDamage, err error) {
	if err := wduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(weapondamage.Table, weapondamage.Columns, sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt))
	id, ok := wduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WeaponDamage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, weapondamage.FieldID)
		for _, f := range fields {
			if !weapondamage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != weapondamage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wduo.mutation.Dice(); ok {
		_spec.SetField(weapondamage.FieldDice, field.TypeString, value)
	}
	if wduo.mutation.WeaponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   weapondamage.WeaponTable,
			Columns: []string{weapondamage.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wduo.mutation.WeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   weapondamage.WeaponTable,
			Columns: []string{weapondamage.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wduo.mutation.DamageTypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   weapondamage.DamageTypeTable,
			Columns: []string{weapondamage.DamageTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wduo.mutation.DamageTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   weapondamage.DamageTypeTable,
			Columns: []string{weapondamage.DamageTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WeaponDamage{config: wduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{weapondamage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wduo.mutation.done = true
	return _node, nil
}
