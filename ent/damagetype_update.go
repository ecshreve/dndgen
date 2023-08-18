// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/weapondamage"
)

// DamageTypeUpdate is the builder for updating DamageType entities.
type DamageTypeUpdate struct {
	config
	hooks    []Hook
	mutation *DamageTypeMutation
}

// Where appends a list predicates to the DamageTypeUpdate builder.
func (dtu *DamageTypeUpdate) Where(ps ...predicate.DamageType) *DamageTypeUpdate {
	dtu.mutation.Where(ps...)
	return dtu
}

// SetIndx sets the "indx" field.
func (dtu *DamageTypeUpdate) SetIndx(s string) *DamageTypeUpdate {
	dtu.mutation.SetIndx(s)
	return dtu
}

// SetName sets the "name" field.
func (dtu *DamageTypeUpdate) SetName(s string) *DamageTypeUpdate {
	dtu.mutation.SetName(s)
	return dtu
}

// SetDesc sets the "desc" field.
func (dtu *DamageTypeUpdate) SetDesc(s string) *DamageTypeUpdate {
	dtu.mutation.SetDesc(s)
	return dtu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (dtu *DamageTypeUpdate) SetNillableDesc(s *string) *DamageTypeUpdate {
	if s != nil {
		dtu.SetDesc(*s)
	}
	return dtu
}

// ClearDesc clears the value of the "desc" field.
func (dtu *DamageTypeUpdate) ClearDesc() *DamageTypeUpdate {
	dtu.mutation.ClearDesc()
	return dtu
}

// SetWeaponDamageID sets the "weapon_damage" edge to the WeaponDamage entity by ID.
func (dtu *DamageTypeUpdate) SetWeaponDamageID(id int) *DamageTypeUpdate {
	dtu.mutation.SetWeaponDamageID(id)
	return dtu
}

// SetNillableWeaponDamageID sets the "weapon_damage" edge to the WeaponDamage entity by ID if the given value is not nil.
func (dtu *DamageTypeUpdate) SetNillableWeaponDamageID(id *int) *DamageTypeUpdate {
	if id != nil {
		dtu = dtu.SetWeaponDamageID(*id)
	}
	return dtu
}

// SetWeaponDamage sets the "weapon_damage" edge to the WeaponDamage entity.
func (dtu *DamageTypeUpdate) SetWeaponDamage(w *WeaponDamage) *DamageTypeUpdate {
	return dtu.SetWeaponDamageID(w.ID)
}

// Mutation returns the DamageTypeMutation object of the builder.
func (dtu *DamageTypeUpdate) Mutation() *DamageTypeMutation {
	return dtu.mutation
}

// ClearWeaponDamage clears the "weapon_damage" edge to the WeaponDamage entity.
func (dtu *DamageTypeUpdate) ClearWeaponDamage() *DamageTypeUpdate {
	dtu.mutation.ClearWeaponDamage()
	return dtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dtu *DamageTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, DamageTypeMutation](ctx, dtu.sqlSave, dtu.mutation, dtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dtu *DamageTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := dtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dtu *DamageTypeUpdate) Exec(ctx context.Context) error {
	_, err := dtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtu *DamageTypeUpdate) ExecX(ctx context.Context) {
	if err := dtu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dtu *DamageTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(damagetype.Table, damagetype.Columns, sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt))
	if ps := dtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtu.mutation.Indx(); ok {
		_spec.SetField(damagetype.FieldIndx, field.TypeString, value)
	}
	if value, ok := dtu.mutation.Name(); ok {
		_spec.SetField(damagetype.FieldName, field.TypeString, value)
	}
	if value, ok := dtu.mutation.Desc(); ok {
		_spec.SetField(damagetype.FieldDesc, field.TypeString, value)
	}
	if dtu.mutation.DescCleared() {
		_spec.ClearField(damagetype.FieldDesc, field.TypeString)
	}
	if dtu.mutation.WeaponDamageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   damagetype.WeaponDamageTable,
			Columns: []string{damagetype.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtu.mutation.WeaponDamageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   damagetype.WeaponDamageTable,
			Columns: []string{damagetype.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{damagetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dtu.mutation.done = true
	return n, nil
}

// DamageTypeUpdateOne is the builder for updating a single DamageType entity.
type DamageTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DamageTypeMutation
}

// SetIndx sets the "indx" field.
func (dtuo *DamageTypeUpdateOne) SetIndx(s string) *DamageTypeUpdateOne {
	dtuo.mutation.SetIndx(s)
	return dtuo
}

// SetName sets the "name" field.
func (dtuo *DamageTypeUpdateOne) SetName(s string) *DamageTypeUpdateOne {
	dtuo.mutation.SetName(s)
	return dtuo
}

// SetDesc sets the "desc" field.
func (dtuo *DamageTypeUpdateOne) SetDesc(s string) *DamageTypeUpdateOne {
	dtuo.mutation.SetDesc(s)
	return dtuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (dtuo *DamageTypeUpdateOne) SetNillableDesc(s *string) *DamageTypeUpdateOne {
	if s != nil {
		dtuo.SetDesc(*s)
	}
	return dtuo
}

// ClearDesc clears the value of the "desc" field.
func (dtuo *DamageTypeUpdateOne) ClearDesc() *DamageTypeUpdateOne {
	dtuo.mutation.ClearDesc()
	return dtuo
}

// SetWeaponDamageID sets the "weapon_damage" edge to the WeaponDamage entity by ID.
func (dtuo *DamageTypeUpdateOne) SetWeaponDamageID(id int) *DamageTypeUpdateOne {
	dtuo.mutation.SetWeaponDamageID(id)
	return dtuo
}

// SetNillableWeaponDamageID sets the "weapon_damage" edge to the WeaponDamage entity by ID if the given value is not nil.
func (dtuo *DamageTypeUpdateOne) SetNillableWeaponDamageID(id *int) *DamageTypeUpdateOne {
	if id != nil {
		dtuo = dtuo.SetWeaponDamageID(*id)
	}
	return dtuo
}

// SetWeaponDamage sets the "weapon_damage" edge to the WeaponDamage entity.
func (dtuo *DamageTypeUpdateOne) SetWeaponDamage(w *WeaponDamage) *DamageTypeUpdateOne {
	return dtuo.SetWeaponDamageID(w.ID)
}

// Mutation returns the DamageTypeMutation object of the builder.
func (dtuo *DamageTypeUpdateOne) Mutation() *DamageTypeMutation {
	return dtuo.mutation
}

// ClearWeaponDamage clears the "weapon_damage" edge to the WeaponDamage entity.
func (dtuo *DamageTypeUpdateOne) ClearWeaponDamage() *DamageTypeUpdateOne {
	dtuo.mutation.ClearWeaponDamage()
	return dtuo
}

// Where appends a list predicates to the DamageTypeUpdate builder.
func (dtuo *DamageTypeUpdateOne) Where(ps ...predicate.DamageType) *DamageTypeUpdateOne {
	dtuo.mutation.Where(ps...)
	return dtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dtuo *DamageTypeUpdateOne) Select(field string, fields ...string) *DamageTypeUpdateOne {
	dtuo.fields = append([]string{field}, fields...)
	return dtuo
}

// Save executes the query and returns the updated DamageType entity.
func (dtuo *DamageTypeUpdateOne) Save(ctx context.Context) (*DamageType, error) {
	return withHooks[*DamageType, DamageTypeMutation](ctx, dtuo.sqlSave, dtuo.mutation, dtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dtuo *DamageTypeUpdateOne) SaveX(ctx context.Context) *DamageType {
	node, err := dtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dtuo *DamageTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := dtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtuo *DamageTypeUpdateOne) ExecX(ctx context.Context) {
	if err := dtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dtuo *DamageTypeUpdateOne) sqlSave(ctx context.Context) (_node *DamageType, err error) {
	_spec := sqlgraph.NewUpdateSpec(damagetype.Table, damagetype.Columns, sqlgraph.NewFieldSpec(damagetype.FieldID, field.TypeInt))
	id, ok := dtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DamageType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, damagetype.FieldID)
		for _, f := range fields {
			if !damagetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != damagetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dtuo.mutation.Indx(); ok {
		_spec.SetField(damagetype.FieldIndx, field.TypeString, value)
	}
	if value, ok := dtuo.mutation.Name(); ok {
		_spec.SetField(damagetype.FieldName, field.TypeString, value)
	}
	if value, ok := dtuo.mutation.Desc(); ok {
		_spec.SetField(damagetype.FieldDesc, field.TypeString, value)
	}
	if dtuo.mutation.DescCleared() {
		_spec.ClearField(damagetype.FieldDesc, field.TypeString)
	}
	if dtuo.mutation.WeaponDamageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   damagetype.WeaponDamageTable,
			Columns: []string{damagetype.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dtuo.mutation.WeaponDamageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   damagetype.WeaponDamageTable,
			Columns: []string{damagetype.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DamageType{config: dtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{damagetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dtuo.mutation.done = true
	return _node, nil
}
