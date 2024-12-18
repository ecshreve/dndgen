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
	"github.com/ecshreve/dndgen/ent/equipmententry"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// EquipmentEntryUpdate is the builder for updating EquipmentEntry entities.
type EquipmentEntryUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentEntryMutation
}

// Where appends a list predicates to the EquipmentEntryUpdate builder.
func (eeu *EquipmentEntryUpdate) Where(ps ...predicate.EquipmentEntry) *EquipmentEntryUpdate {
	eeu.mutation.Where(ps...)
	return eeu
}

// SetQuantity sets the "quantity" field.
func (eeu *EquipmentEntryUpdate) SetQuantity(i int) *EquipmentEntryUpdate {
	eeu.mutation.ResetQuantity()
	eeu.mutation.SetQuantity(i)
	return eeu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (eeu *EquipmentEntryUpdate) SetNillableQuantity(i *int) *EquipmentEntryUpdate {
	if i != nil {
		eeu.SetQuantity(*i)
	}
	return eeu
}

// AddQuantity adds i to the "quantity" field.
func (eeu *EquipmentEntryUpdate) AddQuantity(i int) *EquipmentEntryUpdate {
	eeu.mutation.AddQuantity(i)
	return eeu
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (eeu *EquipmentEntryUpdate) AddClasIDs(ids ...int) *EquipmentEntryUpdate {
	eeu.mutation.AddClasIDs(ids...)
	return eeu
}

// AddClass adds the "class" edges to the Class entity.
func (eeu *EquipmentEntryUpdate) AddClass(c ...*Class) *EquipmentEntryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eeu.AddClasIDs(ids...)
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eeu *EquipmentEntryUpdate) SetEquipmentID(id int) *EquipmentEntryUpdate {
	eeu.mutation.SetEquipmentID(id)
	return eeu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eeu *EquipmentEntryUpdate) SetEquipment(e *Equipment) *EquipmentEntryUpdate {
	return eeu.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentEntryMutation object of the builder.
func (eeu *EquipmentEntryUpdate) Mutation() *EquipmentEntryMutation {
	return eeu.mutation
}

// ClearClass clears all "class" edges to the Class entity.
func (eeu *EquipmentEntryUpdate) ClearClass() *EquipmentEntryUpdate {
	eeu.mutation.ClearClass()
	return eeu
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (eeu *EquipmentEntryUpdate) RemoveClasIDs(ids ...int) *EquipmentEntryUpdate {
	eeu.mutation.RemoveClasIDs(ids...)
	return eeu
}

// RemoveClass removes "class" edges to Class entities.
func (eeu *EquipmentEntryUpdate) RemoveClass(c ...*Class) *EquipmentEntryUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eeu.RemoveClasIDs(ids...)
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eeu *EquipmentEntryUpdate) ClearEquipment() *EquipmentEntryUpdate {
	eeu.mutation.ClearEquipment()
	return eeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eeu *EquipmentEntryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eeu.sqlSave, eeu.mutation, eeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eeu *EquipmentEntryUpdate) SaveX(ctx context.Context) int {
	affected, err := eeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eeu *EquipmentEntryUpdate) Exec(ctx context.Context) error {
	_, err := eeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eeu *EquipmentEntryUpdate) ExecX(ctx context.Context) {
	if err := eeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eeu *EquipmentEntryUpdate) check() error {
	if v, ok := eeu.mutation.Quantity(); ok {
		if err := equipmententry.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "EquipmentEntry.quantity": %w`, err)}
		}
	}
	if eeu.mutation.EquipmentCleared() && len(eeu.mutation.EquipmentIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EquipmentEntry.equipment"`)
	}
	return nil
}

func (eeu *EquipmentEntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eeu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(equipmententry.Table, equipmententry.Columns, sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt))
	if ps := eeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eeu.mutation.Quantity(); ok {
		_spec.SetField(equipmententry.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := eeu.mutation.AddedQuantity(); ok {
		_spec.AddField(equipmententry.FieldQuantity, field.TypeInt, value)
	}
	if eeu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeu.mutation.RemovedClassIDs(); len(nodes) > 0 && !eeu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
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
	if nodes := eeu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
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
	if eeu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmententry.EquipmentTable,
			Columns: []string{equipmententry.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmententry.EquipmentTable,
			Columns: []string{equipmententry.EquipmentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, eeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmententry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eeu.mutation.done = true
	return n, nil
}

// EquipmentEntryUpdateOne is the builder for updating a single EquipmentEntry entity.
type EquipmentEntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentEntryMutation
}

// SetQuantity sets the "quantity" field.
func (eeuo *EquipmentEntryUpdateOne) SetQuantity(i int) *EquipmentEntryUpdateOne {
	eeuo.mutation.ResetQuantity()
	eeuo.mutation.SetQuantity(i)
	return eeuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (eeuo *EquipmentEntryUpdateOne) SetNillableQuantity(i *int) *EquipmentEntryUpdateOne {
	if i != nil {
		eeuo.SetQuantity(*i)
	}
	return eeuo
}

// AddQuantity adds i to the "quantity" field.
func (eeuo *EquipmentEntryUpdateOne) AddQuantity(i int) *EquipmentEntryUpdateOne {
	eeuo.mutation.AddQuantity(i)
	return eeuo
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (eeuo *EquipmentEntryUpdateOne) AddClasIDs(ids ...int) *EquipmentEntryUpdateOne {
	eeuo.mutation.AddClasIDs(ids...)
	return eeuo
}

// AddClass adds the "class" edges to the Class entity.
func (eeuo *EquipmentEntryUpdateOne) AddClass(c ...*Class) *EquipmentEntryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eeuo.AddClasIDs(ids...)
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eeuo *EquipmentEntryUpdateOne) SetEquipmentID(id int) *EquipmentEntryUpdateOne {
	eeuo.mutation.SetEquipmentID(id)
	return eeuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eeuo *EquipmentEntryUpdateOne) SetEquipment(e *Equipment) *EquipmentEntryUpdateOne {
	return eeuo.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentEntryMutation object of the builder.
func (eeuo *EquipmentEntryUpdateOne) Mutation() *EquipmentEntryMutation {
	return eeuo.mutation
}

// ClearClass clears all "class" edges to the Class entity.
func (eeuo *EquipmentEntryUpdateOne) ClearClass() *EquipmentEntryUpdateOne {
	eeuo.mutation.ClearClass()
	return eeuo
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (eeuo *EquipmentEntryUpdateOne) RemoveClasIDs(ids ...int) *EquipmentEntryUpdateOne {
	eeuo.mutation.RemoveClasIDs(ids...)
	return eeuo
}

// RemoveClass removes "class" edges to Class entities.
func (eeuo *EquipmentEntryUpdateOne) RemoveClass(c ...*Class) *EquipmentEntryUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eeuo.RemoveClasIDs(ids...)
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eeuo *EquipmentEntryUpdateOne) ClearEquipment() *EquipmentEntryUpdateOne {
	eeuo.mutation.ClearEquipment()
	return eeuo
}

// Where appends a list predicates to the EquipmentEntryUpdate builder.
func (eeuo *EquipmentEntryUpdateOne) Where(ps ...predicate.EquipmentEntry) *EquipmentEntryUpdateOne {
	eeuo.mutation.Where(ps...)
	return eeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eeuo *EquipmentEntryUpdateOne) Select(field string, fields ...string) *EquipmentEntryUpdateOne {
	eeuo.fields = append([]string{field}, fields...)
	return eeuo
}

// Save executes the query and returns the updated EquipmentEntry entity.
func (eeuo *EquipmentEntryUpdateOne) Save(ctx context.Context) (*EquipmentEntry, error) {
	return withHooks(ctx, eeuo.sqlSave, eeuo.mutation, eeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eeuo *EquipmentEntryUpdateOne) SaveX(ctx context.Context) *EquipmentEntry {
	node, err := eeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eeuo *EquipmentEntryUpdateOne) Exec(ctx context.Context) error {
	_, err := eeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eeuo *EquipmentEntryUpdateOne) ExecX(ctx context.Context) {
	if err := eeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eeuo *EquipmentEntryUpdateOne) check() error {
	if v, ok := eeuo.mutation.Quantity(); ok {
		if err := equipmententry.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "EquipmentEntry.quantity": %w`, err)}
		}
	}
	if eeuo.mutation.EquipmentCleared() && len(eeuo.mutation.EquipmentIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "EquipmentEntry.equipment"`)
	}
	return nil
}

func (eeuo *EquipmentEntryUpdateOne) sqlSave(ctx context.Context) (_node *EquipmentEntry, err error) {
	if err := eeuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(equipmententry.Table, equipmententry.Columns, sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt))
	id, ok := eeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EquipmentEntry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmententry.FieldID)
		for _, f := range fields {
			if !equipmententry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipmententry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eeuo.mutation.Quantity(); ok {
		_spec.SetField(equipmententry.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := eeuo.mutation.AddedQuantity(); ok {
		_spec.AddField(equipmententry.FieldQuantity, field.TypeInt, value)
	}
	if eeuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeuo.mutation.RemovedClassIDs(); len(nodes) > 0 && !eeuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
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
	if nodes := eeuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   equipmententry.ClassTable,
			Columns: equipmententry.ClassPrimaryKey,
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
	if eeuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmententry.EquipmentTable,
			Columns: []string{equipmententry.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eeuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   equipmententry.EquipmentTable,
			Columns: []string{equipmententry.EquipmentColumn},
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
	_node = &EquipmentEntry{config: eeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmententry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eeuo.mutation.done = true
	return _node, nil
}
