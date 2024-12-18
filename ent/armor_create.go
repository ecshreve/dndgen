// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/equipment"
)

// ArmorCreate is the builder for creating a Armor entity.
type ArmorCreate struct {
	config
	mutation *ArmorMutation
	hooks    []Hook
}

// SetArmorCategory sets the "armor_category" field.
func (ac *ArmorCreate) SetArmorCategory(value armor.ArmorCategory) *ArmorCreate {
	ac.mutation.SetArmorCategory(value)
	return ac
}

// SetStrMinimum sets the "str_minimum" field.
func (ac *ArmorCreate) SetStrMinimum(i int) *ArmorCreate {
	ac.mutation.SetStrMinimum(i)
	return ac
}

// SetStealthDisadvantage sets the "stealth_disadvantage" field.
func (ac *ArmorCreate) SetStealthDisadvantage(b bool) *ArmorCreate {
	ac.mutation.SetStealthDisadvantage(b)
	return ac
}

// SetAcBase sets the "ac_base" field.
func (ac *ArmorCreate) SetAcBase(i int) *ArmorCreate {
	ac.mutation.SetAcBase(i)
	return ac
}

// SetAcDexBonus sets the "ac_dex_bonus" field.
func (ac *ArmorCreate) SetAcDexBonus(b bool) *ArmorCreate {
	ac.mutation.SetAcDexBonus(b)
	return ac
}

// SetNillableAcDexBonus sets the "ac_dex_bonus" field if the given value is not nil.
func (ac *ArmorCreate) SetNillableAcDexBonus(b *bool) *ArmorCreate {
	if b != nil {
		ac.SetAcDexBonus(*b)
	}
	return ac
}

// SetAcMaxBonus sets the "ac_max_bonus" field.
func (ac *ArmorCreate) SetAcMaxBonus(i int) *ArmorCreate {
	ac.mutation.SetAcMaxBonus(i)
	return ac
}

// SetNillableAcMaxBonus sets the "ac_max_bonus" field if the given value is not nil.
func (ac *ArmorCreate) SetNillableAcMaxBonus(i *int) *ArmorCreate {
	if i != nil {
		ac.SetAcMaxBonus(*i)
	}
	return ac
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (ac *ArmorCreate) SetEquipmentID(id int) *ArmorCreate {
	ac.mutation.SetEquipmentID(id)
	return ac
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (ac *ArmorCreate) SetEquipment(e *Equipment) *ArmorCreate {
	return ac.SetEquipmentID(e.ID)
}

// Mutation returns the ArmorMutation object of the builder.
func (ac *ArmorCreate) Mutation() *ArmorMutation {
	return ac.mutation
}

// Save creates the Armor in the database.
func (ac *ArmorCreate) Save(ctx context.Context) (*Armor, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArmorCreate) SaveX(ctx context.Context) *Armor {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArmorCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArmorCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArmorCreate) defaults() {
	if _, ok := ac.mutation.AcDexBonus(); !ok {
		v := armor.DefaultAcDexBonus
		ac.mutation.SetAcDexBonus(v)
	}
	if _, ok := ac.mutation.AcMaxBonus(); !ok {
		v := armor.DefaultAcMaxBonus
		ac.mutation.SetAcMaxBonus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArmorCreate) check() error {
	if _, ok := ac.mutation.ArmorCategory(); !ok {
		return &ValidationError{Name: "armor_category", err: errors.New(`ent: missing required field "Armor.armor_category"`)}
	}
	if v, ok := ac.mutation.ArmorCategory(); ok {
		if err := armor.ArmorCategoryValidator(v); err != nil {
			return &ValidationError{Name: "armor_category", err: fmt.Errorf(`ent: validator failed for field "Armor.armor_category": %w`, err)}
		}
	}
	if _, ok := ac.mutation.StrMinimum(); !ok {
		return &ValidationError{Name: "str_minimum", err: errors.New(`ent: missing required field "Armor.str_minimum"`)}
	}
	if _, ok := ac.mutation.StealthDisadvantage(); !ok {
		return &ValidationError{Name: "stealth_disadvantage", err: errors.New(`ent: missing required field "Armor.stealth_disadvantage"`)}
	}
	if _, ok := ac.mutation.AcBase(); !ok {
		return &ValidationError{Name: "ac_base", err: errors.New(`ent: missing required field "Armor.ac_base"`)}
	}
	if v, ok := ac.mutation.AcBase(); ok {
		if err := armor.AcBaseValidator(v); err != nil {
			return &ValidationError{Name: "ac_base", err: fmt.Errorf(`ent: validator failed for field "Armor.ac_base": %w`, err)}
		}
	}
	if _, ok := ac.mutation.AcDexBonus(); !ok {
		return &ValidationError{Name: "ac_dex_bonus", err: errors.New(`ent: missing required field "Armor.ac_dex_bonus"`)}
	}
	if _, ok := ac.mutation.AcMaxBonus(); !ok {
		return &ValidationError{Name: "ac_max_bonus", err: errors.New(`ent: missing required field "Armor.ac_max_bonus"`)}
	}
	if len(ac.mutation.EquipmentIDs()) == 0 {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "Armor.equipment"`)}
	}
	return nil
}

func (ac *ArmorCreate) sqlSave(ctx context.Context) (*Armor, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArmorCreate) createSpec() (*Armor, *sqlgraph.CreateSpec) {
	var (
		_node = &Armor{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(armor.Table, sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.ArmorCategory(); ok {
		_spec.SetField(armor.FieldArmorCategory, field.TypeEnum, value)
		_node.ArmorCategory = value
	}
	if value, ok := ac.mutation.StrMinimum(); ok {
		_spec.SetField(armor.FieldStrMinimum, field.TypeInt, value)
		_node.StrMinimum = value
	}
	if value, ok := ac.mutation.StealthDisadvantage(); ok {
		_spec.SetField(armor.FieldStealthDisadvantage, field.TypeBool, value)
		_node.StealthDisadvantage = value
	}
	if value, ok := ac.mutation.AcBase(); ok {
		_spec.SetField(armor.FieldAcBase, field.TypeInt, value)
		_node.AcBase = value
	}
	if value, ok := ac.mutation.AcDexBonus(); ok {
		_spec.SetField(armor.FieldAcDexBonus, field.TypeBool, value)
		_node.AcDexBonus = value
	}
	if value, ok := ac.mutation.AcMaxBonus(); ok {
		_spec.SetField(armor.FieldAcMaxBonus, field.TypeInt, value)
		_node.AcMaxBonus = value
	}
	if nodes := ac.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: []string{armor.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.equipment_armor = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArmorCreateBulk is the builder for creating many Armor entities in bulk.
type ArmorCreateBulk struct {
	config
	err      error
	builders []*ArmorCreate
}

// Save creates the Armor entities in the database.
func (acb *ArmorCreateBulk) Save(ctx context.Context) ([]*Armor, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Armor, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArmorMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArmorCreateBulk) SaveX(ctx context.Context) []*Armor {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArmorCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArmorCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
