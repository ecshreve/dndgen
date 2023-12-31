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
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/ent/weapondamage"
	"github.com/ecshreve/dndgen/ent/weaponproperty"
)

// WeaponUpdate is the builder for updating Weapon entities.
type WeaponUpdate struct {
	config
	hooks    []Hook
	mutation *WeaponMutation
}

// Where appends a list predicates to the WeaponUpdate builder.
func (wu *WeaponUpdate) Where(ps ...predicate.Weapon) *WeaponUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetIndx sets the "indx" field.
func (wu *WeaponUpdate) SetIndx(s string) *WeaponUpdate {
	wu.mutation.SetIndx(s)
	return wu
}

// SetName sets the "name" field.
func (wu *WeaponUpdate) SetName(s string) *WeaponUpdate {
	wu.mutation.SetName(s)
	return wu
}

// SetEquipmentID sets the "equipment_id" field.
func (wu *WeaponUpdate) SetEquipmentID(i int) *WeaponUpdate {
	wu.mutation.SetEquipmentID(i)
	return wu
}

// SetWeaponCategory sets the "weapon_category" field.
func (wu *WeaponUpdate) SetWeaponCategory(s string) *WeaponUpdate {
	wu.mutation.SetWeaponCategory(s)
	return wu
}

// SetWeaponRange sets the "weapon_range" field.
func (wu *WeaponUpdate) SetWeaponRange(s string) *WeaponUpdate {
	wu.mutation.SetWeaponRange(s)
	return wu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (wu *WeaponUpdate) SetEquipment(e *Equipment) *WeaponUpdate {
	return wu.SetEquipmentID(e.ID)
}

// AddWeaponDamageIDs adds the "weapon_damage" edge to the WeaponDamage entity by IDs.
func (wu *WeaponUpdate) AddWeaponDamageIDs(ids ...int) *WeaponUpdate {
	wu.mutation.AddWeaponDamageIDs(ids...)
	return wu
}

// AddWeaponDamage adds the "weapon_damage" edges to the WeaponDamage entity.
func (wu *WeaponUpdate) AddWeaponDamage(w ...*WeaponDamage) *WeaponUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWeaponDamageIDs(ids...)
}

// AddWeaponPropertyIDs adds the "weapon_properties" edge to the WeaponProperty entity by IDs.
func (wu *WeaponUpdate) AddWeaponPropertyIDs(ids ...int) *WeaponUpdate {
	wu.mutation.AddWeaponPropertyIDs(ids...)
	return wu
}

// AddWeaponProperties adds the "weapon_properties" edges to the WeaponProperty entity.
func (wu *WeaponUpdate) AddWeaponProperties(w ...*WeaponProperty) *WeaponUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.AddWeaponPropertyIDs(ids...)
}

// Mutation returns the WeaponMutation object of the builder.
func (wu *WeaponUpdate) Mutation() *WeaponMutation {
	return wu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (wu *WeaponUpdate) ClearEquipment() *WeaponUpdate {
	wu.mutation.ClearEquipment()
	return wu
}

// ClearWeaponDamage clears all "weapon_damage" edges to the WeaponDamage entity.
func (wu *WeaponUpdate) ClearWeaponDamage() *WeaponUpdate {
	wu.mutation.ClearWeaponDamage()
	return wu
}

// RemoveWeaponDamageIDs removes the "weapon_damage" edge to WeaponDamage entities by IDs.
func (wu *WeaponUpdate) RemoveWeaponDamageIDs(ids ...int) *WeaponUpdate {
	wu.mutation.RemoveWeaponDamageIDs(ids...)
	return wu
}

// RemoveWeaponDamage removes "weapon_damage" edges to WeaponDamage entities.
func (wu *WeaponUpdate) RemoveWeaponDamage(w ...*WeaponDamage) *WeaponUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWeaponDamageIDs(ids...)
}

// ClearWeaponProperties clears all "weapon_properties" edges to the WeaponProperty entity.
func (wu *WeaponUpdate) ClearWeaponProperties() *WeaponUpdate {
	wu.mutation.ClearWeaponProperties()
	return wu
}

// RemoveWeaponPropertyIDs removes the "weapon_properties" edge to WeaponProperty entities by IDs.
func (wu *WeaponUpdate) RemoveWeaponPropertyIDs(ids ...int) *WeaponUpdate {
	wu.mutation.RemoveWeaponPropertyIDs(ids...)
	return wu
}

// RemoveWeaponProperties removes "weapon_properties" edges to WeaponProperty entities.
func (wu *WeaponUpdate) RemoveWeaponProperties(w ...*WeaponProperty) *WeaponUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wu.RemoveWeaponPropertyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WeaponUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WeaponUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WeaponUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WeaponUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WeaponUpdate) check() error {
	if v, ok := wu.mutation.Indx(); ok {
		if err := weapon.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Weapon.indx": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Name(); ok {
		if err := weapon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Weapon.name": %w`, err)}
		}
	}
	if _, ok := wu.mutation.EquipmentID(); wu.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Weapon.equipment"`)
	}
	return nil
}

func (wu *WeaponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(weapon.Table, weapon.Columns, sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Indx(); ok {
		_spec.SetField(weapon.FieldIndx, field.TypeString, value)
	}
	if value, ok := wu.mutation.Name(); ok {
		_spec.SetField(weapon.FieldName, field.TypeString, value)
	}
	if value, ok := wu.mutation.WeaponCategory(); ok {
		_spec.SetField(weapon.FieldWeaponCategory, field.TypeString, value)
	}
	if value, ok := wu.mutation.WeaponRange(); ok {
		_spec.SetField(weapon.FieldWeaponRange, field.TypeString, value)
	}
	if wu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   weapon.EquipmentTable,
			Columns: []string{weapon.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wu.mutation.WeaponDamageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weapon.WeaponDamageTable,
			Columns: []string{weapon.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWeaponDamageIDs(); len(nodes) > 0 && !wu.mutation.WeaponDamageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weapon.WeaponDamageTable,
			Columns: []string{weapon.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WeaponDamageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weapon.WeaponDamageTable,
			Columns: []string{weapon.WeaponDamageColumn},
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
	if wu.mutation.WeaponPropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.WeaponPropertiesTable,
			Columns: weapon.WeaponPropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedWeaponPropertiesIDs(); len(nodes) > 0 && !wu.mutation.WeaponPropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.WeaponPropertiesTable,
			Columns: weapon.WeaponPropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.WeaponPropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.WeaponPropertiesTable,
			Columns: weapon.WeaponPropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{weapon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WeaponUpdateOne is the builder for updating a single Weapon entity.
type WeaponUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WeaponMutation
}

// SetIndx sets the "indx" field.
func (wuo *WeaponUpdateOne) SetIndx(s string) *WeaponUpdateOne {
	wuo.mutation.SetIndx(s)
	return wuo
}

// SetName sets the "name" field.
func (wuo *WeaponUpdateOne) SetName(s string) *WeaponUpdateOne {
	wuo.mutation.SetName(s)
	return wuo
}

// SetEquipmentID sets the "equipment_id" field.
func (wuo *WeaponUpdateOne) SetEquipmentID(i int) *WeaponUpdateOne {
	wuo.mutation.SetEquipmentID(i)
	return wuo
}

// SetWeaponCategory sets the "weapon_category" field.
func (wuo *WeaponUpdateOne) SetWeaponCategory(s string) *WeaponUpdateOne {
	wuo.mutation.SetWeaponCategory(s)
	return wuo
}

// SetWeaponRange sets the "weapon_range" field.
func (wuo *WeaponUpdateOne) SetWeaponRange(s string) *WeaponUpdateOne {
	wuo.mutation.SetWeaponRange(s)
	return wuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (wuo *WeaponUpdateOne) SetEquipment(e *Equipment) *WeaponUpdateOne {
	return wuo.SetEquipmentID(e.ID)
}

// AddWeaponDamageIDs adds the "weapon_damage" edge to the WeaponDamage entity by IDs.
func (wuo *WeaponUpdateOne) AddWeaponDamageIDs(ids ...int) *WeaponUpdateOne {
	wuo.mutation.AddWeaponDamageIDs(ids...)
	return wuo
}

// AddWeaponDamage adds the "weapon_damage" edges to the WeaponDamage entity.
func (wuo *WeaponUpdateOne) AddWeaponDamage(w ...*WeaponDamage) *WeaponUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWeaponDamageIDs(ids...)
}

// AddWeaponPropertyIDs adds the "weapon_properties" edge to the WeaponProperty entity by IDs.
func (wuo *WeaponUpdateOne) AddWeaponPropertyIDs(ids ...int) *WeaponUpdateOne {
	wuo.mutation.AddWeaponPropertyIDs(ids...)
	return wuo
}

// AddWeaponProperties adds the "weapon_properties" edges to the WeaponProperty entity.
func (wuo *WeaponUpdateOne) AddWeaponProperties(w ...*WeaponProperty) *WeaponUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.AddWeaponPropertyIDs(ids...)
}

// Mutation returns the WeaponMutation object of the builder.
func (wuo *WeaponUpdateOne) Mutation() *WeaponMutation {
	return wuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (wuo *WeaponUpdateOne) ClearEquipment() *WeaponUpdateOne {
	wuo.mutation.ClearEquipment()
	return wuo
}

// ClearWeaponDamage clears all "weapon_damage" edges to the WeaponDamage entity.
func (wuo *WeaponUpdateOne) ClearWeaponDamage() *WeaponUpdateOne {
	wuo.mutation.ClearWeaponDamage()
	return wuo
}

// RemoveWeaponDamageIDs removes the "weapon_damage" edge to WeaponDamage entities by IDs.
func (wuo *WeaponUpdateOne) RemoveWeaponDamageIDs(ids ...int) *WeaponUpdateOne {
	wuo.mutation.RemoveWeaponDamageIDs(ids...)
	return wuo
}

// RemoveWeaponDamage removes "weapon_damage" edges to WeaponDamage entities.
func (wuo *WeaponUpdateOne) RemoveWeaponDamage(w ...*WeaponDamage) *WeaponUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWeaponDamageIDs(ids...)
}

// ClearWeaponProperties clears all "weapon_properties" edges to the WeaponProperty entity.
func (wuo *WeaponUpdateOne) ClearWeaponProperties() *WeaponUpdateOne {
	wuo.mutation.ClearWeaponProperties()
	return wuo
}

// RemoveWeaponPropertyIDs removes the "weapon_properties" edge to WeaponProperty entities by IDs.
func (wuo *WeaponUpdateOne) RemoveWeaponPropertyIDs(ids ...int) *WeaponUpdateOne {
	wuo.mutation.RemoveWeaponPropertyIDs(ids...)
	return wuo
}

// RemoveWeaponProperties removes "weapon_properties" edges to WeaponProperty entities.
func (wuo *WeaponUpdateOne) RemoveWeaponProperties(w ...*WeaponProperty) *WeaponUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wuo.RemoveWeaponPropertyIDs(ids...)
}

// Where appends a list predicates to the WeaponUpdate builder.
func (wuo *WeaponUpdateOne) Where(ps ...predicate.Weapon) *WeaponUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WeaponUpdateOne) Select(field string, fields ...string) *WeaponUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Weapon entity.
func (wuo *WeaponUpdateOne) Save(ctx context.Context) (*Weapon, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WeaponUpdateOne) SaveX(ctx context.Context) *Weapon {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WeaponUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WeaponUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WeaponUpdateOne) check() error {
	if v, ok := wuo.mutation.Indx(); ok {
		if err := weapon.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Weapon.indx": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Name(); ok {
		if err := weapon.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Weapon.name": %w`, err)}
		}
	}
	if _, ok := wuo.mutation.EquipmentID(); wuo.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Weapon.equipment"`)
	}
	return nil
}

func (wuo *WeaponUpdateOne) sqlSave(ctx context.Context) (_node *Weapon, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(weapon.Table, weapon.Columns, sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Weapon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, weapon.FieldID)
		for _, f := range fields {
			if !weapon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != weapon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Indx(); ok {
		_spec.SetField(weapon.FieldIndx, field.TypeString, value)
	}
	if value, ok := wuo.mutation.Name(); ok {
		_spec.SetField(weapon.FieldName, field.TypeString, value)
	}
	if value, ok := wuo.mutation.WeaponCategory(); ok {
		_spec.SetField(weapon.FieldWeaponCategory, field.TypeString, value)
	}
	if value, ok := wuo.mutation.WeaponRange(); ok {
		_spec.SetField(weapon.FieldWeaponRange, field.TypeString, value)
	}
	if wuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   weapon.EquipmentTable,
			Columns: []string{weapon.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wuo.mutation.WeaponDamageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weapon.WeaponDamageTable,
			Columns: []string{weapon.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWeaponDamageIDs(); len(nodes) > 0 && !wuo.mutation.WeaponDamageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weapon.WeaponDamageTable,
			Columns: []string{weapon.WeaponDamageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapondamage.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WeaponDamageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   weapon.WeaponDamageTable,
			Columns: []string{weapon.WeaponDamageColumn},
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
	if wuo.mutation.WeaponPropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.WeaponPropertiesTable,
			Columns: weapon.WeaponPropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedWeaponPropertiesIDs(); len(nodes) > 0 && !wuo.mutation.WeaponPropertiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.WeaponPropertiesTable,
			Columns: weapon.WeaponPropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.WeaponPropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   weapon.WeaponPropertiesTable,
			Columns: weapon.WeaponPropertiesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weaponproperty.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Weapon{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{weapon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
