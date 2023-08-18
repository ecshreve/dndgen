// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/race"
)

// LanguageCreate is the builder for creating a Language entity.
type LanguageCreate struct {
	config
	mutation *LanguageMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (lc *LanguageCreate) SetIndx(s string) *LanguageCreate {
	lc.mutation.SetIndx(s)
	return lc
}

// SetName sets the "name" field.
func (lc *LanguageCreate) SetName(s string) *LanguageCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetDesc sets the "desc" field.
func (lc *LanguageCreate) SetDesc(s string) *LanguageCreate {
	lc.mutation.SetDesc(s)
	return lc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (lc *LanguageCreate) SetNillableDesc(s *string) *LanguageCreate {
	if s != nil {
		lc.SetDesc(*s)
	}
	return lc
}

// SetCategory sets the "category" field.
func (lc *LanguageCreate) SetCategory(l language.Category) *LanguageCreate {
	lc.mutation.SetCategory(l)
	return lc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (lc *LanguageCreate) SetNillableCategory(l *language.Category) *LanguageCreate {
	if l != nil {
		lc.SetCategory(*l)
	}
	return lc
}

// SetScript sets the "script" field.
func (lc *LanguageCreate) SetScript(l language.Script) *LanguageCreate {
	lc.mutation.SetScript(l)
	return lc
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (lc *LanguageCreate) SetNillableScript(l *language.Script) *LanguageCreate {
	if l != nil {
		lc.SetScript(*l)
	}
	return lc
}

// AddSpeakerIDs adds the "speakers" edge to the Race entity by IDs.
func (lc *LanguageCreate) AddSpeakerIDs(ids ...int) *LanguageCreate {
	lc.mutation.AddSpeakerIDs(ids...)
	return lc
}

// AddSpeakers adds the "speakers" edges to the Race entity.
func (lc *LanguageCreate) AddSpeakers(r ...*Race) *LanguageCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return lc.AddSpeakerIDs(ids...)
}

// Mutation returns the LanguageMutation object of the builder.
func (lc *LanguageCreate) Mutation() *LanguageMutation {
	return lc.mutation
}

// Save creates the Language in the database.
func (lc *LanguageCreate) Save(ctx context.Context) (*Language, error) {
	lc.defaults()
	return withHooks[*Language, LanguageMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LanguageCreate) SaveX(ctx context.Context) *Language {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LanguageCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LanguageCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LanguageCreate) defaults() {
	if _, ok := lc.mutation.Category(); !ok {
		v := language.DefaultCategory
		lc.mutation.SetCategory(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LanguageCreate) check() error {
	if _, ok := lc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Language.indx"`)}
	}
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Language.name"`)}
	}
	if _, ok := lc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Language.category"`)}
	}
	if v, ok := lc.mutation.Category(); ok {
		if err := language.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Language.category": %w`, err)}
		}
	}
	if v, ok := lc.mutation.Script(); ok {
		if err := language.ScriptValidator(v); err != nil {
			return &ValidationError{Name: "script", err: fmt.Errorf(`ent: validator failed for field "Language.script": %w`, err)}
		}
	}
	return nil
}

func (lc *LanguageCreate) sqlSave(ctx context.Context) (*Language, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LanguageCreate) createSpec() (*Language, *sqlgraph.CreateSpec) {
	var (
		_node = &Language{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(language.Table, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.Indx(); ok {
		_spec.SetField(language.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.SetField(language.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lc.mutation.Desc(); ok {
		_spec.SetField(language.FieldDesc, field.TypeString, value)
		_node.Desc = &value
	}
	if value, ok := lc.mutation.Category(); ok {
		_spec.SetField(language.FieldCategory, field.TypeEnum, value)
		_node.Category = value
	}
	if value, ok := lc.mutation.Script(); ok {
		_spec.SetField(language.FieldScript, field.TypeEnum, value)
		_node.Script = &value
	}
	if nodes := lc.mutation.SpeakersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   language.SpeakersTable,
			Columns: language.SpeakersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LanguageCreateBulk is the builder for creating many Language entities in bulk.
type LanguageCreateBulk struct {
	config
	builders []*LanguageCreate
}

// Save creates the Language entities in the database.
func (lcb *LanguageCreateBulk) Save(ctx context.Context) ([]*Language, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Language, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LanguageMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LanguageCreateBulk) SaveX(ctx context.Context) []*Language {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LanguageCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LanguageCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
