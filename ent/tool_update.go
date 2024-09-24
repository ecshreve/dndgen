// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/tool"
)

// ToolUpdate is the builder for updating Tool entities.
type ToolUpdate struct {
	config
	hooks    []Hook
	mutation *ToolMutation
}

// Where appends a list predicates to the ToolUpdate builder.
func (tu *ToolUpdate) Where(ps ...predicate.Tool) *ToolUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetToolCategory sets the "tool_category" field.
func (tu *ToolUpdate) SetToolCategory(s string) *ToolUpdate {
	tu.mutation.SetToolCategory(s)
	return tu
}

// SetNillableToolCategory sets the "tool_category" field if the given value is not nil.
func (tu *ToolUpdate) SetNillableToolCategory(s *string) *ToolUpdate {
	if s != nil {
		tu.SetToolCategory(*s)
	}
	return tu
}

// ClearToolCategory clears the value of the "tool_category" field.
func (tu *ToolUpdate) ClearToolCategory() *ToolUpdate {
	tu.mutation.ClearToolCategory()
	return tu
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (tu *ToolUpdate) SetEquipmentID(id int) *ToolUpdate {
	tu.mutation.SetEquipmentID(id)
	return tu
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (tu *ToolUpdate) SetNillableEquipmentID(id *int) *ToolUpdate {
	if id != nil {
		tu = tu.SetEquipmentID(*id)
	}
	return tu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (tu *ToolUpdate) SetEquipment(e *Equipment) *ToolUpdate {
	return tu.SetEquipmentID(e.ID)
}

// Mutation returns the ToolMutation object of the builder.
func (tu *ToolUpdate) Mutation() *ToolMutation {
	return tu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (tu *ToolUpdate) ClearEquipment() *ToolUpdate {
	tu.mutation.ClearEquipment()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *ToolUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *ToolUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *ToolUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *ToolUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *ToolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tool.Table, tool.Columns, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.ToolCategory(); ok {
		_spec.SetField(tool.FieldToolCategory, field.TypeString, value)
	}
	if tu.mutation.ToolCategoryCleared() {
		_spec.ClearField(tool.FieldToolCategory, field.TypeString)
	}
	if tu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tool.EquipmentTable,
			Columns: []string{tool.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// ToolUpdateOne is the builder for updating a single Tool entity.
type ToolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ToolMutation
}

// SetToolCategory sets the "tool_category" field.
func (tuo *ToolUpdateOne) SetToolCategory(s string) *ToolUpdateOne {
	tuo.mutation.SetToolCategory(s)
	return tuo
}

// SetNillableToolCategory sets the "tool_category" field if the given value is not nil.
func (tuo *ToolUpdateOne) SetNillableToolCategory(s *string) *ToolUpdateOne {
	if s != nil {
		tuo.SetToolCategory(*s)
	}
	return tuo
}

// ClearToolCategory clears the value of the "tool_category" field.
func (tuo *ToolUpdateOne) ClearToolCategory() *ToolUpdateOne {
	tuo.mutation.ClearToolCategory()
	return tuo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (tuo *ToolUpdateOne) SetEquipmentID(id int) *ToolUpdateOne {
	tuo.mutation.SetEquipmentID(id)
	return tuo
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (tuo *ToolUpdateOne) SetNillableEquipmentID(id *int) *ToolUpdateOne {
	if id != nil {
		tuo = tuo.SetEquipmentID(*id)
	}
	return tuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (tuo *ToolUpdateOne) SetEquipment(e *Equipment) *ToolUpdateOne {
	return tuo.SetEquipmentID(e.ID)
}

// Mutation returns the ToolMutation object of the builder.
func (tuo *ToolUpdateOne) Mutation() *ToolMutation {
	return tuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (tuo *ToolUpdateOne) ClearEquipment() *ToolUpdateOne {
	tuo.mutation.ClearEquipment()
	return tuo
}

// Where appends a list predicates to the ToolUpdate builder.
func (tuo *ToolUpdateOne) Where(ps ...predicate.Tool) *ToolUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *ToolUpdateOne) Select(field string, fields ...string) *ToolUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tool entity.
func (tuo *ToolUpdateOne) Save(ctx context.Context) (*Tool, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *ToolUpdateOne) SaveX(ctx context.Context) *Tool {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *ToolUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *ToolUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *ToolUpdateOne) sqlSave(ctx context.Context) (_node *Tool, err error) {
	_spec := sqlgraph.NewUpdateSpec(tool.Table, tool.Columns, sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tool.FieldID)
		for _, f := range fields {
			if !tool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.ToolCategory(); ok {
		_spec.SetField(tool.FieldToolCategory, field.TypeString, value)
	}
	if tuo.mutation.ToolCategoryCleared() {
		_spec.ClearField(tool.FieldToolCategory, field.TypeString)
	}
	if tuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tool.EquipmentTable,
			Columns: []string{tool.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tool{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
