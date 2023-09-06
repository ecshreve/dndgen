// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentchoice"
)

// EquipmentChoiceCreate is the builder for creating a EquipmentChoice entity.
type EquipmentChoiceCreate struct {
	config
	mutation *EquipmentChoiceMutation
	hooks    []Hook
}

// SetClassID sets the "class_id" field.
func (ecc *EquipmentChoiceCreate) SetClassID(i int) *EquipmentChoiceCreate {
	ecc.mutation.SetClassID(i)
	return ecc
}

// SetChoose sets the "choose" field.
func (ecc *EquipmentChoiceCreate) SetChoose(i int) *EquipmentChoiceCreate {
	ecc.mutation.SetChoose(i)
	return ecc
}

// SetDesc sets the "desc" field.
func (ecc *EquipmentChoiceCreate) SetDesc(s string) *EquipmentChoiceCreate {
	ecc.mutation.SetDesc(s)
	return ecc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (ecc *EquipmentChoiceCreate) SetNillableDesc(s *string) *EquipmentChoiceCreate {
	if s != nil {
		ecc.SetDesc(*s)
	}
	return ecc
}

// SetClass sets the "class" edge to the Class entity.
func (ecc *EquipmentChoiceCreate) SetClass(c *Class) *EquipmentChoiceCreate {
	return ecc.SetClassID(c.ID)
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (ecc *EquipmentChoiceCreate) AddEquipmentIDs(ids ...int) *EquipmentChoiceCreate {
	ecc.mutation.AddEquipmentIDs(ids...)
	return ecc
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (ecc *EquipmentChoiceCreate) AddEquipment(e ...*Equipment) *EquipmentChoiceCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecc.AddEquipmentIDs(ids...)
}

// Mutation returns the EquipmentChoiceMutation object of the builder.
func (ecc *EquipmentChoiceCreate) Mutation() *EquipmentChoiceMutation {
	return ecc.mutation
}

// Save creates the EquipmentChoice in the database.
func (ecc *EquipmentChoiceCreate) Save(ctx context.Context) (*EquipmentChoice, error) {
	return withHooks(ctx, ecc.sqlSave, ecc.mutation, ecc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ecc *EquipmentChoiceCreate) SaveX(ctx context.Context) *EquipmentChoice {
	v, err := ecc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecc *EquipmentChoiceCreate) Exec(ctx context.Context) error {
	_, err := ecc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecc *EquipmentChoiceCreate) ExecX(ctx context.Context) {
	if err := ecc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecc *EquipmentChoiceCreate) check() error {
	if _, ok := ecc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class_id", err: errors.New(`ent: missing required field "EquipmentChoice.class_id"`)}
	}
	if _, ok := ecc.mutation.Choose(); !ok {
		return &ValidationError{Name: "choose", err: errors.New(`ent: missing required field "EquipmentChoice.choose"`)}
	}
	if _, ok := ecc.mutation.ClassID(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required edge "EquipmentChoice.class"`)}
	}
	return nil
}

func (ecc *EquipmentChoiceCreate) sqlSave(ctx context.Context) (*EquipmentChoice, error) {
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

func (ecc *EquipmentChoiceCreate) createSpec() (*EquipmentChoice, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentChoice{config: ecc.config}
		_spec = sqlgraph.NewCreateSpec(equipmentchoice.Table, sqlgraph.NewFieldSpec(equipmentchoice.FieldID, field.TypeInt))
	)
	if value, ok := ecc.mutation.Choose(); ok {
		_spec.SetField(equipmentchoice.FieldChoose, field.TypeInt, value)
		_node.Choose = value
	}
	if value, ok := ecc.mutation.Desc(); ok {
		_spec.SetField(equipmentchoice.FieldDesc, field.TypeString, value)
		_node.Desc = value
	}
	if nodes := ecc.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: []string{equipmentchoice.ClassColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClassID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ecc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.EquipmentTable,
			Columns: equipmentchoice.EquipmentPrimaryKey,
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

// EquipmentChoiceCreateBulk is the builder for creating many EquipmentChoice entities in bulk.
type EquipmentChoiceCreateBulk struct {
	config
	builders []*EquipmentChoiceCreate
}

// Save creates the EquipmentChoice entities in the database.
func (eccb *EquipmentChoiceCreateBulk) Save(ctx context.Context) ([]*EquipmentChoice, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eccb.builders))
	nodes := make([]*EquipmentChoice, len(eccb.builders))
	mutators := make([]Mutator, len(eccb.builders))
	for i := range eccb.builders {
		func(i int, root context.Context) {
			builder := eccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentChoiceMutation)
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
func (eccb *EquipmentChoiceCreateBulk) SaveX(ctx context.Context) []*EquipmentChoice {
	v, err := eccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eccb *EquipmentChoiceCreateBulk) Exec(ctx context.Context) error {
	_, err := eccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eccb *EquipmentChoiceCreateBulk) ExecX(ctx context.Context) {
	if err := eccb.Exec(ctx); err != nil {
		panic(err)
	}
}