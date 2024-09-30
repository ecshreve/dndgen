// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/characterabilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// CharacterAbilityScoreQuery is the builder for querying CharacterAbilityScore entities.
type CharacterAbilityScoreQuery struct {
	config
	ctx              *QueryContext
	order            []characterabilityscore.OrderOption
	inters           []Interceptor
	predicates       []predicate.CharacterAbilityScore
	withCharacter    *CharacterQuery
	withAbilityScore *AbilityScoreQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*CharacterAbilityScore) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CharacterAbilityScoreQuery builder.
func (casq *CharacterAbilityScoreQuery) Where(ps ...predicate.CharacterAbilityScore) *CharacterAbilityScoreQuery {
	casq.predicates = append(casq.predicates, ps...)
	return casq
}

// Limit the number of records to be returned by this query.
func (casq *CharacterAbilityScoreQuery) Limit(limit int) *CharacterAbilityScoreQuery {
	casq.ctx.Limit = &limit
	return casq
}

// Offset to start from.
func (casq *CharacterAbilityScoreQuery) Offset(offset int) *CharacterAbilityScoreQuery {
	casq.ctx.Offset = &offset
	return casq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (casq *CharacterAbilityScoreQuery) Unique(unique bool) *CharacterAbilityScoreQuery {
	casq.ctx.Unique = &unique
	return casq
}

// Order specifies how the records should be ordered.
func (casq *CharacterAbilityScoreQuery) Order(o ...characterabilityscore.OrderOption) *CharacterAbilityScoreQuery {
	casq.order = append(casq.order, o...)
	return casq
}

// QueryCharacter chains the current query on the "character" edge.
func (casq *CharacterAbilityScoreQuery) QueryCharacter() *CharacterQuery {
	query := (&CharacterClient{config: casq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := casq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := casq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(characterabilityscore.Table, characterabilityscore.FieldID, selector),
			sqlgraph.To(character.Table, character.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, characterabilityscore.CharacterTable, characterabilityscore.CharacterColumn),
		)
		fromU = sqlgraph.SetNeighbors(casq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAbilityScore chains the current query on the "ability_score" edge.
func (casq *CharacterAbilityScoreQuery) QueryAbilityScore() *AbilityScoreQuery {
	query := (&AbilityScoreClient{config: casq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := casq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := casq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(characterabilityscore.Table, characterabilityscore.FieldID, selector),
			sqlgraph.To(abilityscore.Table, abilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, characterabilityscore.AbilityScoreTable, characterabilityscore.AbilityScoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(casq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CharacterAbilityScore entity from the query.
// Returns a *NotFoundError when no CharacterAbilityScore was found.
func (casq *CharacterAbilityScoreQuery) First(ctx context.Context) (*CharacterAbilityScore, error) {
	nodes, err := casq.Limit(1).All(setContextOp(ctx, casq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{characterabilityscore.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) FirstX(ctx context.Context) *CharacterAbilityScore {
	node, err := casq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CharacterAbilityScore ID from the query.
// Returns a *NotFoundError when no CharacterAbilityScore ID was found.
func (casq *CharacterAbilityScoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = casq.Limit(1).IDs(setContextOp(ctx, casq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{characterabilityscore.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) FirstIDX(ctx context.Context) int {
	id, err := casq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CharacterAbilityScore entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CharacterAbilityScore entity is found.
// Returns a *NotFoundError when no CharacterAbilityScore entities are found.
func (casq *CharacterAbilityScoreQuery) Only(ctx context.Context) (*CharacterAbilityScore, error) {
	nodes, err := casq.Limit(2).All(setContextOp(ctx, casq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{characterabilityscore.Label}
	default:
		return nil, &NotSingularError{characterabilityscore.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) OnlyX(ctx context.Context) *CharacterAbilityScore {
	node, err := casq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CharacterAbilityScore ID in the query.
// Returns a *NotSingularError when more than one CharacterAbilityScore ID is found.
// Returns a *NotFoundError when no entities are found.
func (casq *CharacterAbilityScoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = casq.Limit(2).IDs(setContextOp(ctx, casq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{characterabilityscore.Label}
	default:
		err = &NotSingularError{characterabilityscore.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := casq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CharacterAbilityScores.
func (casq *CharacterAbilityScoreQuery) All(ctx context.Context) ([]*CharacterAbilityScore, error) {
	ctx = setContextOp(ctx, casq.ctx, ent.OpQueryAll)
	if err := casq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CharacterAbilityScore, *CharacterAbilityScoreQuery]()
	return withInterceptors[[]*CharacterAbilityScore](ctx, casq, qr, casq.inters)
}

// AllX is like All, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) AllX(ctx context.Context) []*CharacterAbilityScore {
	nodes, err := casq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CharacterAbilityScore IDs.
func (casq *CharacterAbilityScoreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if casq.ctx.Unique == nil && casq.path != nil {
		casq.Unique(true)
	}
	ctx = setContextOp(ctx, casq.ctx, ent.OpQueryIDs)
	if err = casq.Select(characterabilityscore.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) IDsX(ctx context.Context) []int {
	ids, err := casq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (casq *CharacterAbilityScoreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, casq.ctx, ent.OpQueryCount)
	if err := casq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, casq, querierCount[*CharacterAbilityScoreQuery](), casq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) CountX(ctx context.Context) int {
	count, err := casq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (casq *CharacterAbilityScoreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, casq.ctx, ent.OpQueryExist)
	switch _, err := casq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (casq *CharacterAbilityScoreQuery) ExistX(ctx context.Context) bool {
	exist, err := casq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CharacterAbilityScoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (casq *CharacterAbilityScoreQuery) Clone() *CharacterAbilityScoreQuery {
	if casq == nil {
		return nil
	}
	return &CharacterAbilityScoreQuery{
		config:           casq.config,
		ctx:              casq.ctx.Clone(),
		order:            append([]characterabilityscore.OrderOption{}, casq.order...),
		inters:           append([]Interceptor{}, casq.inters...),
		predicates:       append([]predicate.CharacterAbilityScore{}, casq.predicates...),
		withCharacter:    casq.withCharacter.Clone(),
		withAbilityScore: casq.withAbilityScore.Clone(),
		// clone intermediate query.
		sql:  casq.sql.Clone(),
		path: casq.path,
	}
}

// WithCharacter tells the query-builder to eager-load the nodes that are connected to
// the "character" edge. The optional arguments are used to configure the query builder of the edge.
func (casq *CharacterAbilityScoreQuery) WithCharacter(opts ...func(*CharacterQuery)) *CharacterAbilityScoreQuery {
	query := (&CharacterClient{config: casq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	casq.withCharacter = query
	return casq
}

// WithAbilityScore tells the query-builder to eager-load the nodes that are connected to
// the "ability_score" edge. The optional arguments are used to configure the query builder of the edge.
func (casq *CharacterAbilityScoreQuery) WithAbilityScore(opts ...func(*AbilityScoreQuery)) *CharacterAbilityScoreQuery {
	query := (&AbilityScoreClient{config: casq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	casq.withAbilityScore = query
	return casq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CharacterAbilityScore.Query().
//		GroupBy(characterabilityscore.FieldScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (casq *CharacterAbilityScoreQuery) GroupBy(field string, fields ...string) *CharacterAbilityScoreGroupBy {
	casq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CharacterAbilityScoreGroupBy{build: casq}
	grbuild.flds = &casq.ctx.Fields
	grbuild.label = characterabilityscore.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score,omitempty"`
//	}
//
//	client.CharacterAbilityScore.Query().
//		Select(characterabilityscore.FieldScore).
//		Scan(ctx, &v)
func (casq *CharacterAbilityScoreQuery) Select(fields ...string) *CharacterAbilityScoreSelect {
	casq.ctx.Fields = append(casq.ctx.Fields, fields...)
	sbuild := &CharacterAbilityScoreSelect{CharacterAbilityScoreQuery: casq}
	sbuild.label = characterabilityscore.Label
	sbuild.flds, sbuild.scan = &casq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CharacterAbilityScoreSelect configured with the given aggregations.
func (casq *CharacterAbilityScoreQuery) Aggregate(fns ...AggregateFunc) *CharacterAbilityScoreSelect {
	return casq.Select().Aggregate(fns...)
}

func (casq *CharacterAbilityScoreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range casq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, casq); err != nil {
				return err
			}
		}
	}
	for _, f := range casq.ctx.Fields {
		if !characterabilityscore.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if casq.path != nil {
		prev, err := casq.path(ctx)
		if err != nil {
			return err
		}
		casq.sql = prev
	}
	return nil
}

func (casq *CharacterAbilityScoreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CharacterAbilityScore, error) {
	var (
		nodes       = []*CharacterAbilityScore{}
		_spec       = casq.querySpec()
		loadedTypes = [2]bool{
			casq.withCharacter != nil,
			casq.withAbilityScore != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CharacterAbilityScore).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CharacterAbilityScore{config: casq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(casq.modifiers) > 0 {
		_spec.Modifiers = casq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, casq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := casq.withCharacter; query != nil {
		if err := casq.loadCharacter(ctx, query, nodes, nil,
			func(n *CharacterAbilityScore, e *Character) { n.Edges.Character = e }); err != nil {
			return nil, err
		}
	}
	if query := casq.withAbilityScore; query != nil {
		if err := casq.loadAbilityScore(ctx, query, nodes, nil,
			func(n *CharacterAbilityScore, e *AbilityScore) { n.Edges.AbilityScore = e }); err != nil {
			return nil, err
		}
	}
	for i := range casq.loadTotal {
		if err := casq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (casq *CharacterAbilityScoreQuery) loadCharacter(ctx context.Context, query *CharacterQuery, nodes []*CharacterAbilityScore, init func(*CharacterAbilityScore), assign func(*CharacterAbilityScore, *Character)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CharacterAbilityScore)
	for i := range nodes {
		fk := nodes[i].CharacterID
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
			return fmt.Errorf(`unexpected foreign-key "character_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (casq *CharacterAbilityScoreQuery) loadAbilityScore(ctx context.Context, query *AbilityScoreQuery, nodes []*CharacterAbilityScore, init func(*CharacterAbilityScore), assign func(*CharacterAbilityScore, *AbilityScore)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CharacterAbilityScore)
	for i := range nodes {
		fk := nodes[i].AbilityScoreID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(abilityscore.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ability_score_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (casq *CharacterAbilityScoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := casq.querySpec()
	if len(casq.modifiers) > 0 {
		_spec.Modifiers = casq.modifiers
	}
	_spec.Node.Columns = casq.ctx.Fields
	if len(casq.ctx.Fields) > 0 {
		_spec.Unique = casq.ctx.Unique != nil && *casq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, casq.driver, _spec)
}

func (casq *CharacterAbilityScoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(characterabilityscore.Table, characterabilityscore.Columns, sqlgraph.NewFieldSpec(characterabilityscore.FieldID, field.TypeInt))
	_spec.From = casq.sql
	if unique := casq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if casq.path != nil {
		_spec.Unique = true
	}
	if fields := casq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, characterabilityscore.FieldID)
		for i := range fields {
			if fields[i] != characterabilityscore.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if casq.withCharacter != nil {
			_spec.Node.AddColumnOnce(characterabilityscore.FieldCharacterID)
		}
		if casq.withAbilityScore != nil {
			_spec.Node.AddColumnOnce(characterabilityscore.FieldAbilityScoreID)
		}
	}
	if ps := casq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := casq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := casq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := casq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (casq *CharacterAbilityScoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(casq.driver.Dialect())
	t1 := builder.Table(characterabilityscore.Table)
	columns := casq.ctx.Fields
	if len(columns) == 0 {
		columns = characterabilityscore.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if casq.sql != nil {
		selector = casq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if casq.ctx.Unique != nil && *casq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range casq.predicates {
		p(selector)
	}
	for _, p := range casq.order {
		p(selector)
	}
	if offset := casq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := casq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CharacterAbilityScoreGroupBy is the group-by builder for CharacterAbilityScore entities.
type CharacterAbilityScoreGroupBy struct {
	selector
	build *CharacterAbilityScoreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (casgb *CharacterAbilityScoreGroupBy) Aggregate(fns ...AggregateFunc) *CharacterAbilityScoreGroupBy {
	casgb.fns = append(casgb.fns, fns...)
	return casgb
}

// Scan applies the selector query and scans the result into the given value.
func (casgb *CharacterAbilityScoreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, casgb.build.ctx, ent.OpQueryGroupBy)
	if err := casgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterAbilityScoreQuery, *CharacterAbilityScoreGroupBy](ctx, casgb.build, casgb, casgb.build.inters, v)
}

func (casgb *CharacterAbilityScoreGroupBy) sqlScan(ctx context.Context, root *CharacterAbilityScoreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(casgb.fns))
	for _, fn := range casgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*casgb.flds)+len(casgb.fns))
		for _, f := range *casgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*casgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := casgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CharacterAbilityScoreSelect is the builder for selecting fields of CharacterAbilityScore entities.
type CharacterAbilityScoreSelect struct {
	*CharacterAbilityScoreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cass *CharacterAbilityScoreSelect) Aggregate(fns ...AggregateFunc) *CharacterAbilityScoreSelect {
	cass.fns = append(cass.fns, fns...)
	return cass
}

// Scan applies the selector query and scans the result into the given value.
func (cass *CharacterAbilityScoreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cass.ctx, ent.OpQuerySelect)
	if err := cass.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CharacterAbilityScoreQuery, *CharacterAbilityScoreSelect](ctx, cass.CharacterAbilityScoreQuery, cass, cass.inters, v)
}

func (cass *CharacterAbilityScoreSelect) sqlScan(ctx context.Context, root *CharacterAbilityScoreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cass.fns))
	for _, fn := range cass.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cass.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}