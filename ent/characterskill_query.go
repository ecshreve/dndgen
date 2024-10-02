// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterproficiency"
	"github.com/ecshreve/dndgen/ent/characterskill"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/skill"
)

// CharacterSkillQuery is the builder for querying CharacterSkill entities.
type CharacterSkillQuery struct {
	config
	ctx                       *QueryContext
	order                     []characterskill.OrderOption
	inters                    []Interceptor
	predicates                []predicate.CharacterSkill
	withCharacter             *CharacterQuery
	withSkill                 *SkillQuery
	withCharacterAbilityScore *CharacterAbilityScoreQuery
	withCharacterProficiency  *CharacterProficiencyQuery
	withFKs                   bool
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*CharacterSkill) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CharacterSkillQuery builder.
func (csq *CharacterSkillQuery) Where(ps ...predicate.CharacterSkill) *CharacterSkillQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit the number of records to be returned by this query.
func (csq *CharacterSkillQuery) Limit(limit int) *CharacterSkillQuery {
	csq.ctx.Limit = &limit
	return csq
}

// Offset to start from.
func (csq *CharacterSkillQuery) Offset(offset int) *CharacterSkillQuery {
	csq.ctx.Offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CharacterSkillQuery) Unique(unique bool) *CharacterSkillQuery {
	csq.ctx.Unique = &unique
	return csq
}

// Order specifies how the records should be ordered.
func (csq *CharacterSkillQuery) Order(o ...characterskill.OrderOption) *CharacterSkillQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryCharacter chains the current query on the "character" edge.
func (csq *CharacterSkillQuery) QueryCharacter() *CharacterQuery {
	query := (&CharacterClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(characterskill.Table, characterskill.FieldID, selector),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, characterskill.CharacterTable, characterskill.CharacterColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySkill chains the current query on the "skill" edge.
func (csq *CharacterSkillQuery) QuerySkill() *SkillQuery {
	query := (&SkillClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(characterskill.Table, characterskill.FieldID, selector),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, characterskill.SkillTable, characterskill.SkillColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCharacterAbilityScore chains the current query on the "character_ability_score" edge.
func (csq *CharacterSkillQuery) QueryCharacterAbilityScore() *CharacterAbilityScoreQuery {
	query := (&CharacterAbilityScoreClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(characterskill.Table, characterskill.FieldID, selector),
			sqlgraph.To(characterabilityscore.Table, characterabilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, characterskill.CharacterAbilityScoreTable, characterskill.CharacterAbilityScoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCharacterProficiency chains the current query on the "character_proficiency" edge.
func (csq *CharacterSkillQuery) QueryCharacterProficiency() *CharacterProficiencyQuery {
	query := (&CharacterProficiencyClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(characterskill.Table, characterskill.FieldID, selector),
			sqlgraph.To(characterproficiency.Table, characterproficiency.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, characterskill.CharacterProficiencyTable, characterskill.CharacterProficiencyColumn),
		)
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CharacterSkill entity from the query.
// Returns a *NotFoundError when no CharacterSkill was found.
func (csq *CharacterSkillQuery) First(ctx context.Context) (*CharacterSkill, error) {
	nodes, err := csq.Limit(1).All(setContextOp(ctx, csq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{characterskill.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CharacterSkillQuery) FirstX(ctx context.Context) *CharacterSkill {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CharacterSkill ID from the query.
// Returns a *NotFoundError when no CharacterSkill ID was found.
func (csq *CharacterSkillQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(1).IDs(setContextOp(ctx, csq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{characterskill.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CharacterSkillQuery) FirstIDX(ctx context.Context) int {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CharacterSkill entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CharacterSkill entity is found.
// Returns a *NotFoundError when no CharacterSkill entities are found.
func (csq *CharacterSkillQuery) Only(ctx context.Context) (*CharacterSkill, error) {
	nodes, err := csq.Limit(2).All(setContextOp(ctx, csq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{characterskill.Label}
	default:
		return nil, &NotSingularError{characterskill.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CharacterSkillQuery) OnlyX(ctx context.Context) *CharacterSkill {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CharacterSkill ID in the query.
// Returns a *NotSingularError when more than one CharacterSkill ID is found.
// Returns a *NotFoundError when no entities are found.
func (csq *CharacterSkillQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = csq.Limit(2).IDs(setContextOp(ctx, csq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{characterskill.Label}
	default:
		err = &NotSingularError{characterskill.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CharacterSkillQuery) OnlyIDX(ctx context.Context) int {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CharacterSkills.
func (csq *CharacterSkillQuery) All(ctx context.Context) ([]*CharacterSkill, error) {
	ctx = setContextOp(ctx, csq.ctx, ent.OpQueryAll)
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CharacterSkill, *CharacterSkillQuery]()
	return withInterceptors[[]*CharacterSkill](ctx, csq, qr, csq.inters)
}

// AllX is like All, but panics if an error occurs.
func (csq *CharacterSkillQuery) AllX(ctx context.Context) []*CharacterSkill {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CharacterSkill IDs.
func (csq *CharacterSkillQuery) IDs(ctx context.Context) (ids []int, err error) {
	if csq.ctx.Unique == nil && csq.path != nil {
		csq.Unique(true)
	}
	ctx = setContextOp(ctx, csq.ctx, ent.OpQueryIDs)
	if err = csq.Select(characterskill.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CharacterSkillQuery) IDsX(ctx context.Context) []int {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CharacterSkillQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, csq.ctx, ent.OpQueryCount)
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, csq, querierCount[*CharacterSkillQuery](), csq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CharacterSkillQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CharacterSkillQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, csq.ctx, ent.OpQueryExist)
	switch _, err := csq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CharacterSkillQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CharacterSkillQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CharacterSkillQuery) Clone() *CharacterSkillQuery {
	if csq == nil {
		return nil
	}
	return &CharacterSkillQuery{
		config:                    csq.config,
		ctx:                       csq.ctx.Clone(),
		order:                     append([]characterskill.OrderOption{}, csq.order...),
		inters:                    append([]Interceptor{}, csq.inters...),
		predicates:                append([]predicate.CharacterSkill{}, csq.predicates...),
		withCharacter:             csq.withCharacter.Clone(),
		withSkill:                 csq.withSkill.Clone(),
		withCharacterAbilityScore: csq.withCharacterAbilityScore.Clone(),
		withCharacterProficiency:  csq.withCharacterProficiency.Clone(),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// WithCharacter tells the query-builder to eager-load the nodes that are connected to
// the "character" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CharacterSkillQuery) WithCharacter(opts ...func(*CharacterQuery)) *CharacterSkillQuery {
	query := (&CharacterClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withCharacter = query
	return csq
}

// WithSkill tells the query-builder to eager-load the nodes that are connected to
// the "skill" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CharacterSkillQuery) WithSkill(opts ...func(*SkillQuery)) *CharacterSkillQuery {
	query := (&SkillClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withSkill = query
	return csq
}

// WithCharacterAbilityScore tells the query-builder to eager-load the nodes that are connected to
// the "character_ability_score" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CharacterSkillQuery) WithCharacterAbilityScore(opts ...func(*CharacterAbilityScoreQuery)) *CharacterSkillQuery {
	query := (&CharacterAbilityScoreClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withCharacterAbilityScore = query
	return csq
}

// WithCharacterProficiency tells the query-builder to eager-load the nodes that are connected to
// the "character_proficiency" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CharacterSkillQuery) WithCharacterProficiency(opts ...func(*CharacterProficiencyQuery)) *CharacterSkillQuery {
	query := (&CharacterProficiencyClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withCharacterProficiency = query
	return csq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Proficient bool `json:"proficient,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CharacterSkill.Query().
//		GroupBy(characterskill.FieldProficient).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (csq *CharacterSkillQuery) GroupBy(field string, fields ...string) *CharacterSkillGroupBy {
	csq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CharacterSkillGroupBy{build: csq}
	grbuild.flds = &csq.ctx.Fields
	grbuild.label = characterskill.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Proficient bool `json:"proficient,omitempty"`
//	}
//
//	client.CharacterSkill.Query().
//		Select(characterskill.FieldProficient).
//		Scan(ctx, &v)
func (csq *CharacterSkillQuery) Select(fields ...string) *CharacterSkillSelect {
	csq.ctx.Fields = append(csq.ctx.Fields, fields...)
	sbuild := &CharacterSkillSelect{CharacterSkillQuery: csq}
	sbuild.label = characterskill.Label
	sbuild.flds, sbuild.scan = &csq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CharacterSkillSelect configured with the given aggregations.
func (csq *CharacterSkillQuery) Aggregate(fns ...AggregateFunc) *CharacterSkillSelect {
	return csq.Select().Aggregate(fns...)
}

func (csq *CharacterSkillQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range csq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, csq); err != nil {
				return err
			}
		}
	}
	for _, f := range csq.ctx.Fields {
		if !characterskill.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CharacterSkillQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CharacterSkill, error) {
	var (
		nodes       = []*CharacterSkill{}
		withFKs     = csq.withFKs
		_spec       = csq.querySpec()
		loadedTypes = [4]bool{
			csq.withCharacter != nil,
			csq.withSkill != nil,
			csq.withCharacterAbilityScore != nil,
			csq.withCharacterProficiency != nil,
		}
	)
	if csq.withCharacter != nil || csq.withSkill != nil || csq.withCharacterAbilityScore != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, characterskill.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CharacterSkill).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CharacterSkill{config: csq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(csq.modifiers) > 0 {
		_spec.Modifiers = csq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := csq.withCharacter; query != nil {
		if err := csq.loadCharacter(ctx, query, nodes, nil,
			func(n *CharacterSkill, e *Character) { n.Edges.Character = e }); err != nil {
			return nil, err
		}
	}
	if query := csq.withSkill; query != nil {
		if err := csq.loadSkill(ctx, query, nodes, nil,
			func(n *CharacterSkill, e *Skill) { n.Edges.Skill = e }); err != nil {
			return nil, err
		}
	}
	if query := csq.withCharacterAbilityScore; query != nil {
		if err := csq.loadCharacterAbilityScore(ctx, query, nodes, nil,
			func(n *CharacterSkill, e *CharacterAbilityScore) { n.Edges.CharacterAbilityScore = e }); err != nil {
			return nil, err
		}
	}
	if query := csq.withCharacterProficiency; query != nil {
		if err := csq.loadCharacterProficiency(ctx, query, nodes, nil,
			func(n *CharacterSkill, e *CharacterProficiency) { n.Edges.CharacterProficiency = e }); err != nil {
			return nil, err
		}
	}
	for i := range csq.loadTotal {
		if err := csq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (csq *CharacterSkillQuery) loadCharacter(ctx context.Context, query *CharacterQuery, nodes []*CharacterSkill, init func(*CharacterSkill), assign func(*CharacterSkill, *Character)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CharacterSkill)
	for i := range nodes {
		if nodes[i].character_character_skills == nil {
			continue
		}
		fk := *nodes[i].character_character_skills
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(character.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "character_character_skills" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (csq *CharacterSkillQuery) loadSkill(ctx context.Context, query *SkillQuery, nodes []*CharacterSkill, init func(*CharacterSkill), assign func(*CharacterSkill, *Skill)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CharacterSkill)
	for i := range nodes {
		if nodes[i].character_skill_skill == nil {
			continue
		}
		fk := *nodes[i].character_skill_skill
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(skill.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "character_skill_skill" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (csq *CharacterSkillQuery) loadCharacterAbilityScore(ctx context.Context, query *CharacterAbilityScoreQuery, nodes []*CharacterSkill, init func(*CharacterSkill), assign func(*CharacterSkill, *CharacterAbilityScore)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CharacterSkill)
	for i := range nodes {
		if nodes[i].character_skill_character_ability_score == nil {
			continue
		}
		fk := *nodes[i].character_skill_character_ability_score
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(characterabilityscore.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "character_skill_character_ability_score" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (csq *CharacterSkillQuery) loadCharacterProficiency(ctx context.Context, query *CharacterProficiencyQuery, nodes []*CharacterSkill, init func(*CharacterSkill), assign func(*CharacterSkill, *CharacterProficiency)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*CharacterSkill)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.CharacterProficiency(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(characterskill.CharacterProficiencyColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.character_skill_character_proficiency
		if fk == nil {
			return fmt.Errorf(`foreign-key "character_skill_character_proficiency" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "character_skill_character_proficiency" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (csq *CharacterSkillQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	if len(csq.modifiers) > 0 {
		_spec.Modifiers = csq.modifiers
	}
	_spec.Node.Columns = csq.ctx.Fields
	if len(csq.ctx.Fields) > 0 {
		_spec.Unique = csq.ctx.Unique != nil && *csq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CharacterSkillQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(characterskill.Table, characterskill.Columns, sqlgraph.NewFieldSpec(characterskill.FieldID, field.TypeInt))
	_spec.From = csq.sql
	if unique := csq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if csq.path != nil {
		_spec.Unique = true
	}
	if fields := csq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, characterskill.FieldID)
		for i := range fields {
			if fields[i] != characterskill.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *CharacterSkillQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(characterskill.Table)
	columns := csq.ctx.Fields
	if len(columns) == 0 {
		columns = characterskill.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.ctx.Unique != nil && *csq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CharacterSkillGroupBy is the group-by builder for CharacterSkill entities.
type CharacterSkillGroupBy struct {
	selector
	build *CharacterSkillQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CharacterSkillGroupBy) Aggregate(fns ...AggregateFunc) *CharacterSkillGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the selector query and scans the result into the given value.
func (csgb *CharacterSkillGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csgb.build.ctx, ent.OpQueryGroupBy)
	if err := csgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterSkillQuery, *CharacterSkillGroupBy](ctx, csgb.build, csgb, csgb.build.inters, v)
}

func (csgb *CharacterSkillGroupBy) sqlScan(ctx context.Context, root *CharacterSkillQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*csgb.flds)+len(csgb.fns))
		for _, f := range *csgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*csgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CharacterSkillSelect is the builder for selecting fields of CharacterSkill entities.
type CharacterSkillSelect struct {
	*CharacterSkillQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (css *CharacterSkillSelect) Aggregate(fns ...AggregateFunc) *CharacterSkillSelect {
	css.fns = append(css.fns, fns...)
	return css
}

// Scan applies the selector query and scans the result into the given value.
func (css *CharacterSkillSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, css.ctx, ent.OpQuerySelect)
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterSkillQuery, *CharacterSkillSelect](ctx, css.CharacterSkillQuery, css, css.inters, v)
}

func (css *CharacterSkillSelect) sqlScan(ctx context.Context, root *CharacterSkillQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(css.fns))
	for _, fn := range css.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*css.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
