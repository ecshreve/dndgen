// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentcategory"
)

// EquipmentCategoryCreate is the builder for creating a EquipmentCategory entity.
type EquipmentCategoryCreate struct {
	config
	mutation *EquipmentCategoryMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (ecc *EquipmentCategoryCreate) SetIndx(s string) *EquipmentCategoryCreate {
	ecc.mutation.SetIndx(s)
	return ecc
}

// SetName sets the "name" field.
func (ecc *EquipmentCategoryCreate) SetName(s string) *EquipmentCategoryCreate {
	ecc.mutation.SetName(s)
	return ecc
}

// SetDesc sets the "desc" field.
func (ecc *EquipmentCategoryCreate) SetDesc(s []string) *EquipmentCategoryCreate {
	ecc.mutation.SetDesc(s)
	return ecc
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (ecc *EquipmentCategoryCreate) AddEquipmentIDs(ids ...int) *EquipmentCategoryCreate {
	ecc.mutation.AddEquipmentIDs(ids...)
	return ecc
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (ecc *EquipmentCategoryCreate) AddEquipment(e ...*Equipment) *EquipmentCategoryCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddEquipmentIDs(ids...)
}

// Mutation returns the EquipmentCategoryMutation object of the builder.
func (ecc *EquipmentCategoryCreate) Mutation() *EquipmentCategoryMutation {
	return ecc.mutation
}

// Save creates the EquipmentCategory in the database.
func (ecc *EquipmentCategoryCreate) Save(ctx context.Context) (*EquipmentCategory, error) {
	return withHooks[*EquipmentCategory, EquipmentCategoryMutation](ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *EquipmentCategoryCreate) SaveX(ctx context.Context) *EquipmentCategory {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *EquipmentCategoryCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *EquipmentCategoryCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *EquipmentCategoryCreate) check() error {
	if _, ok := ecc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "EquipmentCategory.indx"`)}
	}
	if _, ok := ecc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EquipmentCategory.name"`)}
	}
	return nil
}

func (ecc *EquipmentCategoryCreate) sqlSave(ctx context.Context) (*EquipmentCategory, error) {
	if err := ecc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ecc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ecc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ecc.mutation.id = &_node.ID
	ecc.mutation.done = true
	return _node, nil
}

func (ecc *EquipmentCategoryCreate) createSpec() (*EquipmentCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentCategory{config: ecc.config}
		_spec = sqlgraph.NewCreateSpec(equipmentcategory.Table, sqlgraph.NewFieldSpec(equipmentcategory.FieldID, field.TypeInt))
	)
	if value, ok := ecc.mutation.Indx(); ok {
		_spec.SetField(equipmentcategory.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := ecc.mutation.Name(); ok {
		_spec.SetField(equipmentcategory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ecc.mutation.Desc(); ok {
		_spec.SetField(equipmentcategory.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if nodes := ecc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentcategory.EquipmentTable,
			Columns: equipmentcategory.EquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentCategoryCreateBulk is the builder for creating many EquipmentCategory entities in bulk.
type EquipmentCategoryCreateBulk struct {
	config
	builders []*EquipmentCategoryCreate
}

// Save creates the EquipmentCategory entities in the database.
func (eccb *EquipmentCategoryCreateBulk) Save(ctx context.Context) ([]*EquipmentCategory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*EquipmentCategory, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, eccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eccb *EquipmentCategoryCreateBulk) SaveX(ctx context.Context) []*EquipmentCategory {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *EquipmentCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *EquipmentCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}
