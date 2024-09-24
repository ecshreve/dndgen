// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/armorclass"
)

// ArmorClassCreate is the builder for creating a ArmorClass entity.
type ArmorClassCreate struct {
	config
	mutation *ArmorClassMutation
	hooks    []Hook
}

// SetBase sets the "base" field.
func (acc *ArmorClassCreate) SetBase(i int) *ArmorClassCreate {
	acc.mutation.SetBase(i)
	return acc
}

// SetDexBonus sets the "dex_bonus" field.
func (acc *ArmorClassCreate) SetDexBonus(b bool) *ArmorClassCreate {
	acc.mutation.SetDexBonus(b)
	return acc
}

// SetNillableDexBonus sets the "dex_bonus" field if the given value is not nil.
func (acc *ArmorClassCreate) SetNillableDexBonus(b *bool) *ArmorClassCreate {
	if b != nil {
		acc.SetDexBonus(*b)
	}
	return acc
}

// SetArmorID sets the "armor" edge to the Armor entity by ID.
func (acc *ArmorClassCreate) SetArmorID(id int) *ArmorClassCreate {
	acc.mutation.SetArmorID(id)
	return acc
}

// SetNillableArmorID sets the "armor" edge to the Armor entity by ID if the given value is not nil.
func (acc *ArmorClassCreate) SetNillableArmorID(id *int) *ArmorClassCreate {
	if id != nil {
		acc = acc.SetArmorID(*id)
	}
	return acc
}

// SetArmor sets the "armor" edge to the Armor entity.
func (acc *ArmorClassCreate) SetArmor(a *Armor) *ArmorClassCreate {
	return acc.SetArmorID(a.ID)
}

// Mutation returns the ArmorClassMutation object of the builder.
func (acc *ArmorClassCreate) Mutation() *ArmorClassMutation {
	return acc.mutation
}

// Save creates the ArmorClass in the database.
func (acc *ArmorClassCreate) Save(ctx context.Context) (*ArmorClass, error) {
	acc.defaults()
	return withHooks(ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *ArmorClassCreate) SaveX(ctx context.Context) *ArmorClass {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *ArmorClassCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *ArmorClassCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *ArmorClassCreate) defaults() {
	if _, ok := acc.mutation.DexBonus(); !ok {
		v := armorclass.DefaultDexBonus
		acc.mutation.SetDexBonus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *ArmorClassCreate) check() error {
	if _, ok := acc.mutation.Base(); !ok {
		return &ValidationError{Name: "base", err: errors.New(`ent: missing required field "ArmorClass.base"`)}
	}
	if v, ok := acc.mutation.Base(); ok {
		if err := armorclass.BaseValidator(v); err != nil {
			return &ValidationError{Name: "base", err: fmt.Errorf(`ent: validator failed for field "ArmorClass.base": %w`, err)}
		}
	}
	if _, ok := acc.mutation.DexBonus(); !ok {
		return &ValidationError{Name: "dex_bonus", err: errors.New(`ent: missing required field "ArmorClass.dex_bonus"`)}
	}
	return nil
}

func (acc *ArmorClassCreate) sqlSave(ctx context.Context) (*ArmorClass, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *ArmorClassCreate) createSpec() (*ArmorClass, *sqlgraph.CreateSpec) {
	var (
		_node = &ArmorClass{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(armorclass.Table, sqlgraph.NewFieldSpec(armorclass.FieldID, field.TypeInt))
	)
	if value, ok := acc.mutation.Base(); ok {
		_spec.SetField(armorclass.FieldBase, field.TypeInt, value)
		_node.Base = value
	}
	if value, ok := acc.mutation.DexBonus(); ok {
		_spec.SetField(armorclass.FieldDexBonus, field.TypeBool, value)
		_node.DexBonus = value
	}
	if nodes := acc.mutation.ArmorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   armorclass.ArmorTable,
			Columns: []string{armorclass.ArmorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.armor_armor_class = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArmorClassCreateBulk is the builder for creating many ArmorClass entities in bulk.
type ArmorClassCreateBulk struct {
	config
	err      error
	builders []*ArmorClassCreate
}

// Save creates the ArmorClass entities in the database.
func (accb *ArmorClassCreateBulk) Save(ctx context.Context) ([]*ArmorClass, error) {
	if accb.err != nil {
		return nil, accb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*ArmorClass, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArmorClassMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *ArmorClassCreateBulk) SaveX(ctx context.Context) []*ArmorClass {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *ArmorClassCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *ArmorClassCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
