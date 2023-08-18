// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// DamageTypeCreate is the builder for creating a DamageType entity.
type DamageTypeCreate struct {
	config
	mutation *DamageTypeMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (dtc *DamageTypeCreate) SetIndx(s string) *DamageTypeCreate {
	dtc.mutation.SetIndx(s)
	return dtc
}

// SetName sets the "name" field.
func (dtc *DamageTypeCreate) SetName(s string) *DamageTypeCreate {
	dtc.mutation.SetName(s)
	return dtc
}

// SetDesc sets the "desc" field.
func (dtc *DamageTypeCreate) SetDesc(s string) *DamageTypeCreate {
	dtc.mutation.SetDesc(s)
	return dtc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (dtc *DamageTypeCreate) SetNillableDesc(s *string) *DamageTypeCreate {
	if s != nil {
		dtc.SetDesc(*s)
	}
	return dtc
}

// SetWeaponDamageID sets the "weapon_damage" edge to the WeaponDamage entity by ID.
func (dtc *DamageTypeCreate) SetWeaponDamageID(id int) *DamageTypeCreate {
	dtc.mutation.SetWeaponDamageID(id)
	return dtc
}

// SetNillableWeaponDamageID sets the "weapon_damage" edge to the WeaponDamage entity by ID if the given value is not nil.
func (dtc *DamageTypeCreate) SetNillableWeaponDamageID(id *int) *DamageTypeCreate {
	if id != nil {
		dtc = dtc.SetWeaponDamageID(*id)
	}
	return dtc
}

// SetWeaponDamage sets the "weapon_damage" edge to the WeaponDamage entity.
func (dtc *DamageTypeCreate) SetWeaponDamage(w *WeaponDamage) *DamageTypeCreate {
	return dtc.SetWeaponDamageID(w.ID)
}

// Mutation returns the DamageTypeMutation object of the builder.
func (dtc *DamageTypeCreate) Mutation() *DamageTypeMutation {
	return dtc.mutation
}

// Save creates the DamageType in the database.
func (dtc *DamageTypeCreate) Save(ctx context.Context) (*DamageType, error) {
	return withHooks[*DamageType, DamageTypeMutation](ctx, dtc.sqlSave, dtc.mutation, dtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dtc *DamageTypeCreate) SaveX(ctx context.Context) *DamageType {
	v, err := dtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtc *DamageTypeCreate) Exec(ctx context.Context) error {
	_, err := dtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtc *DamageTypeCreate) ExecX(ctx context.Context) {
	if err := dtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtc *DamageTypeCreate) check() error {
	if _, ok := dtc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "DamageType.indx"`)}
	}
	if _, ok := dtc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DamageType.name"`)}
	}
	return nil
}

func (dtc *DamageTypeCreate) sqlSave(ctx context.Context) (*DamageType, error) {
	if err := dtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dtc.mutation.id = &_node.ID
	dtc.mutation.done = true
	return _node, nil
}

func (dtc *DamageTypeCreate) createSpec() (*DamageType, *sqlgraph.CreateSpec) {
	var (
		_node = &DamageType{config: dtc.config}
		_spec = sqlgraph.NewCreateSpec(damagetype.Table, sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt))
	)
	if value, ok := dtc.mutation.Indx(); ok {
		_spec.SetField(damagetype.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := dtc.mutation.Name(); ok {
		_spec.SetField(damagetype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dtc.mutation.Desc(); ok {
		_spec.SetField(damagetype.FieldDesc, field.TypeString, value)
		_node.Desc = &value
	}
	if nodes := dtc.mutation.WeaponDamageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   damagetype.WeaponDamageTable,
			Columns: []string{damagetype.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.weapon_damage_damage_type = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DamageTypeCreateBulk is the builder for creating many DamageType entities in bulk.
type DamageTypeCreateBulk struct {
	config
	builders []*DamageTypeCreate
}

// Save creates the DamageType entities in the database.
func (dtcb *DamageTypeCreateBulk) Save(ctx context.Context) ([]*DamageType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dtcb.builders))
	nodes := make([]*DamageType, len(dtcb.builders))
	mutators := make([]Mutator, len(dtcb.builders))
	for i := range dtcb.builders {
		func(i int, root context.Context) {
			builder := dtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DamageTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, dtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtcb *DamageTypeCreateBulk) SaveX(ctx context.Context) []*DamageType {
	v, err := dtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtcb *DamageTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := dtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtcb *DamageTypeCreateBulk) ExecX(ctx context.Context) {
	if err := dtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
