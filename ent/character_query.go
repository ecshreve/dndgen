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
	"github.com/ecshreve/dndgen/ent/alignment"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/characterproficiency"
	"github.com/ecshreve/dndgen/ent/characterskill"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/race"
)

// CharacterQuery is the builder for querying Character entities.
type CharacterQuery struct {
	config
	ctx                             *QueryContext
	order                           []character.OrderOption
	inters                          []Interceptor
	predicates                      []predicate.Character
	withRace                        *RaceQuery
	withClass                       *ClassQuery
	withAlignment                   *AlignmentQuery
	withCharacterAbilityScores      *CharacterAbilityScoreQuery
	withCharacterSkills             *CharacterSkillQuery
	withCharacterProficiencies      *CharacterProficiencyQuery
	withFKs                         bool
	modifiers                       []func(*sql.Selector)
	loadTotal                       []func(context.Context, []*Character) error
	withNamedCharacterAbilityScores map[string]*CharacterAbilityScoreQuery
	withNamedCharacterSkills        map[string]*CharacterSkillQuery
	withNamedCharacterProficiencies map[string]*CharacterProficiencyQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CharacterQuery builder.
func (cq *CharacterQuery) Where(ps ...predicate.Character) *CharacterQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CharacterQuery) Limit(limit int) *CharacterQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CharacterQuery) Offset(offset int) *CharacterQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CharacterQuery) Unique(unique bool) *CharacterQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CharacterQuery) Order(o ...character.OrderOption) *CharacterQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryRace chains the current query on the "race" edge.
func (cq *CharacterQuery) QueryRace() *RaceQuery {
	query := (&RaceClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(race.Table, race.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, character.RaceTable, character.RaceColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryClass chains the current query on the "class" edge.
func (cq *CharacterQuery) QueryClass() *ClassQuery {
	query := (&ClassClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(class.Table, class.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, character.ClassTable, character.ClassColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAlignment chains the current query on the "alignment" edge.
func (cq *CharacterQuery) QueryAlignment() *AlignmentQuery {
	query := (&AlignmentClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(alignment.Table, alignment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, character.AlignmentTable, character.AlignmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCharacterAbilityScores chains the current query on the "character_ability_scores" edge.
func (cq *CharacterQuery) QueryCharacterAbilityScores() *CharacterAbilityScoreQuery {
	query := (&CharacterAbilityScoreClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(characterabilityscore.Table, characterabilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, character.CharacterAbilityScoresTable, character.CharacterAbilityScoresColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCharacterSkills chains the current query on the "character_skills" edge.
func (cq *CharacterQuery) QueryCharacterSkills() *CharacterSkillQuery {
	query := (&CharacterSkillClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(characterskill.Table, characterskill.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, character.CharacterSkillsTable, character.CharacterSkillsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCharacterProficiencies chains the current query on the "character_proficiencies" edge.
func (cq *CharacterQuery) QueryCharacterProficiencies() *CharacterProficiencyQuery {
	query := (&CharacterProficiencyClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(character.Table, character.FieldID, selector),
			sqlgraph.To(characterproficiency.Table, characterproficiency.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, character.CharacterProficienciesTable, character.CharacterProficienciesColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Character entity from the query.
// Returns a *NotFoundError when no Character was found.
func (cq *CharacterQuery) First(ctx context.Context) (*Character, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{character.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CharacterQuery) FirstX(ctx context.Context) *Character {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Character ID from the query.
// Returns a *NotFoundError when no Character ID was found.
func (cq *CharacterQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{character.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CharacterQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Character entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Character entity is found.
// Returns a *NotFoundError when no Character entities are found.
func (cq *CharacterQuery) Only(ctx context.Context) (*Character, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{character.Label}
	default:
		return nil, &NotSingularError{character.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CharacterQuery) OnlyX(ctx context.Context) *Character {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Character ID in the query.
// Returns a *NotSingularError when more than one Character ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CharacterQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{character.Label}
	default:
		err = &NotSingularError{character.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CharacterQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Characters.
func (cq *CharacterQuery) All(ctx context.Context) ([]*Character, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Character, *CharacterQuery]()
	return withInterceptors[[]*Character](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CharacterQuery) AllX(ctx context.Context) []*Character {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Character IDs.
func (cq *CharacterQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(character.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CharacterQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CharacterQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CharacterQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CharacterQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CharacterQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CharacterQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CharacterQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CharacterQuery) Clone() *CharacterQuery {
	if cq == nil {
		return nil
	}
	return &CharacterQuery{
		config:                     cq.config,
		ctx:                        cq.ctx.Clone(),
		order:                      append([]character.OrderOption{}, cq.order...),
		inters:                     append([]Interceptor{}, cq.inters...),
		predicates:                 append([]predicate.Character{}, cq.predicates...),
		withRace:                   cq.withRace.Clone(),
		withClass:                  cq.withClass.Clone(),
		withAlignment:              cq.withAlignment.Clone(),
		withCharacterAbilityScores: cq.withCharacterAbilityScores.Clone(),
		withCharacterSkills:        cq.withCharacterSkills.Clone(),
		withCharacterProficiencies: cq.withCharacterProficiencies.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithRace tells the query-builder to eager-load the nodes that are connected to
// the "race" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithRace(opts ...func(*RaceQuery)) *CharacterQuery {
	query := (&RaceClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withRace = query
	return cq
}

// WithClass tells the query-builder to eager-load the nodes that are connected to
// the "class" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithClass(opts ...func(*ClassQuery)) *CharacterQuery {
	query := (&ClassClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withClass = query
	return cq
}

// WithAlignment tells the query-builder to eager-load the nodes that are connected to
// the "alignment" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithAlignment(opts ...func(*AlignmentQuery)) *CharacterQuery {
	query := (&AlignmentClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withAlignment = query
	return cq
}

// WithCharacterAbilityScores tells the query-builder to eager-load the nodes that are connected to
// the "character_ability_scores" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithCharacterAbilityScores(opts ...func(*CharacterAbilityScoreQuery)) *CharacterQuery {
	query := (&CharacterAbilityScoreClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCharacterAbilityScores = query
	return cq
}

// WithCharacterSkills tells the query-builder to eager-load the nodes that are connected to
// the "character_skills" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithCharacterSkills(opts ...func(*CharacterSkillQuery)) *CharacterQuery {
	query := (&CharacterSkillClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCharacterSkills = query
	return cq
}

// WithCharacterProficiencies tells the query-builder to eager-load the nodes that are connected to
// the "character_proficiencies" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithCharacterProficiencies(opts ...func(*CharacterProficiencyQuery)) *CharacterQuery {
	query := (&CharacterProficiencyClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCharacterProficiencies = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Character.Query().
//		GroupBy(character.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CharacterQuery) GroupBy(field string, fields ...string) *CharacterGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CharacterGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = character.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Character.Query().
//		Select(character.FieldName).
//		Scan(ctx, &v)
func (cq *CharacterQuery) Select(fields ...string) *CharacterSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CharacterSelect{CharacterQuery: cq}
	sbuild.label = character.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CharacterSelect configured with the given aggregations.
func (cq *CharacterQuery) Aggregate(fns ...AggregateFunc) *CharacterSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CharacterQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !character.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CharacterQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Character, error) {
	var (
		nodes       = []*Character{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [6]bool{
			cq.withRace != nil,
			cq.withClass != nil,
			cq.withAlignment != nil,
			cq.withCharacterAbilityScores != nil,
			cq.withCharacterSkills != nil,
			cq.withCharacterProficiencies != nil,
		}
	)
	if cq.withRace != nil || cq.withClass != nil || cq.withAlignment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, character.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Character).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Character{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withRace; query != nil {
		if err := cq.loadRace(ctx, query, nodes, nil,
			func(n *Character, e *Race) { n.Edges.Race = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withClass; query != nil {
		if err := cq.loadClass(ctx, query, nodes, nil,
			func(n *Character, e *Class) { n.Edges.Class = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withAlignment; query != nil {
		if err := cq.loadAlignment(ctx, query, nodes, nil,
			func(n *Character, e *Alignment) { n.Edges.Alignment = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCharacterAbilityScores; query != nil {
		if err := cq.loadCharacterAbilityScores(ctx, query, nodes,
			func(n *Character) { n.Edges.CharacterAbilityScores = []*CharacterAbilityScore{} },
			func(n *Character, e *CharacterAbilityScore) {
				n.Edges.CharacterAbilityScores = append(n.Edges.CharacterAbilityScores, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := cq.withCharacterSkills; query != nil {
		if err := cq.loadCharacterSkills(ctx, query, nodes,
			func(n *Character) { n.Edges.CharacterSkills = []*CharacterSkill{} },
			func(n *Character, e *CharacterSkill) { n.Edges.CharacterSkills = append(n.Edges.CharacterSkills, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withCharacterProficiencies; query != nil {
		if err := cq.loadCharacterProficiencies(ctx, query, nodes,
			func(n *Character) { n.Edges.CharacterProficiencies = []*CharacterProficiency{} },
			func(n *Character, e *CharacterProficiency) {
				n.Edges.CharacterProficiencies = append(n.Edges.CharacterProficiencies, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedCharacterAbilityScores {
		if err := cq.loadCharacterAbilityScores(ctx, query, nodes,
			func(n *Character) { n.appendNamedCharacterAbilityScores(name) },
			func(n *Character, e *CharacterAbilityScore) { n.appendNamedCharacterAbilityScores(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedCharacterSkills {
		if err := cq.loadCharacterSkills(ctx, query, nodes,
			func(n *Character) { n.appendNamedCharacterSkills(name) },
			func(n *Character, e *CharacterSkill) { n.appendNamedCharacterSkills(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedCharacterProficiencies {
		if err := cq.loadCharacterProficiencies(ctx, query, nodes,
			func(n *Character) { n.appendNamedCharacterProficiencies(name) },
			func(n *Character, e *CharacterProficiency) { n.appendNamedCharacterProficiencies(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CharacterQuery) loadRace(ctx context.Context, query *RaceQuery, nodes []*Character, init func(*Character), assign func(*Character, *Race)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Character)
	for i := range nodes {
		if nodes[i].character_race == nil {
			continue
		}
		fk := *nodes[i].character_race
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(race.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "character_race" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CharacterQuery) loadClass(ctx context.Context, query *ClassQuery, nodes []*Character, init func(*Character), assign func(*Character, *Class)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Character)
	for i := range nodes {
		if nodes[i].character_class == nil {
			continue
		}
		fk := *nodes[i].character_class
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(class.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "character_class" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CharacterQuery) loadAlignment(ctx context.Context, query *AlignmentQuery, nodes []*Character, init func(*Character), assign func(*Character, *Alignment)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Character)
	for i := range nodes {
		if nodes[i].character_alignment == nil {
			continue
		}
		fk := *nodes[i].character_alignment
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(alignment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "character_alignment" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CharacterQuery) loadCharacterAbilityScores(ctx context.Context, query *CharacterAbilityScoreQuery, nodes []*Character, init func(*Character), assign func(*Character, *CharacterAbilityScore)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Character)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CharacterAbilityScore(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(character.CharacterAbilityScoresColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.character_character_ability_scores
		if fk == nil {
			return fmt.Errorf(`foreign-key "character_character_ability_scores" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "character_character_ability_scores" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CharacterQuery) loadCharacterSkills(ctx context.Context, query *CharacterSkillQuery, nodes []*Character, init func(*Character), assign func(*Character, *CharacterSkill)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Character)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CharacterSkill(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(character.CharacterSkillsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.character_character_skills
		if fk == nil {
			return fmt.Errorf(`foreign-key "character_character_skills" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "character_character_skills" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *CharacterQuery) loadCharacterProficiencies(ctx context.Context, query *CharacterProficiencyQuery, nodes []*Character, init func(*Character), assign func(*Character, *CharacterProficiency)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Character)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.CharacterProficiency(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(character.CharacterProficienciesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.character_character_proficiencies
		if fk == nil {
			return fmt.Errorf(`foreign-key "character_character_proficiencies" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "character_character_proficiencies" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CharacterQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CharacterQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(character.Table, character.Columns, sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, character.FieldID)
		for i := range fields {
			if fields[i] != character.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CharacterQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(character.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = character.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedCharacterAbilityScores tells the query-builder to eager-load the nodes that are connected to the "character_ability_scores"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithNamedCharacterAbilityScores(name string, opts ...func(*CharacterAbilityScoreQuery)) *CharacterQuery {
	query := (&CharacterAbilityScoreClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedCharacterAbilityScores == nil {
		cq.withNamedCharacterAbilityScores = make(map[string]*CharacterAbilityScoreQuery)
	}
	cq.withNamedCharacterAbilityScores[name] = query
	return cq
}

// WithNamedCharacterSkills tells the query-builder to eager-load the nodes that are connected to the "character_skills"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithNamedCharacterSkills(name string, opts ...func(*CharacterSkillQuery)) *CharacterQuery {
	query := (&CharacterSkillClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedCharacterSkills == nil {
		cq.withNamedCharacterSkills = make(map[string]*CharacterSkillQuery)
	}
	cq.withNamedCharacterSkills[name] = query
	return cq
}

// WithNamedCharacterProficiencies tells the query-builder to eager-load the nodes that are connected to the "character_proficiencies"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CharacterQuery) WithNamedCharacterProficiencies(name string, opts ...func(*CharacterProficiencyQuery)) *CharacterQuery {
	query := (&CharacterProficiencyClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedCharacterProficiencies == nil {
		cq.withNamedCharacterProficiencies = make(map[string]*CharacterProficiencyQuery)
	}
	cq.withNamedCharacterProficiencies[name] = query
	return cq
}

// CharacterGroupBy is the group-by builder for Character entities.
type CharacterGroupBy struct {
	selector
	build *CharacterQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CharacterGroupBy) Aggregate(fns ...AggregateFunc) *CharacterGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CharacterGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterQuery, *CharacterGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CharacterGroupBy) sqlScan(ctx context.Context, root *CharacterQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CharacterSelect is the builder for selecting fields of Character entities.
type CharacterSelect struct {
	*CharacterQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CharacterSelect) Aggregate(fns ...AggregateFunc) *CharacterSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CharacterSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterQuery, *CharacterSelect](ctx, cs.CharacterQuery, cs, cs.inters, v)
}

func (cs *CharacterSelect) sqlScan(ctx context.Context, root *CharacterQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
