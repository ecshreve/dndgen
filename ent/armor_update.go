// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ArmorUpdate is the builder for updating Armor entities.
type ArmorUpdate struct {
	config
	hooks    []Hook
	mutation *ArmorMutation
}

// Where appends a list predicates to the ArmorUpdate builder.
func (au *ArmorUpdate) Where(ps ...predicate.Armor) *ArmorUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetStealthDisadvantage sets the "stealth_disadvantage" field.
func (au *ArmorUpdate) SetStealthDisadvantage(b bool) *ArmorUpdate {
	au.mutation.SetStealthDisadvantage(b)
	return au
}

// SetArmorClass sets the "armor_class" field.
func (au *ArmorUpdate) SetArmorClass(s string) *ArmorUpdate {
	au.mutation.SetArmorClass(s)
	return au
}

// SetMinStrength sets the "min_strength" field.
func (au *ArmorUpdate) SetMinStrength(i int) *ArmorUpdate {
	au.mutation.ResetMinStrength()
	au.mutation.SetMinStrength(i)
	return au
}

// AddMinStrength adds i to the "min_strength" field.
func (au *ArmorUpdate) AddMinStrength(i int) *ArmorUpdate {
	au.mutation.AddMinStrength(i)
	return au
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (au *ArmorUpdate) AddEquipmentIDs(ids ...int) *ArmorUpdate {
	au.mutation.AddEquipmentIDs(ids...)
	return au
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (au *ArmorUpdate) AddEquipment(e ...*Equipment) *ArmorUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddEquipmentIDs(ids...)
}

// Mutation returns the ArmorMutation object of the builder.
func (au *ArmorUpdate) Mutation() *ArmorMutation {
	return au.mutation
}

// ClearEquipment clears all "equipment" edges to the Equipment entity.
func (au *ArmorUpdate) ClearEquipment() *ArmorUpdate {
	au.mutation.ClearEquipment()
	return au
}

// RemoveEquipmentIDs removes the "equipment" edge to Equipment entities by IDs.
func (au *ArmorUpdate) RemoveEquipmentIDs(ids ...int) *ArmorUpdate {
	au.mutation.RemoveEquipmentIDs(ids...)
	return au
}

// RemoveEquipment removes "equipment" edges to Equipment entities.
func (au *ArmorUpdate) RemoveEquipment(e ...*Equipment) *ArmorUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveEquipmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArmorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ArmorMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArmorUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArmorUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArmorUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ArmorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(armor.Table, armor.Columns, sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.StealthDisadvantage(); ok {
		_spec.SetField(armor.FieldStealthDisadvantage, field.TypeBool, value)
	}
	if value, ok := au.mutation.ArmorClass(); ok {
		_spec.SetField(armor.FieldArmorClass, field.TypeString, value)
	}
	if value, ok := au.mutation.MinStrength(); ok {
		_spec.SetField(armor.FieldMinStrength, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedMinStrength(); ok {
		_spec.AddField(armor.FieldMinStrength, field.TypeInt, value)
	}
	if au.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: armor.EquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedEquipmentIDs(); len(nodes) > 0 && !au.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: armor.EquipmentPrimaryKey,
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
	if nodes := au.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: armor.EquipmentPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{armor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArmorUpdateOne is the builder for updating a single Armor entity.
type ArmorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArmorMutation
}

// SetStealthDisadvantage sets the "stealth_disadvantage" field.
func (auo *ArmorUpdateOne) SetStealthDisadvantage(b bool) *ArmorUpdateOne {
	auo.mutation.SetStealthDisadvantage(b)
	return auo
}

// SetArmorClass sets the "armor_class" field.
func (auo *ArmorUpdateOne) SetArmorClass(s string) *ArmorUpdateOne {
	auo.mutation.SetArmorClass(s)
	return auo
}

// SetMinStrength sets the "min_strength" field.
func (auo *ArmorUpdateOne) SetMinStrength(i int) *ArmorUpdateOne {
	auo.mutation.ResetMinStrength()
	auo.mutation.SetMinStrength(i)
	return auo
}

// AddMinStrength adds i to the "min_strength" field.
func (auo *ArmorUpdateOne) AddMinStrength(i int) *ArmorUpdateOne {
	auo.mutation.AddMinStrength(i)
	return auo
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (auo *ArmorUpdateOne) AddEquipmentIDs(ids ...int) *ArmorUpdateOne {
	auo.mutation.AddEquipmentIDs(ids...)
	return auo
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (auo *ArmorUpdateOne) AddEquipment(e ...*Equipment) *ArmorUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddEquipmentIDs(ids...)
}

// Mutation returns the ArmorMutation object of the builder.
func (auo *ArmorUpdateOne) Mutation() *ArmorMutation {
	return auo.mutation
}

// ClearEquipment clears all "equipment" edges to the Equipment entity.
func (auo *ArmorUpdateOne) ClearEquipment() *ArmorUpdateOne {
	auo.mutation.ClearEquipment()
	return auo
}

// RemoveEquipmentIDs removes the "equipment" edge to Equipment entities by IDs.
func (auo *ArmorUpdateOne) RemoveEquipmentIDs(ids ...int) *ArmorUpdateOne {
	auo.mutation.RemoveEquipmentIDs(ids...)
	return auo
}

// RemoveEquipment removes "equipment" edges to Equipment entities.
func (auo *ArmorUpdateOne) RemoveEquipment(e ...*Equipment) *ArmorUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveEquipmentIDs(ids...)
}

// Where appends a list predicates to the ArmorUpdate builder.
func (auo *ArmorUpdateOne) Where(ps ...predicate.Armor) *ArmorUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArmorUpdateOne) Select(field string, fields ...string) *ArmorUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Armor entity.
func (auo *ArmorUpdateOne) Save(ctx context.Context) (*Armor, error) {
	return withHooks[*Armor, ArmorMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArmorUpdateOne) SaveX(ctx context.Context) *Armor {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArmorUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArmorUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ArmorUpdateOne) sqlSave(ctx context.Context) (_node *Armor, err error) {
	_spec := sqlgraph.NewUpdateSpec(armor.Table, armor.Columns, sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Armor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, armor.FieldID)
		for _, f := range fields {
			if !armor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != armor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.StealthDisadvantage(); ok {
		_spec.SetField(armor.FieldStealthDisadvantage, field.TypeBool, value)
	}
	if value, ok := auo.mutation.ArmorClass(); ok {
		_spec.SetField(armor.FieldArmorClass, field.TypeString, value)
	}
	if value, ok := auo.mutation.MinStrength(); ok {
		_spec.SetField(armor.FieldMinStrength, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedMinStrength(); ok {
		_spec.AddField(armor.FieldMinStrength, field.TypeInt, value)
	}
	if auo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: armor.EquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedEquipmentIDs(); len(nodes) > 0 && !auo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: armor.EquipmentPrimaryKey,
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
	if nodes := auo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   armor.EquipmentTable,
			Columns: armor.EquipmentPrimaryKey,
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
	_node = &Armor{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{armor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}