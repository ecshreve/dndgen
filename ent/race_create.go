// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/choice"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/trait"
)

// RaceCreate is the builder for creating a Race entity.
type RaceCreate struct {
	config
	mutation *RaceMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (rc *RaceCreate) SetIndx(s string) *RaceCreate {
	rc.mutation.SetIndx(s)
	return rc
}

// SetName sets the "name" field.
func (rc *RaceCreate) SetName(s string) *RaceCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetAlignment sets the "alignment" field.
func (rc *RaceCreate) SetAlignment(s string) *RaceCreate {
	rc.mutation.SetAlignment(s)
	return rc
}

// SetAge sets the "age" field.
func (rc *RaceCreate) SetAge(s string) *RaceCreate {
	rc.mutation.SetAge(s)
	return rc
}

// SetSize sets the "size" field.
func (rc *RaceCreate) SetSize(s string) *RaceCreate {
	rc.mutation.SetSize(s)
	return rc
}

// SetSizeDescription sets the "size_description" field.
func (rc *RaceCreate) SetSizeDescription(s string) *RaceCreate {
	rc.mutation.SetSizeDescription(s)
	return rc
}

// SetLanguageDesc sets the "language_desc" field.
func (rc *RaceCreate) SetLanguageDesc(s string) *RaceCreate {
	rc.mutation.SetLanguageDesc(s)
	return rc
}

// SetSpeed sets the "speed" field.
func (rc *RaceCreate) SetSpeed(i int) *RaceCreate {
	rc.mutation.SetSpeed(i)
	return rc
}

// AddLanguageIDs adds the "languages" edge to the Language entity by IDs.
func (rc *RaceCreate) AddLanguageIDs(ids ...int) *RaceCreate {
	rc.mutation.AddLanguageIDs(ids...)
	return rc
}

// AddLanguages adds the "languages" edges to the Language entity.
func (rc *RaceCreate) AddLanguages(l ...*Language) *RaceCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return rc.AddLanguageIDs(ids...)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (rc *RaceCreate) AddProficiencyIDs(ids ...int) *RaceCreate {
	rc.mutation.AddProficiencyIDs(ids...)
	return rc
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (rc *RaceCreate) AddProficiencies(p ...*Proficiency) *RaceCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return rc.AddProficiencyIDs(ids...)
}

// AddSubraceIDs adds the "subraces" edge to the Subrace entity by IDs.
func (rc *RaceCreate) AddSubraceIDs(ids ...int) *RaceCreate {
	rc.mutation.AddSubraceIDs(ids...)
	return rc
}

// AddSubraces adds the "subraces" edges to the Subrace entity.
func (rc *RaceCreate) AddSubraces(s ...*Subrace) *RaceCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return rc.AddSubraceIDs(ids...)
}

// AddTraitIDs adds the "traits" edge to the Trait entity by IDs.
func (rc *RaceCreate) AddTraitIDs(ids ...int) *RaceCreate {
	rc.mutation.AddTraitIDs(ids...)
	return rc
}

// AddTraits adds the "traits" edges to the Trait entity.
func (rc *RaceCreate) AddTraits(t ...*Trait) *RaceCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddTraitIDs(ids...)
}

// AddAbilityBonuseIDs adds the "ability_bonuses" edge to the AbilityBonus entity by IDs.
func (rc *RaceCreate) AddAbilityBonuseIDs(ids ...int) *RaceCreate {
	rc.mutation.AddAbilityBonuseIDs(ids...)
	return rc
}

// AddAbilityBonuses adds the "ability_bonuses" edges to the AbilityBonus entity.
func (rc *RaceCreate) AddAbilityBonuses(a ...*AbilityBonus) *RaceCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAbilityBonuseIDs(ids...)
}

// SetStartingProficiencyOptionsID sets the "starting_proficiency_options" edge to the Choice entity by ID.
func (rc *RaceCreate) SetStartingProficiencyOptionsID(id int) *RaceCreate {
	rc.mutation.SetStartingProficiencyOptionsID(id)
	return rc
}

// SetNillableStartingProficiencyOptionsID sets the "starting_proficiency_options" edge to the Choice entity by ID if the given value is not nil.
func (rc *RaceCreate) SetNillableStartingProficiencyOptionsID(id *int) *RaceCreate {
	if id != nil {
		rc = rc.SetStartingProficiencyOptionsID(*id)
	}
	return rc
}

// SetStartingProficiencyOptions sets the "starting_proficiency_options" edge to the Choice entity.
func (rc *RaceCreate) SetStartingProficiencyOptions(c *Choice) *RaceCreate {
	return rc.SetStartingProficiencyOptionsID(c.ID)
}

// Mutation returns the RaceMutation object of the builder.
func (rc *RaceCreate) Mutation() *RaceMutation {
	return rc.mutation
}

// Save creates the Race in the database.
func (rc *RaceCreate) Save(ctx context.Context) (*Race, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RaceCreate) SaveX(ctx context.Context) *Race {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RaceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RaceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RaceCreate) check() error {
	if _, ok := rc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Race.indx"`)}
	}
	if v, ok := rc.mutation.Indx(); ok {
		if err := race.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Race.indx": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Race.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := race.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Race.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Alignment(); !ok {
		return &ValidationError{Name: "alignment", err: errors.New(`ent: missing required field "Race.alignment"`)}
	}
	if _, ok := rc.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Race.age"`)}
	}
	if _, ok := rc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Race.size"`)}
	}
	if _, ok := rc.mutation.SizeDescription(); !ok {
		return &ValidationError{Name: "size_description", err: errors.New(`ent: missing required field "Race.size_description"`)}
	}
	if _, ok := rc.mutation.LanguageDesc(); !ok {
		return &ValidationError{Name: "language_desc", err: errors.New(`ent: missing required field "Race.language_desc"`)}
	}
	if _, ok := rc.mutation.Speed(); !ok {
		return &ValidationError{Name: "speed", err: errors.New(`ent: missing required field "Race.speed"`)}
	}
	return nil
}

func (rc *RaceCreate) sqlSave(ctx context.Context) (*Race, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RaceCreate) createSpec() (*Race, *sqlgraph.CreateSpec) {
	var (
		_node = &Race{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(race.Table, sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Indx(); ok {
		_spec.SetField(race.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(race.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Alignment(); ok {
		_spec.SetField(race.FieldAlignment, field.TypeString, value)
		_node.Alignment = value
	}
	if value, ok := rc.mutation.Age(); ok {
		_spec.SetField(race.FieldAge, field.TypeString, value)
		_node.Age = value
	}
	if value, ok := rc.mutation.Size(); ok {
		_spec.SetField(race.FieldSize, field.TypeString, value)
		_node.Size = value
	}
	if value, ok := rc.mutation.SizeDescription(); ok {
		_spec.SetField(race.FieldSizeDescription, field.TypeString, value)
		_node.SizeDescription = value
	}
	if value, ok := rc.mutation.LanguageDesc(); ok {
		_spec.SetField(race.FieldLanguageDesc, field.TypeString, value)
		_node.LanguageDesc = value
	}
	if value, ok := rc.mutation.Speed(); ok {
		_spec.SetField(race.FieldSpeed, field.TypeInt, value)
		_node.Speed = value
	}
	if nodes := rc.mutation.LanguagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.SubracesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TraitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.AbilityBonusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.StartingProficiencyOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   race.StartingProficiencyOptionsTable,
			Columns: []string{race.StartingProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.race_starting_proficiency_options = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RaceCreateBulk is the builder for creating many Race entities in bulk.
type RaceCreateBulk struct {
	config
	builders []*RaceCreate
}

// Save creates the Race entities in the database.
func (rcb *RaceCreateBulk) Save(ctx context.Context) ([]*Race, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Race, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RaceMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RaceCreateBulk) SaveX(ctx context.Context) []*Race {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RaceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RaceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
