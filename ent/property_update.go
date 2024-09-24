// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/property"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// PropertyUpdate is the builder for updating Property entities.
type PropertyUpdate struct {
	config
	hooks    []Hook
	mutation *PropertyMutation
}

// Where appends a list predicates to the PropertyUpdate builder.
func (pu *PropertyUpdate) Where(ps ...predicate.Property) *PropertyUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetIndx sets the "indx" field.
func (pu *PropertyUpdate) SetIndx(s string) *PropertyUpdate {
	pu.mutation.SetIndx(s)
	return pu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableIndx(s *string) *PropertyUpdate {
	if s != nil {
		pu.SetIndx(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PropertyUpdate) SetName(s string) *PropertyUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PropertyUpdate) SetNillableName(s *string) *PropertyUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDesc sets the "desc" field.
func (pu *PropertyUpdate) SetDesc(s []string) *PropertyUpdate {
	pu.mutation.SetDesc(s)
	return pu
}

// AppendDesc appends s to the "desc" field.
func (pu *PropertyUpdate) AppendDesc(s []string) *PropertyUpdate {
	pu.mutation.AppendDesc(s)
	return pu
}

// ClearDesc clears the value of the "desc" field.
func (pu *PropertyUpdate) ClearDesc() *PropertyUpdate {
	pu.mutation.ClearDesc()
	return pu
}

// AddWeaponIDs adds the "weapons" edge to the Weapon entity by IDs.
func (pu *PropertyUpdate) AddWeaponIDs(ids ...int) *PropertyUpdate {
	pu.mutation.AddWeaponIDs(ids...)
	return pu
}

// AddWeapons adds the "weapons" edges to the Weapon entity.
func (pu *PropertyUpdate) AddWeapons(w ...*Weapon) *PropertyUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pu.AddWeaponIDs(ids...)
}

// Mutation returns the PropertyMutation object of the builder.
func (pu *PropertyUpdate) Mutation() *PropertyMutation {
	return pu.mutation
}

// ClearWeapons clears all "weapons" edges to the Weapon entity.
func (pu *PropertyUpdate) ClearWeapons() *PropertyUpdate {
	pu.mutation.ClearWeapons()
	return pu
}

// RemoveWeaponIDs removes the "weapons" edge to Weapon entities by IDs.
func (pu *PropertyUpdate) RemoveWeaponIDs(ids ...int) *PropertyUpdate {
	pu.mutation.RemoveWeaponIDs(ids...)
	return pu
}

// RemoveWeapons removes "weapons" edges to Weapon entities.
func (pu *PropertyUpdate) RemoveWeapons(w ...*Weapon) *PropertyUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return pu.RemoveWeaponIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PropertyUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PropertyUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PropertyUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PropertyUpdate) check() error {
	if v, ok := pu.mutation.Indx(); ok {
		if err := property.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Property.indx": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := property.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Property.name": %w`, err)}
		}
	}
	return nil
}

func (pu *PropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(property.Table, property.Columns, sqlgraph.NewFieldSpec(property.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Indx(); ok {
		_spec.SetField(property.FieldIndx, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(property.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Desc(); ok {
		_spec.SetField(property.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := pu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, property.FieldDesc, value)
		})
	}
	if pu.mutation.DescCleared() {
		_spec.ClearField(property.FieldDesc, field.TypeJSON)
	}
	if pu.mutation.WeaponsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   property.WeaponsTable,
			Columns: property.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedWeaponsIDs(); len(nodes) > 0 && !pu.mutation.WeaponsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   property.WeaponsTable,
			Columns: property.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.WeaponsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   property.WeaponsTable,
			Columns: property.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PropertyUpdateOne is the builder for updating a single Property entity.
type PropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PropertyMutation
}

// SetIndx sets the "indx" field.
func (puo *PropertyUpdateOne) SetIndx(s string) *PropertyUpdateOne {
	puo.mutation.SetIndx(s)
	return puo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableIndx(s *string) *PropertyUpdateOne {
	if s != nil {
		puo.SetIndx(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PropertyUpdateOne) SetName(s string) *PropertyUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PropertyUpdateOne) SetNillableName(s *string) *PropertyUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDesc sets the "desc" field.
func (puo *PropertyUpdateOne) SetDesc(s []string) *PropertyUpdateOne {
	puo.mutation.SetDesc(s)
	return puo
}

// AppendDesc appends s to the "desc" field.
func (puo *PropertyUpdateOne) AppendDesc(s []string) *PropertyUpdateOne {
	puo.mutation.AppendDesc(s)
	return puo
}

// ClearDesc clears the value of the "desc" field.
func (puo *PropertyUpdateOne) ClearDesc() *PropertyUpdateOne {
	puo.mutation.ClearDesc()
	return puo
}

// AddWeaponIDs adds the "weapons" edge to the Weapon entity by IDs.
func (puo *PropertyUpdateOne) AddWeaponIDs(ids ...int) *PropertyUpdateOne {
	puo.mutation.AddWeaponIDs(ids...)
	return puo
}

// AddWeapons adds the "weapons" edges to the Weapon entity.
func (puo *PropertyUpdateOne) AddWeapons(w ...*Weapon) *PropertyUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return puo.AddWeaponIDs(ids...)
}

// Mutation returns the PropertyMutation object of the builder.
func (puo *PropertyUpdateOne) Mutation() *PropertyMutation {
	return puo.mutation
}

// ClearWeapons clears all "weapons" edges to the Weapon entity.
func (puo *PropertyUpdateOne) ClearWeapons() *PropertyUpdateOne {
	puo.mutation.ClearWeapons()
	return puo
}

// RemoveWeaponIDs removes the "weapons" edge to Weapon entities by IDs.
func (puo *PropertyUpdateOne) RemoveWeaponIDs(ids ...int) *PropertyUpdateOne {
	puo.mutation.RemoveWeaponIDs(ids...)
	return puo
}

// RemoveWeapons removes "weapons" edges to Weapon entities.
func (puo *PropertyUpdateOne) RemoveWeapons(w ...*Weapon) *PropertyUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return puo.RemoveWeaponIDs(ids...)
}

// Where appends a list predicates to the PropertyUpdate builder.
func (puo *PropertyUpdateOne) Where(ps ...predicate.Property) *PropertyUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PropertyUpdateOne) Select(field string, fields ...string) *PropertyUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Property entity.
func (puo *PropertyUpdateOne) Save(ctx context.Context) (*Property, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PropertyUpdateOne) SaveX(ctx context.Context) *Property {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PropertyUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PropertyUpdateOne) check() error {
	if v, ok := puo.mutation.Indx(); ok {
		if err := property.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Property.indx": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := property.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Property.name": %w`, err)}
		}
	}
	return nil
}

func (puo *PropertyUpdateOne) sqlSave(ctx context.Context) (_node *Property, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(property.Table, property.Columns, sqlgraph.NewFieldSpec(property.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Property.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, property.FieldID)
		for _, f := range fields {
			if !property.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != property.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Indx(); ok {
		_spec.SetField(property.FieldIndx, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(property.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Desc(); ok {
		_spec.SetField(property.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := puo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, property.FieldDesc, value)
		})
	}
	if puo.mutation.DescCleared() {
		_spec.ClearField(property.FieldDesc, field.TypeJSON)
	}
	if puo.mutation.WeaponsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   property.WeaponsTable,
			Columns: property.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedWeaponsIDs(); len(nodes) > 0 && !puo.mutation.WeaponsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   property.WeaponsTable,
			Columns: property.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.WeaponsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   property.WeaponsTable,
			Columns: property.WeaponsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Property{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{property.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}