// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentchoice"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// EquipmentChoiceUpdate is the builder for updating EquipmentChoice entities.
type EquipmentChoiceUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentChoiceMutation
}

// Where appends a list predicates to the EquipmentChoiceUpdate builder.
func (ecu *EquipmentChoiceUpdate) Where(ps ...predicate.EquipmentChoice) *EquipmentChoiceUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetChoose sets the "choose" field.
func (ecu *EquipmentChoiceUpdate) SetChoose(i int) *EquipmentChoiceUpdate {
	ecu.mutation.ResetChoose()
	ecu.mutation.SetChoose(i)
	return ecu
}

// AddChoose adds i to the "choose" field.
func (ecu *EquipmentChoiceUpdate) AddChoose(i int) *EquipmentChoiceUpdate {
	ecu.mutation.AddChoose(i)
	return ecu
}

// SetDesc sets the "desc" field.
func (ecu *EquipmentChoiceUpdate) SetDesc(s string) *EquipmentChoiceUpdate {
	ecu.mutation.SetDesc(s)
	return ecu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (ecu *EquipmentChoiceUpdate) SetNillableDesc(s *string) *EquipmentChoiceUpdate {
	if s != nil {
		ecu.SetDesc(*s)
	}
	return ecu
}

// ClearDesc clears the value of the "desc" field.
func (ecu *EquipmentChoiceUpdate) ClearDesc() *EquipmentChoiceUpdate {
	ecu.mutation.ClearDesc()
	return ecu
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (ecu *EquipmentChoiceUpdate) AddClasIDs(ids ...int) *EquipmentChoiceUpdate {
	ecu.mutation.AddClasIDs(ids...)
	return ecu
}

// AddClass adds the "class" edges to the Class entity.
func (ecu *EquipmentChoiceUpdate) AddClass(c ...*Class) *EquipmentChoiceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ecu.AddClasIDs(ids...)
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (ecu *EquipmentChoiceUpdate) AddEquipmentIDs(ids ...int) *EquipmentChoiceUpdate {
	ecu.mutation.AddEquipmentIDs(ids...)
	return ecu
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (ecu *EquipmentChoiceUpdate) AddEquipment(e ...*Equipment) *EquipmentChoiceUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.AddEquipmentIDs(ids...)
}

// Mutation returns the EquipmentChoiceMutation object of the builder.
func (ecu *EquipmentChoiceUpdate) Mutation() *EquipmentChoiceMutation {
	return ecu.mutation
}

// ClearClass clears all "class" edges to the Class entity.
func (ecu *EquipmentChoiceUpdate) ClearClass() *EquipmentChoiceUpdate {
	ecu.mutation.ClearClass()
	return ecu
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (ecu *EquipmentChoiceUpdate) RemoveClasIDs(ids ...int) *EquipmentChoiceUpdate {
	ecu.mutation.RemoveClasIDs(ids...)
	return ecu
}

// RemoveClass removes "class" edges to Class entities.
func (ecu *EquipmentChoiceUpdate) RemoveClass(c ...*Class) *EquipmentChoiceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ecu.RemoveClasIDs(ids...)
}

// ClearEquipment clears all "equipment" edges to the Equipment entity.
func (ecu *EquipmentChoiceUpdate) ClearEquipment() *EquipmentChoiceUpdate {
	ecu.mutation.ClearEquipment()
	return ecu
}

// RemoveEquipmentIDs removes the "equipment" edge to Equipment entities by IDs.
func (ecu *EquipmentChoiceUpdate) RemoveEquipmentIDs(ids ...int) *EquipmentChoiceUpdate {
	ecu.mutation.RemoveEquipmentIDs(ids...)
	return ecu
}

// RemoveEquipment removes "equipment" edges to Equipment entities.
func (ecu *EquipmentChoiceUpdate) RemoveEquipment(e ...*Equipment) *EquipmentChoiceUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecu.RemoveEquipmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EquipmentChoiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EquipmentChoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EquipmentChoiceUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EquipmentChoiceUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ecu *EquipmentChoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(equipmentchoice.Table, equipmentchoice.Columns, sqlgraph.NewFieldSpec(equipmentchoice.FieldID, field.TypeInt))
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Choose(); ok {
		_spec.SetField(equipmentchoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := ecu.mutation.AddedChoose(); ok {
		_spec.AddField(equipmentchoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := ecu.mutation.Desc(); ok {
		_spec.SetField(equipmentchoice.FieldDesc, field.TypeString, value)
	}
	if ecu.mutation.DescCleared() {
		_spec.ClearField(equipmentchoice.FieldDesc, field.TypeString)
	}
	if ecu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: equipmentchoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedClassIDs(); len(nodes) > 0 && !ecu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: equipmentchoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: equipmentchoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ecu.mutation.EquipmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.RemovedEquipmentIDs(); len(nodes) > 0 && !ecu.mutation.EquipmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentchoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// EquipmentChoiceUpdateOne is the builder for updating a single EquipmentChoice entity.
type EquipmentChoiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentChoiceMutation
}

// SetChoose sets the "choose" field.
func (ecuo *EquipmentChoiceUpdateOne) SetChoose(i int) *EquipmentChoiceUpdateOne {
	ecuo.mutation.ResetChoose()
	ecuo.mutation.SetChoose(i)
	return ecuo
}

// AddChoose adds i to the "choose" field.
func (ecuo *EquipmentChoiceUpdateOne) AddChoose(i int) *EquipmentChoiceUpdateOne {
	ecuo.mutation.AddChoose(i)
	return ecuo
}

// SetDesc sets the "desc" field.
func (ecuo *EquipmentChoiceUpdateOne) SetDesc(s string) *EquipmentChoiceUpdateOne {
	ecuo.mutation.SetDesc(s)
	return ecuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (ecuo *EquipmentChoiceUpdateOne) SetNillableDesc(s *string) *EquipmentChoiceUpdateOne {
	if s != nil {
		ecuo.SetDesc(*s)
	}
	return ecuo
}

// ClearDesc clears the value of the "desc" field.
func (ecuo *EquipmentChoiceUpdateOne) ClearDesc() *EquipmentChoiceUpdateOne {
	ecuo.mutation.ClearDesc()
	return ecuo
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (ecuo *EquipmentChoiceUpdateOne) AddClasIDs(ids ...int) *EquipmentChoiceUpdateOne {
	ecuo.mutation.AddClasIDs(ids...)
	return ecuo
}

// AddClass adds the "class" edges to the Class entity.
func (ecuo *EquipmentChoiceUpdateOne) AddClass(c ...*Class) *EquipmentChoiceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ecuo.AddClasIDs(ids...)
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (ecuo *EquipmentChoiceUpdateOne) AddEquipmentIDs(ids ...int) *EquipmentChoiceUpdateOne {
	ecuo.mutation.AddEquipmentIDs(ids...)
	return ecuo
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (ecuo *EquipmentChoiceUpdateOne) AddEquipment(e ...*Equipment) *EquipmentChoiceUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.AddEquipmentIDs(ids...)
}

// Mutation returns the EquipmentChoiceMutation object of the builder.
func (ecuo *EquipmentChoiceUpdateOne) Mutation() *EquipmentChoiceMutation {
	return ecuo.mutation
}

// ClearClass clears all "class" edges to the Class entity.
func (ecuo *EquipmentChoiceUpdateOne) ClearClass() *EquipmentChoiceUpdateOne {
	ecuo.mutation.ClearClass()
	return ecuo
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (ecuo *EquipmentChoiceUpdateOne) RemoveClasIDs(ids ...int) *EquipmentChoiceUpdateOne {
	ecuo.mutation.RemoveClasIDs(ids...)
	return ecuo
}

// RemoveClass removes "class" edges to Class entities.
func (ecuo *EquipmentChoiceUpdateOne) RemoveClass(c ...*Class) *EquipmentChoiceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ecuo.RemoveClasIDs(ids...)
}

// ClearEquipment clears all "equipment" edges to the Equipment entity.
func (ecuo *EquipmentChoiceUpdateOne) ClearEquipment() *EquipmentChoiceUpdateOne {
	ecuo.mutation.ClearEquipment()
	return ecuo
}

// RemoveEquipmentIDs removes the "equipment" edge to Equipment entities by IDs.
func (ecuo *EquipmentChoiceUpdateOne) RemoveEquipmentIDs(ids ...int) *EquipmentChoiceUpdateOne {
	ecuo.mutation.RemoveEquipmentIDs(ids...)
	return ecuo
}

// RemoveEquipment removes "equipment" edges to Equipment entities.
func (ecuo *EquipmentChoiceUpdateOne) RemoveEquipment(e ...*Equipment) *EquipmentChoiceUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ecuo.RemoveEquipmentIDs(ids...)
}

// Where appends a list predicates to the EquipmentChoiceUpdate builder.
func (ecuo *EquipmentChoiceUpdateOne) Where(ps ...predicate.EquipmentChoice) *EquipmentChoiceUpdateOne {
	ecuo.mutation.Where(ps...)
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EquipmentChoiceUpdateOne) Select(field string, fields ...string) *EquipmentChoiceUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EquipmentChoice entity.
func (ecuo *EquipmentChoiceUpdateOne) Save(ctx context.Context) (*EquipmentChoice, error) {
	return withHooks(ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EquipmentChoiceUpdateOne) SaveX(ctx context.Context) *EquipmentChoice {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EquipmentChoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EquipmentChoiceUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ecuo *EquipmentChoiceUpdateOne) sqlSave(ctx context.Context) (_node *EquipmentChoice, err error) {
	_spec := sqlgraph.NewUpdateSpec(equipmentchoice.Table, equipmentchoice.Columns, sqlgraph.NewFieldSpec(equipmentchoice.FieldID, field.TypeInt))
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EquipmentChoice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentchoice.FieldID)
		for _, f := range fields {
			if !equipmentchoice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipmentchoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Choose(); ok {
		_spec.SetField(equipmentchoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := ecuo.mutation.AddedChoose(); ok {
		_spec.AddField(equipmentchoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := ecuo.mutation.Desc(); ok {
		_spec.SetField(equipmentchoice.FieldDesc, field.TypeString, value)
	}
	if ecuo.mutation.DescCleared() {
		_spec.ClearField(equipmentchoice.FieldDesc, field.TypeString)
	}
	if ecuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: equipmentchoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedClassIDs(); len(nodes) > 0 && !ecuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: equipmentchoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmentchoice.ClassTable,
			Columns: equipmentchoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ecuo.mutation.EquipmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.RemovedEquipmentIDs(); len(nodes) > 0 && !ecuo.mutation.EquipmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EquipmentChoice{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentchoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}
