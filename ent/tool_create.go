// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/tool"
)

// ToolCreate is the builder for creating a Tool entity.
type ToolCreate struct {
	config
	mutation *ToolMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (tc *ToolCreate) SetIndx(s string) *ToolCreate {
	tc.mutation.SetIndx(s)
	return tc
}

// SetName sets the "name" field.
func (tc *ToolCreate) SetName(s string) *ToolCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetToolCategory sets the "tool_category" field.
func (tc *ToolCreate) SetToolCategory(s string) *ToolCreate {
	tc.mutation.SetToolCategory(s)
	return tc
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (tc *ToolCreate) SetEquipmentID(id int) *ToolCreate {
	tc.mutation.SetEquipmentID(id)
	return tc
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (tc *ToolCreate) SetNillableEquipmentID(id *int) *ToolCreate {
	if id != nil {
		tc = tc.SetEquipmentID(*id)
	}
	return tc
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (tc *ToolCreate) SetEquipment(e *Equipment) *ToolCreate {
	return tc.SetEquipmentID(e.ID)
}

// Mutation returns the ToolMutation object of the builder.
func (tc *ToolCreate) Mutation() *ToolMutation {
	return tc.mutation
}

// Save creates the Tool in the database.
func (tc *ToolCreate) Save(ctx context.Context) (*Tool, error) {
	return withHooks[*Tool, ToolMutation](ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *ToolCreate) SaveX(ctx context.Context) *Tool {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *ToolCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *ToolCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *ToolCreate) check() error {
	if _, ok := tc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Tool.indx"`)}
	}
	if v, ok := tc.mutation.Indx(); ok {
		if err := tool.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Tool.indx": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Tool.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := tool.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tool.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ToolCategory(); !ok {
		return &ValidationError{Name: "tool_category", err: errors.New(`ent: missing required field "Tool.tool_category"`)}
	}
	return nil
}

func (tc *ToolCreate) sqlSave(ctx context.Context) (*Tool, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *ToolCreate) createSpec() (*Tool, *sqlgraph.CreateSpec) {
	var (
		_node = &Tool{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tool.Table, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.Indx(); ok {
		_spec.SetField(tool.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(tool.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.ToolCategory(); ok {
		_spec.SetField(tool.FieldToolCategory, field.TypeString, value)
		_node.ToolCategory = value
	}
	if nodes := tc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tool.EquipmentTable,
			Columns: []string{tool.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.equipment_tool = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ToolCreateBulk is the builder for creating many Tool entities in bulk.
type ToolCreateBulk struct {
	config
	builders []*ToolCreate
}

// Save creates the Tool entities in the database.
func (tcb *ToolCreateBulk) Save(ctx context.Context) ([]*Tool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tool, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToolMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *ToolCreateBulk) SaveX(ctx context.Context) []*Tool {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *ToolCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *ToolCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
