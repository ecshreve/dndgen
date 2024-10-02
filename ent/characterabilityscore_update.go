// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterskill"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// CharacterAbilityScoreUpdate is the builder for updating CharacterAbilityScore entities.
type CharacterAbilityScoreUpdate struct {
	config
	hooks    []Hook
	mutation *CharacterAbilityScoreMutation
}

// Where appends a list predicates to the CharacterAbilityScoreUpdate builder.
func (casu *CharacterAbilityScoreUpdate) Where(ps ...predicate.CharacterAbilityScore) *CharacterAbilityScoreUpdate {
	casu.mutation.Where(ps...)
	return casu
}

// SetScore sets the "score" field.
func (casu *CharacterAbilityScoreUpdate) SetScore(i int) *CharacterAbilityScoreUpdate {
	casu.mutation.ResetScore()
	casu.mutation.SetScore(i)
	return casu
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (casu *CharacterAbilityScoreUpdate) SetNillableScore(i *int) *CharacterAbilityScoreUpdate {
	if i != nil {
		casu.SetScore(*i)
	}
	return casu
}

// AddScore adds i to the "score" field.
func (casu *CharacterAbilityScoreUpdate) AddScore(i int) *CharacterAbilityScoreUpdate {
	casu.mutation.AddScore(i)
	return casu
}

// SetModifier sets the "modifier" field.
func (casu *CharacterAbilityScoreUpdate) SetModifier(i int) *CharacterAbilityScoreUpdate {
	casu.mutation.ResetModifier()
	casu.mutation.SetModifier(i)
	return casu
}

// SetNillableModifier sets the "modifier" field if the given value is not nil.
func (casu *CharacterAbilityScoreUpdate) SetNillableModifier(i *int) *CharacterAbilityScoreUpdate {
	if i != nil {
		casu.SetModifier(*i)
	}
	return casu
}

// AddModifier adds i to the "modifier" field.
func (casu *CharacterAbilityScoreUpdate) AddModifier(i int) *CharacterAbilityScoreUpdate {
	casu.mutation.AddModifier(i)
	return casu
}

// SetCharacterID sets the "character" edge to the Character entity by ID.
func (casu *CharacterAbilityScoreUpdate) SetCharacterID(id int) *CharacterAbilityScoreUpdate {
	casu.mutation.SetCharacterID(id)
	return casu
}

// SetCharacter sets the "character" edge to the Character entity.
func (casu *CharacterAbilityScoreUpdate) SetCharacter(c *Character) *CharacterAbilityScoreUpdate {
	return casu.SetCharacterID(c.ID)
}

// SetAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID.
func (casu *CharacterAbilityScoreUpdate) SetAbilityScoreID(id int) *CharacterAbilityScoreUpdate {
	casu.mutation.SetAbilityScoreID(id)
	return casu
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (casu *CharacterAbilityScoreUpdate) SetAbilityScore(a *AbilityScore) *CharacterAbilityScoreUpdate {
	return casu.SetAbilityScoreID(a.ID)
}

// AddCharacterSkillIDs adds the "character_skills" edge to the CharacterSkill entity by IDs.
func (casu *CharacterAbilityScoreUpdate) AddCharacterSkillIDs(ids ...int) *CharacterAbilityScoreUpdate {
	casu.mutation.AddCharacterSkillIDs(ids...)
	return casu
}

// AddCharacterSkills adds the "character_skills" edges to the CharacterSkill entity.
func (casu *CharacterAbilityScoreUpdate) AddCharacterSkills(c ...*CharacterSkill) *CharacterAbilityScoreUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return casu.AddCharacterSkillIDs(ids...)
}

// Mutation returns the CharacterAbilityScoreMutation object of the builder.
func (casu *CharacterAbilityScoreUpdate) Mutation() *CharacterAbilityScoreMutation {
	return casu.mutation
}

// ClearCharacter clears the "character" edge to the Character entity.
func (casu *CharacterAbilityScoreUpdate) ClearCharacter() *CharacterAbilityScoreUpdate {
	casu.mutation.ClearCharacter()
	return casu
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (casu *CharacterAbilityScoreUpdate) ClearAbilityScore() *CharacterAbilityScoreUpdate {
	casu.mutation.ClearAbilityScore()
	return casu
}

// ClearCharacterSkills clears all "character_skills" edges to the CharacterSkill entity.
func (casu *CharacterAbilityScoreUpdate) ClearCharacterSkills() *CharacterAbilityScoreUpdate {
	casu.mutation.ClearCharacterSkills()
	return casu
}

// RemoveCharacterSkillIDs removes the "character_skills" edge to CharacterSkill entities by IDs.
func (casu *CharacterAbilityScoreUpdate) RemoveCharacterSkillIDs(ids ...int) *CharacterAbilityScoreUpdate {
	casu.mutation.RemoveCharacterSkillIDs(ids...)
	return casu
}

// RemoveCharacterSkills removes "character_skills" edges to CharacterSkill entities.
func (casu *CharacterAbilityScoreUpdate) RemoveCharacterSkills(c ...*CharacterSkill) *CharacterAbilityScoreUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return casu.RemoveCharacterSkillIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (casu *CharacterAbilityScoreUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, casu.sqlSave, casu.mutation, casu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (casu *CharacterAbilityScoreUpdate) SaveX(ctx context.Context) int {
	affected, err := casu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (casu *CharacterAbilityScoreUpdate) Exec(ctx context.Context) error {
	_, err := casu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (casu *CharacterAbilityScoreUpdate) ExecX(ctx context.Context) {
	if err := casu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (casu *CharacterAbilityScoreUpdate) check() error {
	if v, ok := casu.mutation.Score(); ok {
		if err := characterabilityscore.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "CharacterAbilityScore.score": %w`, err)}
		}
	}
	if casu.mutation.CharacterCleared() && len(casu.mutation.CharacterIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterAbilityScore.character"`)
	}
	if casu.mutation.AbilityScoreCleared() && len(casu.mutation.AbilityScoreIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterAbilityScore.ability_score"`)
	}
	return nil
}

func (casu *CharacterAbilityScoreUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := casu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(characterabilityscore.Table, characterabilityscore.Columns, sqlgraph.NewFieldSpec(characterabilityscore.FieldID, field.TypeInt))
	if ps := casu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := casu.mutation.Score(); ok {
		_spec.SetField(characterabilityscore.FieldScore, field.TypeInt, value)
	}
	if value, ok := casu.mutation.AddedScore(); ok {
		_spec.AddField(characterabilityscore.FieldScore, field.TypeInt, value)
	}
	if value, ok := casu.mutation.Modifier(); ok {
		_spec.SetField(characterabilityscore.FieldModifier, field.TypeInt, value)
	}
	if value, ok := casu.mutation.AddedModifier(); ok {
		_spec.AddField(characterabilityscore.FieldModifier, field.TypeInt, value)
	}
	if casu.mutation.CharacterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   characterabilityscore.CharacterTable,
			Columns: []string{characterabilityscore.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casu.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   characterabilityscore.CharacterTable,
			Columns: []string{characterabilityscore.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if casu.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterabilityscore.AbilityScoreTable,
			Columns: []string{characterabilityscore.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casu.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterabilityscore.AbilityScoreTable,
			Columns: []string{characterabilityscore.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if casu.mutation.CharacterSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casu.mutation.RemovedCharacterSkillsIDs(); len(nodes) > 0 && !casu.mutation.CharacterSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casu.mutation.CharacterSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, casu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{characterabilityscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	casu.mutation.done = true
	return n, nil
}

// CharacterAbilityScoreUpdateOne is the builder for updating a single CharacterAbilityScore entity.
type CharacterAbilityScoreUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CharacterAbilityScoreMutation
}

// SetScore sets the "score" field.
func (casuo *CharacterAbilityScoreUpdateOne) SetScore(i int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.ResetScore()
	casuo.mutation.SetScore(i)
	return casuo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (casuo *CharacterAbilityScoreUpdateOne) SetNillableScore(i *int) *CharacterAbilityScoreUpdateOne {
	if i != nil {
		casuo.SetScore(*i)
	}
	return casuo
}

// AddScore adds i to the "score" field.
func (casuo *CharacterAbilityScoreUpdateOne) AddScore(i int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.AddScore(i)
	return casuo
}

// SetModifier sets the "modifier" field.
func (casuo *CharacterAbilityScoreUpdateOne) SetModifier(i int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.ResetModifier()
	casuo.mutation.SetModifier(i)
	return casuo
}

// SetNillableModifier sets the "modifier" field if the given value is not nil.
func (casuo *CharacterAbilityScoreUpdateOne) SetNillableModifier(i *int) *CharacterAbilityScoreUpdateOne {
	if i != nil {
		casuo.SetModifier(*i)
	}
	return casuo
}

// AddModifier adds i to the "modifier" field.
func (casuo *CharacterAbilityScoreUpdateOne) AddModifier(i int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.AddModifier(i)
	return casuo
}

// SetCharacterID sets the "character" edge to the Character entity by ID.
func (casuo *CharacterAbilityScoreUpdateOne) SetCharacterID(id int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.SetCharacterID(id)
	return casuo
}

// SetCharacter sets the "character" edge to the Character entity.
func (casuo *CharacterAbilityScoreUpdateOne) SetCharacter(c *Character) *CharacterAbilityScoreUpdateOne {
	return casuo.SetCharacterID(c.ID)
}

// SetAbilityScoreID sets the "ability_score" edge to the AbilityScore entity by ID.
func (casuo *CharacterAbilityScoreUpdateOne) SetAbilityScoreID(id int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.SetAbilityScoreID(id)
	return casuo
}

// SetAbilityScore sets the "ability_score" edge to the AbilityScore entity.
func (casuo *CharacterAbilityScoreUpdateOne) SetAbilityScore(a *AbilityScore) *CharacterAbilityScoreUpdateOne {
	return casuo.SetAbilityScoreID(a.ID)
}

// AddCharacterSkillIDs adds the "character_skills" edge to the CharacterSkill entity by IDs.
func (casuo *CharacterAbilityScoreUpdateOne) AddCharacterSkillIDs(ids ...int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.AddCharacterSkillIDs(ids...)
	return casuo
}

// AddCharacterSkills adds the "character_skills" edges to the CharacterSkill entity.
func (casuo *CharacterAbilityScoreUpdateOne) AddCharacterSkills(c ...*CharacterSkill) *CharacterAbilityScoreUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return casuo.AddCharacterSkillIDs(ids...)
}

// Mutation returns the CharacterAbilityScoreMutation object of the builder.
func (casuo *CharacterAbilityScoreUpdateOne) Mutation() *CharacterAbilityScoreMutation {
	return casuo.mutation
}

// ClearCharacter clears the "character" edge to the Character entity.
func (casuo *CharacterAbilityScoreUpdateOne) ClearCharacter() *CharacterAbilityScoreUpdateOne {
	casuo.mutation.ClearCharacter()
	return casuo
}

// ClearAbilityScore clears the "ability_score" edge to the AbilityScore entity.
func (casuo *CharacterAbilityScoreUpdateOne) ClearAbilityScore() *CharacterAbilityScoreUpdateOne {
	casuo.mutation.ClearAbilityScore()
	return casuo
}

// ClearCharacterSkills clears all "character_skills" edges to the CharacterSkill entity.
func (casuo *CharacterAbilityScoreUpdateOne) ClearCharacterSkills() *CharacterAbilityScoreUpdateOne {
	casuo.mutation.ClearCharacterSkills()
	return casuo
}

// RemoveCharacterSkillIDs removes the "character_skills" edge to CharacterSkill entities by IDs.
func (casuo *CharacterAbilityScoreUpdateOne) RemoveCharacterSkillIDs(ids ...int) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.RemoveCharacterSkillIDs(ids...)
	return casuo
}

// RemoveCharacterSkills removes "character_skills" edges to CharacterSkill entities.
func (casuo *CharacterAbilityScoreUpdateOne) RemoveCharacterSkills(c ...*CharacterSkill) *CharacterAbilityScoreUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return casuo.RemoveCharacterSkillIDs(ids...)
}

// Where appends a list predicates to the CharacterAbilityScoreUpdate builder.
func (casuo *CharacterAbilityScoreUpdateOne) Where(ps ...predicate.CharacterAbilityScore) *CharacterAbilityScoreUpdateOne {
	casuo.mutation.Where(ps...)
	return casuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (casuo *CharacterAbilityScoreUpdateOne) Select(field string, fields ...string) *CharacterAbilityScoreUpdateOne {
	casuo.fields = append([]string{field}, fields...)
	return casuo
}

// Save executes the query and returns the updated CharacterAbilityScore entity.
func (casuo *CharacterAbilityScoreUpdateOne) Save(ctx context.Context) (*CharacterAbilityScore, error) {
	return withHooks(ctx, casuo.sqlSave, casuo.mutation, casuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (casuo *CharacterAbilityScoreUpdateOne) SaveX(ctx context.Context) *CharacterAbilityScore {
	node, err := casuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (casuo *CharacterAbilityScoreUpdateOne) Exec(ctx context.Context) error {
	_, err := casuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (casuo *CharacterAbilityScoreUpdateOne) ExecX(ctx context.Context) {
	if err := casuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (casuo *CharacterAbilityScoreUpdateOne) check() error {
	if v, ok := casuo.mutation.Score(); ok {
		if err := characterabilityscore.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "CharacterAbilityScore.score": %w`, err)}
		}
	}
	if casuo.mutation.CharacterCleared() && len(casuo.mutation.CharacterIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterAbilityScore.character"`)
	}
	if casuo.mutation.AbilityScoreCleared() && len(casuo.mutation.AbilityScoreIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "CharacterAbilityScore.ability_score"`)
	}
	return nil
}

func (casuo *CharacterAbilityScoreUpdateOne) sqlSave(ctx context.Context) (_node *CharacterAbilityScore, err error) {
	if err := casuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(characterabilityscore.Table, characterabilityscore.Columns, sqlgraph.NewFieldSpec(characterabilityscore.FieldID, field.TypeInt))
	id, ok := casuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CharacterAbilityScore.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := casuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, characterabilityscore.FieldID)
		for _, f := range fields {
			if !characterabilityscore.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != characterabilityscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := casuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := casuo.mutation.Score(); ok {
		_spec.SetField(characterabilityscore.FieldScore, field.TypeInt, value)
	}
	if value, ok := casuo.mutation.AddedScore(); ok {
		_spec.AddField(characterabilityscore.FieldScore, field.TypeInt, value)
	}
	if value, ok := casuo.mutation.Modifier(); ok {
		_spec.SetField(characterabilityscore.FieldModifier, field.TypeInt, value)
	}
	if value, ok := casuo.mutation.AddedModifier(); ok {
		_spec.AddField(characterabilityscore.FieldModifier, field.TypeInt, value)
	}
	if casuo.mutation.CharacterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   characterabilityscore.CharacterTable,
			Columns: []string{characterabilityscore.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casuo.mutation.CharacterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   characterabilityscore.CharacterTable,
			Columns: []string{characterabilityscore.CharacterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if casuo.mutation.AbilityScoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterabilityscore.AbilityScoreTable,
			Columns: []string{characterabilityscore.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casuo.mutation.AbilityScoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   characterabilityscore.AbilityScoreTable,
			Columns: []string{characterabilityscore.AbilityScoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if casuo.mutation.CharacterSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casuo.mutation.RemovedCharacterSkillsIDs(); len(nodes) > 0 && !casuo.mutation.CharacterSkillsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := casuo.mutation.CharacterSkillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   characterabilityscore.CharacterSkillsTable,
			Columns: []string{characterabilityscore.CharacterSkillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CharacterAbilityScore{config: casuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, casuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{characterabilityscore.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	casuo.mutation.done = true
	return _node, nil
}
