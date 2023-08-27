// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// WeaponDamageCreate is the builder for creating a WeaponDamage entity.
type WeaponDamageCreate struct {
	config
	mutation *WeaponDamageMutation
	hooks    []Hook
}

// SetWeaponID sets the "weapon_id" field.
func (wdc *WeaponDamageCreate) SetWeaponID(i int) *WeaponDamageCreate {
	wdc.mutation.SetWeaponID(i)
	return wdc
}

// SetDamageTypeID sets the "damage_type_id" field.
func (wdc *WeaponDamageCreate) SetDamageTypeID(i int) *WeaponDamageCreate {
	wdc.mutation.SetDamageTypeID(i)
	return wdc
}

// SetDice sets the "dice" field.
func (wdc *WeaponDamageCreate) SetDice(s string) *WeaponDamageCreate {
	wdc.mutation.SetDice(s)
	return wdc
}

// SetWeapon sets the "weapon" edge to the Weapon entity.
func (wdc *WeaponDamageCreate) SetWeapon(w *Weapon) *WeaponDamageCreate {
	return wdc.SetWeaponID(w.ID)
}

// SetDamageType sets the "damage_type" edge to the DamageType entity.
func (wdc *WeaponDamageCreate) SetDamageType(d *DamageType) *WeaponDamageCreate {
	return wdc.SetDamageTypeID(d.ID)
}

// Mutation returns the WeaponDamageMutation object of the builder.
func (wdc *WeaponDamageCreate) Mutation() *WeaponDamageMutation {
	return wdc.mutation
}

// Save creates the WeaponDamage in the database.
func (wdc *WeaponDamageCreate) Save(ctx context.Context) (*WeaponDamage, error) {
	return withHooks[*WeaponDamage, WeaponDamageMutation](ctx, wdc.sqlSave, wdc.mutation, wdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wdc *WeaponDamageCreate) SaveX(ctx context.Context) *WeaponDamage {
	v, err := wdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdc *WeaponDamageCreate) Exec(ctx context.Context) error {
	_, err := wdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdc *WeaponDamageCreate) ExecX(ctx context.Context) {
	if err := wdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wdc *WeaponDamageCreate) check() error {
	if _, ok := wdc.mutation.WeaponID(); !ok {
		return &ValidationError{Name: "weapon_id", err: errors.New(`ent: missing required field "WeaponDamage.weapon_id"`)}
	}
	if _, ok := wdc.mutation.DamageTypeID(); !ok {
		return &ValidationError{Name: "damage_type_id", err: errors.New(`ent: missing required field "WeaponDamage.damage_type_id"`)}
	}
	if _, ok := wdc.mutation.Dice(); !ok {
		return &ValidationError{Name: "dice", err: errors.New(`ent: missing required field "WeaponDamage.dice"`)}
	}
	if _, ok := wdc.mutation.WeaponID(); !ok {
		return &ValidationError{Name: "weapon", err: errors.New(`ent: missing required edge "WeaponDamage.weapon"`)}
	}
	if _, ok := wdc.mutation.DamageTypeID(); !ok {
		return &ValidationError{Name: "damage_type", err: errors.New(`ent: missing required edge "WeaponDamage.damage_type"`)}
	}
	return nil
}

func (wdc *WeaponDamageCreate) sqlSave(ctx context.Context) (*WeaponDamage, error) {
	if err := wdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wdc.mutation.id = &_node.ID
	wdc.mutation.done = true
	return _node, nil
}

func (wdc *WeaponDamageCreate) createSpec() (*WeaponDamage, *sqlgraph.CreateSpec) {
	var (
		_node = &WeaponDamage{config: wdc.config}
		_spec = sqlgraph.NewCreateSpec(weapondamage.Table, sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt))
	)
	if value, ok := wdc.mutation.Dice(); ok {
		_spec.SetField(weapondamage.FieldDice, field.TypeString, value)
		_node.Dice = value
	}
	if nodes := wdc.mutation.WeaponIDs(); len(nodes) > 0 {
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
		_node.WeaponID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wdc.mutation.DamageTypeIDs(); len(nodes) > 0 {
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
		_node.DamageTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WeaponDamageCreateBulk is the builder for creating many WeaponDamage entities in bulk.
type WeaponDamageCreateBulk struct {
	config
	builders []*WeaponDamageCreate
}

// Save creates the WeaponDamage entities in the database.
func (wdcb *WeaponDamageCreateBulk) Save(ctx context.Context) ([]*WeaponDamage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wdcb.builders))
	nodes := make([]*WeaponDamage, len(wdcb.builders))
	mutators := make([]Mutator, len(wdcb.builders))
	for i := range wdcb.builders {
		func(i int, root context.Context) {
			builder := wdcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WeaponDamageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wdcb *WeaponDamageCreateBulk) SaveX(ctx context.Context) []*WeaponDamage {
	v, err := wdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wdcb *WeaponDamageCreateBulk) Exec(ctx context.Context) error {
	_, err := wdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wdcb *WeaponDamageCreateBulk) ExecX(ctx context.Context) {
	if err := wdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
