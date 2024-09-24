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
	"github.com/ecshreve/dndgen/ent/cost"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/tool"
	"github.com/ecshreve/dndgen/ent/vehicle"
	"github.com/ecshreve/dndgen/ent/weapon"
)

// EquipmentUpdate is the builder for updating Equipment entities.
type EquipmentUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentMutation
}

// Where appends a list predicates to the EquipmentUpdate builder.
func (eu *EquipmentUpdate) Where(ps ...predicate.Equipment) *EquipmentUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetIndx sets the "indx" field.
func (eu *EquipmentUpdate) SetIndx(s string) *EquipmentUpdate {
	eu.mutation.SetIndx(s)
	return eu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableIndx(s *string) *EquipmentUpdate {
	if s != nil {
		eu.SetIndx(*s)
	}
	return eu
}

// SetName sets the "name" field.
func (eu *EquipmentUpdate) SetName(s string) *EquipmentUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableName(s *string) *EquipmentUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetEquipmentCategory sets the "equipment_category" field.
func (eu *EquipmentUpdate) SetEquipmentCategory(ec equipment.EquipmentCategory) *EquipmentUpdate {
	eu.mutation.SetEquipmentCategory(ec)
	return eu
}

// SetNillableEquipmentCategory sets the "equipment_category" field if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableEquipmentCategory(ec *equipment.EquipmentCategory) *EquipmentUpdate {
	if ec != nil {
		eu.SetEquipmentCategory(*ec)
	}
	return eu
}

// SetWeight sets the "weight" field.
func (eu *EquipmentUpdate) SetWeight(f float64) *EquipmentUpdate {
	eu.mutation.ResetWeight()
	eu.mutation.SetWeight(f)
	return eu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableWeight(f *float64) *EquipmentUpdate {
	if f != nil {
		eu.SetWeight(*f)
	}
	return eu
}

// AddWeight adds f to the "weight" field.
func (eu *EquipmentUpdate) AddWeight(f float64) *EquipmentUpdate {
	eu.mutation.AddWeight(f)
	return eu
}

// ClearWeight clears the value of the "weight" field.
func (eu *EquipmentUpdate) ClearWeight() *EquipmentUpdate {
	eu.mutation.ClearWeight()
	return eu
}

// SetCostID sets the "cost" edge to the Cost entity by ID.
func (eu *EquipmentUpdate) SetCostID(id int) *EquipmentUpdate {
	eu.mutation.SetCostID(id)
	return eu
}

// SetNillableCostID sets the "cost" edge to the Cost entity by ID if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableCostID(id *int) *EquipmentUpdate {
	if id != nil {
		eu = eu.SetCostID(*id)
	}
	return eu
}

// SetCost sets the "cost" edge to the Cost entity.
func (eu *EquipmentUpdate) SetCost(c *Cost) *EquipmentUpdate {
	return eu.SetCostID(c.ID)
}

// SetGearID sets the "gear" edge to the Gear entity by ID.
func (eu *EquipmentUpdate) SetGearID(id int) *EquipmentUpdate {
	eu.mutation.SetGearID(id)
	return eu
}

// SetNillableGearID sets the "gear" edge to the Gear entity by ID if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableGearID(id *int) *EquipmentUpdate {
	if id != nil {
		eu = eu.SetGearID(*id)
	}
	return eu
}

// SetGear sets the "gear" edge to the Gear entity.
func (eu *EquipmentUpdate) SetGear(g *Gear) *EquipmentUpdate {
	return eu.SetGearID(g.ID)
}

// SetToolID sets the "tool" edge to the Tool entity by ID.
func (eu *EquipmentUpdate) SetToolID(id int) *EquipmentUpdate {
	eu.mutation.SetToolID(id)
	return eu
}

// SetNillableToolID sets the "tool" edge to the Tool entity by ID if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableToolID(id *int) *EquipmentUpdate {
	if id != nil {
		eu = eu.SetToolID(*id)
	}
	return eu
}

// SetTool sets the "tool" edge to the Tool entity.
func (eu *EquipmentUpdate) SetTool(t *Tool) *EquipmentUpdate {
	return eu.SetToolID(t.ID)
}

// SetWeaponID sets the "weapon" edge to the Weapon entity by ID.
func (eu *EquipmentUpdate) SetWeaponID(id int) *EquipmentUpdate {
	eu.mutation.SetWeaponID(id)
	return eu
}

// SetNillableWeaponID sets the "weapon" edge to the Weapon entity by ID if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableWeaponID(id *int) *EquipmentUpdate {
	if id != nil {
		eu = eu.SetWeaponID(*id)
	}
	return eu
}

// SetWeapon sets the "weapon" edge to the Weapon entity.
func (eu *EquipmentUpdate) SetWeapon(w *Weapon) *EquipmentUpdate {
	return eu.SetWeaponID(w.ID)
}

// SetVehicleID sets the "vehicle" edge to the Vehicle entity by ID.
func (eu *EquipmentUpdate) SetVehicleID(id int) *EquipmentUpdate {
	eu.mutation.SetVehicleID(id)
	return eu
}

// SetNillableVehicleID sets the "vehicle" edge to the Vehicle entity by ID if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableVehicleID(id *int) *EquipmentUpdate {
	if id != nil {
		eu = eu.SetVehicleID(*id)
	}
	return eu
}

// SetVehicle sets the "vehicle" edge to the Vehicle entity.
func (eu *EquipmentUpdate) SetVehicle(v *Vehicle) *EquipmentUpdate {
	return eu.SetVehicleID(v.ID)
}

// SetArmorID sets the "armor" edge to the Armor entity by ID.
func (eu *EquipmentUpdate) SetArmorID(id int) *EquipmentUpdate {
	eu.mutation.SetArmorID(id)
	return eu
}

// SetNillableArmorID sets the "armor" edge to the Armor entity by ID if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableArmorID(id *int) *EquipmentUpdate {
	if id != nil {
		eu = eu.SetArmorID(*id)
	}
	return eu
}

// SetArmor sets the "armor" edge to the Armor entity.
func (eu *EquipmentUpdate) SetArmor(a *Armor) *EquipmentUpdate {
	return eu.SetArmorID(a.ID)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (eu *EquipmentUpdate) AddProficiencyIDs(ids ...int) *EquipmentUpdate {
	eu.mutation.AddProficiencyIDs(ids...)
	return eu
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (eu *EquipmentUpdate) AddProficiencies(p ...*Proficiency) *EquipmentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.AddProficiencyIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (eu *EquipmentUpdate) Mutation() *EquipmentMutation {
	return eu.mutation
}

// ClearCost clears the "cost" edge to the Cost entity.
func (eu *EquipmentUpdate) ClearCost() *EquipmentUpdate {
	eu.mutation.ClearCost()
	return eu
}

// ClearGear clears the "gear" edge to the Gear entity.
func (eu *EquipmentUpdate) ClearGear() *EquipmentUpdate {
	eu.mutation.ClearGear()
	return eu
}

// ClearTool clears the "tool" edge to the Tool entity.
func (eu *EquipmentUpdate) ClearTool() *EquipmentUpdate {
	eu.mutation.ClearTool()
	return eu
}

// ClearWeapon clears the "weapon" edge to the Weapon entity.
func (eu *EquipmentUpdate) ClearWeapon() *EquipmentUpdate {
	eu.mutation.ClearWeapon()
	return eu
}

// ClearVehicle clears the "vehicle" edge to the Vehicle entity.
func (eu *EquipmentUpdate) ClearVehicle() *EquipmentUpdate {
	eu.mutation.ClearVehicle()
	return eu
}

// ClearArmor clears the "armor" edge to the Armor entity.
func (eu *EquipmentUpdate) ClearArmor() *EquipmentUpdate {
	eu.mutation.ClearArmor()
	return eu
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (eu *EquipmentUpdate) ClearProficiencies() *EquipmentUpdate {
	eu.mutation.ClearProficiencies()
	return eu
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (eu *EquipmentUpdate) RemoveProficiencyIDs(ids ...int) *EquipmentUpdate {
	eu.mutation.RemoveProficiencyIDs(ids...)
	return eu
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (eu *EquipmentUpdate) RemoveProficiencies(p ...*Proficiency) *EquipmentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return eu.RemoveProficiencyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EquipmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EquipmentUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EquipmentUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EquipmentUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EquipmentUpdate) check() error {
	if v, ok := eu.mutation.Indx(); ok {
		if err := equipment.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Equipment.indx": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Name(); ok {
		if err := equipment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Equipment.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.EquipmentCategory(); ok {
		if err := equipment.EquipmentCategoryValidator(v); err != nil {
			return &ValidationError{Name: "equipment_category", err: fmt.Errorf(`ent: validator failed for field "Equipment.equipment_category": %w`, err)}
		}
	}
	return nil
}

func (eu *EquipmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(equipment.Table, equipment.Columns, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Indx(); ok {
		_spec.SetField(equipment.FieldIndx, field.TypeString, value)
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(equipment.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.EquipmentCategory(); ok {
		_spec.SetField(equipment.FieldEquipmentCategory, field.TypeEnum, value)
	}
	if value, ok := eu.mutation.Weight(); ok {
		_spec.SetField(equipment.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := eu.mutation.AddedWeight(); ok {
		_spec.AddField(equipment.FieldWeight, field.TypeFloat64, value)
	}
	if eu.mutation.WeightCleared() {
		_spec.ClearField(equipment.FieldWeight, field.TypeFloat64)
	}
	if eu.mutation.CostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.CostTable,
			Columns: []string{equipment.CostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.CostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.CostTable,
			Columns: []string{equipment.CostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.GearCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.GearTable,
			Columns: []string{equipment.GearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.GearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.GearTable,
			Columns: []string{equipment.GearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ToolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ToolTable,
			Columns: []string{equipment.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ToolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ToolTable,
			Columns: []string{equipment.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.WeaponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.WeaponTable,
			Columns: []string{equipment.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.WeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.WeaponTable,
			Columns: []string{equipment.WeaponColumn},
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
	if eu.mutation.VehicleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.VehicleTable,
			Columns: []string{equipment.VehicleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.VehicleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.VehicleTable,
			Columns: []string{equipment.VehicleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ArmorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ArmorTable,
			Columns: []string{equipment.ArmorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ArmorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ArmorTable,
			Columns: []string{equipment.ArmorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipment.ProficienciesTable,
			Columns: []string{equipment.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !eu.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipment.ProficienciesTable,
			Columns: []string{equipment.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipment.ProficienciesTable,
			Columns: []string{equipment.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EquipmentUpdateOne is the builder for updating a single Equipment entity.
type EquipmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentMutation
}

// SetIndx sets the "indx" field.
func (euo *EquipmentUpdateOne) SetIndx(s string) *EquipmentUpdateOne {
	euo.mutation.SetIndx(s)
	return euo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableIndx(s *string) *EquipmentUpdateOne {
	if s != nil {
		euo.SetIndx(*s)
	}
	return euo
}

// SetName sets the "name" field.
func (euo *EquipmentUpdateOne) SetName(s string) *EquipmentUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableName(s *string) *EquipmentUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetEquipmentCategory sets the "equipment_category" field.
func (euo *EquipmentUpdateOne) SetEquipmentCategory(ec equipment.EquipmentCategory) *EquipmentUpdateOne {
	euo.mutation.SetEquipmentCategory(ec)
	return euo
}

// SetNillableEquipmentCategory sets the "equipment_category" field if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableEquipmentCategory(ec *equipment.EquipmentCategory) *EquipmentUpdateOne {
	if ec != nil {
		euo.SetEquipmentCategory(*ec)
	}
	return euo
}

// SetWeight sets the "weight" field.
func (euo *EquipmentUpdateOne) SetWeight(f float64) *EquipmentUpdateOne {
	euo.mutation.ResetWeight()
	euo.mutation.SetWeight(f)
	return euo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableWeight(f *float64) *EquipmentUpdateOne {
	if f != nil {
		euo.SetWeight(*f)
	}
	return euo
}

// AddWeight adds f to the "weight" field.
func (euo *EquipmentUpdateOne) AddWeight(f float64) *EquipmentUpdateOne {
	euo.mutation.AddWeight(f)
	return euo
}

// ClearWeight clears the value of the "weight" field.
func (euo *EquipmentUpdateOne) ClearWeight() *EquipmentUpdateOne {
	euo.mutation.ClearWeight()
	return euo
}

// SetCostID sets the "cost" edge to the Cost entity by ID.
func (euo *EquipmentUpdateOne) SetCostID(id int) *EquipmentUpdateOne {
	euo.mutation.SetCostID(id)
	return euo
}

// SetNillableCostID sets the "cost" edge to the Cost entity by ID if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableCostID(id *int) *EquipmentUpdateOne {
	if id != nil {
		euo = euo.SetCostID(*id)
	}
	return euo
}

// SetCost sets the "cost" edge to the Cost entity.
func (euo *EquipmentUpdateOne) SetCost(c *Cost) *EquipmentUpdateOne {
	return euo.SetCostID(c.ID)
}

// SetGearID sets the "gear" edge to the Gear entity by ID.
func (euo *EquipmentUpdateOne) SetGearID(id int) *EquipmentUpdateOne {
	euo.mutation.SetGearID(id)
	return euo
}

// SetNillableGearID sets the "gear" edge to the Gear entity by ID if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableGearID(id *int) *EquipmentUpdateOne {
	if id != nil {
		euo = euo.SetGearID(*id)
	}
	return euo
}

// SetGear sets the "gear" edge to the Gear entity.
func (euo *EquipmentUpdateOne) SetGear(g *Gear) *EquipmentUpdateOne {
	return euo.SetGearID(g.ID)
}

// SetToolID sets the "tool" edge to the Tool entity by ID.
func (euo *EquipmentUpdateOne) SetToolID(id int) *EquipmentUpdateOne {
	euo.mutation.SetToolID(id)
	return euo
}

// SetNillableToolID sets the "tool" edge to the Tool entity by ID if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableToolID(id *int) *EquipmentUpdateOne {
	if id != nil {
		euo = euo.SetToolID(*id)
	}
	return euo
}

// SetTool sets the "tool" edge to the Tool entity.
func (euo *EquipmentUpdateOne) SetTool(t *Tool) *EquipmentUpdateOne {
	return euo.SetToolID(t.ID)
}

// SetWeaponID sets the "weapon" edge to the Weapon entity by ID.
func (euo *EquipmentUpdateOne) SetWeaponID(id int) *EquipmentUpdateOne {
	euo.mutation.SetWeaponID(id)
	return euo
}

// SetNillableWeaponID sets the "weapon" edge to the Weapon entity by ID if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableWeaponID(id *int) *EquipmentUpdateOne {
	if id != nil {
		euo = euo.SetWeaponID(*id)
	}
	return euo
}

// SetWeapon sets the "weapon" edge to the Weapon entity.
func (euo *EquipmentUpdateOne) SetWeapon(w *Weapon) *EquipmentUpdateOne {
	return euo.SetWeaponID(w.ID)
}

// SetVehicleID sets the "vehicle" edge to the Vehicle entity by ID.
func (euo *EquipmentUpdateOne) SetVehicleID(id int) *EquipmentUpdateOne {
	euo.mutation.SetVehicleID(id)
	return euo
}

// SetNillableVehicleID sets the "vehicle" edge to the Vehicle entity by ID if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableVehicleID(id *int) *EquipmentUpdateOne {
	if id != nil {
		euo = euo.SetVehicleID(*id)
	}
	return euo
}

// SetVehicle sets the "vehicle" edge to the Vehicle entity.
func (euo *EquipmentUpdateOne) SetVehicle(v *Vehicle) *EquipmentUpdateOne {
	return euo.SetVehicleID(v.ID)
}

// SetArmorID sets the "armor" edge to the Armor entity by ID.
func (euo *EquipmentUpdateOne) SetArmorID(id int) *EquipmentUpdateOne {
	euo.mutation.SetArmorID(id)
	return euo
}

// SetNillableArmorID sets the "armor" edge to the Armor entity by ID if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableArmorID(id *int) *EquipmentUpdateOne {
	if id != nil {
		euo = euo.SetArmorID(*id)
	}
	return euo
}

// SetArmor sets the "armor" edge to the Armor entity.
func (euo *EquipmentUpdateOne) SetArmor(a *Armor) *EquipmentUpdateOne {
	return euo.SetArmorID(a.ID)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (euo *EquipmentUpdateOne) AddProficiencyIDs(ids ...int) *EquipmentUpdateOne {
	euo.mutation.AddProficiencyIDs(ids...)
	return euo
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (euo *EquipmentUpdateOne) AddProficiencies(p ...*Proficiency) *EquipmentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.AddProficiencyIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (euo *EquipmentUpdateOne) Mutation() *EquipmentMutation {
	return euo.mutation
}

// ClearCost clears the "cost" edge to the Cost entity.
func (euo *EquipmentUpdateOne) ClearCost() *EquipmentUpdateOne {
	euo.mutation.ClearCost()
	return euo
}

// ClearGear clears the "gear" edge to the Gear entity.
func (euo *EquipmentUpdateOne) ClearGear() *EquipmentUpdateOne {
	euo.mutation.ClearGear()
	return euo
}

// ClearTool clears the "tool" edge to the Tool entity.
func (euo *EquipmentUpdateOne) ClearTool() *EquipmentUpdateOne {
	euo.mutation.ClearTool()
	return euo
}

// ClearWeapon clears the "weapon" edge to the Weapon entity.
func (euo *EquipmentUpdateOne) ClearWeapon() *EquipmentUpdateOne {
	euo.mutation.ClearWeapon()
	return euo
}

// ClearVehicle clears the "vehicle" edge to the Vehicle entity.
func (euo *EquipmentUpdateOne) ClearVehicle() *EquipmentUpdateOne {
	euo.mutation.ClearVehicle()
	return euo
}

// ClearArmor clears the "armor" edge to the Armor entity.
func (euo *EquipmentUpdateOne) ClearArmor() *EquipmentUpdateOne {
	euo.mutation.ClearArmor()
	return euo
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (euo *EquipmentUpdateOne) ClearProficiencies() *EquipmentUpdateOne {
	euo.mutation.ClearProficiencies()
	return euo
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (euo *EquipmentUpdateOne) RemoveProficiencyIDs(ids ...int) *EquipmentUpdateOne {
	euo.mutation.RemoveProficiencyIDs(ids...)
	return euo
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (euo *EquipmentUpdateOne) RemoveProficiencies(p ...*Proficiency) *EquipmentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return euo.RemoveProficiencyIDs(ids...)
}

// Where appends a list predicates to the EquipmentUpdate builder.
func (euo *EquipmentUpdateOne) Where(ps ...predicate.Equipment) *EquipmentUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EquipmentUpdateOne) Select(field string, fields ...string) *EquipmentUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Equipment entity.
func (euo *EquipmentUpdateOne) Save(ctx context.Context) (*Equipment, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EquipmentUpdateOne) SaveX(ctx context.Context) *Equipment {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EquipmentUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EquipmentUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EquipmentUpdateOne) check() error {
	if v, ok := euo.mutation.Indx(); ok {
		if err := equipment.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Equipment.indx": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Name(); ok {
		if err := equipment.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Equipment.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.EquipmentCategory(); ok {
		if err := equipment.EquipmentCategoryValidator(v); err != nil {
			return &ValidationError{Name: "equipment_category", err: fmt.Errorf(`ent: validator failed for field "Equipment.equipment_category": %w`, err)}
		}
	}
	return nil
}

func (euo *EquipmentUpdateOne) sqlSave(ctx context.Context) (_node *Equipment, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(equipment.Table, equipment.Columns, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Equipment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipment.FieldID)
		for _, f := range fields {
			if !equipment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Indx(); ok {
		_spec.SetField(equipment.FieldIndx, field.TypeString, value)
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(equipment.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.EquipmentCategory(); ok {
		_spec.SetField(equipment.FieldEquipmentCategory, field.TypeEnum, value)
	}
	if value, ok := euo.mutation.Weight(); ok {
		_spec.SetField(equipment.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := euo.mutation.AddedWeight(); ok {
		_spec.AddField(equipment.FieldWeight, field.TypeFloat64, value)
	}
	if euo.mutation.WeightCleared() {
		_spec.ClearField(equipment.FieldWeight, field.TypeFloat64)
	}
	if euo.mutation.CostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.CostTable,
			Columns: []string{equipment.CostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.CostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.CostTable,
			Columns: []string{equipment.CostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.GearCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.GearTable,
			Columns: []string{equipment.GearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.GearIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.GearTable,
			Columns: []string{equipment.GearColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gear.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ToolCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ToolTable,
			Columns: []string{equipment.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ToolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ToolTable,
			Columns: []string{equipment.ToolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.WeaponCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.WeaponTable,
			Columns: []string{equipment.WeaponColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(weapon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.WeaponIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.WeaponTable,
			Columns: []string{equipment.WeaponColumn},
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
	if euo.mutation.VehicleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.VehicleTable,
			Columns: []string{equipment.VehicleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.VehicleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.VehicleTable,
			Columns: []string{equipment.VehicleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ArmorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ArmorTable,
			Columns: []string{equipment.ArmorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ArmorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   equipment.ArmorTable,
			Columns: []string{equipment.ArmorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(armor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipment.ProficienciesTable,
			Columns: []string{equipment.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !euo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipment.ProficienciesTable,
			Columns: []string{equipment.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   equipment.ProficienciesTable,
			Columns: []string{equipment.ProficienciesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Equipment{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
