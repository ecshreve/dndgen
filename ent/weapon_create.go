// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damage"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/property"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// WeaponCreate is the builder for creating a Weapon entity.
type WeaponCreate struct {
	config
	mutation *WeaponMutation
	hooks    []Hook
}

// SetWeaponCategory sets the "weapon_category" field.
func (wc *WeaponCreate) SetWeaponCategory(value weapon.WeaponCategory) *WeaponCreate {
	wc.mutation.SetWeaponCategory(value)
	return wc
}

// SetWeaponSubcategory sets the "weapon_subcategory" field.
func (wc *WeaponCreate) SetWeaponSubcategory(ws weapon.WeaponSubcategory) *WeaponCreate {
	wc.mutation.SetWeaponSubcategory(ws)
	return wc
}

// SetDamageID sets the "damage" edge to the Damage entity by ID.
func (wc *WeaponCreate) SetDamageID(id int) *WeaponCreate {
	wc.mutation.SetDamageID(id)
	return wc
}

// SetNillableDamageID sets the "damage" edge to the Damage entity by ID if the given value is not nil.
func (wc *WeaponCreate) SetNillableDamageID(id *int) *WeaponCreate {
	if id != nil {
		wc = wc.SetDamageID(*id)
	}
	return wc
}

// SetDamage sets the "damage" edge to the Damage entity.
func (wc *WeaponCreate) SetDamage(d *Damage) *WeaponCreate {
	return wc.SetDamageID(d.ID)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (wc *WeaponCreate) AddPropertyIDs(ids ...int) *WeaponCreate {
	wc.mutation.AddPropertyIDs(ids...)
	return wc
}

// AddProperties adds the "properties" edges to the Property entity.
func (wc *WeaponCreate) AddProperties(p ...*Property) *WeaponCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return wc.AddPropertyIDs(ids...)
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (wc *WeaponCreate) SetEquipmentID(id int) *WeaponCreate {
	wc.mutation.SetEquipmentID(id)
	return wc
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (wc *WeaponCreate) SetNillableEquipmentID(id *int) *WeaponCreate {
	if id != nil {
		wc = wc.SetEquipmentID(*id)
	}
	return wc
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (wc *WeaponCreate) SetEquipment(e *Equipment) *WeaponCreate {
	return wc.SetEquipmentID(e.ID)
}

// Mutation returns the WeaponMutation object of the builder.
func (wc *WeaponCreate) Mutation() *WeaponMutation {
	return wc.mutation
}

// Save creates the Weapon in the database.
func (wc *WeaponCreate) Save(ctx context.Context) (*Weapon, error) {
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WeaponCreate) SaveX(ctx context.Context) *Weapon {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WeaponCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WeaponCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WeaponCreate) check() error {
	if _, ok := wc.mutation.WeaponCategory(); !ok {
		return &ValidationError{Name: "weapon_category", err: errors.New(`ent: missing required field "Weapon.weapon_category"`)}
	}
	if v, ok := wc.mutation.WeaponCategory(); ok {
		if err := weapon.WeaponCategoryValidator(v); err != nil {
			return &ValidationError{Name: "weapon_category", err: fmt.Errorf(`ent: validator failed for field "Weapon.weapon_category": %w`, err)}
		}
	}
	if _, ok := wc.mutation.WeaponSubcategory(); !ok {
		return &ValidationError{Name: "weapon_subcategory", err: errors.New(`ent: missing required field "Weapon.weapon_subcategory"`)}
	}
	if v, ok := wc.mutation.WeaponSubcategory(); ok {
		if err := weapon.WeaponSubcategoryValidator(v); err != nil {
			return &ValidationError{Name: "weapon_subcategory", err: fmt.Errorf(`ent: validator failed for field "Weapon.weapon_subcategory": %w`, err)}
		}
	}
	return nil
}

func (wc *WeaponCreate) sqlSave(ctx context.Context) (*Weapon, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WeaponCreate) createSpec() (*Weapon, *sqlgraph.CreateSpec) {
	var (
		_node = &Weapon{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(weapon.Table, sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.WeaponCategory(); ok {
		_spec.SetField(weapon.FieldWeaponCategory, field.TypeEnum, value)
		_node.WeaponCategory = value
	}
	if value, ok := wc.mutation.WeaponSubcategory(); ok {
		_spec.SetField(weapon.FieldWeaponSubcategory, field.TypeEnum, value)
		_node.WeaponSubcategory = value
	}
	if nodes := wc.mutation.DamageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   weapon.DamageTable,
			Columns: []string{weapon.DamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(damage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.weapon_damage = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.PropertiesTable,
			Columns: weapon.PropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(property.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   weapon.EquipmentTable,
			Columns: []string{weapon.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.weapon_equipment = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WeaponCreateBulk is the builder for creating many Weapon entities in bulk.
type WeaponCreateBulk struct {
	config
	err      error
	builders []*WeaponCreate
}

// Save creates the Weapon entities in the database.
func (wcb *WeaponCreateBulk) Save(ctx context.Context) ([]*Weapon, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Weapon, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WeaponMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WeaponCreateBulk) SaveX(ctx context.Context) []*Weapon {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WeaponCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WeaponCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
