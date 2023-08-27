// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weaponproperty"
)

// WeaponPropertyCreate is the builder for creating a WeaponProperty entity.
type WeaponPropertyCreate struct {
	config
	mutation *WeaponPropertyMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (wpc *WeaponPropertyCreate) SetIndx(s string) *WeaponPropertyCreate {
	wpc.mutation.SetIndx(s)
	return wpc
}

// SetName sets the "name" field.
func (wpc *WeaponPropertyCreate) SetName(s string) *WeaponPropertyCreate {
	wpc.mutation.SetName(s)
	return wpc
}

// SetDesc sets the "desc" field.
func (wpc *WeaponPropertyCreate) SetDesc(s []string) *WeaponPropertyCreate {
	wpc.mutation.SetDesc(s)
	return wpc
}

// AddWeaponIDs adds the "weapons" edge to the Weapon entity by IDs.
func (wpc *WeaponPropertyCreate) AddWeaponIDs(ids ...int) *WeaponPropertyCreate {
	wpc.mutation.AddWeaponIDs(ids...)
	return wpc
}

// AddWeapons adds the "weapons" edges to the Weapon entity.
func (wpc *WeaponPropertyCreate) AddWeapons(w ...*Weapon) *WeaponPropertyCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wpc.AddWeaponIDs(ids...)
}

// Mutation returns the WeaponPropertyMutation object of the builder.
func (wpc *WeaponPropertyCreate) Mutation() *WeaponPropertyMutation {
	return wpc.mutation
}

// Save creates the WeaponProperty in the database.
func (wpc *WeaponPropertyCreate) Save(ctx context.Context) (*WeaponProperty, error) {
	return withHooks[*WeaponProperty, WeaponPropertyMutation](ctx, wpc.sqlSave, wpc.mutation, wpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wpc *WeaponPropertyCreate) SaveX(ctx context.Context) *WeaponProperty {
	v, err := wpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wpc *WeaponPropertyCreate) Exec(ctx context.Context) error {
	_, err := wpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpc *WeaponPropertyCreate) ExecX(ctx context.Context) {
	if err := wpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wpc *WeaponPropertyCreate) check() error {
	if _, ok := wpc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "WeaponProperty.indx"`)}
	}
	if v, ok := wpc.mutation.Indx(); ok {
		if err := weaponproperty.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "WeaponProperty.indx": %w`, err)}
		}
	}
	if _, ok := wpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WeaponProperty.name"`)}
	}
	if v, ok := wpc.mutation.Name(); ok {
		if err := weaponproperty.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "WeaponProperty.name": %w`, err)}
		}
	}
	if _, ok := wpc.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "WeaponProperty.desc"`)}
	}
	return nil
}

func (wpc *WeaponPropertyCreate) sqlSave(ctx context.Context) (*WeaponProperty, error) {
	if err := wpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wpc.mutation.id = &_node.ID
	wpc.mutation.done = true
	return _node, nil
}

func (wpc *WeaponPropertyCreate) createSpec() (*WeaponProperty, *sqlgraph.CreateSpec) {
	var (
		_node = &WeaponProperty{config: wpc.config}
		_spec = sqlgraph.NewCreateSpec(weaponproperty.Table, sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt))
	)
	if value, ok := wpc.mutation.Indx(); ok {
		_spec.SetField(weaponproperty.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := wpc.mutation.Name(); ok {
		_spec.SetField(weaponproperty.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wpc.mutation.Desc(); ok {
		_spec.SetField(weaponproperty.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if nodes := wpc.mutation.WeaponsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   weaponproperty.WeaponsTable,
			Columns: weaponproperty.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WeaponPropertyCreateBulk is the builder for creating many WeaponProperty entities in bulk.
type WeaponPropertyCreateBulk struct {
	config
	builders []*WeaponPropertyCreate
}

// Save creates the WeaponProperty entities in the database.
func (wpcb *WeaponPropertyCreateBulk) Save(ctx context.Context) ([]*WeaponProperty, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wpcb.builders))
	nodes := make([]*WeaponProperty, len(wpcb.builders))
	mutators := make([]Mutator, len(wpcb.builders))
	for i := range wpcb.builders {
		func(i int, root context.Context) {
			builder := wpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WeaponPropertyMutation)
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
					_, err = mutators[i+1].Mutate(root, wpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wpcb *WeaponPropertyCreateBulk) SaveX(ctx context.Context) []*WeaponProperty {
	v, err := wpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wpcb *WeaponPropertyCreateBulk) Exec(ctx context.Context) error {
	_, err := wpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wpcb *WeaponPropertyCreateBulk) ExecX(ctx context.Context) {
	if err := wpcb.Exec(ctx); err != nil {
		panic(err)
	}
}